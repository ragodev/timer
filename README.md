# timer

## Usage:

```
timer, a general-purpose countdown timer.

Usage:

    timer [flags] time[ s | m | h ]

Flags:

    --persistent, -p   continue beeping until user hits return.
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
```
