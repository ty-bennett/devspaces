// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package main

type LanguageProfile struct {
    Name         string
    DefaultImage string
    Workdir      string
    DepFiles     []string  // e.g. ["requirements.txt", "pyproject.toml"]
}

type Language

var languageVersions = map[string][]string{
	"python": {"3.11", "3.12", "3.13", "3.14"},

	"go": {"1.22", "1.23", "1.24"},

	// Node.js LTS releases; JavaScript has no separate runtime image
	"javascript": {"18", "20", "22", "23"},

	// TypeScript compiles to JS — uses Node.js images at runtime
	"typescript": {"18", "20", "22", "23"},

	// GCC versions; no official "C++" image, uses gcc base image
	"c++": {"12", "13", "14"},

	"php": {"8.1", "8.2", "8.3", "8.4"},

	// 11 and 17 are LTS; 21 is current LTS; 23 is latest non-LTS
	"java": {"11", "17", "21", "23"},

	"rust": {"1.80", "1.81", "1.82", "1.83"},
}
var languageProfiles = map[string]LanguageProfile{ /* python, go, node, java */ }


func profileForLanguage(language, version string) (LanguageProfile, error)
func generateContainerfileContent(profile LanguageProfile) string