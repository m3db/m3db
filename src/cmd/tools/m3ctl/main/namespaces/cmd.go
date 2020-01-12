package namespaces

import (
	"flag"
	"fmt"
	"github.com/m3db/m3/src/cmd/tools/m3ctl/main/errors"
	"os"
)

const (
	defaultPath = "/api/v1/namespace"
	debugQS     = "debug=true"
)

type NamespaceArgs struct {
	showAll *bool
	delete  *string
}

type NamespaceFlags struct {
	Namespace     *flag.FlagSet
	NamespaceDoer func(*NamespaceArgs, string)
	Delete        *flag.FlagSet
	DeleteDoer    func(*NamespaceArgs, string)
}

func SetupFlags(flags *NamespaceArgs) NamespaceFlags {
	namespaceFlags := flag.NewFlagSet("ns", flag.ExitOnError)
	deleteFlags := flag.NewFlagSet("delete", flag.ExitOnError)
	flags.delete = deleteFlags.String("name", "", "name of namespace to delete")

	flags.showAll = namespaceFlags.Bool("all", false, "show all the standard info for namespaces (otherwise default behaviour lists only the names)")
	namespaceFlags.Usage = func() {
		fmt.Fprintf(namespaceFlags.Output(), `
This is the subcommand for acting on namespaces.

Description:

The namespaces subcommand "%s"" provides the ability to:

* list all namespaces (default)
* verbosely list all the available information about the namespaces (-all)
* delete a specific namespace (see the delete subcommand)

Default behaviour (no arguments) is to print out the names of the namespaces.

Specify only one action at a time.

It has the following subcommands:

	%s

Usage:

`, namespaceFlags.Name(), deleteFlags.Name())
		namespaceFlags.PrintDefaults()
	}
	deleteFlags.Usage = func() {
		fmt.Fprintf(deleteFlags.Output(), `
This is the "%s" subcommand for %s scoped operations.

Description:

This subcommand allows the creation of a new database from a yaml specification.

Usage of %s:

`, deleteFlags.Name(), namespaceFlags.Name(), deleteFlags.Name())
		deleteFlags.PrintDefaults()
	}
	return NamespaceFlags{Namespace: namespaceFlags, NamespaceDoer: Show, Delete: deleteFlags, DeleteDoer: Delete}
}

func ParseAndDo(args *NamespaceArgs, flags *NamespaceFlags, ep string) {
	osArgs := flag.Args()
	// right here args should be like "ns delete -name someName"
	if len(osArgs) < 1 {
		flags.Namespace.Usage()
		os.Exit(1)
	}
	// pop and parse
	if err := parseAndDo(osArgs[1:], args, flags, ep); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func parseAndDo(args []string, finalArgs *NamespaceArgs, flags *NamespaceFlags, ep string) error {
	if err := flags.Namespace.Parse(args); err != nil {
		flags.Namespace.Usage()
		return &errors.FlagsError{}
	}
	// maybe do "ns -all"
	if flags.Namespace.NArg() == 0 {
		flags.NamespaceDoer(finalArgs, ep)
		return nil
	}
	nextArgs := flags.Namespace.Args()
	switch nextArgs[0] {
	case flags.Delete.Name():
		if err := flags.Delete.Parse(nextArgs[1:]); err != nil {
			flags.Delete.Usage()
			return &errors.FlagsError{}
		}
		if flags.Delete.NFlag() == 0 {
			flags.Delete.Usage()
			return &errors.FlagsError{}
		}
		flags.DeleteDoer(finalArgs, ep)
		return nil
	default:
		flags.Namespace.Usage()
		return &errors.FlagsError{}
	}
}
