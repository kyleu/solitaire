// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"github.com/valyala/fasthttp"

	"github.com/kyleu/solitaire/app"
	"github.com/kyleu/solitaire/app/controller/cutil"
	"github.com/kyleu/solitaire/app/lib/sandbox"
	"github.com/kyleu/solitaire/app/lib/telemetry"
	"github.com/kyleu/solitaire/views/vsandbox"
)

func SandboxList(rc *fasthttp.RequestCtx) {
	Act("sandbox.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.SetTitleAndData("Sandboxes", sandbox.AllSandboxes)
		return Render(rc, as, &vsandbox.List{}, ps, "sandbox")
	})
}

func SandboxRun(rc *fasthttp.RequestCtx) {
	Act("sandbox.run", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}

		sb := sandbox.AllSandboxes.Get(key)
		if sb == nil {
			return ERsp("no sandbox with key [%s]", key)
		}

		ctx, span, logger := telemetry.StartSpan(ps.Context, "sandbox."+key, ps.Logger)
		defer span.Complete()

		ret, err := sb.Run(ctx, as, logger.With("sandbox", key))
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(sb.Title, ret)
		if sb.Key == "testbed" {
			return Render(rc, as, &vsandbox.Testbed{}, ps, "sandbox", sb.Key)
		}
		if sb.Key == "wasm" {
			return Render(rc, as, &vsandbox.WASM{}, ps, "sandbox", sb.Key)
		}
		return Render(rc, as, &vsandbox.Run{Key: key, Title: sb.Title, Icon: sb.Icon, Result: ret}, ps, "sandbox", sb.Key)
	})
}
