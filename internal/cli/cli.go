// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Package cli implements the devspaces command-line interface.
package cli

// Run runs devspaces with the arguments that follow the program name.
// There are no subcommands yet, so it always starts the wizard.
func Run(args []string) error {
	return runWizard()
}
