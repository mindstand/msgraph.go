// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsArchitecture undocumented
type WindowsArchitecture string

const (
	// WindowsArchitectureVNone undocumented
	WindowsArchitectureVNone WindowsArchitecture = "None"
	// WindowsArchitectureVX86 undocumented
	WindowsArchitectureVX86 WindowsArchitecture = "X86"
	// WindowsArchitectureVX64 undocumented
	WindowsArchitectureVX64 WindowsArchitecture = "X64"
	// WindowsArchitectureVArm undocumented
	WindowsArchitectureVArm WindowsArchitecture = "Arm"
	// WindowsArchitectureVNeutral undocumented
	WindowsArchitectureVNeutral WindowsArchitecture = "Neutral"
)

// WindowsArchitecturePNone returns a pointer to WindowsArchitectureVNone
func WindowsArchitecturePNone() *WindowsArchitecture {
	v := WindowsArchitectureVNone
	return &v
}

// WindowsArchitecturePX86 returns a pointer to WindowsArchitectureVX86
func WindowsArchitecturePX86() *WindowsArchitecture {
	v := WindowsArchitectureVX86
	return &v
}

// WindowsArchitecturePX64 returns a pointer to WindowsArchitectureVX64
func WindowsArchitecturePX64() *WindowsArchitecture {
	v := WindowsArchitectureVX64
	return &v
}

// WindowsArchitecturePArm returns a pointer to WindowsArchitectureVArm
func WindowsArchitecturePArm() *WindowsArchitecture {
	v := WindowsArchitectureVArm
	return &v
}

// WindowsArchitecturePNeutral returns a pointer to WindowsArchitectureVNeutral
func WindowsArchitecturePNeutral() *WindowsArchitecture {
	v := WindowsArchitectureVNeutral
	return &v
}
