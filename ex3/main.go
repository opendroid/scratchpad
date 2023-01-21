package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// ex3: Test the arguments and flags of the go command.
// flag.Parse() parses the command-line flags from os.Args[1:] until a non-flag argument is found.
// Must be called after all flags are defined and before flags are accessed by the program.

// config is the configuration for the program.
type config struct {
	projectID string
	agentID   string
	verbose   bool
	filename  string
	timeout   time.Duration
}

// initFlags initializes the flags and binds variable to a command
func initFlags(conf *config) {
	flag.StringVar(&conf.projectID, "project", "", "Project ID\nLong\ndescription")
	flag.StringVar(&conf.agentID, "agent", "", "Agent ID")
	flag.BoolVar(&conf.verbose, "verbose", false, "Verbose")
	flag.StringVar(&conf.filename, "file", "", "Input File name")
	flag.DurationVar(&conf.timeout, "duration", 99*time.Millisecond, "Duration")
}

// main is the entry point
func main() {
	args := append([]string{}, os.Args[1:]...)
	var conf config
	initFlags(&conf)
	flag.Parse()

	if x, err := time.ParseDuration("199s"); err != nil {
		fmt.Println("ParseDuration Error:", err)
	} else {
		fmt.Printf("x: %d, %[1]s\n", x)
	}

	fmt.Println("Args:", strings.Join(args, ","))
	fmt.Printf("Flags: projectID:%s agentID:%s verbose:%t filename:%s, timeout:%s\n", conf.projectID, conf.agentID, conf.verbose, conf.filename, conf.timeout)
}
