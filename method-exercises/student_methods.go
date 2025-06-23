package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	Grades    []int
}

func (s Student) CalculateAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s Student) GetStatus() string {
	average := s.getAverageGrade()
	switch {
	case average >= 4.5:
		return "отличник"
	case average >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}

func (s *Student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}

func (s Student) getAverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

func main() {
	student := Student{
		Name:      "Alexander Ivanov",
		BirthYear: 2000,
		Grades:    []int{5, 4, 5, 3},
	}

	fmt.Printf("Student: %s\n", student.Name)
	fmt.Printf("Age: %d лет\n", student.CalculateAge())
	fmt.Printf("Status: %s\n", student.GetStatus())

	student.AddGrade(5)
	fmt.Printf("New average grade: %.2f\n", student.getAverageGrade())
	fmt.Printf("Updated status: %s\n", student.GetStatus())
}
