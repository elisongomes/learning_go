package main

import "fmt"

func weekDayName1(day int8) string {
	switch day {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturnday"
	default:
		return "Invalid number"
	}
}

func weekDayName2(day int8) string {
	var weekDayName string

	// It has no "break" instruction because it automatically exits on eval true condition
	switch {
	case day == 1:
		weekDayName = "Sunday"
		//fallthrough // Go to next case, without eval condition
	case day == 2:
		weekDayName = "Monday"
	case day == 3:
		weekDayName = "Tuesday"
	case day == 4:
		weekDayName = "Wednesday"
	case day == 5:
		weekDayName = "Thursday"
	case day == 6:
		weekDayName = "Friday"
	case day == 7:
		weekDayName = "Saturnday"
	default:
		weekDayName = "Invalid number"
	}

	return weekDayName
}

func main()  {
	fmt.Println("Control structures: Switch")

	day1 := weekDayName1(3)
	fmt.Println(day1)

	day2 := weekDayName2(6)
	fmt.Println(day2)
}
