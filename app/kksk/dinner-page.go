package main

import . "github.com/maxence-charriere/go-app/v9/pkg/app"

type dinnerPage struct {
	Compo

	fallacy rhetologicalFallacy
}

func newDinnerPage() *dinnerPage {
	return &dinnerPage{}
}

func (c *dinnerPage) Render() UI {
	return newPage().
		Title("Dinner").
		Icon(imgFolderSVG).
		Content(
			newDinner(),
		)
}
