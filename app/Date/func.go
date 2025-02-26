package Date

var days = [...]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// IsLeap true 是闰年，366天
func (date Date) IsLeap() bool {
	return date.Year%400 == 0 || date.Year%100 != 0 && date.Year%4 == 0
}

// DayOfMonth 获取某年某月的天数
func (date Date) DayOfMonth() int {
	if date.Month == 2 && date.IsLeap() { // 润年2月为 28 + 1天
		return days[date.Month] + 1
	}
	return days[date.Month]
}

// WeekOfTheFirstDay 获取某年某月的第一天是星期几
// 0-6 对应周日-周六
/*
Teller(蔡勒)公式 快速算日期对应星期  从1582开始
expression: w=y+[y/4]+[c/4]-2c+[26(m+1)/10]+d-1
explain: w星期；c世纪-1；y年（后两位数）m月（m大于等于3，小于等于14，即在蔡勒公式中，某年的1、2月要看作上一年的13、14月来计算
eg: 比如2003年1月1日要看作2002年的13月1日来计算
ans: 0星期天 5星期五
*/
func (date Date) WeekOfTheFirstDay() int {
	c := 0
	y := date.Year
	m := date.Month
	d := date.Day
	if m <= 2 {
		y -= 1
		m += 12
	}
	c = y / 100
	y = y % 100
	return (y + y/4 + c/4 - 2*c + 13*(m+1)/5 + d - 1) % 7
}
