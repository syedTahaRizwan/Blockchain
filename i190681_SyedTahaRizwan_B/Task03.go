package main

import (
	"fmt"	
	"strings"
)
type Student struct {
	
	rollnumber int
	name 	   string
	address    string

}
func NewStudent(rollno int, name string, address string) *Student {
	s:= new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}
type StudentList struct {
	list []*Student
}
func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st:= NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}
func Print(ls *StudentList) {
	for i:= 0; i < 2; i++ {
		fmt.Printf("\n\n%s List %d %s\n\n", strings.Repeat("=" , 25 ) , i , strings.Repeat("=" , 25 ))
		fmt.Printf("Student roll no	    %d\n",ls.list[i].rollnumber)
		fmt.Printf("Student name	    %s\n",ls.list[i].name)
		fmt.Printf("Student address	    %s\n",ls.list[i].address)
	}
}

func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAAAA")
	student.CreateStudent(25, "Naveed", "BBBBBB")
	Print(student)

}