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
// Example: ex3 --project="MyFav" --agent "GCP_ID" get --verbose --duration=100s
//  flag.Parse() parses till "get" and flag.Args() returns the rest of the arguments.
//  Means that --verbose --duration=100s are not parsed by flag.Parse() and are available in flag.Args().

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
	flag.Parse() // Parsed till a non-flag argument is found.

	// cmdline: ex2  --project="MyFav" --agent "GCP_ID" get --verbose --duration=100s

	// Prints: Args: --project=MyFav, --agent, GCP_ID, get, --verbose, --duration=100s
	fmt.Println("Args:", strings.Join(args, ", ")) // ParseDuration parses a duration string.
	// Prints: Flags: projectID:MyFav agentID:GCP_ID verbose:false filename:, timeout:99ms
	fmt.Printf("Flags: projectID:%s agentID:%s verbose:%t filename:%s, timeout:%s\n", conf.projectID, conf.agentID, conf.verbose, conf.filename, conf.timeout)
	// Prints: Args: get, --verbose, --duration=100s
	fmt.Printf("Args: %s\n", strings.Join(flag.Args(), ", "))

}
