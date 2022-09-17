package main

import . "github.com/maxence-charriere/go-app/v9/pkg/app"

type home struct {
	Compo
}

func newHome() *home {
	return &home{}
}

func (c *home) Render() UI {
	return H1().Text("Hello, world!")
}
