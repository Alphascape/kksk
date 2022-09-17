package main

import (
	. "github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

type menu struct {
	Compo

	Iclass string

	appInstallable bool
}

func newMenu() *menu {
	return &menu{}
}

func (m *menu) Class(v string) *menu {
	m.Iclass = AppendClass(m.Iclass, v)
	return m
}

func (m *menu) OnNav(ctx Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *menu) OnAppInstallChange(ctx Context) {
	m.appInstallable = ctx.IsAppInstallable()
}

func (m *menu) Render() UI {
	linkClass := "link heading fit unselectable"

	isFocus := func(path string) string {
		if Window().URL().Path == path {
			return "focus"
		}
		return ""
	}

	return ui.Scroll().
		Class("menu").
		Class(m.Iclass).
		HeaderHeight(headerHeight).
		Header(
			ui.Stack().
				Class("fill").
				Middle().
				Content(
					Header().Body(
						A().
							Class("heading").
							Class("app-title").
							Href("/").
							Text("KKSK"),
					),
				),
		).
		Content(
			Nav().Body(
				Div().Class("separator"),

				ui.Link().
					Class(linkClass).
					Icon(homeSVG).
					Label("Home").
					Href("/").
					Class(isFocus("/")),

				Div().Class("separator"),

				ui.Link().
					Class(linkClass).
					Icon(imgFolderSVG).
					Label("Dinner").
					Href("/dinner").
					Class(isFocus("/dinner")),

				Div().Class("separator"),

				ui.Link().
					Class(linkClass).
					Icon(twitterSVG).
					Label("Twitter").
					Href(twitterURL),
				ui.Link().
					Class(linkClass).
					Icon(githubSVG).
					Label("GitHub").
					Href(githubURL),

				Div().Class("separator"),

				If(m.appInstallable,
					ui.Link().
						Class(linkClass).
						Icon(downloadSVG).
						Label("Install").
						OnClick(m.installApp),
				),

				Div().Class("separator"),
			),
		)
}

func (m *menu) installApp(ctx Context, e Event) {
	ctx.NewAction(installApp)
}
