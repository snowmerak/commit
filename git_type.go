package main

type GitTag struct {
	Tag   string
	Emoji string
}

var GitCommitType = map[string]GitTag{
	"initial":          {Tag: "INIT", Emoji: "🎉"},
	"feature":          {Tag: "FEAT", Emoji: "✨"},
	"work in progress": {Tag: "WIP", Emoji: "🚧"},
	"improvement":      {Tag: "IMPROVE", Emoji: "🚀"},
	"breaking change":  {Tag: "BREAK", Emoji: "💥"},
	"remove code":      {Tag: "REMOVE", Emoji: "🔥"},
	"typo fix":         {Tag: "TYPO", Emoji: "✏️"},
	"bug fix":          {Tag: "FIX", Emoji: "🐛"},
	"docs or comments": {Tag: "DOCS", Emoji: "📝"},
	"style":            {Tag: "STYLE", Emoji: "💄"},
	"refactor":         {Tag: "REFACTOR", Emoji: "♻️"},
	"ci":               {Tag: "CI", Emoji: "🔧"},
	"build":            {Tag: "BUILD", Emoji: "👷"},
	"test":             {Tag: "TEST", Emoji: "🚨"},
}

func GetGitCommitType() []string {
	keys := make([]string, 0, len(GitCommitType))
	for k := range GitCommitType {
		keys = append(keys, k)
	}
	return keys
}
