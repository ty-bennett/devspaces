// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newUpCmd() *cobra.Command {
	cfg := Config{}

	cmd := &cobra.Command{
		Use:     "up <image>",
		Short:   "Start an interactive container from a development image",
		Example: `  devspaces up --name api-dev api:dev`,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg.ImageTag = args[0]
			return runUp(cfg)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&cfg.ContainerName, "name", "n", "", "name for the container")
	f.StringVarP(&cfg.Runtime, "runtime", "r", defaultContainerRuntime, "container runtime (podman or docker)")
	f.IntVarP(&cfg.Port, "port", "p", 0, "port to publish from the container")

	return cmd
}

// runUp is the command boundary for starting an existing development image.
func runUp(cfg Config) error {
	if err := cfg.ValidateRun(); err != nil {
		return err
	}
	return fmt.Errorf("up command is not implemented")
}
