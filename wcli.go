package main

import (
	"fmt"
	"github.com/qeesung/image2ascii/convert"
	"github.com/imroc/req/v3"
	"github.com/probandula/figlet4go"
	"time"
	_ "image/jpeg"
	_ "image/png"
)

const welcomeText = `Welcome to terminal-based weather CLI! Just type the name of your location below and the program will fetch weather data from OpenWeatherMaps API. The data will be accurate and you will need internet connectivity to use this app.`
const sunshineImg = "sunshine.jpg"
const hazeImg = "haze.jpg"
const rainImg = "rain.jpg"

func main() {
	textRenderer("WINDY", "b")
	fmt.Println("\n\n" + welcomeText)
	time.Sleep(1000000)
}

func textRenderer(text string, color string) {
	ascii := figlet4go.NewAsciiRender()
	ASCIIOptions := figlet4go.NewRenderOptions()
	if color == "b" {
		ASCIIOptions.FontColor = []figlet4go.Color{
			figlet4go.ColorBlue,
		}
	} else if color == "g" {
		ASCIIOptions.FontColor = []figlet4go.Color{
			figlet4go.ColorGreen,
		}
	} else {
		ASCIIOptions.FontColor = []figlet4go.Color{
			figlet4go.ColorWhite,
		}
	} 
	renderStr, _ := ascii.RenderOpts(text, ASCIIOptions)
	fmt.Print(renderStr)
}

func imgRenderer(imgRelativePath string) {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 100
	convertOptions.Colored = true
	converter := convert.NewImageConverter()
	fmt.Print(converter.ImageFile2ASCIIString(imgRelativePath, &convertOptions))
}

func dataFetcher() {
	req.DevMode()

}