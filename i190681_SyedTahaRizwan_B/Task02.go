package main

import "fmt"

type company struct {
        companyName string
        employees   []employee
}

type employee struct {
        name     string
        salary   int
        position string
}

func main() {
        Ali := employee{"Ali => ", 10000, " Full-Stack Developer\n"}
		Taha:= employee{"Taha => ", 7000, " Back-end Developer\n"}
		Moiz:= employee{"Moiz => ",90500," WebScrapper\n"}

        employees := []employee{Ali, Taha, Moiz}
        company := company{"\nTesla\nEmployees Details : \n", employees}

        fmt.Printf("Company Name :  %v\n", company)
}