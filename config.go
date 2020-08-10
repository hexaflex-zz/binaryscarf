// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Config defines application configuration.
type Config struct {
	Output       string         // Output file name.
	Columns      int            // Number of 7-bit columns to use.
	Spacing      int            // Spacing between columns.
	Border       int            // Numer of rows in optional top/bottom borders.
	StitchWidth  int            // Stitch width, in pixels.
	StitchHeight int            // Stitch height, in pixels.
	Colors       [2]color.NRGBA // Foreground and background colors.
}

func (c *Config) BorderWidth() int   { return c.Border * c.StitchWidth }
func (c *Config) BorderHeight() int  { return c.Border * c.StitchHeight }
func (c *Config) SpacingWidth() int  { return c.Spacing * c.StitchWidth }
func (c *Config) SpacingHeight() int { return c.Spacing * c.StitchHeight }
func (c *Config) ColumnWidth() int   { return 7 * c.StitchWidth }

// ParseConfig parses and validates command line arguments into a config
// struct. This function prints an error and exits the program if invalid
// data is encountered.
//
// Additionally, it returns the text contents. Read either from a file, or
// stdin.
func ParseConfig() (*Config, []byte) {
	var c Config
	c.Columns = 3
	c.Spacing = 2
	c.Border = 2
	c.Output = "out.png"
	c.StitchWidth = 2
	c.StitchHeight = 3

	clrA := "0xffffff"
	clrB := "0x647384"
	repeat := 1

	flag.Usage = func() {
		fmt.Println("usage:", os.Args[0], "[options] <textfile>")
		fmt.Println("   or: cat <textfile> |", os.Args[0], "[options]")
		flag.PrintDefaults()
	}

	flag.StringVar(&c.Output, "out", c.Output, "Filename for resulting PNG image.")
	flag.StringVar(&clrA, "color-a", clrA, "Foreground color in 24-bit hexadecimal notation.")
	flag.StringVar(&clrB, "color-b", clrB, "Background color in 24-bit hexadecimal notation.")
	flag.IntVar(&c.StitchWidth, "stitch-width", c.StitchWidth, "Width of a single stitch, in pixels: [1..n]")
	flag.IntVar(&c.StitchHeight, "stitch-height", c.StitchHeight, "Height of a single stitch, in pixels: [1..n]")
	flag.IntVar(&c.Columns, "columns", c.Columns, "Number of 7-bit columns to generate: [1..n]")
	flag.IntVar(&c.Spacing, "spacing", c.Spacing, "Number of stitches to leave blank between columns: [0..n]")
	flag.IntVar(&c.Border, "border", c.Border, "Optional decorative n-row border at top and bottom of work: [0..n]")
	flag.IntVar(&repeat, "repeat", repeat, "Number of times to repeat the input text: [1..n]")
	version := flag.Bool("version", false, "Display version information.")
	flag.Parse()

	if *version {
		fmt.Println(Version())
		os.Exit(0)
	}

	c.Output = filepath.Clean(c.Output)
	if len(c.Output) == 0 || c.Columns < 1 || c.Spacing < 0 || c.Border < 0 ||
		c.StitchWidth < 1 || c.StitchHeight < 1 || repeat < 1 {
		flag.Usage()
		os.Exit(1)
	}

	err := parseColor(&c.Colors[0], clrA)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = parseColor(&c.Colors[1], clrB)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return &c, filterText(data, repeat)
	}

	data, err := ioutil.ReadFile(filepath.Clean(flag.Arg(0)))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return &c, filterText(data, repeat)
}

var (
	bCR     = []byte{'\r'}
	bLF     = []byte{'\n'}
	bTab    = []byte{'\t'}
	bSpace  = []byte{' '}
	b2Space = []byte{' ', ' '}
)

// filterText returns n times the given text, with excessive
// whitespace and formatting removed.
func filterText(v []byte, n int) []byte {
	v = bytes.Replace(v, bCR, nil, -1)
	v = bytes.Replace(v, bLF, bSpace, -1)
	v = bytes.Replace(v, bTab, bSpace, -1)
	v = bytes.Replace(v, b2Space, bSpace, -1)
	v = bytes.TrimSpace(v)

	if len(v) == 0 {
		fmt.Fprintln(os.Stderr, "input text is empty after filtering")
		os.Exit(1)
	}

	return bytes.Repeat(v, n)
}

func parseColor(out *color.NRGBA, v string) error {
	pickError := func(a, b, c error) error {
		if a != nil {
			return a
		}
		if b != nil {
			return b
		}
		return c
	}

	v = strings.ToLower(v)
	if len(v) != 8 || !strings.HasPrefix(v, "0x") {
		return fmt.Errorf("invalid or missing color value: %q", v)
	}

	v = v[2:]

	r, er := strconv.ParseUint(v[:2], 16, 8)
	g, eg := strconv.ParseUint(v[2:4], 16, 8)
	b, eb := strconv.ParseUint(v[4:], 16, 8)

	out.R = byte(r)
	out.G = byte(g)
	out.B = byte(b)
	out.A = 0xff
	return pickError(er, eg, eb)
}
