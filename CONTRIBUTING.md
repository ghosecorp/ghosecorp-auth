# Contributing to ghosecorp-auth

First off, thank you for taking the time to contribute! Contributions are what make the open source community such an amazing place to learn, inspire, and create.

This guide outlines our development workflow and standard guidelines.

## Code of Conduct

By participating in this project, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md).

## Getting Started

1. **Fork** the repository and clone it locally.
2. Ensure you have the required prerequisites:
   - **Go** (v1.25+ recommended)
   - **Python** (v3.9+ for database setup script)
   - **PostgreSQL** database
3. Copy environment configurations and setup your environment.

## Branching Guidelines

We follow a simple branching model:
* `main` / `master` is the production-ready code.
* Create a feature branch off of the development or main branch for your work:
  ```bash
  git checkout -b feat/your-feature-name
  # or
  git checkout -b fix/bug-description
  ```

## Code Style & Standards

### Go Code Guidelines
* Format your code using `go fmt` before committing.
* Run `go vet ./...` to verify there are no common correctness issues.
* Organize files following the Clean Architecture folders (`internal/domain`, `internal/repository`, `internal/usercase`).
* Use meaningful variable and function names. Add documentation comments on public symbols.

### SQL Guidelines
* Write ANSI-compliant SQL.
* Ensure keys, constraints, and indexes are clearly documented in [auth_schemas.sql](setup/auth_schemas.sql).

## Pull Request Process

1. Create a Pull Request (PR) against the main repository.
2. Provide a detailed PR description explaining the **why** and **what** of your changes, along with verification logs.
3. Ensure all tests and static analysis pass locally:
   - For Go services: `go test ./...`
4. Request reviews from the maintainers. Once approved, it will be merged.

## Reporting Bugs or Feature Requests

* Use GitHub Issues to report bugs or submit feature requests.
* Describe the bug clearly, listing the exact steps to reproduce, expected behavior, and actual behavior.
