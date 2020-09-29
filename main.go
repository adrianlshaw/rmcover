package main

import (
	"fmt"
	"image/color"
	"net/http"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

func getImage(url string) {

	if url == "" {
		fmt.Println("Can't fetch empty URL")
		return
	}

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Can't get image")
	}

	img, err := imaging.Decode(response.Body)

	if err != nil {
		fmt.Println("Failed to decode image")
	}

	width := 1404
	height := 1872

	img = resize.Resize(uint(width), 0, img, resize.Bilinear)

	background := imaging.New(
		width,
		height,
		color.RGBA{255, 255, 255, 255},
	)
	img = imaging.PasteCenter(background, img)

	imaging.Save(img, "/usr/share/remarkable/suspended.png")

}

func scrape(url string) string {

	doc, err := htmlquery.LoadURL(url)

	if err != nil {
		fmt.Println("Failed to load URL")
		return ""
	}

	list := htmlquery.Find(doc, "//div")

	for _, n := range list {

		if htmlquery.SelectAttr(n, "id") == "comic" {

			imgtag := htmlquery.FindOne(n, "//img")
			imgurl := htmlquery.SelectAttr(imgtag, "src")
			imgurl = "https:" + imgurl

			fmt.Println(imgurl)

			return imgurl
		}
	}

	return ""
}

func main() {

	url := "https://c.xkcd.com/random/comic/"

	for {

		img := scrape(url)

		if img != "" {
			getImage(img)
		}

		fmt.Println("Sleeping")

		time.Sleep(60 * time.Minute)
	}
}
