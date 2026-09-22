# devspaces
A cli tool to create containers for different development environments. Utilizes Podman in the backend for rootless runtime. Written in Go.


## Overview

The goal is to create a CLI tool that I can use to spin up different development containers that reflect different projets I am working on. I am sure that something like this already exists but I wanted to do it for me and for my own cluster. I am using Golang because it works well with cloud-native tooling such as Docker, Podman, K8s, etc. I also am using a cloud hosted CR (Elastic Container Registry), because I have credits for AWS, and I can authenticate first before running the jobs to spin-up the container images, and pull them down using my information and identity; so super secure 😎. 

## Installation

Idk how to make an app so this will be blank for now. The goal would be to curl the latest tarball from my domain. 
E.g. `https://tybennett.net/tools/devspaces/install.sh` Then that will download the zip for you to install and add to your PATH. 

## How to use

First, you'll want to add the files to your PATH so you can call the `devspaces` command from anywhere.

### Example
```
curl -sSL https://tybennett.net/tools/devspaces/install.sh | bash -
export PATH="$PATH:$(go env GOPATH)/bin"
# then update your shell config
# linux
source ~/.bashrc
# mac (me)
source ~/.zshrc

# Then, to create a container using current dir we pass the first arg. Then all proceeding flags will be for customization.
~/> devspaces .
# For a full list of options, run `devspaces --help`
~/> devspaces --help, -h, or help
# For more info on a specific command, run `devspaces <command> --help`
~/> devspaces <command> --help
# For a specific runtime
~/> devspaces . --runtime=java --version=25.0, shorthand is -r=java, -v=25.0
```
Ok that's it for now.

##
