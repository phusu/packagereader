// Package packagereader reads package information from given file location and creates
// an internal representation/data structure based on file contents.
// File must follow "Syntax of control files" of the "Debian Policy Manual".
// In Debian/Ubuntu Linux system package information is given in
// /var/lib/dpkg/status.
package packagereader

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"regexp"
	"strings"
)

// PackageReader contains the packages that are read from the file.
type PackageReader struct {
	packages map[string]*LinuxPackage
}

// Constants for recognizing different prefixes in the file
const packagePrefix = "Package:"
const descriptionPrefix = "Description:"
const dependsPrefix = "Depends:"
const whitespace = " "
const blankLine = " ."
const lineBreak = "\n"

// NewPackageReader constructs a new package reader object.
// Recommended instead of creating directly the object, as this function initializes underlying
// data structures correctly.
func NewPackageReader() *PackageReader {
	p := new(PackageReader)
	p.packages = make(map[string]*LinuxPackage)
	return p
}

// Packages returns a map of the packages. Key to the map
// is the package name as a string, value is a pointer to the LinuxPackage struct
// describing the package contents.
func (pr *PackageReader) Packages() map[string]*LinuxPackage {
	return pr.packages
}

// ParseFile reads and parses the contents of a given file.
// Returns an error if parsing was not successful. Common errors: wrong filename / missing file.
func (pr *PackageReader) ParseFile(fileName string) error {
	err := pr.readFileContents(fileName, false)
	if err != nil {
		return err
	}
	err = pr.readFileContents(fileName, true)
	return err
}

// Reads and parses the contents of a given file.
//
// Note that you need to call this method twice:
// 1) Call with packagesAlreadyScanned=false. This will read and parse the file contents and
// create the packages that really exist in the system. Reverse dependencies are not handled this time.
// 2) Call with packagesAlreadyScanned=true. This will read and parse the file contents and handle
// the reverse dependencies of each package, filling the package information if and only if the package
// exists.
//
// This behavior is needed because there might be alternates in the dependency list and when reading
// first time the file contents, we don't know which ones of the alternates exist in the system (and
// we don't want to create packages that don't exist).
//
// Example: Package A depends on Package B or C. When reading through the information for package A,
// we don't know which one of the alternate packages B or C really exist in the system.
func (pr *PackageReader) readFileContents(fileName string, packagesAlreadyScanned bool) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var packageName string
	var simpleDescription string
	var readingDescription = false
	var dependencies []string
	var packageDescriptionBuffer bytes.Buffer

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, packagePrefix):
			packageName = strings.TrimSpace(line[len(packagePrefix):])
			readingDescription = false
		case strings.HasPrefix(line, descriptionPrefix):
			simpleDescription = strings.TrimSpace(line[len(descriptionPrefix):])
			readingDescription = true
		case strings.HasPrefix(line, dependsPrefix):
			dependencies = parseDependencies(strings.TrimSpace(line[len(dependsPrefix):]))
		case len(line) == 0:
			readingDescription = false

			p := NewLinuxPackage(packageName, simpleDescription, packageDescriptionBuffer.String())
			for _, dependency := range dependencies {
				if packagesAlreadyScanned {
					_, packageExists := pr.packages[dependency]
					p.AddDependency(dependency, packageExists)
				} else {
					p.AddDependency(dependency, false)
				}
			}

			pr.packages[packageName] = p

			if packagesAlreadyScanned && dependencies != nil {
				pr.handleReverseDependencies(packageName, dependencies)
			}
			packageDescriptionBuffer.Reset()
			dependencies = nil
		case readingDescription:
			if strings.HasPrefix(line, whitespace) {
				if line == blankLine {
					packageDescriptionBuffer.WriteString(lineBreak)
				} else {
					packageDescriptionBuffer.WriteString(strings.TrimSpace(line))
					packageDescriptionBuffer.WriteString(lineBreak)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

// Parses package dependencies stripping out the version information.
// Example input: "libc6 (>= 2.14), debconf (>= 0.5) | debconf-2.0"
// Example output: [libc6, debconf, debconf-2.0]
func parseDependencies(line string) []string {
	regExp := regexp.MustCompile(`,|\|`)
	dependencies := regExp.Split(line, -1)
	cleanedDependencies := make([]string, 0, len(dependencies))
	for _, item := range dependencies {
		idx := strings.Index(item, "(")
		if idx != -1 {
			item = item[:idx]
		}
		item = strings.TrimSpace(item)
		cleanedDependencies = append(cleanedDependencies, item)
	}
	return cleanedDependencies
}

// Handles reverse dependencies for a given package and its dependencies.
// Updates the reverse dependency information for the package names listed in dependencies
// by adding packageName as a reverse dependency for each package in dependencies.
func (pr *PackageReader) handleReverseDependencies(packageName string, dependencies []string) {
	for _, item := range dependencies {
		elem, ok := pr.packages[item]
		if ok {
			_, packageExists := pr.packages[packageName]
			elem.AddReverseDependency(packageName, packageExists)
		}
	}
}
