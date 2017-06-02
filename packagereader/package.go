package packagereader

import "strings"

// PackageInfo represents an installed package in a Debian/Ubuntu Linux environment, formed
// by data in /var/lib/dpkg/status file. Contains attributes:
// * Simple description (one line)
// * The names of the packages that depend on the current package
// * Extended description (multiple lines)
// * The names of the packages the current package depends on (skips version numbers)
// * Name
// * Maintainer info
// * Architecture
// * Version number
type PackageInfo struct {
	name                string
	simpleDescription   string
	extendedDescription string
	maintainer          string
	architecture        string
	version             string
	depends             map[string]bool
	reverseDepends      map[string]bool
}

// NewPackageInfo creates a new PackageInfo with given details.
func NewPackageInfo(name, simpleDescription, extendedDescription, maintainer, architecture, version string) *PackageInfo {
	p := new(PackageInfo)
	p.name = name
	p.simpleDescription = simpleDescription
	p.extendedDescription = extendedDescription
	p.maintainer = maintainer
	p.architecture = architecture
	p.version = version
	p.depends = make(map[string]bool)
	p.reverseDepends = make(map[string]bool)
	return p
}

// Name returns the name of the package
func (p *PackageInfo) Name() string {
	return p.name
}

// SimpleDescription returns the simple description of the package.
func (p *PackageInfo) SimpleDescription() string {
	return p.simpleDescription
}

// SetSimpleDescription sets the simple description of the package.
func (p *PackageInfo) SetSimpleDescription(description string) {
	p.simpleDescription = description
}

// ExtendedDescription returns the extended description of the package as a string slice, each slice representing one line.
// When this string slice is passed to the HTML template, it is easier to insert a line break tag between each line
func (p *PackageInfo) ExtendedDescription() []string {
	return strings.Split(p.extendedDescription, "\n")
}

// SetExtendedDescription sets the extended description of the package.
func (p *PackageInfo) SetExtendedDescription(description string) {
	p.extendedDescription = description
}

// Maintainer returns the package maintainer.
func (p *PackageInfo) Maintainer() string {
	return p.maintainer
}

// Architecture returns the package architecture.
func (p *PackageInfo) Architecture() string {
	return p.architecture
}

// Version returns the package version.
func (p *PackageInfo) Version() string {
	return p.version
}

// AddDependency adds a dependency to the package.
func (p *PackageInfo) AddDependency(packageName string, exists bool) {
	p.depends[packageName] = exists
}

// AddReverseDependency adds a reverse dependency to the package.
func (p *PackageInfo) AddReverseDependency(packageName string, exists bool) {
	p.reverseDepends[packageName] = exists
}

// Dependencies returns a slice of the dependencies of the package.
func (p *PackageInfo) Dependencies() map[string]bool {
	return p.depends
}

// ReverseDependencies returns a slice of the reverse dependencies of the package.
func (p *PackageInfo) ReverseDependencies() map[string]bool {
	return p.reverseDepends
}
