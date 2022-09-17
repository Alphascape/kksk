package main

import (
	"context"
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/cli"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
	"github.com/maxence-charriere/go-app/v9/pkg/logs"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
	"net/http"
	"os"
	"syscall"
)

const (
	defaultTitle       = "KKSK"
	defaultDescription = "kksk.app 是一个渐进式网络应用程序(PWA)，你可以在这找到一些乐子(大概)"
	backgroundColor    = "#2e343a"

	githubURL  = "https://github.com/Alphascape/kksk"
	twitterURL = "https://twitter.com/nangcr"
)

type localOptions struct {
	Port int `env:"PORT" help:"The port used to listen connections."`
}

type githubOptions struct {
	Output string `cli:"o" env:"-" help:"The directory where static resources are saved."`
}

func main() {
	ui.BaseHPadding = 42
	ui.BlockPadding = 18

	app.Route("/", newHomePage())
	app.Route("/dinner", newDinnerPage())
	app.Handle(installApp, handleAppInstall)
	app.Handle(updateApp, handleAppUpdate)

	app.RunWhenOnBrowser()

	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	defer exit()

	localOpts := localOptions{Port: 8080}
	cli.Register("local").
		Help(`Launches a server that serves the documentation app in a local environment.`).
		Options(&localOpts)

	githubOpts := githubOptions{}
	cli.Register("github").
		Help(`Generates the required resources to run kksk app on GitHub Pages.`).
		Options(&githubOpts)

	h := app.Handler{
		Name:        "KKSK App",
		Title:       defaultTitle,
		Description: defaultDescription,
		Author:      "Alphascape",

		Icon: app.Icon{
			Default: "/web/logo.png",
		},
		Keywords: []string{
			"kksk",
			"app",
			"pwa",
		},
		BackgroundColor: backgroundColor,
		ThemeColor:      backgroundColor,
		LoadingLabel:    "仙德祈祷中...",
		Lang:            "zh-cmn-Hans",

		Styles: []string{
			"https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
			"/web/css/docs.css",
		},
	}

	cli.Load()

	switch cli.Load() {
	case "local":
		runLocal(ctx, &h, localOpts)

	case "github":
		generateGitHubPages(ctx, &h, githubOpts)
	}

}

func runLocal(ctx context.Context, h http.Handler, opts localOptions) {
	app.Logf("%s", logs.New("starting kksk app server").
		Tag("port", opts.Port),
	)

	s := http.Server{
		Addr:    fmt.Sprintf(":%v", opts.Port),
		Handler: h,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func generateGitHubPages(ctx context.Context, h *app.Handler, opts githubOptions) {
	if err := app.GenerateStaticWebsite(opts.Output, h); err != nil {
		panic(err)
	}
}

func exit() {
	err := recover()
	if err != nil {
		app.Logf("command failed: %s", errors.Newf("%v", err))
		os.Exit(-1)
	}
}
