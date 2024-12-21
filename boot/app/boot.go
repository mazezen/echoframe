package app

import (
	"flag"
	"github.com/mazezen/echoframe/internal/app/router"
	"github.com/mazezen/echoframe/vm"
	"github.com/mazezen/itools"
)

type App struct {
	c    *echo.Echo
	port string
}

var cf string

func init() {
	flag.StringVar(&cf, "cf", "config", "config file path")
}

func NewApp() *App {
	flag.Parse()
	ap := new(App)
	if ap.c == nil {
		ap.c = echo.New()
		itools.LoadConfig(cf)
	}
	val := itools.Gm.Get("port")
	if val.(string) == "" {
		ap.port = ":8090"
	}
	ap.port = val.(string)

	return ap
}

func (a *App) Start() {
	a.do()
	a.Route()
}

func (a *App) do() {
	vm.BootStore()
}

func (a *App) Route() {
	router.BootApp(a.c)
	a.c.Logger.Fatal(a.c.Start(a.port))
}
