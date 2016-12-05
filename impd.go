package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/blacksails/pdevent"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	key := flag.String("k", "", "Pagerduty service key")
	desc := flag.String("d", "Intermapper incident", "Incident description")
	event := flag.String("e", "Down", "Intermapper event type")
	iKey := flag.String("i", "", "Incident key")
	flag.Parse()

	// Check if key is defined
	if *key == "" {
		return errors.New("The k flag is required")
	}

	// Create Pagerduty event API client
	c := pdevent.NewClient(*key, "Intermapper")

	// Determine what event to send
	switch *event {
	case "Down", "Alarm", "Warning":
		return c.Trigger(*desc, *iKey)
	case "Up", "OK":
		return c.Resolve(*iKey)
	case "ACK":
		return c.Acknowledge(*iKey)
	default:
		return c.Trigger(*desc, *iKey)
	}
}
