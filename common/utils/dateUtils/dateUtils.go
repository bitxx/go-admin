package dateUtils

import "time"

//
//  GetDistanceOfTwoDate
//  @param timeStart
//  @param timeEnd
//  @return int64
//
func GetDistanceOfTwoDate(timeStart, timeEnd time.Time) int64 {
	before := timeStart.Unix()
	after := timeEnd.Unix()
	return (after - before) / 86400
}

func BeginTime(param time.Time) time.Time {
	timeStr := param.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return time.Unix(t.Unix(), 0)
}

func EndTimeNum(param time.Time) time.Time {
	timeStr := param.Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	return time.Unix(t.Unix()+86399, 999)
}

func ParseTimestrToTimestamp(timeStr string, flag int) int64 {
	var t int64
	loc, _ := time.LoadLocation("Local")
	if flag == 1 {
		t1, _ := time.ParseInLocation("2006.01.02 15:04:05", timeStr, loc)
		t = t1.Unix()
	} else if flag == 2 {
		t1, _ := time.ParseInLocation("2006-01-02 15:04", timeStr, loc)
		t = t1.Unix()
	} else if flag == 3 {
		t1, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
		t = t1.Unix()
	} else if flag == 4 {
		t1, _ := time.ParseInLocation("2006.01.02", timeStr, loc)
		t = t1.Unix()
	} else {
		t1, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
		t = t1.Unix()
	}
	return t
}

func ParseTimestrToTime(timeStr string, flag int) time.Time {
	//loc, _ := time.LoadLocation("Local")
	if flag == 1 {
		t1, _ := time.Parse("2006.01.02 15:04:05", timeStr)
		return t1
	} else if flag == 2 {
		t1, _ := time.Parse("2006-01-02 15:04", timeStr)
		return t1
	} else if flag == 3 {
		t1, _ := time.Parse("2006-01-02", timeStr)
		return t1
	} else if flag == 4 {
		t1, _ := time.Parse("2006.01.02", timeStr)
		return t1
	}
	t1, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	return t1
}

//
//  ConvertToStrByPrt
//  @Description:
//  @param dateTime
//  @param flag
//  @return string
//
func ConvertToStrByPrt(dateTime *time.Time, flag int) string {
	if dateTime == nil {
		return ""
	}
	switch flag {
	case 1:
		return dateTime.Format("2006-01-02")
	case 2:
		return dateTime.Format("2006-01-02 15:04")
	}
	return dateTime.Format("2006-01-02 15:04:05")
}

func ConvertToStr(dateTime time.Time, flag int) string {
	switch flag {
	case 1:
		return dateTime.Format("2006-01-02")
	case 2:
		return dateTime.Format("2006-01-02 15:04")
	case 3:
		return dateTime.Format("2006_01_02_15_04_05")
	}
	return dateTime.Format("2006-01-02 15:04:05")
}

//
//  ConvertToStrByPrt
//  @Description: 传入的地址是指针，避免外部频繁判断是否为空
//  @param dateTime
//  @param flag
//  @return string
//
/*func ConvertToStrByPrt(dateTime *time.Time, flag int) string {
	if dateTime==nil{
		return ""
	}
	switch flag {
	case 1:
		return dateTime.Format("2006-01-02")
	case 2:
		return dateTime.Format("2006-01-02 15:04")
	}
	return dateTime.Format("2006-01-02 15:04:05")
}*/
