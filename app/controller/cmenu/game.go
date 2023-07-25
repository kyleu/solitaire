package cmenu

import (
	"context"

	"github.com/kyleu/solitaire/app/lib/menu"
)

func gameMenu(_ context.Context) *menu.Item {
	kids := menu.Items{
		&menu.Item{Key: "html", Title: "HTML", Description: "Renders a basic game", Icon: "code", Route: "/game/test/html"},
		&menu.Item{Key: "json", Title: "JSON", Description: "Renders a game to JSON", Icon: "graph", Route: "/game/test/json"},
		&menu.Item{Key: "wasm", Title: "WASM", Description: "Renders a game using WebAssembly", Icon: "cog", Route: "/game/test/wasm"},
	}
	return &menu.Item{Key: "game", Title: "Game Tests", Description: "Good luck!", Icon: "play", Route: "/game", Children: kids}
}
