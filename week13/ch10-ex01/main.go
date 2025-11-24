package main

import (
	"fmt"
	"log"
	"week13/pkg/calendar"
)

func main() {
	today := calendar.Event{} // Date가 임베딩되어 있음
	today.SetTitle("Final Exam D-14")
	// today.year = 2025

	err := today.SetYear(2025)
	if err != nil {
		log.Fatal(err)
	}
	err = today.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}
	err = today.SetDay(25)
	if err != nil {
		log.Fatal(err)
	}

	//today.SetYear(2025)
	//today.SetMonth(11)
	//today.SetDay(24)
	// fmt.Println(today.Year(), "년 ", today.Month(), "월 ", today.Day(), "일 ")
	fmt.Printf("%d년 %d월 %d일\n", today.Title(), today.Year(), today.Month(), today.Day())
}
