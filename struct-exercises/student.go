package main

import (
	"fmt"
)

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func NewStudent(name string, age, course int, avgGrade float64) Student {
	return Student{
		Name:     name,
		Age:      age,
		Course:   course,
		AvgGrade: avgGrade,
	}
}

func (s *Student) UpgradeCourse() {
	s.Course++
}

func (s Student) PrintInfo() {
	fmt.Printf("Student: %s\n", s.Name)
	fmt.Printf("Age: %d\n", s.Age)
	fmt.Printf("Course: %d\n", s.Course)
	fmt.Printf("Average score: %.2f\n\n", s.AvgGrade)
}

func main() {
	student1 := NewStudent("Ivan Ivanov", 20, 2, 4.5)
	student2 := Student{"Petr Petrov", 21, 3, 4.2}
	var student3 Student
	student3.Name = "Fedor Fedorovich"
	student3.Age = 19
	student3.Course = 1
	student3.AvgGrade = 4.8

	student1.PrintInfo()
	student2.PrintInfo()
	student3.PrintInfo()

	student1.UpgradeCourse()
	fmt.Println("After the course increase:")
	student1.PrintInfo()

	students := []Student{student1, student2, student3}
	fmt.Println("All students:")
	for _, s := range students {
		s.PrintInfo()
	}
}
