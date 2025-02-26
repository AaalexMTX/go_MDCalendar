package Date

type Date struct {
	Day   int
	Month int
	Year  int
}

func NewDate(d, m, y int) Date {
	return Date{
		Day:   d,
		Month: m,
		Year:  y,
	}
}
