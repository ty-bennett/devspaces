// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Package container builds images and runs containers.
package container

import (
	"context"
	"fmt"
)

// Runtime is the only interface the CLI needs from Docker or Podman. It also
// makes command construction testable without invoking a real daemon.
type Runtime interface {
	Build(context.Context, BuildOptions) error
	Run(context.Context, RunOptions) error
}

type BuildOptions struct {
	Tag           string
	Containerfile string
	ContextDir    string
	Runtime       string
}

type RunOptions struct {
	Image         string
	ContainerName string
	Port          int
	Workdir       string
	RemoveOnExit  bool
}

// ValidateRuntime verifies that a supported command-line runtime was chosen.
func ValidateRuntime(name string) error {
	switch name {
	case "docker", "podman":
		return nil
	default:
		return fmt.Errorf("unsupported container runtime %q", name)
	}
}

// NewRuntime returns the implementation selected by name. The command-backed
// implementation is intentionally left for the build phase.
func NewRuntime(name string) (Runtime, error) {
	if err := ValidateRuntime(name); err != nil {
		return nil, err
	}
	return commandRuntime{name: name}, nil
}

type commandRuntime struct{ name string }

func (r commandRuntime) Build(context.Context, BuildOptions) error {
	return fmt.Errorf("%s image builds are not implemented", r.name)
}

func (r commandRuntime) Run(context.Context, RunOptions) error {
	return fmt.Errorf("%s container runs are not implemented", r.name)
}
