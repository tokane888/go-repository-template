# go-repository-template

Template repository for Go monorepo architecture

## Directory Structure

```sh
.
├── services/          # Services
│   └── sample/       # Sample service
│       ├── cmd/
│       ├── .env/
│       └── go.mod
├── pkg/              # Shared modules
│   └── logger/       # Logging
│       └── go.mod
└── README.md
```

## Development Environment Setup

- Start the devcontainer

## Design Principles

- Directory structure follows [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ja.md#standard-go-project-layout)
- Multiple services managed in a Go monorepo
- Each service has its own go.mod
- Shared modules placed in `pkg/` directory
  - Referenced locally via replace directives
- Architecture follows Clean Architecture
- Commit messages follow Conventional Commits

## TODO When Using This Template

- If not using devcontainer
  - Delete the .devcontainer directory
- Rename `services/sample/` to the actual service name
- Create new services under `services/`
- Search for "TODO: " in the repository and address them
- Search for "go-repository-template" in the repository and update
- Delete CLAUDE.md and regenerate with `/init` inside claude
- If not using Claude Code, find and delete related files with:
  - `find . -name '*claude*' -not -path './.git/*'`
- Remove unused services and README.md under services as needed
- To enable code review by Claude Code, run the `claude` command and then install the GitHub app with:
  - `/install-github-app`
    - See the [official docs](https://docs.anthropic.com/en/docs/claude-code/github-actions) for details

## Service Run Example

```bash
# Run the sample service
cd services/sample
go run ./cmd/sample
```

## Service Debug Run Example

- Open the "RUN AND DEBUG" menu with ctrl+shift+d
- Select the service to debug from the top menu
- Press F5 to start debugging
