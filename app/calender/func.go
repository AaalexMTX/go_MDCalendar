package calender

import (
	"fmt"
	"math"
)

// PrintCalendarMon 打印月历（以周一开始）
func (calender Calender) PrintCalendarMon() {
	allDays := calender.Data.DayOfMonth()                // 该月份总日期
	week := calender.Data.WeekOfTheFirstDay()            // 起始星期
	rowNum := math.Ceil(float64((week+6)%7+allDays) / 7) // 计算行数28天可能就4行
	// print table header
	fmt.Println("总日期 起始星期 行数")
	fmt.Printf("   %2d     %d     %d\n\n", allDays, week, int(rowNum))
	for i := 0; i < 7; i++ {
		switch i {
		case 2:
			fmt.Printf("|%4d", calender.Data.Year)
		case 3:
			fmt.Print("| -- ")
		case 4:
			fmt.Printf("| %2d ", calender.Data.Month)
		case 6:
			fmt.Println("|    |")
		default:
			fmt.Print("|    ")
		}
	}
	for i := 0; i < 7; i++ {
		if i != 6 {
			fmt.Print("|:--:")
		} else {
			fmt.Print("|:--:|\n")
		}
	}
	fmt.Println("| 一 | 二 | 三 | 四 | 五 | 六 | 日 |")
	// print table
	dayPointer := 1
	for i := 0; i < int(rowNum); i++ {
		fmt.Print("|")
		for j := 0; j < 7; j++ {
			if i == 0 && j >= (week+6)%7 || i > 0 && dayPointer <= allDays { // (week+6)%7 即 0-6对应周一到周日
				fmt.Printf("%2d<br><br>|", dayPointer)
				dayPointer++
			} else {
				fmt.Printf("    |")
			}
		}
		fmt.Println()
	}
	// print ending
	for i := 0; i < 7; i++ {
		fmt.Print("|")
		for j := 0; j < 15; j++ {
			fmt.Print("-")
		}
		if i == 6 {
			fmt.Print("|\n")
		}
	}
}

// PrintCalendarSun 打印月历（以周日开始）
func (calender Calender) PrintCalendarSun() {
	allDays := calender.Data.DayOfMonth()          // 该月份总日期
	week := calender.Data.WeekOfTheFirstDay()      // 起始星期
	rowNum := math.Ceil(float64(week+allDays) / 7) // 空格数+总日期数来计算行数,因为28天可能就4行
	// print table header
	fmt.Println("总日期 起始星期 行数")
	fmt.Printf("   %2d     %d     %d\n\n", allDays, week, int(rowNum))
	for i := 0; i < 7; i++ {
		switch i {
		case 2:
			fmt.Printf("|%4d", calender.Data.Year)
		case 3:
			fmt.Print("| -- ")
		case 4:
			fmt.Printf("| %2d ", calender.Data.Month)
		case 6:
			fmt.Println("|    |")
		default:
			fmt.Print("|    ")
		}
	}
	for i := 0; i < 7; i++ {
		if i != 6 {
			fmt.Print("|:--:")
		} else {
			fmt.Print("|:--:|\n")
		}
	}
	fmt.Println("| 日 | 一 | 二 | 三 | 四 | 五 | 六 |")

	// print table
	dayPointer := 1
	for i := 0; i < int(rowNum); i++ {
		fmt.Print("|")
		for j := 0; j < 7; j++ {
			if i == 0 && j >= week || i > 0 && dayPointer <= allDays { // week 即 0-6对应周日到周六
				fmt.Printf("%2d<br><br>|", dayPointer)
				dayPointer++
			} else {
				fmt.Printf("    |")
			}
		}
		fmt.Println()
	}
	// print ending
	for i := 0; i < 7; i++ {
		fmt.Print("|")
		for j := 0; j < 15; j++ {
			fmt.Print("-")
		}
		if i == 6 {
			fmt.Println("|")
		}
	}
}
