package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/scraper/", func(c *fiber.Ctx) {
		c.JSON(map[string]string{
			"res": "Passe a query que voce deseja apos a barra",
		})
		c.Status(200)
	})
	app.Get("/scraper/:query/:trans?/:size?", scraper)
	app.Listen(":3000")

}

func scraper(f *fiber.Ctx) {
	var res []string
	query := f.Params("query")
	trans := f.Params("trans")
	size := f.Params("size")
	if trans != "true" {
		trans = "false"
	}
	if size != "l" && size != "i" && size != "m" {
		size = "false"
	}
	var url string
	url = "https://www.google.com/search?q=" + query + "%20icon&hl=pt-BR&tbm=isch"
	switch {
	case trans == "true" && size != "false":
		if size == "l" {
			url = url + "&tbs=ic:trans%2Cisz:l"
		}
		if size == "m" {
			url = url + "&tbs=ic:trans%2Cisz:m"
		}
		if size == "i" {
			url = url + "&tbs=ic:trans%2Cisz:i"
		}
	case trans == "false" && size != "false":
		if size == "l" {
			url = url + "&tbs=isz:l"
		}
		if size == "m" {
			url = url + "&tbs=isz:m"
		}
		if size == "i" {
			url = url + "&tbs=isz:i"
		}
	case trans == "true" && size == "false":
		url = url + "&tbs=ic:trans"
	default:
		break
	}
	c := colly.NewCollector()

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		res = append(res, e.Attr("src"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visitando", r.URL)
	})

	c.Visit(url)
	fmt.Println("Params", trans, " ", size)
	f.JSON(map[string]interface{}{
		"res":          res[1:],
		"size":         size,
		"transparency": trans,
	})
	f.Status(200)
}
