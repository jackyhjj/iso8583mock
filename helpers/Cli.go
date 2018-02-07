// Package helpers implements commonly used functions (command line interface) //
package helpers

import (
	"time"

	"github.com/urfave/cli"
)

// StringFlag - get string flag
func StringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

// BoolFlag - get boolean flag
func BoolFlag(name, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

// IntFlag - get integer flag
func IntFlag(name string, value int, usage string) cli.IntFlag {
	return cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

// DurationFlag - get duration flag
func DurationFlag(name string, value time.Duration, usage string) cli.DurationFlag {
	return cli.DurationFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}
