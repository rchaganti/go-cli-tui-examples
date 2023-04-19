package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type TInfo struct {
	Location string
	Time     string
	offset   string
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
	res = append(res, TInfo{"Local", target.Format("2/1/2006 15:04"), "0 hours 0 minutes"})

	for _, v := range tz {
		loc, _ := time.LoadLocation(v)
		tzTime := target.In(loc)
		offsetStr := getOffset(target, tzTime)
		res = append(res, TInfo{v, tzTime.Format("2/1/2006 15:04"), offsetStr})
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	fmt.Fprintln(w, "Location\tTime\tOffset\t")
	for _, v := range res {
		fmt.Fprintln(w, v.Location, "\t", v.Time, "\t", v.offset, "\t")
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

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-targetTime] <timezones...>\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	nArgs := flag.NArg()
	if nArgs < 1 {
		fmt.Println("You must specify at least one timezone string!")
		flag.Usage()
		os.Exit(1)
	}

	args := flag.Args()
	for _, v := range args {
		_, err := time.LoadLocation(v)

		if err != nil {
			fmt.Printf("%s is not a valid timezone\n", v)
			os.Exit(1)
		}
		tz = append(tz, v)
	}

	convertTime(targetTime)
}
