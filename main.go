package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// CONSTANTS for container defaults
const DIRECTORY = '.'
const DEFAULT_CONTAINER_RUNTIME = "podman"
const DEFAULT_LANGUAGE = "python"

// CONSTANTS for deployment defaults
const DEFAULT_NAMESPACE = "default"
const DEFAULT_DEPLOYMENT_REPLICAS = 1
const DEFAULT_CONTAINER_IMAGE = "docker.io/library/python:3.14-slim"
const DEFAULT_CONTAINER_ACCESS_PORT = 22222

// variables with proper types
var containerRuntime string
var language string

// variables for deployment with types
var namespace string
var replicas int
var containerImage string
var containerAccessPort int

func main() {

	// Parse the language flags provided by user
	flag.StringVar(&language, "language", DEFAULT_LANGUAGE, "The language to use for the container")
	flag.StringVar(&language, "l", DEFAULT_LANGUAGE, "The language to use for the container")
	// Parse the container runtime flags provided by user
	flag.StringVar(&containerRuntime, "container-runtime", DEFAULT_CONTAINER_RUNTIME, "The container runtime to use")
	flag.StringVar(&containerRuntime, "c", DEFAULT_CONTAINER_RUNTIME, "The container runtime to use")

	// Parse the deployment flags provided by user
	flag.StringVar(&namespace, "namespace", DEFAULT_NAMESPACE, "The namespace to deploy to")
	flag.StringVar(&namespace, "n", DEFAULT_NAMESPACE, "The namespace to deploy to")

	flag.IntVar(&replicas, "replicas", DEFAULT_DEPLOYMENT_REPLICAS, "The number of replicas to deploy")
	flag.IntVar(&replicas, "r", DEFAULT_DEPLOYMENT_REPLICAS, "The number of replicas to deploy")

	flag.StringVar(&containerImage, "image", DEFAULT_CONTAINER_IMAGE, "The container image to use")
	flag.StringVar(&containerImage, "i", DEFAULT_CONTAINER_IMAGE, "The container image to use")

	//
	flag.IntVar(&containerAccessPort, "port", DEFAULT_CONTAINER_ACCESS_PORT, "The port to access the container on")
	flag.IntVar(&containerAccessPort, "p", DEFAULT_CONTAINER_ACCESS_PORT, "The port to access the container on")

	// Parse the args provided by the user
	flag.Parse()

	directory := flag.Arg(0)
	directory = configureDirectory(directory)

	language = configureLanguage(language)
	containerRuntime = configureContainerRuntime(containerRuntime)
	namespace = configureNamespace(namespace)
	replicas = configureReplicas(replicas)
	containerImage = configureContainerImage(containerImage)
	containerAccessPort = configureContainerAccessPort(containerAccessPort)

	//

	PrintDir := fmt.Sprintf("Creating container with directory: %s", directory)
	PrintLanguage := fmt.Sprintf("Language: %s", language)
	PrintContainerRuntime := fmt.Sprintf("Container runtime: %s", containerRuntime)
	PrintNamespace := fmt.Sprintf("Namespace: %s", namespace)
	PrintReplicas := fmt.Sprintf("Replicas: %d", replicas)
	PrintContainerImage := fmt.Sprintf("Container image: %s", containerImage)
	PrintContainerAccessPort := fmt.Sprintf("Container access port: %d", containerAccessPort)

	fmt.Println(PrintDir)
	fmt.Println(PrintLanguage)
	fmt.Println(PrintContainerRuntime)
	fmt.Println(PrintNamespace)
	fmt.Println(PrintReplicas)
	fmt.Println(PrintContainerImage)
	fmt.Println(PrintContainerAccessPort)
}

func configureDirectory(directory string) string {
	directory = flag.Arg(0)
	if strings.TrimSpace(directory) == "" || directory == "." {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			panic(err)
		}
		directory = pwd
	}
	return directory
}

func findContainerDefinition(directory string) (string, error) {
	definitionNames := []string{"Dockerfile", "Containerfile"}

	for _, definitionName := range definitionNames {
		definitionPath := filepath.Join(directory, definitionName)
		info, err := os.Stat(definitionPath)
		if err == nil && !info.IsDir() {
			return definitionPath, nil
		}
		if err != nil && !os.IsNotExist(err) {
			return "", fmt.Errorf("check %s: %w", definitionPath, err)
		}
	}

	return "", fmt.Errorf("no Dockerfile or Containerfile found in %q", directory)
}

func configureLanguage(language string) string {
	if strings.TrimSpace(language) == "" {
		return DEFAULT_LANGUAGE
	}
	return language
}

func configureContainerRuntime(containerRuntime string) string {
	if strings.TrimSpace(containerRuntime) == "" {
		return DEFAULT_CONTAINER_RUNTIME
	}
	return containerRuntime
}

func configureNamespace(namespace string) string {
	if strings.TrimSpace(namespace) == "" {
		return DEFAULT_NAMESPACE
	}
	return namespace
}

func configureReplicas(replicas int) int {
	if strconv.Itoa(replicas) == "" || replicas < 1 || replicas > 1 {
		return DEFAULT_DEPLOYMENT_REPLICAS
	}
	return replicas
}

func configureContainerImage(containerImage string) string {
	if strings.TrimSpace(containerImage) == "" {
		return DEFAULT_CONTAINER_IMAGE
	}
	return containerImage
}

func configureContainerAccessPort(containerAccessPort int) int {
	if containerAccessPort < 32767 ||
		containerAccessPort > 65535 ||
		strconv.Itoa(containerAccessPort) == "" {
		return DEFAULT_CONTAINER_ACCESS_PORT
	}
	return containerAccessPort
}
