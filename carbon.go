package carbon

import (
	"strings"
	"time"
)

type carbon struct {
	Time time.Time
	loc  *time.Location
}

// format转layout
func format2layout(format string) string {
	format = strings.Trim(format, " ")
	layout := strings.Replace(format, "Y", "2006", 1)
	layout = strings.Replace(layout, "y", "06", 1)
	layout = strings.Replace(layout, "m", "01", 1)
	layout = strings.Replace(layout, "n", "1", 1)
	layout = strings.Replace(layout, "d", "02", 1)
	layout = strings.Replace(layout, "j", "2", 1)
	layout = strings.Replace(layout, "H", "15", 1)
	layout = strings.Replace(layout, "h", "03", 1)
	layout = strings.Replace(layout, "g", "3", 1)
	layout = strings.Replace(layout, "i", "04", 1)
	layout = strings.Replace(layout, "s", "05", 1)
	return layout
}

// New 初始化,设置默认时区
func New() *carbon {
	c := &carbon{}
	loc, _ := time.LoadLocation(Local)
	c.loc = loc
	return c
}

// Timezone 设置时区
func (c *carbon) Timezone(name string) *carbon {
	loc, _ := time.LoadLocation(name)
	c.loc = loc
	return c
}

// CreateFromTimestamp 时间戳转时间对象
func (c *carbon) CreateFromTimestamp(timestamp int64) *carbon {
	c.Time = time.Unix(timestamp, 0)
	return c
}

// CreateFromDateTime 年月日时分秒转时间对象
func (c *carbon) CreateFromDateTime(year int, month time.Month, day int, hour int, minute int, second int) *carbon {
	c.Time = time.Date(year, month, day, hour, minute, second, 0, c.loc)
	return c
}

// CreateFromDate 年月日转时间对象
func (c *carbon) CreateFromDate(year int, month time.Month, day int) *carbon {
	c.Time = time.Date(year, month, day, 0, 0, 0, 0, c.loc)
	return c
}

// CreateFromTime 时分秒转时间对象
func (c *carbon) CreateFromTime(hour int, minute int, second int) *carbon {
	c.Time = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minute, second, 0, c.loc)
	return c
}

// Now 当前
func (c *carbon) Now() *carbon {
	c.Time = time.Now()
	return c
}

// AddYears N年后
func (c *carbon) AddYears(years int) *carbon {
	c.Time = c.Time.AddDate(years, 0, 0)
	return c
}

// AddYear 1年后
func (c *carbon) AddYear() *carbon {
	c.Time = c.Time.AddDate(1, 0, 0)
	return c
}

// SubYears N年前
func (c *carbon) SubYears(years int) *carbon {
	c.Time = c.Time.AddDate(-years, 0, 0)
	return c
}

// SubYear 1年前
func (c *carbon) SubYear() *carbon {
	c.Time = c.Time.AddDate(-1, 0, 0)
	return c
}

// AddMonths N月后
func (c *carbon) AddMonths(months int) *carbon {
	c.Time = c.Time.AddDate(0, months, 0)
	return c
}

// AddMonth 1月后
func (c *carbon) AddMonth() *carbon {
	c.Time = c.Time.AddDate(0, 1, 0)
	return c
}

// SubMonths N月前
func (c *carbon) SubMonths(months int) *carbon {
	c.Time = c.Time.AddDate(0, -months, 0)
	return c
}

// SubMonth 减少1月
func (c *carbon) SubMonth() *carbon {
	c.Time = c.Time.AddDate(0, -1, 0)
	return c
}

// AddDays N天后
func (c *carbon) AddDays(days int) *carbon {
	c.Time = c.Time.AddDate(0, 0, days)
	return c
}

// AddDay 1天后
func (c *carbon) AddDay() *carbon {
	c.Time = c.Time.AddDate(0, 0, 1)
	return c
}

// SubDays N天前
func (c *carbon) SubDays(days int) *carbon {
	c.Time = c.Time.AddDate(0, 0, -days)
	return c
}

// SubDay 1天前
func (c *carbon) SubDay() *carbon {
	c.Time = c.Time.AddDate(0, 0, -1)
	return c
}

// AddHours N小时后
func (c *carbon) AddHours(hours int) *carbon {
	c.Time = c.Time.AddDate(0, 0, hours/24)
	return c
}

