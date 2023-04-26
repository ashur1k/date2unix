package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var utc bool
	var milliseconds bool
	var format string
	var help bool
	flag.BoolVar(&utc, "utc", false, "It's UTC time")
	flag.BoolVar(&milliseconds, "milli", false, "Print milliseconds")
	flag.StringVar(&format, "format", "", "Date format as golang 'time.Time'. Example: 2006-01-02T15:04:05Z07:00")
	flag.BoolVar(&help, "help", false, "Help")
	flag.Parse()

	var userDateTime string
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from pipe")
			os.Exit(1)
		}
		userDateTime = strings.TrimSpace(string(bytes))
	} else {
		userDateTime = strings.Join(flag.Args(), " ")
	}

	if help || userDateTime == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	dateFormats := []string{
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		time.DateTime,
		time.DateOnly,
		time.TimeOnly,
		"2006-01-02 15:04",
		"2006-01-02T15:04",
	}

	if format != "" {
		dateFormats = append(dateFormats, format)
	}

	userTimestamp, err := strconv.ParseInt(userDateTime, 10, 64)
	if err == nil {
		var thetime time.Time
		if userTimestamp > 9999999999 {
			milliseconds = true
		}
		if milliseconds {
			thetime = time.UnixMilli(userTimestamp)
		} else {
			thetime = time.Unix(userTimestamp, 0)
		}
		var datetime string
		if utc {
			datetime = thetime.UTC().String()
		} else {
			datetime = thetime.String()
		}
		fmt.Println(datetime)
		os.Exit(0)
	} else {
		for _, dateFormat := range dateFormats {
			var thetime time.Time
			var err error
			if utc {
				thetime, err = time.Parse(dateFormat, userDateTime)
			} else {
				loc, err := time.LoadLocation("Local")
				if err != nil {
					panic(err)
				}
				thetime, err = time.ParseInLocation(dateFormat, userDateTime, loc)
			}

			if err == nil && thetime.Unix() != -62135596800 {
				var timestamp int64
				if milliseconds {
					timestamp = thetime.UnixMilli()
				} else {
					timestamp = thetime.Unix()
				}
				fmt.Println("Timestamp:", timestamp)
				os.Exit(0)
			}
		}
	}
	fmt.Println("Can't parse date or timestamp!")
	os.Exit(1)
}
