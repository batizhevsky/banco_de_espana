package cli

import (
	"fmt"
	"os"
)

func CLIApp() {
	args := os.Args[1:]
	fmt.Println(args)

}
