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

func convertTime(target TargetTime) error {
	// Convert time to each timezone
	res = append(res, TInfo{"Local", target.Format("2/1/2006 15:04"), "0 hours 0 minutes"})

	for _, v := range tz {
		loc, err := time.LoadLocation(v)

		if err != nil {
			return (fmt.Errorf("%s is not a valid timezone", v))
		}

		tzTime := target.In(loc)
		offsetStr := getOffset(target.Time, tzTime)
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

// custom type for timezones flag
type Timezones []string

func (s *Timezones) String() string {
	return fmt.Sprint(*s)
}

func (s *Timezones) Set(value string) error {
	parts := strings.Split(value, ",")
	*s = append(*s, parts...)
	return nil
}

var tz Timezones

// custom type for targetTime flag
type TargetTime struct {
	time.Time
}

func (s *TargetTime) String() string {
	return s.Time.Format("2/1/2006 15:04")
}

func (s *TargetTime) Set(value string) error {
	localLoc, err := time.LoadLocation("Local")
	if err != nil {
		return fmt.Errorf("Cannot load local location: %v", err)
	}

	parsedTime, err := time.ParseInLocation("2/1/2006 15:04", value, localLoc)
	if err != nil {
		return fmt.Errorf("Error parsing input as time: %v", err)
	}

	s.Time = parsedTime
	return nil
}

func main() {
	flag.Var(&tz, "timezones", "Comma-separated list of timezones in the form Asia/Shanghai")

	target := TargetTime{time.Now()}

	flag.Var(&target, "targetTime", "Target time in the format 2/1/2006 15:04 (DD/MM/YYYY HH:MM)")
	flag.Parse()

	convertTime(target)
}
