// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func createCmd() *cobra.Command {
	cfg := Config{}

	cmd := &cobra.Command{
		Use:   "create [dir]",
		Short: "Build a development image from a local directory or Git repository",
		Example: `  devspaces create --language python --tag my-app:dev .
  devspaces create --repo https://github.com/example/api.git --language go --tag api:dev`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				cfg.Source = args[0]
			} else if cfg.Repository == "" {
				cfg.Source = "."
			}
			return runCreate(cfg)
		},
	}

	f := cmd.Flags()
	f.StringVarP(&cfg.Language, "language", "l", defaultLanguage, "language of the project")
	f.StringVar(&cfg.LanguageVersion, "language-version", "", "language version (defaults to the profile's image)")
	f.StringVarP(&cfg.Runtime, "runtime", "r", defaultContainerRuntime, "container runtime (podman or docker)")
	f.StringVarP(&cfg.ImageTag, "tag", "t", "", "tag for the built image, e.g. my-app:dev")
	f.StringVar(&cfg.Repository, "repo", "", "Git repository to build instead of a local directory")
	f.StringVar(&cfg.GitRefHash, "ref", "", "branch, tag, or commit to check out with --repo")
	f.StringVarP(&cfg.Containerfile, "file", "f", "", "path to a Containerfile or Dockerfile")
	f.IntVarP(&cfg.Port, "port", "p", 0, "port to expose from the image")

	return cmd
}

// runCreate is the command boundary for image creation. It orchestrates the
// source, language, and container packages.
func runCreate(cfg Config) error {
	if err := cfg.ValidateBuild(); err != nil {
		return err
	}
	return fmt.Errorf("create command is not implemented")
}
