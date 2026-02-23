# sdd-hello-world

Minimal SDD test fixture for cobbler-scaffold.

## Architectural Thesis

Cobbler-scaffold end-to-end tests need a stable, minimal Go project to scaffold against. Using a production SDD project couples test stability to ongoing development. This repository provides the simplest possible Go binary wrapped in the SDD project structure, isolating cobbler-scaffold validation from production project changes.

## Scope and Status

The project is complete. It consists of a single Go binary that prints "Hello, World!" and a VISION specification document. The cobbler-scaffold orchestrator scaffolds this project and runs mage targets against it during e2e testing.

## Repository Structure

```
sdd-hello-world/
  cmd/sdd-hello-world/main.go   -- Hello World binary
  docs/VISION.yaml               -- SDD vision specification
  docs/constitutions/             -- Scaffolded design constitutions
  docs/prompts/                   -- Scaffolded prompt templates
  magefiles/orchestrator.go       -- Cobbler-scaffold mage targets
  configuration.yaml              -- Orchestrator configuration
```

## Build and Test

```bash
go build -o bin/sdd-hello-world ./cmd/sdd-hello-world
./bin/sdd-hello-world
# Output: Hello, World!
```

With the orchestrator scaffolded:

```bash
mage build
mage init
mage -l
```
