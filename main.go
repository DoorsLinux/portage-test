package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"
)

func main() {
	parser := argparse.NewParser("portage-test", "Tests the current directory for what portage can build.")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

  if (current_dir_can_build_ebs()) {
     fmt.Println("current directory can be built using `portage build .`")
  } 
  if current_dir_can_build_pbs() {
    fmt.Println("current directory can be used online");
  }

  fmt.Println("Dir evaluation complete.");
}
