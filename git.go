package main

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	Untracked = "untracked"
	Modified  = "modified"
	Deleted   = "deleted"
	Renamed   = "renamed"
	Copied    = "copied"
)

type GitFiles struct {
	Untracked []string
	Modified  []string
	Deleted   []string
	Renamed   []string
	Copied    []string
}

func GetUnstagedFiles(dir string) (GitFiles, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return GitFiles{}, err
	}

	files := GitFiles{}
	for _, line := range strings.Split(string(out), "\n") {
		if len(line) < 4 {
			continue
		}
		switch line[0] {
		case ' ':
			continue
		case '?':
			files.Untracked = append(files.Untracked, line[3:])
		case 'M':
			files.Modified = append(files.Modified, line[3:])
		case 'D':
			files.Deleted = append(files.Deleted, line[3:])
		case 'R':
			files.Renamed = append(files.Renamed, line[3:])
		case 'C':
			files.Copied = append(files.Copied, line[3:])
		}
	}

	return files, nil
}

const prefixLength = 11

func FormatUnsagedFiles(files GitFiles) []string {
	const format = "%-10s %s"
	length := len(files.Untracked) + len(files.Modified) + len(files.Deleted) + len(files.Renamed) + len(files.Copied)
	formatted := make([]string, 0, length)
	for _, file := range files.Untracked {
		formatted = append(formatted, fmt.Sprintf(format, Untracked, file))
	}
	for _, file := range files.Modified {
		formatted = append(formatted, fmt.Sprintf(format, Modified, file))
	}
	for _, file := range files.Deleted {
		formatted = append(formatted, fmt.Sprintf(format, Deleted, file))
	}
	for _, file := range files.Renamed {
		formatted = append(formatted, fmt.Sprintf(format, Renamed, file))
	}
	for _, file := range files.Copied {
		formatted = append(formatted, fmt.Sprintf(format, Copied, file))
	}
	return formatted
}

type CommitMessage struct {
	Emoji   string
	Type    string
	Message string
}

func (c CommitMessage) String() string {
	return fmt.Sprintf("%s %s: %s", c.Emoji, c.Type, c.Message)
}

func GitAdd(files []string) error {
	args := append([]string{"add"}, files...)
	cmd := exec.Command("git", args...)
	return cmd.Run()
}

func GitCommit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	return cmd.Run()
}

func GitPush() error {
	cmd := exec.Command("git", "push")
	return cmd.Run()
}
