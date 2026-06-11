# UnCONSULT

**A compliance automation toolkit for building a complete ISMS — conforming to both ISO 27001:2022 and SOC2.**

UnCONSULT is built on top of [Comply](https://github.com/strongdm/comply) by [StrongDM](https://www.strongdm.com/), an open-source SOC2 compliance automation tool. Full credit to the StrongDM team for the foundation this project extends — see [NOTICE](./NOTICE) and [LICENSE.txt](./LICENSE.txt) (Apache 2.0).

## What it does

- **Policy Generator**: markdown-powered document pipeline that publishes auditor-friendly policy PDFs and a static HTML dashboard
- **Dual-standard coverage**: ships with both **ISO 27001:2022 Annex A** (all 93 controls) and **SOC2 Trust Services Criteria**, with every control mapped to a satisfying policy or narrative out of the box
- **ISMS scaffolding**: includes ISMS scope, information security objectives, and Statement of Applicability narratives alongside 39 policy templates
- **Clause-grouped dashboard**: controls are organized by ISO clause family (A.5 Organizational, A.6 People, A.7 Physical, A.8 Technological) and TSC category, with satisfied/unsatisfied status per control
- **Ticketing integration**: track procedure completion through GitHub, Jira, or GitLab

## What's different from upstream Comply

- ISO 27001:2022 Annex A standard definition and full `satisfies:` mappings across all templates
- 12 additional policies and 3 ISMS narratives closing every Annex A gap (93/93 coverage)
- Dashboard renders one table per control family per standard, instead of a flat list
- 13 bug fixes (goroutine deadlocks, nil dereferences, swallowed errors, WebSocket origin checking, deprecated `io/ioutil`)
- `comply init` always provisions the dual-standard theme — no SOC2-vs-blank prompt

## Installation

Build from source (requires Go and `make`):

```bash
git clone <this-repo>
cd complyISO27K
PATH="$PATH:$(go env GOPATH)/bin" make
sudo cp comply /opt/homebrew/bin/comply   # or anywhere on your PATH
```

The binary is named `comply`, same as upstream.

### Dependencies

Rendering PDFs relies on [pandoc](https://pandoc.org/), installed either as an OS package or invoked via Docker:

```bash
brew install pandoc basictex   # macOS — recommended over Docker on Apple Silicon
```

## Get started

```bash
mkdir my-company
cd my-company
comply init      # interactive setup
comply build     # generate output/ with HTML dashboard and PDFs
comply serve     # live-reload server at :4000
```

Once `comply init` is complete, `git init` and push your project to a new repository, then begin editing the included policy text.

## CLI

```
COMMANDS:
     init             initialize a new compliance repository (interactive)
     build, b         generate a static website summarizing the compliance program
     procedure, proc  create ticket by procedure ID
     scheduler        create tickets based on procedure schedule
     serve            live updating version of the build command
     sync             sync ticket status to local cache
     todo             list declared vs satisfied compliance controls
```

## Ticketing integrations

GitHub, Jira, and GitLab are supported. Configure in `comply.yml`:

```yaml
tickets:
  github:
    repo: <repo-name>
    token: <token>
    username: org or personal username
```

Environment variables (`GITHUB_REPO`, `GITHUB_TOKEN`, `GITHUB_USERNAME`, `JIRA_USERNAME`, etc.) override values from the YAML file.

For Jira, ensure the default _Create Screen_ includes assignee, description, issuetype, labels, project key, reporter, and summary, with no other required fields. Authenticate with an [API token](https://id.atlassian.com/manage-profile/security/api-tokens).

## Development

```bash
PATH="$PATH:$(go env GOPATH)/bin" make    # regenerates embedded theme assets, then builds
go test ./...                              # run tests
```

After editing anything under `themes/`, `make assets` (or plain `make`) must be re-run to regenerate the embedded assets before the change takes effect.

## Credits and license

UnCONSULT is a derivative of [strongdm/comply](https://github.com/strongdm/comply), © strongDM, Inc., used under the [Apache License 2.0](./LICENSE.txt). Modifications are documented in [NOTICE](./NOTICE). StrongDM does not endorse this derivative work.
