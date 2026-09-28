// Copyright (c) 2026 Ty Bennett. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"

	"github.com/ty-bennett/devspaces/internal/language"
)

// CONSTANTS for container defaults
const defaultContainerRuntime = "podman"
const defaultLanguage = "python"

// CONSTANTS for deployment defaults
const defaultNamespace = "default"
const defaultDeploymentReplicas = 1
const defaultContainerAccessPort = 32222

// answers collects everything gathered from the wizard
type answers struct {
	directory           string
	language            string
	containerRuntime    string
	namespace           string
	replicas            string
	containerImage      string
	containerAccessPort string
}

func runWizard() error {
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting current directory: %w", err)
	}

	a := answers{
		directory:           pwd,
		language:            defaultLanguage,
		containerRuntime:    defaultContainerRuntime,
		namespace:           defaultNamespace,
		replicas:            strconv.Itoa(defaultDeploymentReplicas),
		containerAccessPort: strconv.Itoa(defaultContainerAccessPort),
	}

	// see about other themes
	theme := huh.ThemeCatppuccin()

	intro := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().
				Title("devspaces").
				Description("Spin up a development container for this project.\nAnswer a few questions to get started."),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Project directory").
				Description("The directory to mount into the container").
				Value(&a.directory).
				Validate(notEmpty("directory")),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Language").
				Description("Base language for the container image").
				Options(
					huh.NewOption("Python", "python"),
					huh.NewOption("Node.js", "node"),
					huh.NewOption("Go", "go"),
					huh.NewOption("Java", "java"),
					huh.NewOption("Rust", "rust"),
				).
				Value(&a.language),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Runtime Environment").
				Description("Base runtime for the container").
				Options(
					huh.NewOption("Docker", "docker"),
					huh.NewOption("Podman", "podman"),
				).
				Value(&a.containerRuntime),
		),
	).WithTheme(theme)

	err = intro.Run()
	if err != nil {
		return formError(err)
	}

	// Now that the language is known, seed the image default before asking.
	a.containerImage = language.DefaultImage(a.language)

	rest := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Container image").
				Description("Overrides the default image chosen for the selected language").
				Value(&a.containerImage).
				Validate(notEmpty("container image")),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Namespace").
				Value(&a.namespace).
				Validate(notEmpty("namespace")),
			huh.NewInput().
				Title("Replicas").
				Value(&a.replicas).
				Validate(positiveInt("replicas")),
			huh.NewInput().
				Title("Container access port").
				Description("Must be between 32768 and 65535").
				Value(&a.containerAccessPort).
				Validate(portInRange("port", 32768, 65535)),
		),
	).WithTheme(theme)

	if err := rest.Run(); err != nil {
		return formError(err)
	}

	replicas, _ := strconv.Atoi(a.replicas)
	port, _ := strconv.Atoi(a.containerAccessPort)

	fmt.Println()
	fmt.Println(huh.ThemeCharm().Focused.Title.Render("Summary"))
	fmt.Printf("Directory:       %s\n", a.directory)
	fmt.Printf("Language:        %s\n", a.language)
	fmt.Printf("Container runtime: %s\n", a.containerRuntime)
	fmt.Printf("Container image: %s\n", a.containerImage)
	fmt.Printf("Namespace:       %s\n", a.namespace)
	fmt.Printf("Replicas:        %d\n", replicas)
	fmt.Printf("Access port:     %d\n", port)
	return nil
}

// formError turns a wizard error into the error runWizard returns.
// Cancelling isn't a failure, so it prints a message and returns nil.
func formError(err error) error {
	if errors.Is(err, huh.ErrUserAborted) {
		fmt.Println("Cancelled.")
		return nil
	}
	return fmt.Errorf("running wizard: %w", err)
}

func notEmpty(field string) func(string) error {
	return func(s string) error {
		if s == "" {
			return fmt.Errorf("%s cannot be empty", field)
		}
		return nil
	}
}

func positiveInt(field string) func(string) error {
	return func(s string) error {
		n, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("%s must be a number", field)
		}
		if n < 1 {
			return fmt.Errorf("%s must be at least 1", field)
		}
		return nil
	}
}

func portInRange(field string, min, max int) func(string) error {
	return func(s string) error {
		n, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("%s must be a number", field)
		}
		if n < min || n > max {
			return fmt.Errorf("%s must be between %d and %d", field, min, max)
		}
		return nil
	}
}
