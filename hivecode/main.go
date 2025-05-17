package main

import (
	"fmt"

	"github.com/devops-cyberlab/hivebox/hivecode/git-handler/tagging"
)

func main() {
	version, err := tagging.GetCurrentGitTag()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Git Version:", version)
}
