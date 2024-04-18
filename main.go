// package mad

// import (
// 	"runtime"

// 	"github.com/gofiber/fiber/v3"
// )

// func mad() {

// 	numCPU := runtime.NumCPU()

// 	// Set the maximum number of operating system threads (P) to use
// 	runtime.GOMAXPROCS(numCPU)

// 	app := fiber.New()

// 	app.Get("/", func(c fiber.Ctx) error {
// 		return c.SendString("Hello, World!")
// 	})

// 	app.Listen(":3000")
// }
