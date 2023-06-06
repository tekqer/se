package api

import (
	"github.com/tekqer/se/internal/hr"
	"github.com/tekqer/se/proto"
)

var svc Service

func Init(r *hr.Router) {
	app := r.Prefix("/app")
	app.POST("", createApp)
}

func createApp(c *hr.Ctx) error {
	var app proto.App
	if err := c.Bind(&app); err != nil {
		return err
	}
	return svc.CreateApp(c, &app)
}
