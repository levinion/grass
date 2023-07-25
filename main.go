package main

import (
	"github.com/levinion/grass/app"
	"github.com/levinion/grass/client"
	"github.com/levinion/grass/flax"
	"github.com/levinion/grass/message"
	"github.com/levinion/grass/server"
)

func main() {
	flax.App("").Func(func(c *flax.Context) {
		applist := app.CreateAppList()
		server.Serve(applist)
	}).ExactArgs(0)

	flax.App("add").Func(func(c *flax.Context) {
		for _, v := range c.Args {
			client.SendMsg(message.Add, v)
		}
	}) //add Expect feature

	flax.App("stop").Func(func(c *flax.Context) {
		for _, v := range c.Args {
			client.SendMsg(message.Stop, v)
		}
	})

	flax.App("remove").Func(func(c *flax.Context) {
		for _, v := range c.Args {
			client.SendMsg(message.Remove, v)
		}
	})

	flax.App("start").Func(func(c *flax.Context) {
		for _, v := range c.Args {
			client.SendMsg(message.Start, v)
		}
	})

	flax.App("reload").Func(func(c *flax.Context) {
		for _, v := range c.Args {
			client.SendMsg(message.Reload, v)
		}
	})

	flax.App("show").Func(func(c *flax.Context) {
		client.SendMsg(message.Show, "")
	})

	flax.Run()
}
