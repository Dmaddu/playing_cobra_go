package main

import (
	"fmt"
	"os"

	"github.com/Dmaddu/playing_cobra_go/cmd"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
