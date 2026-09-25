// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

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

type BuildConfig struct {
	Tag           string
	Containerfile string
	ContextDir    string
	Runtime       string
}

type RunConfig struct {
	Image         string
	ContainerName string
	Port          int
	Workdir       string
	RemoveOnExit  bool
}
