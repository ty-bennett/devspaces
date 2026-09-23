package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
)

// CONSTANTS for container defaults
const DEFAULT_CONTAINER_RUNTIME = "podman"
const DEFAULT_LANGUAGE = "python"

// CONSTANTS for deployment defaults
const DEFAULT_NAMESPACE = "default"
const DEFAULT_DEPLOYMENT_REPLICAS = 1
const DEFAULT_CONTAINER_ACCESS_PORT = 22222

var languageImages = map[string]string{
	"python": "docker.io/library/python:3.14-slim",
	"node":   "docker.io/library/node:22-slim",
	"go":     "docker.io/library/golang:1.23-alpine",
	"java":   "docker.io/library/eclipse-temurin:21-jdk",
	"rust":   "docker.io/library/rust:1.82-slim",
}

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

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	a := answers{
		directory:           pwd,
		language:            DEFAULT_LANGUAGE,
		containerRuntime:    DEFAULT_CONTAINER_RUNTIME,
		namespace:           DEFAULT_NAMESPACE,
		replicas:            strconv.Itoa(DEFAULT_DEPLOYMENT_REPLICAS),
		containerAccessPort: strconv.Itoa(DEFAULT_CONTAINER_ACCESS_PORT),
	}
	theme := huh.ThemeCharm()

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
				Title("Language / runtime").
				Description("Base language for the container image").
				Options(
					huh.NewOption("Python", "python"),
					huh.NewOption("Node.js", "node"),
					huh.NewOption("Go", "go"),
					huh.NewOption("Java", "java"),
					huh.NewOption("Rust", "rust"),
				).
				Value(&a.language),
			huh.NewSelect[string]().
				Title("Container runtime").
				Description("Which engine should build and run the container").
				Options(
					huh.NewOption("Podman", "podman"),
					huh.NewOption("Docker", "docker"),
				).
				Value(&a.containerRuntime),
		),
	).WithTheme(theme)

	if err := intro.Run(); err != nil {
		reportFormError(err)
	}

	// Now that the language is known, seed the image default before asking.
	a.containerImage = languageImages[a.language]

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
		reportFormError(err)
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
}

func reportFormError(err error) {
	if err == huh.ErrUserAborted {
		fmt.Println("Cancelled.")
		os.Exit(0)
	}
	fmt.Println("Error running wizard:", err)
	os.Exit(1)
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
