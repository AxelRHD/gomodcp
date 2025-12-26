# gomodcp

`gomodcp` is a **local-first CLI tool** for copying Go projects and **renaming the Go module cleanly** — including `go.mod` and all Go imports.

> **Everything happens locally.**
> No GitHub access, no remote templates, no network calls, no hidden magic.

---

## Why gomodcp?

Many tools that "clone" or "rename" Go projects:

- require GitHub repositories
- fetch remote templates
- modify Git history
- or only work online

`gomodcp` intentionally takes a different approach:

> **Copy an existing Go project on your local machine and adapt it to a new module path — transparently and reproducibly.**

Typical use cases:

- duplicating internal tools
- starting a new project from an existing codebase
- refactoring locally without touching Git history
- working offline or in restricted environments

---

## Features

- ✅ **100% local operation**
- ✅ no GitHub / no internet required
- ✅ AST-based import rewriting (no regex hacks)
- ✅ optional Git-aware mode (`--git`)
- ✅ `.gitignore` respected automatically
- ✅ clear CLI UX with helpful error messages
- ✅ build-time versioning
- ✅ Fish shell completion

---

## Installation

### Using `go install`

```bash
go install github.com/axelrhd/gomodcp/cmd/gomodcp@latest
```

### Build locally

```bash
task build
task save
```

---

## Usage

```text
gomodcp [flags] <src-dir> <dest-mod>
```

### Example

```bash
gomodcp ./old-project github.com/axelrhd/new-project
```

Result:

```text
./new-project/
├── go.mod        # module github.com/axelrhd/new-project
├── cmd/
├── internal/
└── ...
```

All Go imports are rewritten automatically.

---

## Arguments

### `<src-dir>`

Path to the **source Go project**.  
Must contain a `go.mod` file.

### `<dest-mod>`

Target Go module path, for example:

```text
github.com/axelrhd/my-new-tool
```

If `--dst` is not provided, the destination directory defaults to the last
path element (`my-new-tool`).

---

## Flags

### `--dst`

Explicitly set the destination directory:

```bash
gomodcp ./old github.com/axelrhd/new --dst ./sandbox/new
```

---

### `--git`

Copy **only Git-tracked files**:

```bash
gomodcp ./old github.com/axelrhd/new --git
```

Benefits:

- `.gitignore` is automatically respected
- no `vendor/`, `bin/`, `dist/`, …
- ideal for clean project copies

> ℹ️ Git is **read-only**. `gomodcp` never modifies your repository.

---

### `--version`

```bash
gomodcp --version
```

Example output:

```text
gomodcp v0.1.0 (95bbde2-dirty)
```

- app version → set at build time
- git version → derived from `git describe`
- `-dirty` → working tree has uncommitted changes

---

## Examples

### Minimal

```bash
gomodcp ./tool github.com/axelrhd/tool-v2
```

### Git-aware copy

```bash
gomodcp ./tool github.com/axelrhd/tool-v2 --git
```

### Explicit destination directory

```bash
gomodcp ./tool github.com/axelrhd/tool-v2 --dst ./experiments/tool-v2
```

---

## Versioning & Releases

`gomodcp` uses **build-time versioning** (no hardcoded versions).

### Development build

```bash
task build
```

```text
gomodcp v0.0.0-dev (<commit>-dirty)
```

### Release build

```bash
task build:0.1.0
```

```text
gomodcp v0.1.0 (<commit>)
```

Recommended release flow:

```bash
git status
task build:0.1.0
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```

---

## Why not `git mv`?

`git mv` is meant for **renaming files inside a repository**.
It is not suitable for:

- copying projects
- creating new modules
- cleanly starting a new codebase

`gomodcp` copies files without altering Git history.

---

## Why not `gonew`?

`gonew`:

- assumes GitHub repositories
- fetches remote templates
- is template-oriented, not project-oriented

`gomodcp` works directly with **your existing local code**.

---

## Non-goals (by design)

- ❌ no Git operations (commit, tag, push)
- ❌ no `go mod tidy`
- ❌ no network access

The goal is **predictability and transparency**.

---

## License

MIT License

Feel free to fork, adapt, and integrate `gomodcp` into your own workflows.
