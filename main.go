package main

import (
	"embed"
	"paas-spider/biz"
	"paas-spider/toolkit"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	/* if !toolkit.Signature("10-F6-0A-AD-A3-81,04-42-1A-D7-FE-8B", "2025-03-01") {
		return
	} */
	httpClient := toolkit.NewHttpClient("http://localhost:8080")
	taskBiz := biz.NewTaskBiz()
	recordBiz := biz.NewRecordBiz()
	planBiz := biz.NewPlanBiz()
	app := NewApp()

	err := wails.Run(&options.App{
		Title:  "Web自动化工具",
		Width:  1200,
		Height: 900,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			httpClient,
			taskBiz,
			recordBiz,
			planBiz,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              false,
			WindowIsTranslucent:               false,
			BackdropType:                      windows.Mica,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.Light,
			/* CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:   windows.RGB(20, 20, 20),
				DarkModeTitleText:  windows.RGB(200, 200, 200),
				DarkModeBorder:     windows.RGB(20, 0, 20),
				LightModeTitleBar:  windows.RGB(200, 200, 200),
				LightModeTitleText: windows.RGB(20, 20, 20),
				LightModeBorder:    windows.RGB(200, 200, 200),
			}, */
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
