package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

var config Config

func main() {
	cfg, err := ReadConfig("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config = cfg

	app := fiber.New()
	app.Get("/*", HandleLink)

	addr := fmt.Sprintf(":%d", config.Cmd.Port)
	log.Fatal(app.Listen(addr))
}

// HandleLink handles the reference to an entity and redirects to its dynamic links
func HandleLink(ctx *fiber.Ctx) error {
	fullURL, err := url.Parse(ctx.BaseURL() + ctx.OriginalURL())
	if err != nil {
		log.Printf("[ERR] %s - %v", ctx.IP(), err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal error")
	}

	for _, replacement := range config.Replacements {
		if !LinkMatches(fullURL, replacement) {
			continue
		}

		// Found
		redirectLink := RewriteLink(fullURL, replacement)

		if config.Cmd.Verbose {
			log.Printf("[%s] %s://%s%s -> %s", ctx.IP(), fullURL.Scheme, fullURL.Host, fullURL.RequestURI(), redirectLink)
		}
		ctx.Set("Location", redirectLink)
		return ctx.SendStatus(fiber.StatusFound) // 302
	}

	log.Printf("[ERR] [%s] The URL does not match any replacement (%s)", ctx.IP(), ctx.BaseURL())
	return ctx.Status(fiber.StatusNotFound).SendString("The domain is not recognized")
}
