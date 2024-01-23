package main

type GitTag struct {
	Tag   string
	Emoji string
}

var GitCommitType = map[string]GitTag{
	"initial":                    {Tag: "INIT", Emoji: "🎉"},
	"feature":                    {Tag: "FEAT", Emoji: "✨"},
	"breaking change":            {Tag: "BREAK", Emoji: "💥"},
	"append/fix code or feature": {Tag: "CODE", Emoji: "🔧"},
	"remove code or feature":     {Tag: "CODE", Emoji: "🔥"},
	"refactor or reformat code":  {Tag: "CODE", Emoji: "🎨"},
	"improve code or feature":    {Tag: "CODE", Emoji: "🚀"},
	"typo fix":                   {Tag: "FIX", Emoji: "✏️"},
	"bug fix":                    {Tag: "FIX", Emoji: "🐛"},
	"critical hotfix":            {Tag: "FIX", Emoji: "🚑"},
	"security fix":               {Tag: "FIX", Emoji: "🔒"},
	"docs or comments":           {Tag: "DOCS", Emoji: "📝"},
	"add or update test":         {Tag: "TEST", Emoji: "🚨"},
	"ci script":                  {Tag: "META", Emoji: "⚙️"},
	"build script":               {Tag: "META", Emoji: "👷"},
	"metadata or package":        {Tag: "META", Emoji: "📦"},
	"add dependency":             {Tag: "DEP", Emoji: "➕"},
	"remove dependency":          {Tag: "DEP", Emoji: "➖"},
	"update dependency":          {Tag: "DEP", Emoji: "⬆️"},
	"downgrade dependency":       {Tag: "DEP", Emoji: "⬇️"},
}

func GetGitCommitType() []string {
	keys := make([]string, 0, len(GitCommitType))
	for k := range GitCommitType {
		keys = append(keys, k)
	}
	return keys
}
