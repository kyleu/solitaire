// Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"github.com/fasthttp/router"

	"github.com/kyleu/solitaire/app/controller/clib"
)

func execRoutes(r *router.Router) {
	r.GET("/admin/exec", clib.ExecList)
	r.GET("/admin/exec/new", clib.ExecForm)
	r.POST("/admin/exec/new", clib.ExecNew)
	r.GET("/admin/exec/{key}/{idx}", clib.ExecDetail)
	r.GET("/admin/exec/{key}/{idx}/connect", clib.ExecSocket)
	r.GET("/admin/exec/{key}/{idx}/kill", clib.ExecKill)
}
