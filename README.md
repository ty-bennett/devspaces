# devspaces

> THIS IS A WIP and will be filled out as the project progresses

A CLI tool to create containers for different development environments. Utilizes Podman in the backend for rootless runtime. Written in Go.


## Overview

The goal is to create a CLI tool that I can use to spin up different development containers that reflect different projets I am working on. I am sure that something like this already exists but I wanted to do it for me and for my own cluster. I am using Golang because it works well with cloud-native tooling such as Docker, Podman, K8s, etc. I also am using a cloud hosted CR (Elastic Container Registry), because I have credits for AWS, and I can authenticate first before running the jobs to spin-up the container images, and pull them down using my information and identity; so super secure 😎. 

## Installation

The goal would be to curl the latest install script from my site. \
```https://tybennett.net/tools/devspaces/install.sh```\
This runs the shell script to install the program and add it to your PATH. 

## How to use

First, you'll want to add the files to your PATH so you can call the `devspaces` command from anywhere.
```
curl -sSL https://tybennett.net/tools/devspaces/install.sh | bash -
```
Then, add the executable to your path
```
export PATH="$PATH:$(go env GOPATH)/bin"
```
Source your config to apply the $PATH changes.
```
# linux
source ~/.bashrc
# mac (me)
source ~/.zshrc
```

Then, to create a container definition, use the `create` command to run through the TUI process to define all required variables. Defaults are preloaded for you. \
```
~/> devspaces create
```
For a full list of options, run `devspaces -h or --help or help`
```
~/> devspaces --help, -h, or help
```
For more info on a specific command, run `devspaces <command> --help, -h, or help`
```
~/> devspaces <command> --help, -h, or help
```
For a specific runtime
```
~/> devspaces create --runtime=java --version=25.0, shorthand is -r=java, -v=25.0
```
## TODO
