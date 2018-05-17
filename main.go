package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	nonGit := flag.Bool("g", false, "Display directories that are not git repositories.")
	pull := flag.Bool("l", false, "Detect out of date repositories that require a pull request.")
	quiet := flag.Bool("q", false, "Only display repository names and hide summary.")
	flag.Parse()

	directories := flag.Args()

	if len(directories) == 0 {
		dirRun, err := os.Executable()
		if err != nil {
			directories = []string{"."}
		} else {
			//Strip the application name from path
			directories = []string{filepath.Dir(dirRun)}
		}
	}

	for _, dir := range directories {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if filepath.Dir(path) != dir || info == nil || !info.IsDir() {
				return nil
			}

			if gitInfo, err := os.Stat(filepath.Join(path, ".git")); os.IsNotExist(err) || !gitInfo.IsDir() {
				//Directory is not a git repository
				if *nonGit {
					fmt.Println(filepath.Base(path), "Not a git repository")
				}
				return nil
			}

			var checks []string
			var changes bool
			if *pull {
				out, err := run(path, "git", "remote", "show", "origin")
				if err != nil {
					return err
				}
				if bytes.Contains(out, []byte(" (local out of date)")) {
					checks = append(checks, "pull")
					changes = true
				}
			}

			out, err := run(path, "git", "status")
			out = bytes.TrimSpace(out)

			if bytes.Contains(out, []byte("\nYour branch is ahead of ")) {
				checks = append(checks, "push")
				changes = true
			}
			if bytes.Contains(out, []byte("\nChanges to be committed:")) {
				checks = append(checks, "uncommitted")
				changes = true
			}
			if bytes.Contains(out, []byte("\nChanges not staged for commit:")) {
				checks = append(checks, "local changes")
				changes = true
			}
			if bytes.Contains(out, []byte("\nUntracked files:")) {
				checks = append(checks, "untracked files")
				changes = true
			}

			if changes {
				if *quiet {
					fmt.Println(filepath.Base(path))
				} else {
					fmt.Printf("%v: %v\n", filepath.Base(path), strings.Join(checks, ", "))
				}
			}
			return err
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}

func run(dir, command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	if dir != "" {
		cmd.Dir = dir
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	var errs []string
	var output []byte
	scanOut := bufio.NewScanner(stdout)
	go func() {
		//Gather output from external command
		for scanOut.Scan() {
			output = append(output, []byte(fmt.Sprintf("%v\n", scanOut.Text()))...)
		}
		//Collect stdout scanner error
		if err = scanOut.Err(); err != nil {
			errs = append(errs, "stdout scan err: "+err.Error())
		}
	}()

	scanErr := bufio.NewScanner(stderr)
	go func() {
		//Collect all errors returned from stderr and scanner errors
		for scanErr.Scan() {
			errs = append(errs, scanErr.Text())
		}
		if err = scanErr.Err(); err != nil {
			errs = append(errs, "stderr scan err: "+err.Error())
		}
	}()

	if err = cmd.Start(); err != nil {
		errs = append(errs, err.Error())
	}

	if err = cmd.Wait(); err != nil {
		errs = append(errs, err.Error())
	}

	//Return all the errors from stdout, stderr, start & wait
	if len(errs) > 0 {
		return output, fmt.Errorf(strings.Join(errs, "\n"))
	}
	return output, nil
}
