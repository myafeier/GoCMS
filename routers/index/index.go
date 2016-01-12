package index

import (
	"github.com/go-macaron/renders"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func Index(ctx *macaron.Context, sess session.Store, r renders.Render) {
	ctx.Data["hi"] = "Hello world!"
	ctx.Data["username"] = sess.Get("username")
	r.HTML(200, "index/index.html", ctx.Data)
}
