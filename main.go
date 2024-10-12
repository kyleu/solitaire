package main // import github.com/kyleu/solitaire

import (
	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/cmd"
)

var (
	version = "0.1.27" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(&app.BuildInfo{Version: version, Commit: commit, Date: date})
}
