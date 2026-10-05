// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Package cli implements the devspaces command-line interface.
package cli

import "github.com/spf13/cobra"

// Run runs devspaces with the arguments that follow the program name.
// With no subcommand it starts the interactive wizard.
func Run(args []string) error {
	root := newRootCmd()
	root.SetArgs(args)
	return root.Execute()
}

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:     "devspaces",
		Short:   "Create and run development container images",
		Version: version,
		Args:    cobra.NoArgs,
		// main prints the error; usage is only shown for flag/arg mistakes.
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runWizard()
		},
	}

	root.AddCommand(
		newCreateCmd(),
		newUpCmd(),
		newVersionCmd(),
	)
	return root
}
