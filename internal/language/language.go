// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

// Package language describes the languages devspaces can build images for.
package language

import "fmt"

type Profile struct {
	Name         string
	DefaultImage string
	Workdir      string
	DepFiles     []string // e.g. ["requirements.txt", "pyproject.toml"]
}

var languageImages = map[string]string{
	"python":      "docker.io/library/python:3.14",
	"python-slim": "docker.io/library/python:3.14-slim",
	"node":        "docker.io/library/node:22-slim",

	"go":   "docker.io/library/golang:1.23-alpine",
	"java": "docker.io/library/eclipse-temurin:21-jdk",
	"rust": "docker.io/library/rust:1.82-slim",
}

var languageVersions = map[string][]string{
	"python": {"3.11", "3.12", "3.13", "3.14"},

	"go": {"1.22", "1.23", "1.24"},

	// Node.js LTS releases for TS and JS images
	"javascript": {"18", "20", "22", "23", "24"},
	"typescript": {"18", "20", "22", "23", "24"},

	// GCC versions for C/CPP
	"c++": {"12", "13", "14"},
	"c":   {"12", "13", "14"},

	// 11 and 17 are LTS; 21 is current LTS; 23 is latest non-LTS
	"java": {"11", "17", "21", "23"},

	"rust": {"1.80", "1.81", "1.82", "1.83"},
}
var languageProfiles = map[string]Profile{
	"python": {Name: "python", DefaultImage: "docker.io/library/python:3.14-slim", Workdir: "/workspace", DepFiles: []string{"requirements.txt", "pyproject.toml"}},
	"go":     {Name: "go", DefaultImage: "docker.io/library/golang:1.23-alpine", Workdir: "/workspace", DepFiles: []string{"go.mod", "go.sum"}},
	"node":   {Name: "node", DefaultImage: "docker.io/library/node:22-slim", Workdir: "/workspace", DepFiles: []string{"package.json", "package-lock.json", "pnpm-lock.yaml", "yarn.lock"}},
	"java":   {Name: "java", DefaultImage: "docker.io/library/eclipse-temurin:21-jdk", Workdir: "/workspace", DepFiles: []string{"pom.xml", "build.gradle", "build.gradle.kts"}},
	"rust":   {Name: "rust", DefaultImage: "docker.io/library/rust:1.82-slim", Workdir: "/workspace", DepFiles: []string{"Cargo.toml", "Cargo.lock"}},
}

// DefaultImage returns the default container image for a language,
// or "" if there isn't one.
func DefaultImage(name string) string {
	return languageImages[name]
}

func Validate(name string) error {
	if _, ok := languageProfiles[name]; ok {
		return nil
	}
	return fmt.Errorf("unsupported language %q", name)
}

// Get returns the profile for a supported language.
func Get(name string, version string) (Profile, error) {
	if err := Validate(name); err != nil {
		return Profile{}, err
	}
	profile, ok := languageProfiles[name]
	if !ok {
		return Profile{}, fmt.Errorf("language profile for %q is not implemented", name)
	}
	if version != "" {
		profile.DefaultImage = imageFor(name, version)
	}
	return profile, nil
}

// Containerfile returns a minimal reproducible definition for a profile.
func Containerfile(profile Profile) string {
	return fmt.Sprintf("FROM %s\nWORKDIR %s\nCOPY . .\nCMD [\"/bin/sh\"]\n", profile.DefaultImage, profile.Workdir)
}

func imageFor(name, version string) string {
	imageTemplates := map[string]string{
		"python": "docker.io/library/python:%s-slim",
		"go":     "docker.io/library/golang:%s-alpine",
		"node":   "docker.io/library/node:%s-slim",
		"java":   "docker.io/library/eclipse-temurin:%s-jdk",
		"rust":   "docker.io/library/rust:%s-slim",
	}
	template, ok := imageTemplates[name]
	if !ok {
		return ""
	}
	return fmt.Sprintf(template, version)
}
