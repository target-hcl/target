package targetdir

import (
	"fmt"
	"os"
	"runtime"
)

// HomeFolder returns the users homefolder this will be $HOME on windows and mac and
// USERPROFILE on windows
func HomeFolder() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")

	} else {
		return os.Getenv("HOME")
	}

	// return os.Getenv("HOME")
}

// TargetHome returns the location of the target
// folder, usually $HOME/.targer
func TargetHome() string {
	return fmt.Sprintf("%s/.target/profiles.hcl", HomeFolder())
}

// TargetHomeCreate checks for the target directory
// and profiles.hcl file and creates if they don't exist
func TargetHomeCreate() {
	if _, err := os.Stat(TargetHome()); os.IsNotExist(err) {
		os.Mkdir(TargetHome(), 0755)
	}
	//functional test
	fmt.Println("creating target home")
}
