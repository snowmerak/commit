package main

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func SelectFilesSurvey(files ...string) ([]string, error) {
	q := &survey.MultiSelect{
		Message:  "Select files",
		Options:  files,
		PageSize: 6,
	}

	var selected []string
	if err := survey.AskOne(q, &selected); err != nil {
		return nil, err
	}

	for i, v := range selected {
		selected[i] = v[prefixLength:]
	}

	return selected, nil
}

func SelectGitEmojiSurvey() (string, error) {
	formatted := FormatGitEmoji()

	q := &survey.Select{
		Message:  "Select git emoji",
		Options:  formatted,
		PageSize: 6,
	}

	var selected string
	if err := survey.AskOne(q, &selected); err != nil {
		return "", err
	}

	split := strings.SplitN(selected, ":", 2)

	return split[0], nil
}

func SelectCommitType() (string, error) {
	q := &survey.Select{
		Message:  "Select commit type",
		Options:  GetGitCommitType(),
		PageSize: 6,
	}

	var selected string
	if err := survey.AskOne(q, &selected); err != nil {
		return "", err
	}

	return GitCommitType[selected], nil
}

func WriteCommitMessage() (string, error) {
	q := &survey.Input{
		Message: "Write commit message",
	}

	var message string
	if err := survey.AskOne(q, &message); err != nil {
		return "", err
	}

	return message, nil
}

func ConfirmPush() (bool, error) {
	q := &survey.Confirm{
		Message: "Push commit?",
	}

	var confirm bool
	if err := survey.AskOne(q, &confirm); err != nil {
		return false, err
	}

	return confirm, nil
}
