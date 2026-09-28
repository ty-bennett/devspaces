// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ty-bennett/devspaces/internal/container"
	"github.com/ty-bennett/devspaces/internal/language"
)

// Config is the command-level configuration. Command parsing should populate
// this type; the packages below cli should receive their own smaller option
// types instead.
type Config struct {
	Source          string
	Repository      string
	GitRefHash      string
	Language        string
	LanguageVersion string
	Runtime         string
	ImageTag        string
	ContainerName   string
	Containerfile   string
	Port            int
}

// Validate preserves the original build-oriented validation entry point.
func (c Config) Validate() error {
	return c.ValidateBuild()
}

// ValidateBuild verifies the inputs required to prepare and build an image.
func (c Config) ValidateBuild() error {
	if (strings.TrimSpace(c.Source) == "") == (strings.TrimSpace(c.Repository) == "") {
		return errors.New("specify exactly one of source or repository")
	}
	if err := language.Validate(c.Language); err != nil {
		return err
	}
	if err := container.ValidateRuntime(c.Runtime); err != nil {
		return err
	}
	if strings.TrimSpace(c.ImageTag) == "" {
		return errors.New("image tag cannot be empty")
	}
	if c.Port < 0 || c.Port > 65535 {
		return fmt.Errorf("port must be between 0 and 65535")
	}
	return nil
}

// ValidateRun verifies the inputs required to start an existing image.
func (c Config) ValidateRun() error {
	if err := container.ValidateRuntime(c.Runtime); err != nil {
		return err
	}
	if strings.TrimSpace(c.ImageTag) == "" {
		return errors.New("image tag cannot be empty")
	}
	if c.Port < 0 || c.Port > 65535 {
		return fmt.Errorf("port must be between 0 and 65535")
	}
	return nil
}