// AddHour 1小时后
func (c *carbon) AddHour() *carbon {
	c.Time = c.Time.AddDate(0, 0, 1/24)
	return c
}

// AddHours N小时前
func (c *carbon) SubHours(hours int) *carbon {
	c.Time = c.Time.AddDate(0, 0, -hours/24)
	return c
}

// AddHour 1小时前
func (c *carbon) SubHour() *carbon {
	c.Time = c.Time.AddDate(0, 0, -1/24)
	return c
}

// AddMinutes N分钟后
func (c *carbon) AddMinutes(minutes int) *carbon {
	c.Time = c.Time.AddDate(0, 0, minutes/1440)
	return c
}

// AddMinute 1分钟后
func (c *carbon) AddMinute() *carbon {
	c.Time = c.Time.AddDate(0, 0, 1/1440)
	return c
}

// SubMinutes N分钟前
func (c *carbon) SubMinutes(minutes int) *carbon {
	c.Time = c.Time.AddDate(0, 0, -minutes/1440)
	return c
}

// SubMinute 1分钟前
func (c *carbon) SubMinute() *carbon {
	c.Time = c.Time.AddDate(0, 0, -1/1440)
	return c
}

// AddSeconds N秒钟后
func (c *carbon) AddSeconds(second int) *carbon {
	c.Time = c.Time.AddDate(0, 0, second/86400)
	return c
}

// AddMinute 1秒钟后
func (c *carbon) AddSecond() *carbon {
	c.Time = c.Time.AddDate(0, 0, 1/86400)
	return c
}

// SubMinutes N秒钟前
func (c *carbon) SubSeconds(second int) *carbon {
	c.Time = c.Time.AddDate(0, 0, -second/86400)
	return c
}

// SubMinute 1秒钟前
func (c *carbon) SubSecond() *carbon {
	c.Time = c.Time.AddDate(0, 0, -1/86400)
	return c
}

// Parse 解析标准时间格式
func (c *carbon) Parse(value string) *carbon {
	value = strings.Trim(value, " ")
	layout := "2006-01-02 15:04:05"

	if len(value) == 10 && strings.Contains(value, "-") {
		layout = "2006-01-02"
	}

	if len(value) == 10 && strings.Contains(value, "/") {
		layout = "2006/01/02"
	}

	if len(value) == 19 && strings.Contains(value, "/") {
		layout = "2006/01/02 15:04:05"
	}

	if len(value) == 14 {
		layout = "20060102150405"
	}

	if len(value) == 8 {
		layout = "20060102"
	}

	t, _ := time.ParseInLocation(layout, value, c.loc)
	c.Time = t
	return c
}

// ParseByCustom 解析自定义时间格式
func (c *carbon) ParseByCustom(value string, format string) *carbon {
	value = strings.Trim(value, " ")
	layout := format2layout(format)
	t, _ := time.ParseInLocation(layout, value, c.loc)
	c.Time = t
	return c
}

// Format 格式化时间
func (c *carbon) Format(format string) string {
	return c.Time.In(c.loc).Format(format2layout(format))
}

// Today 今天
func (c *carbon) Today() string {
	return time.Now().In(c.loc).Format("2006-01-02 00:00:00")
}

// Tomorrow 明天
func (c *carbon) Tomorrow() string {
	return time.Now().AddDate(0, 0, 1).In(c.loc).Format("2006-01-02 00:00:00")
}

// Yesterday 昨天
func (c *carbon) Yesterday() string {
	return time.Now().AddDate(0, 0, -1).In(c.loc).Format("2006-01-02 00:00:00")
}

// FirstDayInYear 年初
func (c *carbon) FirstDayInYear() string {
	return c.CreateFromDate(c.Time.Year(), 01, 01).ToDateTimeString()
}

// LastDayInYear 年末
func (c *carbon) LastDayInYear() string {
	return c.CreateFromDate(c.Time.Year(), 12, 31).ToDateTimeString()
}

// FirstDayInMonth 月初
func (c *carbon) FirstDayInMonth() string {
	return c.CreateFromDate(c.Time.Year(), c.Time.Month(), 1).ToDateTimeString()
}

// LastDayInMonth 月末
func (c *carbon) LastDayInMonth() string {
	return c.CreateFromDate(c.Time.Year(), c.Time.Month(), 1).Time.AddDate(0, 1, -1).Format("2006-01-02 00:00:00")
}

