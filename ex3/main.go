package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// ex3: Add subcommands to the program. Command structure is:
// ex cmd1 --cmd1flag1 --cmd1flag2 subcmd1 --subcmd1flag1 --subcmd1flag2 subcmd2 --subcmd2flag1 --subcmd2flag2
//                                 ^                                     ^
//                                 |                                     |
//                                 |                                     |
//                                 sub command: args ...                 sub command: args ...
//                                 FlagSet.Args()                        FlagSet.Args()

// We achieve this by using flag.NewFlagSet() to create a new flag set for each sub command.

// mainCmd flags
type mainCmd struct {
	cmd1Flag1 bool
	cmd1Flag2 string
	cmd2Flag3 time.Duration
}

// subCmd1 flags
type subCmd1 struct {
	subCmd1Flag1 bool
	subCmd1Flag2 string
	subCmd2Flag3 time.Duration
}

// subCmd2 flags
type subCmd2 struct {
	subCmd2Flag1 bool
	subCmd2Flag2 string
	subCmd2Flag3 time.Duration
}

// config is the configuration for the program
type config struct {
	mainCmd mainCmd // main command configuration
	subCmd1 subCmd1 // sub command 1 configuration
	subCmd2 subCmd2 // sub command 2 configuration
}

// initMainCmdFlags initializes the flags for main command
func initMainCmdFlags(c *mainCmd, set *flag.FlagSet) {
	set.BoolVar(&c.cmd1Flag1, "cmd1flag1", false, "cmd1flag1")
	set.StringVar(&c.cmd1Flag2, "cmd1flag2", "", "cmd1flag2")
	set.DurationVar(&c.cmd2Flag3, "cmd1flag3", 1*time.Millisecond, "cmd2flag3")
}

// initSubCmd1Flags initializes the flags for sub command 1
func initSubCmd1Flags(c *subCmd1, set *flag.FlagSet) {
	set.BoolVar(&c.subCmd1Flag1, "subcmd1flag1", false, "subcmd1flag1")
	set.StringVar(&c.subCmd1Flag2, "subcmd1flag2", "", "subcmd1flag2")
	set.DurationVar(&c.subCmd2Flag3, "subcmd1flag3", 2*time.Millisecond, "subcmd2flag3")
}

// initSubCmd2Flags initializes the flags for sub command 2
func initSubCmd2Flags(c *subCmd2, set *flag.FlagSet) {
	set.BoolVar(&c.subCmd2Flag1, "subcmd2flag1", false, "subcmd2flag1")
	set.StringVar(&c.subCmd2Flag2, "subcmd2flag2", "", "subcmd2flag2")
	set.DurationVar(&c.subCmd2Flag3, "subcmd2flag3", 3*time.Millisecond, "subcmd2flag3")
}

var (
	mainSet    *flag.FlagSet
	subCmd1Set *flag.FlagSet
	subCmd2Set *flag.FlagSet
)

// initFlags initializes the flags for command/subcommand
func initFlags(conf *config) {
	mainSet = flag.NewFlagSet("main", flag.ExitOnError)
	initMainCmdFlags(&conf.mainCmd, mainSet)
	subCmd1Set = flag.NewFlagSet("subcmd1", flag.ExitOnError)
	initSubCmd1Flags(&conf.subCmd1, subCmd1Set)
	subCmd2Set = flag.NewFlagSet("subcmd2", flag.ExitOnError)
	initSubCmd2Flags(&conf.subCmd2, subCmd2Set)
}

// main is the entry point
// Example:
//
//	ex3 -cmd1flag1 --cmd1flag2="main flag2" --cmd1flag3=10ms \
//		sc2 --subcmd1flag1 --subcmd1flag2="sub1 flag2" -subcmd1flag3=20ms \
//	  sc3 --subcmd2flag1 --subcmd2flag2="sub2 flag2" --subcmd2flag3=30ms
func main() {
	var conf config
	initFlags(&conf)

	// Main command
	mainSet.Parse(os.Args[1:])
	args := mainSet.Args() // Remaining args after parsing main command

	// Sub command 1
	subCmd1Set.Parse(args[1:])
	args = subCmd1Set.Args() // Remaining args after parsing sub command 1

	// Sub command 2
	subCmd2Set.Parse(args[1:])

	// Print the configuration
	fmt.Printf("mainCmd: flag1:%t, flag2:%s, flag3:%s\n", conf.mainCmd.cmd1Flag1, conf.mainCmd.cmd1Flag2, conf.mainCmd.cmd2Flag3)
	fmt.Printf("subCmd1: flag1:%t, flag2:%s, flag3:%s\n", conf.subCmd1.subCmd1Flag1, conf.subCmd1.subCmd1Flag2, conf.subCmd1.subCmd2Flag3)
	fmt.Printf("subCmd2: flag1:%t, flag2:%s, flag3:%s\n", conf.subCmd2.subCmd2Flag1, conf.subCmd2.subCmd2Flag2, conf.subCmd2.subCmd2Flag3)
}
