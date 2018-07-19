# checkGit
Flag which Git repositories need to be pulled or have unpushed, unstaged or uncommitted changes.

[![Build Status](https://travis-ci.org/speedyhoon/checkGit.svg?branch=master)](https://travis-ci.org/speedyhoon/checkGit)
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/checkGit)](https://goreportcard.com/report/github.com/speedyhoon/checkGit)

**Install**
```
go get github.com/speedyhoon/checkGit
```

**Help**
```
Usage of checkGit.exe: [flags, ...] [paths, ...]
  -g    Display directories that are not git repositories.
  -l    Detect out of date repositories that require a pull request.
  -q    Only display repository names and hide summary.
```

**Usage:** ```checkGit -g -l path_to_dir_containing_git_repos```

Outputs:
```
BuildIt.ninja: push, uncommitted, local changes, untracked files
EventBucket: pull, local changes
foobar: Not a git repository
mindjholts: uncommitted, local changes
replace: untracked files
session: push, local changes
```

**Quiet mode:** ```checkGit -q path_to_dir_containing_git_repos```

Outputs:
```
BuildIt.ninja
EventBucket
mindjholts
replace
session
```
