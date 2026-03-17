# checkGit

[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/checkGit)](https://goreportcard.com/report/github.com/speedyhoon/checkGit)

Flag which Git repositories need to be pulled or have unpushed, unstaged or uncommitted changes.

### Install

```go install github.com/speedyhoon/checkGit@latest```

### Help

```
Usage of checkGit: [flags, ...] [paths, ...]
  -b    Brief summary. Legned:
                pull: L
                push: P
                uncommitted: U
                local changes: C
                untracked files: F
                not a git repository: G
  -g    Display subdirectories that aren't Git repositories.
  -git string
        Path to git executable if not in $PATH or current directory. (default "git")
  -l    Detect out of date repositories that require a pull request.
  -p    Only display repositories ahead that can be pushed.
  -q    Only display repository names and hide summary.
```

### Usage

```checkGit -g -l path_to_dir_containing_git_repos```

Outputs:

```
BuildIt.ninja: push, uncommitted, local changes, untracked files
EventBucket: pull, local changes
foobar: Not a git repository
mindjholts: uncommitted, local changes
replace: untracked files
session: push, local changes
```

### Quiet mode

```checkGit -q path_to_dir_containing_git_repos```

Outputs:

```
BuildIt.ninja
EventBucket
mindjholts
replace
session
```

### Brief summary mode

```checkGit -b path_to_dir_containing_git_repos```

Outputs:

```
BuildIt.ninja: P  U  C  F
EventBucket: L  C
mindjholts: U  C
replace: F
session: P  C
```
