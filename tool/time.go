package tool

import "time"

// 获取本周周一
func GetFirstDateOfWeek() (weekMonday string) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate.Format("2006-01-02")
	return
}

// 获取比较时间相对于年初的偏移日数（结果取绝对值）
func GetOffSetReferDay(comp string) int {
	compTime, _ := time.ParseInLocation("2006-01-02 15:04:05", comp, time.Local)
	referTime := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.Local)
	offset := compTime.Sub(referTime).Hours() / 24
	return int(offset)
}
