package main

import (
	. "github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

const (
	headerHeight = 72
)

type page struct {
	Compo

	Iclass   string
	Iicon    string
	Ititle   string
	Icontent []UI

	updateAvailable bool
}

func newPage() *page {
	return &page{}
}

func (p *page) Icon(v string) *page {
	p.Iicon = v
	return p
}

func (p *page) Title(v string) *page {
	p.Ititle = v
	return p
}

func (p *page) Content(v ...UI) *page {
	p.Icontent = FilterUIElems(v...)
	return p
}

func (p *page) OnNav(ctx Context) {
	p.updateAvailable = ctx.AppUpdateAvailable()
	ctx.Defer(scrollTo)
}

func (p *page) OnAppUpdate(ctx Context) {
	p.updateAvailable = ctx.AppUpdateAvailable()
}

func (p *page) Render() UI {
	return ui.Shell().
		Class("fill").
		Class("background").
		HamburgerMenu(
			newMenu().
				Class("fill").
				Class("menu-hamburger-background"),
		).
		Menu(
			newMenu().Class("fill"),
		).
		Content(
			ui.Scroll().
				Class("fill").
				Header(
					Nav().
						Class("fill").
						Body(
							ui.Stack().
								Class("fill").
								Right().
								Middle().
								Content(
									If(p.updateAvailable,
										Div().
											Class("link-update").
											Body(
												ui.Link().
													Class("link").
													Class("heading").
													Class("fit").
													Class("unselectable").
													Icon(downloadSVG).
													Label("Update").
													OnClick(p.updateApp),
											),
									),
								),
						),
				).
				HeaderHeight(headerHeight).
				Content(
					Main().Body(
						Article().Body(
							Header().
								ID("page-top").
								Class("page-title").
								Class("center").
								Body(
									ui.Stack().
										Center().
										Middle().
										Content(
											ui.Icon().
												Class("icon-left").
												Class("unselectable").
												Size(90).
												Src(p.Iicon),
											H1().Text(p.Ititle),
										),
								),
							Div().Class("separator"),

							Range(p.Icontent).Slice(func(i int) UI {
								return p.Icontent[i]
							}),
						),
					),
				),
		)
}

func (p *page) updateApp(ctx Context, e Event) {
	ctx.NewAction(updateApp)
}

func scrollTo(ctx Context) {
	id := ctx.Page().URL().Fragment
	if id == "" {
		id = "page-top"
	}
	ctx.ScrollTo(id)
}

func fragmentFocus(fragment string) string {
	if fragment == Window().URL().Fragment {
		return "focus"
	}
	return ""
}
