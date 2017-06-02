package packagereader

import (
	"testing"
)

func TestNewLinuxPackage(t *testing.T) {
	p := NewLinuxPackage("name", "simpledescription", "extendeddescription")
	if p.Name() != "name" {
		t.Error("Expected name")
	}
	if p.SimpleDescription() != "simpledescription" {
		t.Error("Expected simpledescription")
	}
	if p.ExtendedDescription()[0] != "extendeddescription" {
		t.Error("Expected extendeddescription")
	}
}

func TestSetSimpleDescription(t *testing.T) {
	p := NewLinuxPackage("name", "simpledescription", "extendeddescription")
	if p.SimpleDescription() != "simpledescription" {
		t.Error("Expected simpledescription")
	}
	p.SetSimpleDescription("newsimpledescription")
	if p.SimpleDescription() != "newsimpledescription" {
		t.Error("Expected newsimpledescription")
	}
}
func TestSetExtendedDescription(t *testing.T) {
	p := NewLinuxPackage("name", "simpledescription", "extendeddescription")
	if p.ExtendedDescription()[0] != "extendeddescription" {
		t.Error("Expected extendeddescription")
	}
	p.SetExtendedDescription("newextendeddescription")
	if p.ExtendedDescription()[0] != "newextendeddescription" {
		t.Error("Expected newextendeddescription")
	}
}

func TestAddDependency(t *testing.T) {
	p := NewLinuxPackage("name", "simpledescription", "extendeddescription")
	p.AddDependency("dependency", false)
	for key, value := range p.Dependencies() {
		if key == "dependency" && !value {
			return
		}
	}
	t.Error("Expected true")
}

func TestAddReverseDependency(t *testing.T) {
	p := NewLinuxPackage("name", "simpledescription", "extendeddescription")
	p.AddReverseDependency("reversedependency", false)
	for key, value := range p.ReverseDependencies() {
		if key == "reversedependency" && !value {
			return
		}
	}

	t.Error("Expected true")
}
