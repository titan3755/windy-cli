package main

import (
	"bufio"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"
	"github.com/imroc/req/v3"
	"github.com/probandula/figlet4go"
	"github.com/qeesung/image2ascii/convert"
	"github.com/tidwall/gjson"
	"github.com/fatih/color"
)

const welcomeText = `Welcome to terminal-based weather CLI Windy! Just type the name of your location below and the program will fetch weather data from OpenWeatherMaps API. The data will be accurate and you will need internet connectivity to use this app.`
const clearImg = "assets/clear.png"
const hazeImg = "assets/haze.jpg"
const rainImg = "assets/rain.png"
const drizzleImg = "assets/drizzle.jpg"
const cloudImg = "assets/clouds.png"
const thunderImg = "assets/thunder.png"
const snowImg = "assets/snow.png"
const defaultImg = "assets/default.png"
var client = req.C()
var reader = bufio.NewReader(os.Stdin)

type Response struct {
	Title string
	Id string
}

func main() {
	textRenderer("WINDY", "b")
	fmt.Println("\n\n" + welcomeText)
	running := true
	for running {
		fmt.Print("\nLocation -> ")
		query, _ := reader.ReadString('\n')
		query = strings.Replace(query, "\n", "", -1)
		response := dataFetcher("https://api.openweathermap.org/data/2.5/weather?appid=0a828b1755ea9eccaef2aa6c52c745b5&units=metric&q=" + query)
		var printText string
		if gjson.Get(response, "cod").String() == "200" {
			weather := gjson.Get(response, "weather.0.main")
			temp := gjson.Get(response, "main.temp")
			pressure := gjson.Get(response, "main.pressure")
			humidity := gjson.Get(response, "main.humidity")
			wind := gjson.Get(response, "wind.speed")
			visibility := gjson.Get(response, "visibility")			
			location := gjson.Get(response, "name")
			switch weather.String() {
			case "Haze":
				imgRenderer(hazeImg)
			case "Rain":
				imgRenderer(rainImg)
			case "Clouds":
				imgRenderer(cloudImg)
			case "Thunderstorm":
				imgRenderer(thunderImg)
			case "Clear":
				imgRenderer(clearImg)
			case "Drizzle":
				imgRenderer(drizzleImg)
			case "Snow":
				imgRenderer(snowImg)
			default:
				imgRenderer(defaultImg)
			}
			printText = fmt.Sprintf(`
	☁️  Weather: %v
	🌡️  Temperature: %v ℃
	🌫️  Pressure: %v hPa
	💧  Humidity: %v %%
	💨  Wind: %v m/s
	🔭  Visibility: %v m
	🌐  Location: %v
			`, weather, temp, pressure, humidity, wind, visibility, location)
		} else if gjson.Get(response, "cod").String() == "404" {
			printText = `
	Not a valid region!
			`
		} else {
						printText = `
	Something went wrong!
			`
		}
		fmt.Print(printText)
		endPrompt(&running)
	}
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
	convertOptions.FixedWidth = 50
	convertOptions.FixedHeight = 30
	convertOptions.Colored = true
	converter := convert.NewImageConverter()
	fmt.Print(converter.ImageFile2ASCIIString(imgRelativePath, &convertOptions))
}

func dataFetcher(url string) string {
	req.DevMode()
	fixedUrl := strings.TrimSpace(url)
	res, err := client.R().
		Get(fixedUrl)
	if err != nil {
		log.Fatal("Something went wrong! --> \n")
		log.Fatal(err)
	}
	return res.String()
}

func endPrompt(running *bool) {
	fmt.Print("\n\nGet weather of different location? (y/n) ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	if strings.TrimSpace(text) == "y" {
		color.Green("\nContinuing ...")
	} else {
		*running = false
		color.Red("\nExiting program ...")
	}
}