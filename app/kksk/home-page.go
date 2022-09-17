package main

import (
	. "github.com/maxence-charriere/go-app/v9/pkg/app"
)

type homePage struct {
	Compo
}

func newHomePage() *homePage {
	return &homePage{}
}

func (c *homePage) Render() UI {
	return newPage().
		Title("KKSK").
		Icon("/web/logo.png").
		Content(
			Div().Body(
				P().Text("我知道你很急，但是，先别急"),
				P().Text("急也没用"),
			),
		)
}
