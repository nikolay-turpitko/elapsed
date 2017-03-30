package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

var (
	help  bool
	title string

	usage = `
Program intended to be used with 'watch' command to create simple CLI countdown timer.

Program do the following:
- if duration specified, converts specified duration to time as current time plus duration;
- if absolute time specified, uses it without convertion;
- prints to stderr title (if present) in the same color as duration, described below;
- prints specified time to stdout;
- prints into stderr duration remainded to specified time;
  prints it in green if specified time have not elapsed yet, in red otherwise;
- returns 0 if specified time have not elapsed yet, 1 otherwise (and in case of all other errors);

Program parses and prints time in the following format:

  RFC3339     = "2006-01-02T15:04:05Z07:00"

Program parses duration in the Go format such as "300ms", "-1.5h" or "2h45m".

Example:

  watch -bce ./%[1]s $(./%[1]s 10s)
  watch -bce "./elapsed $(./elapsed 10s) || (notify-send Timeout && false)"
  watch -bce "./elapsed $(./elapsed 10s) || (notify-send Timeout && paplay /usr/share/sounds/freedesktop/stereo/complete.oga && false)"

`
)

func init() {
	flag.BoolVar(&help, "h", false, "display usage")
	flag.StringVar(&title, "t", "", "title")
	flag.Usage = func() {
		name := filepath.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s [-h] [-t title] [time|duration]\n\n", name)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, usage, name)
	}
}

func main() {
	flag.Parse()
	if help || flag.NArg() < 1 {
		flag.Usage()
		os.Exit(0)
	}
	s := strings.TrimSpace(flag.Arg(0))
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		d, err2 := time.ParseDuration(s)
		if err2 != nil {
			log.Println(err)
			log.Fatalln(err2)
		}
		t = time.Now().Add(d)
	}
	d := t.Sub(time.Now())
	elapsed := time.Now().After(t)
	if title != "" {
		color.NoColor = false
		if elapsed {
			fmt.Fprintln(os.Stderr, color.RedString("%s", title))
		} else {
			fmt.Fprintln(os.Stderr, color.GreenString("%s", title))
		}
	}
	color.NoColor = true
	fmt.Println(t.Format(time.RFC3339))
	color.NoColor = false
	if elapsed {
		fmt.Fprintln(os.Stderr, color.RedString("%v", d))
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, color.GreenString("%v", d))
}
