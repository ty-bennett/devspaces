// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

func validateBuildConfig(cfg Config) error {
	// TODO:
	// validate that the buildocnfig lines up
	// use linter and make sure all fields are filled out
	return nil
}

func validateLanguage(language string) error {
	// "python"|"go"|"node"|"java"
	// TODO:
	return nil
}

func validateRuntime(runtime string) error {
	// TODO:
	//  "podman"|"docker"
	return nil
}

func validateSource(source string) error {
	// TODO:
	// non-empty, exists if local path
	return nil
}