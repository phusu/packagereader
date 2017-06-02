package packagereader

import (
	"testing"
)

func TestNewPackageReader(t *testing.T) {
	pr := NewPackageInfoReader()

	// Use existing test file
	err := pr.ParseFile("..\\teststatus")
	if err != nil {
		t.Error("File should exists")
	}

	// Package test-a
	elem, ok := pr.Packages()["test-a"]
	if !ok {
		t.Error("Package test-a not found, expected ok")
	}
	if elem.Name() != "test-a" {
		t.Error("Wrong name for package test-a: " + elem.Name())
	}
	if elem.SimpleDescription() != "Description for package test-a." {
		t.Error("Wrong description for package test-a: " + elem.SimpleDescription())
	}
	if elem.ExtendedDescription()[0] != "This is a longer description for package test-a." {
		t.Error("Wrong extended description for package test-a.")
	}
	if elem.ExtendedDescription()[1] != "Package test-a does this and that." {
		t.Error("Wrong extended description for package test-a.")
	}
	if elem.ExtendedDescription()[2] != "" {
		t.Error("Wrong extended description for package test-a.")
	}
	if elem.ExtendedDescription()[3] != "Also test for blank line." {
		t.Error("Wrong extended description for package test-a.")
	}
	if elem.ExtendedDescription()[4] != "" {
		t.Error("Wrong extended description for package test-a.")
	}
	if _, ok := elem.Dependencies()["test-b"]; !ok {
		t.Error("Missing dependency for package test-a")
	}
	if elem.Maintainer() != "Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>" {
		t.Error("Wrong maintainer for package test-a: " + elem.Maintainer())
	}
	if elem.Architecture() != "amd64" {
		t.Error("Wrong architecture for package test-a: " + elem.Architecture())
	}
	if elem.Version() != "3.3" {
		t.Error("Wrong version for package test-a: " + elem.Version())
	}

	// Package test-b
	elem, ok = pr.Packages()["test-b"]
	if !ok {
		t.Error("Package test-b not found, expected ok")
	}
	if elem.Name() != "test-b" {
		t.Error("Wrong name for package test-b: " + elem.Name())
	}
	if elem.SimpleDescription() != "Description for package test-b." {
		t.Error("Wrong description for package test-b: " + elem.SimpleDescription())
	}
	if elem.ExtendedDescription()[0] != "This is a longer description for package test-b." {
		t.Error("Wrong extended description for package test-b.")
	}
	if elem.ExtendedDescription()[1] != "Package test-b does this and that." {
		t.Error("Wrong extended description for package test-b.")
	}
	if elem.ExtendedDescription()[2] != "" {
		t.Error("Wrong extended description for package test-b.")
	}
	if elem.ExtendedDescription()[3] != "Also test for blank line." {
		t.Error("Wrong extended description for package test-b.")
	}
	if elem.ExtendedDescription()[4] != "" {
		t.Error("Wrong extended description for package test-b.")
	}
	if len(elem.Dependencies()) != 0 {
		t.Error("test-b shouldn't have dependencies")
	}
	if elem.Maintainer() != "Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>" {
		t.Error("Wrong maintainer for package test-b: " + elem.Maintainer())
	}
	if elem.Architecture() != "all" {
		t.Error("Wrong architecture for package test-b: " + elem.Architecture())
	}
	if elem.Version() != "1.1" {
		t.Error("Wrong version for package test-b: " + elem.Version())
	}

	// Package test-c
	elem, ok = pr.Packages()["test-c"]
	if !ok {
		t.Error("Package test-c not found, expected ok")
	}
	if elem.Name() != "test-c" {
		t.Error("Wrong name for package test-c: " + elem.Name())
	}
	if elem.SimpleDescription() != "Description for package test-c." {
		t.Error("Wrong description for package test-c: " + elem.SimpleDescription())
	}
	if elem.ExtendedDescription()[0] != "This is a longer description for package test-c." {
		t.Error("Wrong extended description for package test-c.")
	}
	if elem.ExtendedDescription()[1] != "Package test-c does this and that." {
		t.Error("Wrong extended description for package test-c.")
	}
	if elem.ExtendedDescription()[2] != "" {
		t.Error("Wrong extended description for package test-c.")
	}
	if elem.ExtendedDescription()[3] != "Also test for blank line." {
		t.Error("Wrong extended description for package test-c.")
	}
	if elem.ExtendedDescription()[4] != "" {
		t.Error("Wrong extended description for package test-c.")
	}
	if _, ok := elem.Dependencies()["test-b"]; !ok {
		t.Error("Missing dependency for package test-a")
	}
	if elem.Maintainer() != "Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>" {
		t.Error("Wrong maintainer for package test-c: " + elem.Maintainer())
	}
	if elem.Architecture() != "amd64" {
		t.Error("Wrong architecture for package test-c: " + elem.Architecture())
	}
	if elem.Version() != "2.0" {
		t.Error("Wrong version for package test-c: " + elem.Version())
	}

	// Test for missing package
	elem, ok = pr.Packages()["test-d"]
	if ok {
		t.Error("Package test-c should not exist")
	}
}
