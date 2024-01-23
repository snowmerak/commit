package main

type GitTag struct {
	Tag   string
	Emoji string
}

var GitCommitType = map[string]GitTag{
	"initial":                {Tag: "INIT", Emoji: "🎉"},
	"feature":                {Tag: "FEAT", Emoji: "✨"},
	"improvement":            {Tag: "IMPR", Emoji: "🚀"},
	"breaking change":        {Tag: "BREAK", Emoji: "💥"},
	"append code or feature": {Tag: "ADD", Emoji: "🔧"},
	"remove code or feature": {Tag: "REM", Emoji: "🔥"},
	"typo fix":               {Tag: "TYPO", Emoji: "✏️"},
	"bug fix":                {Tag: "FIX", Emoji: "🐛"},
	"critical hotfix":        {Tag: "HOTFIX", Emoji: "🚑"},
	"security fix":           {Tag: "SEC", Emoji: "🔒"},
	"docs or comments":       {Tag: "DOCS", Emoji: "📝"},
	"refactor or reformat":   {Tag: "REFAC", Emoji: "🎨"},
	"ci script":              {Tag: "CI", Emoji: "⚙️"},
	"build script":           {Tag: "BUILD", Emoji: "👷"},
	"add or update test":     {Tag: "TEST", Emoji: "🚨"},
	"metadata or package":    {Tag: "META", Emoji: "📦"},
	"add dependency":         {Tag: "DEP", Emoji: "➕"},
	"remove dependency":      {Tag: "DEP", Emoji: "➖"},
	"update dependency":      {Tag: "DEP", Emoji: "⬆️"},
	"downgrade dependency":   {Tag: "DEP", Emoji: "⬇️"},
}

func GetGitCommitType() []string {
	keys := make([]string, 0, len(GitCommitType))
	for k := range GitCommitType {
		keys = append(keys, k)
	}
	return keys
}
