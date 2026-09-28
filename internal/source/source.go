// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Package source prepares a project's files for an image build.
package source

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func PrepareLocal(path string) (absPath string, err error) {
	if err := Validate(path); err != nil {
		return "", err
	}
	return filepath.Abs(path)
}

// Clone returns the cloned dir, a cleanup func, and any error
func Clone(url, ref string) (dir string, cleanup func(), err error) {
	// TODO:
	return "", nil, errors.New("source.Clone is not implemented yet")
}

func Validate(src string) error {
	if src == "" {
		return errors.New("source path cannot be empty")
	}
	info, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("stat source: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("source %q is not a directory", src)
	}
	return nil
}

// FindContainerDefinition applies the Containerfile precedence rule: explicit
// path, then Containerfile, then Dockerfile. An empty result means generation
// by the language package is required.
func FindContainerDefinition(dir, explicitPath string) (string, error) {
	if explicitPath != "" {
		if _, err := os.Stat(explicitPath); err != nil {
			return "", fmt.Errorf("stat container definition: %w", err)
		}
		return filepath.Abs(explicitPath)
	}
	for _, name := range []string{"Containerfile", "Dockerfile"} {
		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		} else if !errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("stat container definition: %w", err)
		}
	}
	return "", nil
}
