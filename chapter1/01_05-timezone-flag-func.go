package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type TInfo struct {
	Location string
	Time     string
	Offset   string
}

var res []TInfo
var tz []string

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

func convertTime(target time.Time) error {
	// Convert time to each timezone
	res = append(res, TInfo{"Local", target.Format("2/1/2006 15:04"), "0"})

	for _, v := range tz {
		loc, _ := time.LoadLocation(v)

		tzTime := target.In(loc)
		offsetStr := getOffset(target, tzTime)
		res = append(res, TInfo{v, tzTime.Format("2/1/2006 15:04"), offsetStr})
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	fmt.Fprintln(w, "Location\tTime\tOffset\t")
	for _, v := range res {
		fmt.Fprintln(w, v.Location, "\t", v.Time, "\t", v.Offset, "\t")
	}
	w.Flush()

	return nil
}

func main() {
	var targetTime time.Time = time.Now()

	flag.Func("targetTime", "Date-time in '2/1/2006 15:04' format", func(flagValue string) error {
		localLoc, err := time.LoadLocation("Local")
		if err != nil {
			return (fmt.Errorf("Failed parsing local time!"))
		}
		tzTime, err := time.ParseInLocation("2/1/2006 15:04", flagValue, localLoc)
		if err != nil {
			return (fmt.Errorf("Failed parsing time at timezone %s", flagValue))
		}

		if tzTime.Weekday() == time.Saturday || tzTime.Weekday() == time.Sunday {
			return fmt.Errorf("Schedule meetings on a weekday!")
		}

		targetTime = tzTime
		return nil
	})

	flag.Func("timezones", "comma-separated list of timezones", func(value string) error {
		items := strings.Split(value, ",")
		for _, v := range items {
			_, err := time.LoadLocation(v)

			if err != nil {
				return (fmt.Errorf("%s is not a valid timezone", v))
			}
			tz = append(tz, v)
		}

		return nil
	})

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <-timezones> [-targetTime]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	convertTime(targetTime)
}
