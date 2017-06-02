package packagereader

import "strings"

// LinuxPackage represents an installed package in a Debian/Ubuntu Linux environment, formed
// by data in /var/lib/dpkg/status file. Contains attributes that are only required for
// this specific assignment:
// * Simple description (one line)
// * Simple description (one line)
// * The names of the packages that depend on the current package
// * Extended description (multiple lines)
// * The names of the packages the current package depends on (skips version numbers)
// * Package name
type LinuxPackage struct {
	name                string
	simpleDescription   string
	extendedDescription string
	depends             map[string]bool
	reverseDepends      map[string]bool
}

// NewLinuxPackage creates a new linux package with given name, simple description and extended description.
func NewLinuxPackage(name, simpleDescription, extendedDescription string) *LinuxPackage {
	p := new(LinuxPackage)
	p.name = name
	p.simpleDescription = simpleDescription
	p.extendedDescription = extendedDescription
	p.depends = make(map[string]bool)
	p.reverseDepends = make(map[string]bool)
	return p
}

// Name returns the name of the package
func (p *LinuxPackage) Name() string {
	return p.name
}

// SimpleDescription returns the simple description of the package.
func (p *LinuxPackage) SimpleDescription() string {
	return p.simpleDescription
}

// SetSimpleDescription sets the simple description of the package.
func (p *LinuxPackage) SetSimpleDescription(description string) {
	p.simpleDescription = description
}

// ExtendedDescription returns the extended description of the package as a string slice, each slice representing one line.
// When this string slice is passed to the HTML template, it is easier to insert a line break tag between each line
func (p *LinuxPackage) ExtendedDescription() []string {
	return strings.Split(p.extendedDescription, "\n")
}

// SetExtendedDescription sets the extended description of the package.
func (p *LinuxPackage) SetExtendedDescription(description string) {
	p.extendedDescription = description
}

// AddDependency adds a dependency to the package.
func (p *LinuxPackage) AddDependency(packageName string, exists bool) {
	p.depends[packageName] = exists
}

// AddReverseDependency adds a reverse dependency to the package.
func (p *LinuxPackage) AddReverseDependency(packageName string, exists bool) {
	p.reverseDepends[packageName] = exists
}

// Dependencies returns a slice of the dependencies of the package.
func (p *LinuxPackage) Dependencies() map[string]bool {
	return p.depends
}

// ReverseDependencies returns a slice of the reverse dependencies of the package.
func (p *LinuxPackage) ReverseDependencies() map[string]bool {
	return p.reverseDepends
}
