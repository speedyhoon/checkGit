# checkGit
Check all Git subdirectories to see if any have uncommitted changes or need to be pushed or pulled

[![go report card](https://goreportcard.com/badge/github.com/speedyhoon/checkGit)](https://goreportcard.com/report/github.com/speedyhoon/checkGit)


```cmd
checkGit.exe path_to_projects
```

```
BuildIt.ninja
EventBucket
mindjholts
replace
session
```


**Verbose mode**
```cmd
checkGit.exe -v path_to_projects
```

```
BuildIt.ninja: uncommitted, local changes, untracked files
EventBucket: local changes
mindjholts: uncommitted, local changes
replace: untracked files
session: local changes
```
