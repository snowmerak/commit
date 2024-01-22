package main

var GitCommitType = map[string]string{
	"initial":  "INIT",
	"feature":  "FEAT",
	"fix":      "FIX",
	"docs":     "DOCS",
	"style":    "STYLE",
	"refactor": "REFACTOR",
	"test":     "TEST",
	"chore":    "CHORE",
}

func GetGitCommitType() []string {
	keys := make([]string, 0, len(GitCommitType))
	for k := range GitCommitType {
		keys = append(keys, k)
	}
	return keys
}
