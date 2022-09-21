package main
import ("fmt")
type Person struct {
name string
age int
job string
salary int
}

func PassStruct(pers1 Person , pers2 Person) {

	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)
}

func main() {

	pers1 := Person{
		name: "Hege",
		age:  45,
		job: "Teacher",
		salary: 6000,
	 }
	 pers2 := Person{
		name: "Cecilie",
		age:  24,
		job: "Marketing",
		salary: 4500,
	 }
	PassStruct(pers1,pers2)
}