package main

import (
	"fmt"
)

type Student struct {
	name  string
	score int
}

type StudentList []Student

func (sl StudentList) AverageScore() float64 {
	totalScore := 0
	for _, student := range sl {
		totalScore += student.score
	}
	return float64(totalScore) / float64(len(sl))
}

func (sl StudentList) MinScoreStudent() Student {
	if len(sl) == 0 {
		return Student{}
	}

	minStudent := sl[0]
	for _, student := range sl {
		if student.score < minStudent.score {
			minStudent = student
		}
	}
	return minStudent
}

func (sl StudentList) MaxScoreStudent() Student {
	if len(sl) == 0 {
		return Student{}
	}

	maxStudent := sl[0]
	for _, student := range sl {
		if student.score > maxStudent.score {
			maxStudent = student
		}
	}
	return maxStudent
}

func main() {
	students := StudentList{
		{"Yopan", 60},
		{"Rizky", 80},
		{"Anita", 75},
		{"Budi", 70},
		{"Citra", 65},
	}

	average := students.AverageScore()
	minStudent := students.MinScoreStudent()
	maxStudent := students.MaxScoreStudent()

	fmt.Printf("Average Score: %.0f\n", average)
	fmt.Printf("Min Score of Students: %s (%d)\n", minStudent.name, minStudent.score)
	fmt.Printf("Max Score of Students: %s (%d)\n", maxStudent.name, maxStudent.score)
}
