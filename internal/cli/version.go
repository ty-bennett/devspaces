// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// version is set at build time, for example:
//
//	go build -ldflags "-X github.com/ty-bennett/devspaces/internal/cli.version=v0.1.0"

// update when building
var version = "devspaces0.1.0"

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the devspaces version",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), "devspaces version", version)
			return nil
		},
	}
}
