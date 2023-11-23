package main

import (
	"embed"
	"github.com/c3b2a7/decrypotr/backend/constant"
	"github.com/c3b2a7/decrypotr/backend/service"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	cryptoService := service.GetCryptoService()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "EasyCrypto",
		Width:     constant.DEFAULT_WINDOW_WIDTH,
		Height:    constant.DEFAULT_WINDOW_HEIGHT,
		MinWidth:  constant.MIN_WINDOW_WIDTH,
		MinHeight: constant.MIN_WINDOW_HEIGHT,
		MaxWidth:  constant.MAX_WINDOW_WIDTH,
		MaxHeight: constant.MAX_WINDOW_HEIGHT,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Bind: []interface{}{
			cryptoService,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
