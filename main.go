package main

import (
	"fmt"
	"os"
	"time"
)

var (
	fLaundry bool
	fOven    bool
	fTea     bool
	fQuiet   bool
)

const usage = `timer, a general-purpose countdown timer.

Usage:

    timer [flags] time[ s | m | h ]

Flags:

    --quiet, -q        do not send system beep.
    --tea              this is a tea timer.
    --oven             this is an oven timer.
    --laundry          this is a laundry timer.

Examples:

    # Run a timer for 30 seconds:
    timer 30

    # Run a tea timer for 3 minutes:
    timer --tea 3m

    # Run a laundry timer for 1 hour:
    timer --laundry 1h
`

func main() {
	t := parseArgs()

	d, err := time.ParseDuration(t)
	if err != nil {
		d, err = time.ParseDuration(t + "s")
		if err != nil {
			fmt.Printf("Could not parse duration: %s\n", t)
			os.Exit(1)
		}
	}

	ticker := time.NewTicker(time.Second)
	timer := time.NewTimer(d)
	countdown := int64(d.Seconds())

	for {
		fmt.Printf("\r%02d:%02d...", countdown/60, countdown%60)
		select {
		case <-timer.C:
			print()
			alarm()
			persist()
			return
		case <-ticker.C:
			countdown--
		}
	}
}

func alarm() {
	if !fQuiet {
		fmt.Print("\a")
		time.Sleep(time.Millisecond * 250)
		fmt.Print("\a")
		time.Sleep(time.Millisecond * 250)
		fmt.Print("\a")
	}
}

func persist() {
	fmt.Printf("\nHit return to end timer.")
	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(time.Second * 2)
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				alarm()
			}
		}
	}()

	var foo string
	fmt.Scanln(&foo)
	close(done)
}

func usageExit() {
	fmt.Printf(usage)
	os.Exit(0)
}

func parseArgs() string {
	var t string
	if len(os.Args) == 1 {
		usageExit()
	}
	for _, arg := range os.Args[1:] {
		switch arg {
		case "-h", "--help", "-help":
			usageExit()
		case "--tea", "-tea":
			fTea = true
		case "--oven", "-oven":
			fOven = true
		case "--laundry", "-laundry":
			fLaundry = true
		case "-q", "--quiet", "-quiet":
			fQuiet = true
		default:
			t = arg
		}
	}
	if t == "" {
		usageExit()
	}
	return t
}

func print() {
	fmt.Println()
	switch {
	case fTea:
		fmt.Printf(tea)
	case fOven:
		fmt.Printf(oven)
	case fLaundry:
		fmt.Printf(laundry)
	default:
		fmt.Printf("beep! beep! beep!\n")
	}
}
