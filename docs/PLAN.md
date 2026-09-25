# Devspaces Development Plan

## Goal

Build a Codespaces-style CLI that creates reproducible development container
images from either:

- A local project directory.
- A remote Git repository.

The resulting image should contain the project files and a selected language
runtime. Podman will be the default container runtime, with Docker support as an
optional alternative.

## Proposed CLI

```bash
# Build from the current directory
devspaces create --language python --tag my-app:dev .

# Build from another local directory
devspaces create --language go --tag api:dev ./api

# Build from a Git repository
devspaces create \
  --repo https://github.com/example/api.git \
  --language go \
  --tag api:dev

# Start an interactive development container
devspaces up --name api-dev api:dev
```

## Phase 1: Define the Configuration Model

Replace package-level configuration variables with a single configuration
structure:

```go
type Config struct {
  Source          string
  Repository      string
  GitRef          string
  Language        string
  LanguageVersion string
  Runtime         string
  ImageTag        string
  ContainerName   string
  Containerfile   string
  Port            int
}
```

Functions should receive a `Config` or explicit arguments instead of reading
flags directly.

## Phase 2: Introduce Subcommands

Start with these commands:

```text
devspaces create 
devspaces up
devspaces version
```

Add these after the initial workflow works:

```text
devspaces stop
devspaces remove
devspaces push
devspaces deploy
```

Use `flag.NewFlagSet` for each subcommand. Keep Kubernetes settings such as
namespace and replicas under the future `deploy` command because they do not
affect image creation.

## Phase 3: Clean Up Flag Parsing

Keep `main()` focused on orchestration:

1. Determine the subcommand.
2. Parse its flags.
3. Create a `Config` value.
4. Call the appropriate command function.
5. Print errors and exit with the appropriate status.

Resolve the current flag conflicts and validation issues:

- Decide whether `-r` means runtime or replicas; it cannot mean both.
- Remove checks such as `strconv.Itoa(value) == ""`, which can never be true.
- Make the default port agree with the accepted port range.
- Document that the standard Go flag package expects flags before positional
  arguments.

## Phase 4: Validate Configuration

Add focused validation functions:

```go
func validateBuildConfig(config Config) error
func validateLanguage(language string) error
func validateRuntime(runtime string) error
func validateSource(source string) error
```

Initial supported values:

```text
Languages: python, go, node, java,
Runtimes:  podman, docker
```

Return useful errors rather than calling `panic` for expected user mistakes.

## Phase 5: Prepare the Project Source

Support local and remote sources through separate functions:

```go
func prepareLocalSource(path string) (string, error)
func cloneRepository(url, ref string) (directory string, cleanup func(), err error)
```

For local directories:

- Resolve the absolute path.
- Confirm the path exists and is a directory.
- Respect `.containerignore` or `.dockerignore`.
- Avoid including secrets, `.git`, and build artifacts when ignored.

For Git repositories:

- Clone into a temporary directory.
- Optionally check out a branch, tag, or commit supplied through `--ref`.
- Return a cleanup function for removing the temporary checkout.
- Never print credentials or place them inside image layers.

## Phase 6: Find or Generate a Container Definition

Use the following precedence:

1. A file explicitly supplied with `--file`.
2. `Containerfile` or `Dockerfile` in the source directory.
3. A generated Dockerfile/Containerfile for the selected language.

The existing `findContainerDefinition` function is the starting point for this
phase.

An initial generated Containerfile can be as simple as:

```dockerfile
FROM python:3.14-slim
WORKDIR /workspace
COPY . .
CMD ["/bin/sh"]
```

## Phase 7: Add Language Profiles

Represent each supported language with a profile:

```go
type LanguageProfile struct {
  Name         string
  DefaultImage string
  Workdir      string
}
```

Initial profiles:

```text
python -> python:<selected-version>-slim
go     -> golang:<selected-version>
node   -> node:<selected-version>-slim
java   -> eclipse-temurin:<selected-version>-jdk
```

Later, detect dependency files and install dependencies efficiently:

- Python: `requirements.txt` or `pyproject.toml`.
- Go: `go.mod` and `go.sum`.
- Node: `package.json` and its lock file.
- Java: `pom.xml` or Gradle files.

## Phase 8: Build the Image

Create an abstraction for Podman and Docker:

```go
type ContainerRuntime interface {
  Build(config BuildConfig) error
  Run(config RunConfig) error
}
```

The first implementation should run the equivalent of:

```bash
podman build -t my-app:dev -f Containerfile .
```

Use `exec.CommandContext` with each argument passed separately. Connect the
child process to the terminal so users can see build progress:

```go
cmd.Stdout = os.Stdout
cmd.Stderr = os.Stderr
```

Do not construct commands with `sh -c`.

## Phase 9: Start a Development Container

The initial `up` command should:

- Start an interactive shell.
- Use `/workspace` as the working directory.
- Publish the requested port.
- Assign a predictable container name.
- Remove the container on exit when requested.

Example underlying command:

```bash
podman run --rm -it \
  --name api-dev \
  -p 8080:8080 \
  api:dev \
  /bin/sh
```

Later, support a bind-mounted source directory so local edits appear inside the
container without rebuilding the image.

## Phase 10: Add Tests

Create table-driven unit tests for:

- Flag parsing.
- Configuration validation.
- Dockerfile and Containerfile discovery.
- Language profile selection.
- Generated Containerfile contents.
- Local versus Git source selection.

Use `t.TempDir()` for filesystem tests. Put real Podman and Docker builds behind
optional integration tests so normal unit tests remain fast.

## Phase 11: Add Registry Support

After local builds work reliably, add:

```bash
devspaces push --registry <registry> my-app:dev
```

Then add tagging and authentication support for AWS Elastic Container Registry.
Do not accept registry passwords as ordinary command-line arguments.

## Phase 12: Add Kubernetes Deployment

Add deployment only after build, run, and push work independently:

```bash
devspaces deploy \
  --namespace default \
  --replicas 1 \
  my-app:dev
```

Building, running, pushing, and deploying should remain separate operations so
each step can be tested and retried independently.

## First Release Milestone

The first useful release should support:

```bash
devspaces build --language python --tag demo:dev .
devspaces up --name demo demo:dev
```

Acceptance criteria:

- Accept a local project directory.
- Detect an existing Dockerfile or Containerfile.
- Generate a Containerfile when neither exists.
- Support Python and Go language profiles.
- Build the image with Podman.
- Start an interactive development container.
- Show build and runtime output in the terminal.
- Return clear errors without panicking.

Once this works end to end, add remote Git repositories as the next milestone.
