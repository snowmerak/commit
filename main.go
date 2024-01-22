package main

import (
	"log"
)

func main() {
	files, err := GetUnstagedFiles(".")
	if err != nil {
		panic(err)
	}

	formatted := FormatUnsagedFiles(files)

	selectedFiles, err := SelectFilesSurvey(formatted...)
	if err != nil {
		panic(err)
	}

	selectedType, err := SelectCommitType()
	if err != nil {
		panic(err)
	}

	selectedEmoji, err := SelectGitEmojiSurvey()
	if err != nil {
		panic(err)
	}

	message, err := WriteCommitMessage()
	if err != nil {
		panic(err)
	}

	cm := CommitMessage{
		Emoji:   selectedEmoji,
		Type:    selectedType,
		Message: message,
	}

	log.Println("Commit message:", cm.String())
	log.Print("git add: ")
	for _, file := range selectedFiles {
		log.Print(file, "\n")
	}

	if err := GitAdd(selectedFiles); err != nil {
		panic(err)
	}

	log.Println("git commit")
	if err := GitCommit(cm.String()); err != nil {
		panic(err)
	}

	ok, err := ConfirmPush()
	if err != nil {
		panic(err)
	}

	if !ok {
		log.Println("git push canceled")
		return
	}

	log.Println("git push")
	if err := GitPush(); err != nil {
		panic(err)
	}
}
