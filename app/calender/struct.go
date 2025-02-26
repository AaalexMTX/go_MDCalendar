package calender

import (
	"MDCalendar/app/Date"
)

type Calender struct {
	Data Date.Date
}

func NewCalender(d, m, y int) Calender {
	return Calender{
		Data: Date.Date{Day: d, Month: m, Year: y},
	}
}
