package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"api/converters"
)

func main() {

	app := Setup()

	log.Fatal(app.Listen(":6027"))

}

// func ConvertEndpoint(w http.ResponseWriter, r *http.Request) {
// 	// Endpoint işlemleri burada gerçekleştirilir
// }

func Setup() *fiber.App {

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/convert", func(c *fiber.Ctx) error {

		type Result struct {
			Output string `json:"message"`
		}

		type RequestBody struct {
			Value      string `json:"value"`
			SourceType string `json:"sourceType"`
			DestType   string `json:"destType"`
		}

		var reqBody RequestBody

		if err := c.BodyParser(&reqBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Bad Request",
			})
		}

		if reqBody.SourceType == "ascii" && reqBody.DestType == "binary" {
			asciiInput := reqBody.Value
			binaryResult := converters.AsciiToBinary(asciiInput)
			result := Result{Output: binaryResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "ascii" && reqBody.DestType == "octal" {
			asciiInput := reqBody.Value
			octalResult := converters.AsciiToOctal(asciiInput)

			result := Result{Output: octalResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "binary" && reqBody.DestType == "ascii" {
			binaryInput := reqBody.Value
			asciiResult := converters.BinaryToAscii(binaryInput)

			result := Result{Output: asciiResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "binary" && reqBody.DestType == "octal" {
			binaryInput := reqBody.Value
			octalResult := converters.BinaryToOctal(binaryInput)

			result := Result{Output: octalResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "octal" && reqBody.DestType == "ascii" {
			octalInput := reqBody.Value
			asciiResult := converters.OctalToAscii(octalInput)

			result := Result{Output: asciiResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		} else if reqBody.SourceType == "octal" && reqBody.DestType == "binary" {
			octalInput := reqBody.Value
			binaryResult := converters.OctalToBinary(octalInput)

			result := Result{Output: binaryResult}
			jsonResult, _ := json.Marshal(result)

			c.Type("application/json")

			c.Send(jsonResult)

		}

		return nil
	})
	return app
}
