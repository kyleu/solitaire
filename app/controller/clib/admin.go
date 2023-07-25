// Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"fmt"
	"runtime"
	"runtime/pprof"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/util"
	"github.com/kyleu/solitaire/views/vadmin"
)

var AppRoutesList map[string][]string

func Admin(rc *fasthttp.RequestCtx) {
	controller.Act("admin", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		path := util.StringSplitAndTrim(strings.TrimPrefix(string(rc.URI().Path()), "/admin"), "/")
		if len(path) == 0 {
			ps.Title = "Administration"
			return controller.Render(rc, as, &vadmin.Settings{}, ps, "admin")
		}
		switch path[0] {
		case "server":
			info := util.DebugGetInfo(as.BuildInfo.Version, as.Started)
			ps.Data = info
			return controller.Render(rc, as, &vadmin.ServerInfo{Info: info}, ps, "admin", "Server Information")
		case "cpu":
			switch path[1] {
			case "start":
				err := util.DebugStartCPUProfile()
				if err != nil {
					return "", err
				}
				return controller.FlashAndRedir(true, "started CPU profile", "/admin", rc, ps)
			case "stop":
				pprof.StopCPUProfile()
				return controller.FlashAndRedir(true, "stopped CPU profile", "/admin", rc, ps)
			default:
				return "", errors.Errorf("unhandled CPU profile action [%s]", path[1])
			}
		case "gc":
			timer := util.TimerStart()
			runtime.GC()
			msg := fmt.Sprintf("ran garbage collection in [%s]", timer.EndString())
			return controller.FlashAndRedir(true, msg, "/admin", rc, ps)
		case "heap":
			err := util.DebugTakeHeapProfile()
			if err != nil {
				return "", err
			}
			return controller.FlashAndRedir(true, "wrote heap profile", "/admin", rc, ps)
		case "memusage":
			x := util.DebugMemStats()
			ps.Data = x
			return controller.Render(rc, as, &vadmin.MemUsage{Mem: x}, ps, "admin", "Memory Usage")
		case "modules":
			di := util.DebugBuildInfo().Deps
			ps.Title = "Modules"
			ps.Data = di
			return controller.Render(rc, as, &vadmin.Modules{Modules: di}, ps, "admin", "Modules")
		case "request":
			ps.Title = "Request Debug"
			ps.Data = cutil.RequestCtxToMap(rc, nil)
			return controller.Render(rc, as, &vadmin.Request{RC: rc}, ps, "admin", "Request")
		case "routes":
			ps.Title = "HTTP Routes"
			ps.Data = AppRoutesList
			return controller.Render(rc, as, &vadmin.Routes{Routes: AppRoutesList}, ps, "admin", "Request")
		case "session":
			ps.Title = "Session Debug"
			ps.Data = ps.Session
			return controller.Render(rc, as, &vadmin.Session{}, ps, "admin", "Session")
		case "sitemap":
			ps.Title = "Sitemap"
			ps.Data = ps.Menu
			return controller.Render(rc, as, &vadmin.Sitemap{}, ps, "admin", "Sitemap")
		case "sockets":
			return socketRoute(rc, as, ps, path[1:]...)
		// $PF_SECTION_START(admin-actions)$
		// $PF_SECTION_END(admin-actions)$
		default:
			return "", errors.Errorf("unhandled admin action [%s]", path[0])
		}
	})
}
