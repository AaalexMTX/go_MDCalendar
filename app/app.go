package app

import (
	"MDCalendar/app/calender"
	"fmt"
)

func Start() {
	fmt.Println(title)
	var year, month, day, model int
	// 样式
	fmt.Print("select a table style \n(1: start with Mon, 2: start with Sun) -- > ")
	if _, err := fmt.Scanf("%d\n", &model); err != nil {
		fmt.Println("Error reading table style:", err)
		return
	}
	// 日期
	day = 1
	for {
		fmt.Print("input year month (eg: 2025 2) -- > ")
		if _, err := fmt.Scanf("%d %d\n", &year, &month); err != nil {
			fmt.Println("Error reading year and month:", err)
			return
		}
		cal := calender.NewCalender(day, month, year)
		if model == 1 {
			cal.PrintCalendarMon()
		} else {
			cal.PrintCalendarSun()
		}
	}
}

var title string = `
                                                                                                     
    /|    //| |     //    ) )        //   ) )                                                        
   //|   // | |    //    / /        //         ___     //  ___       __      ___   /  ___      __    
  // |  //  | |   //    / /  ____  //        //   ) ) // //___) ) //   ) ) //   ) / //   ) ) //  ) ) 
 //  | //   | |  //    / /        //        //   / / // //       //   / / //   / / //   / / //       
//   |//    | | //____/ /        ((____/ / ((___( ( // ((____   //   / / ((___/ / ((___( ( //        `
