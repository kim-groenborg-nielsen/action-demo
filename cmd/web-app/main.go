package main

import (
	"bytes"
	_ "embed"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kim-groenborg-nielsen/goenv"
	"html/template"
	"log"
)

//go:embed index.html
var indexHTML []byte

var version = "-"
var commit = "-"
var date = "-"

func main() {
	_ = godotenv.Load()
	serverUrl := goenv.GetEnvStr("SERVER_URL", "0.0.0.0:3000")

	app := fiber.New()

	tmpl, err := template.New("index").Parse(string(indexHTML))
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		data := struct {
			Version string
			Commit  string
			Date    string
		}{
			Version: version,
			Commit:  commit,
			Date:    date,
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return c.Status(500).SendString("Error executing template")
		}

		return c.Type("html").SendString(buf.String())
	})

	if err := app.Listen(serverUrl); err != nil {
		log.Fatal(err)
	}
}
