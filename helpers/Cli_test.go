package helpers

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/urfave/cli"
)

func TestStringFlag(t *testing.T) {
	Convey("Creating cli.StringFlag", t, func() {
		name := "test"
		value := "test_value"
		usage := "test usage"
		expected := cli.StringFlag{
			Name:  name,
			Value: value,
			Usage: usage,
		}
		result := StringFlag(name, value, usage)

		Convey("Asserting return type", func() {
			So(result, ShouldHaveSameTypeAs, cli.StringFlag{})
		})

		Convey("Asserting return attr type", func() {
			So(result.Name, ShouldHaveSameTypeAs, name)
			So(result.Value, ShouldHaveSameTypeAs, value)
			So(result.Usage, ShouldHaveSameTypeAs, usage)
		})

		Convey("Asserting return value", func() {
			So(result, ShouldHaveSameTypeAs, expected)
		})
	})
}

func TestBoolFlag(t *testing.T) {
	Convey("Creating cli.BoolFlag", t, func() {
		name := "test"
		usage := "test usage"
		expected := cli.BoolFlag{
			Name:  name,
			Usage: usage,
		}
		result := BoolFlag(name, usage)

		Convey("Asserting return type", func() {
			So(result, ShouldHaveSameTypeAs, cli.BoolFlag{})
		})

		Convey("Asserting return attr type", func() {
			So(result.Name, ShouldHaveSameTypeAs, name)
			So(result.Usage, ShouldHaveSameTypeAs, usage)
		})

		Convey("Asserting return value", func() {
			So(result, ShouldHaveSameTypeAs, expected)
		})
	})
}

func TestIntFlag(t *testing.T) {
	Convey("Creating cli.BoolFlag", t, func() {
		name := "test"
		value := 1
		usage := "test usage"
		expected := cli.IntFlag{
			Name:  name,
			Value: value,
			Usage: usage,
		}
		result := IntFlag(name, value, usage)

		Convey("Asserting return type", func() {
			So(result, ShouldHaveSameTypeAs, cli.IntFlag{})
		})

		Convey("Asserting return attr type", func() {
			So(result.Name, ShouldHaveSameTypeAs, name)
			So(result.Value, ShouldHaveSameTypeAs, value)
			So(result.Usage, ShouldHaveSameTypeAs, usage)
		})

		Convey("Asserting return value", func() {
			So(result, ShouldHaveSameTypeAs, expected)
		})
	})
}

func TestDurationFlag(t *testing.T) {
	Convey("Creating cli.BoolFlag", t, func() {
		name := "test"
		value := time.Since(time.Now().AddDate(0, 0, -1))
		usage := "test usage"
		expected := cli.DurationFlag{
			Name:  name,
			Value: value,
			Usage: usage,
		}
		result := DurationFlag(name, value, usage)

		Convey("Asserting return type", func() {
			So(result, ShouldHaveSameTypeAs, cli.DurationFlag{})
		})

		Convey("Asserting return attr type", func() {
			So(result.Name, ShouldHaveSameTypeAs, name)
			So(result.Value, ShouldHaveSameTypeAs, value)
			So(result.Usage, ShouldHaveSameTypeAs, usage)
		})

		Convey("Asserting return value", func() {
			So(result, ShouldHaveSameTypeAs, expected)
		})
	})
}
