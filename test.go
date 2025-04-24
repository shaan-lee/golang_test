package main

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

// import "fmt"

// func namedFunc(name string) { fmt.Println(name) }

// func main() {
// 	namedFunc("shaan")
// }

// =======================================================

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

func main() {
	app := fiber.new()

	app.Get()
}