// ToDateTimeString 转日期时间字符串
func (c *carbon) ToDateTimeString() string {
	return c.Time.In(c.loc).Format("2006-01-02 15:04:05")
}

// ToDateString 转日期字符串
func (c *carbon) ToDateString() string {
	return c.Time.In(c.loc).Format("2006-01-02")
}

// ToTimeString 转时间字符串
func (c *carbon) ToTimeString() string {
	return c.Time.In(c.loc).Format("15:04:05")
}

// ToTimestamp 转时间戳
func (c *carbon) ToTimestamp() int64 {
	return c.Time.Unix()
}

// IsLeapYear 是否是闰年
func (c *carbon) IsLeapYear() bool {
	year := c.Time.Year()
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}

// IsJanuary 是否是一月
func (c *carbon) IsJanuary() bool {
	return c.Time.In(c.loc).Month().String() == "January"
}

// IsMonday 是否是二月
func (c *carbon) IsFebruary() bool {
	return c.Time.In(c.loc).Month().String() == "February"
}

// IsMarch 是否是三月
func (c *carbon) IsMarch() bool {
	return c.Time.In(c.loc).Month().String() == "March"
}

// IsApril 是否是四月
func (c *carbon) IsApril() bool {
	return c.Time.In(c.loc).Month().String() == "April"
}

// IsMay 是否是五月
func (c *carbon) IsMay() bool {
	return c.Time.In(c.loc).Month().String() == "May"
}

// IsJune 是否是六月
func (c *carbon) IsJune() bool {
	return c.Time.In(c.loc).Month().String() == "June"
}

// IsJuly 是否是七月
func (c *carbon) IsJuly() bool {
	return c.Time.In(c.loc).Month().String() == "July"
}

// IsAugust 是否是八月
func (c *carbon) IsAugust() bool {
	return c.Time.In(c.loc).Month().String() == "August"
}

// IsSeptember 是否是九月
func (c *carbon) IsSeptember() bool {
	return c.Time.In(c.loc).Month().String() == "September"
}

// IsOctober 是否是十月
func (c *carbon) IsOctober() bool {
	return c.Time.In(c.loc).Month().String() == "October"
}

// IsNovember 是否是十一月
func (c *carbon) IsNovember() bool {
	return c.Time.In(c.loc).Month().String() == "November"
}

// IsDecember 是否是十二月
func (c *carbon) IsDecember() bool {
	return c.Time.In(c.loc).Month().String() == "December"
}

// IsMonday 是否是周一
func (c *carbon) IsMonday() bool {
	return c.Time.In(c.loc).Weekday().String() == "Monday"
}

// IsTuesday 是否是周二
func (c *carbon) IsTuesday() bool {
	return c.Time.In(c.loc).Weekday().String() == "Tuesday"
}

// IsWednesday 是否是周三
func (c *carbon) IsWednesday() bool {
	return c.Time.In(c.loc).Weekday().String() == "Wednesday"
}

// IsThursday 是否是周四
func (c *carbon) IsThursday() bool {
	return c.Time.In(c.loc).Weekday().String() == "Thursday"
}

// IsFriday 是否是周五
func (c *carbon) IsFriday() bool {
	return c.Time.In(c.loc).Weekday().String() == "Friday"
}

// IsSaturday 是否是周六
func (c *carbon) IsSaturday() bool {
	return c.Time.In(c.loc).Weekday().String() == "Saturday"
}

// IsSunday 是否是周日
func (c *carbon) IsSunday() bool {
	return c.Time.In(c.loc).Weekday().String() == "Sunday"
}

// IsFirstDayInYear 是否年初
func (c *carbon) IsFirstDayInYear() bool {
	_, month, day := c.Time.Date()
	if month.String() == "January" && day == 1 {
		return true
	}

	return false
}

// IsLastDayInYear 是否是年末
func (c *carbon) IsLastDayInYear() bool {
	_, month, day := c.Time.Date()
	if month.String() == "December" && day == 31 {
		return true
	}

	return false
}

// IsFirstDayInMonth 是否月初
func (c *carbon) IsFirstDayInMonth() bool {
	return c.Time.In(c.loc).Day() == 1
}

// IsLastDayInMonth 是否是月末
func (c *carbon) IsLastDayInMonth() bool {
	return c.Time.In(c.loc).Format("2006-01-02 00:00:00") == c.LastDayInMonth()
}
