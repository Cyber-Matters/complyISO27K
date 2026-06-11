# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What This Is

**Comply** is a Go CLI tool for compliance automation (SOC2/ISO27K). It reads Markdown-based policies, procedures, narratives, and standards, then generates a static HTML dashboard and PDFs via Pandoc. It also integrates with ticketing systems (GitHub, Jira, GitLab) to track procedure completion.

## Commands

### Build & Run
```bash
make            # Build binary (runs assets first)
make assets     # Regenerate embedded theme assets (go-bindata-assetfs) — required after theme changes
make install    # Build and install to $GOPATH/bin
make clean      # Remove bin/ and dist/
make dist       # Cross-compile for darwin/linux/windows with version injection
```

### Testing
```bash
go test ./...                    # Run all tests
go test ./internal/model/...     # Run tests in a specific package
```

### Using the CLI (example project)
```bash
comply init      # Interactive setup — creates comply.yml, directories, example files
comply build     # Generate output/ with HTML and PDFs
comply serve     # Live-reload server at :4000
comply sync      # Pull ticket status from ticketing system
comply scheduler # Create overdue/upcoming procedure tickets
comply todo      # List controls vs satisfied tickets
```

## Architecture

The app has four distinct layers:

**CLI** (`internal/cli/`) — Built on `urfave/cli`. Commands use a `beforeCommand()` decorator chain (e.g. `projectMustExist`, `pandocMustExist`) that run validations before the handler executes.

**Model** (`internal/model/`) — Core data types (`Standard`, `Document`, `Procedure`, `Ticket`, `Audit`). The plugin registry in `model/plugin.go` defines the `TicketPlugin` interface; each ticketing backend implements it. Plugins self-register at startup (e.g. `github.Register()`).

**Render** (`internal/render/`) — Processes Markdown through a Pandoc subprocess (or Docker) to produce PDFs, and uses Hugo + Ace templates for the HTML dashboard. The `serve` command adds WebSocket-based live reload on top of the build pipeline.

**Plugins** (`internal/plugin/github/`, `internal/jira/`, `internal/gitlab/`) — Each plugin manages its own auth and uses `sync.Once` / `sync.Mutex` for thread-safe lazy initialization. Configuration comes from `comply.yml` but env vars take precedence (`GITHUB_TOKEN`, `JIRA_USERNAME`, etc.).

## Configuration

Projects are configured via `comply.yml` at the project root:
```yaml
name: Acme Corp
filePrefix: ACME
pandoc: docker          # or "pandoc" to use local binary
tickets:
  github:               # or jira: or gitlab:
    token: ...
    username: org-name
    repo: repo-name
```

See `example/comply.yml.example` for a full reference.

## Key Conventions

- **Embedded assets**: HTML themes are compiled into the binary using `go-bindata-assetfs`. After editing anything under `themes/`, run `make assets` before building.
- **Pandoc is external**: The render layer shells out to `pandoc` (or `docker run strongdm/pandoc`). Many render failures are Pandoc version or path issues, not Go bugs.
- **Content directories**: A comply project has four subdirectories — `standards/`, `narratives/`, `policies/`, `procedures/`. The `path` package handles discovery; files named `README*` are excluded automatically.
- **Satisfy syntax**: Procedures link to controls using a `satisfies:` YAML front matter field. The `todo` command computes coverage from this mapping.
- **Cron scheduling**: The `scheduler` command uses `robfig/cron` expressions in procedure front matter to determine when tickets should be opened.

## Docker

For cross-platform use or when Pandoc is unavailable locally:
```bash
# macOS Intel
docker run --rm -v "$(pwd)":/comply strongdm/comply build

# macOS M1/Apple Silicon
docker run --rm --platform linux/amd64 -v "$(pwd)":/comply strongdm/comply build
```
