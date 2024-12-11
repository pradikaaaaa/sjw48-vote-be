package lib

import (
	"fmt"
	"time"
)

func GetTimeNow(param string) string {
	location, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		fmt.Println("Error loading location:", err)
		return ""
	}

	currentTime := time.Now().In(location)

	switch param {
	case "timestime":
		return currentTime.Format("2006-01-02 15:04:05")
	case "date":
		return currentTime.Format("2006-01-02")
	case "date2":
		return currentTime.Format("020106")
	case "year":
		return fmt.Sprint(currentTime.Year())
	case "month":
		return fmt.Sprint(int(currentTime.Month()))
	case "month-name":
		return fmt.Sprint(currentTime.Month())
	case "day":
		return fmt.Sprint(currentTime.Day())
	case "hour":
		return string(currentTime.Hour())
	case "minutes":
		return string(currentTime.Minute())
	case "second":
		return string(currentTime.Second())
	case "unixmicro":
		return string(currentTime.UnixMicro())
	default:
		fmt.Println("masukan parameter")
		return ""
	}
}
