package gtime

import (
	"encoding/json"
	"fmt"
	"gtime/styles"
	"time"
)

const (
	sun  = "\U0001F31E"
	moon = "\U0001F319"
)

type TInfo struct {
	Location string    `json:"location"`
	Time     time.Time `json:"time"`
	Offset   string    `json:"offset"`
}

var res []TInfo

func getOffset(target, tzTime time.Time) (offsetStr string) {
	var m int
	var suffix string
	_, tz1 := target.Zone()
	_, tz2 := tzTime.Zone()
	hours := float64(tz1-tz2) / 3600.0
	wHours := int(hours)
	min := int((hours - float64(wHours)) * 60)
	if hours < 0 {
		m = -1
		suffix = "ahead"
	} else {
		m = 1
		suffix = "behind"
	}

	offsetStr = fmt.Sprintf("%d hours %d minutes %s", (wHours * m), (min * m), suffix)
	return
}

func ConvertTime(target string, timezones []string, output string) error {
	targetTime, err := parseTarget(target)
	if err != nil {
		return (fmt.Errorf("failed parsing time %s", target))
	}

	res = append(res, TInfo{"Local", targetTime, "0 hours 0 minutes"})

	for _, v := range timezones {
		loc, err := time.LoadLocation(v)
		if err != nil {
			return (fmt.Errorf("failed loading timezone %s", v))
		}
		tzTime := targetTime.In(loc)
		offsetStr := getOffset(targetTime, tzTime)
		res = append(res, TInfo{v, tzTime, offsetStr})
	}

	if output == "table" {
		header := []string{"Location", "Date", "Time", "Offset"}
		data := [][]string{}

		for _, v := range res {
			var time string
			date := v.Time.Format("2/1/2006")
			if isTimeBetween(v.Time, 6, 18) {
				time = v.Time.Format("2/1/2006 15:04") + " " + sun
			} else {
				time = v.Time.Format("2/1/2006 15:04") + " " + moon
			}
			data = append(data, []string{v.Location, date, time, v.Offset})
		}
		styles.RenderTable(header, data)
	} else if output == "json" {
		// encode res object as JSON
		json, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return (fmt.Errorf("failed encoding to JSON"))
		}

		fmt.Println(string(json))

	}
	return nil
}

func parseTarget(target string) (time.Time, error) {
	var tzTime time.Time

	if target != "" {
		localLoc, err := time.LoadLocation("Local")
		if err != nil {
			return time.Time{}, (fmt.Errorf("failed parsing local time"))
		}
		tzTime, err = time.ParseInLocation("2/1/2006 15:04", target, localLoc)
		if err != nil {
			return time.Time{}, (fmt.Errorf("failed parsing time at timezone %s", target))
		}
	}

	return tzTime, nil
}

func isTimeBetween(t time.Time, startHour, endHour int) bool {
	hour := t.Hour()
	return hour >= startHour && hour < endHour
}
