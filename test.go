package main

// =======================================
// do it first

// go mod init <your git repository url>

// =======================================
// fmt, pointer example

// import "fmt"

// func main() {
// 	fmt.Println("hello world")

// 	str := "hi here is GO"
// 	fmt.Println(str)

// 	str2 := str
// 	str3 := &str

// 	str2 = "changed string str2"
// 	*str3 = "changed string str3"

// 	fmt.Println(str)
// 	fmt.Println(str2)
// 	fmt.Println(str3)
// }

// =====================================================
// named func example

// import "fmt"

// func namedFunc(name string) { fmt.Println(name) }

// func main() {
// 	namedFunc("shaan")
// }

// =======================================================
// struct and interface example

// import "fmt"

// type Job struct {
// 	name     string
// 	position string
// 	salary   float32
// }

// func (job Job) getName() string {
// 	fmt.Println(job.name)
// 	return job.name
// }

// func (job Job) increaseSalary() {
// 	salary := &job.salary
// 	fmt.Println("current salary is ", *salary)

// 	*salary += 100
// 	fmt.Println("increased salary is ", *salary)
// }

// func main() {
// 	var job Job
// 	job = Job{"shaan", "programmer", 500}
// 	job.getName()
// 	job.increaseSalary()
// }

// =======================================================
// fiber example

// import (
// 	"log"

// 	"github.com/gofiber/fiber/v3"
// 	"github.com/gofiber/fiber/v3/middleware/logger"
// )

// func main() {
// 	app := fiber.New()

// 	app.Use(logger.New(logger.ConfigDefault))

// 	app.Get("/", func(c fiber.Ctx) error {
// 		return c.SendString("Hello, GO fiber World!")
// 	})

// 	log.Fatal(app.Listen(":8888"))

// }

// =======================================================
