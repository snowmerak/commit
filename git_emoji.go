package main

import "fmt"

var GitEmoji = map[string]string{
	"reformat code":       "🎨",
	"improve performance": "⚡️",
	"remove code":         "🔥",
	"fix bug":             "🐛",
	"fix critical issue":  "🚑",
	"new feature":         "✨",
	"document":            "📝",
	"test":                "✅",
	"generate files":      "🗃",
	"comments to code":    "💡",
	"merge to upstream":   "🚀",
	"work in progress":    "🚧",
	"fix typo":            "✏️",
	"breaking changes":    "💥",
	"design architecture": "🏗",
}

func FormatGitEmoji() []string {
	const format = "%-3s: %s"
	formatted := make([]string, 0, len(GitEmoji))
	for k, v := range GitEmoji {
		formatted = append(formatted, fmt.Sprintf(format, v, k))
	}
	return formatted
}
