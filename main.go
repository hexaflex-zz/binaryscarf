// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package main

import (
	"fmt"
	"image"
	"math"
	"os"

	"image/color"
	"image/png"
)

func main() {
	cfg, text := ParseConfig()
	charset := buildCharacterSet(cfg, text)
	img := drawPattern(cfg, charset)

	fd, err := os.Create(cfg.Output)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer fd.Close()

	err = png.Encode(fd, img)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// drawPattern draws the pattern and returns the resulting image.
func drawPattern(c *Config, charset []Char) image.Image {
	rect := computePatternRect(c, charset)
	img := image.NewNRGBA(rect)

	fillRect(img, 0, 0, rect.Dx(), rect.Dy(), c.Colors[0])

	// Draw decorative borders, if applicable.
	if c.Border > 0 {
		x := c.SpacingWidth()
		y := c.SpacingHeight()
		w := rect.Dx() - c.SpacingWidth()*2
		h := c.BorderHeight()
		fillRect(img, x, y, w, h, c.Colors[1])

		y = rect.Dy() - c.SpacingHeight() - c.BorderHeight()
		fillRect(img, x, y, w, h, c.Colors[1])
	}

	// Draw words, top down.
	for _, char := range charset {
		char.Plot(img, c.StitchWidth, c.StitchHeight, c.Colors)
	}

	return img
}

// fillRect fills the given rectangular with the specified color.
func fillRect(img *image.NRGBA, x, y, w, h int, clr color.NRGBA) {
	for py := y; py < y+h; py++ {
		for px := x; px < x+w; px++ {
			img.SetNRGBA(px, py, clr)
		}
	}
}

// computeRect returns a rectangle defining the size, in pixels, of the
// entire pattern.
func computePatternRect(c *Config, charset []Char) image.Rectangle {
	w := c.Columns*7 + ((c.Columns-1)+2)*c.Spacing
	h := 0

	// find highest y coordinate in charset. This will help determine
	// the total pattern height.
	for _, c := range charset {
		if h < c.Y {
			h = c.Y
		}
	}

	h += c.StitchHeight

	if c.Border > 0 {
		h += c.SpacingHeight() + c.BorderHeight() + c.SpacingHeight()
	}

	w *= c.StitchWidth
	return image.Rect(0, 0, w, h)
}

// buildCharacterSet converts the given text into a set of Char structs.
// Each one has its respective X/Y coordinate precomputed.
func buildCharacterSet(c *Config, v []byte) []Char {
	rows := int(math.Ceil(float64(len(v)) / float64(c.Columns)))
	out := make([]Char, 0, len(v))
	top := c.SpacingHeight()

	if c.Border > 0 {
		top += c.BorderHeight() + c.SpacingHeight()
	}

	bottom := top + rows*c.StitchHeight
	x, y := c.SpacingWidth(), top
	for _, b := range v {
		if b == ' ' && y == top {
			continue // drop the character
		}

		if b == ' ' && y >= bottom-c.StitchHeight {
			y = top
			x += c.ColumnWidth() + c.SpacingWidth()
			continue // drop the character
		}

		out = append(out, Char{
			Value: b,
			X:     x,
			Y:     y,
		})

		y += c.StitchHeight
		if y >= bottom {
			y = top
			x += c.ColumnWidth() + c.SpacingWidth()
		}
	}

	return out
}

// Char defines a single ASCII character to be drawn.
type Char struct {
	Value byte
	X, Y  int
}

// plotCharr draws the binary representation of the current character.
func (c *Char) Plot(img *image.NRGBA, sw, sh int, colors [2]color.NRGBA) {
	if c.Value == ' ' {
		//fmt.Fprintf(os.Stdout, "  %c: <ignored>\n", c.Value)
		return
	}

	//fmt.Fprintf(os.Stdout, "  %c: %07b\n", c.Value, c.Value)
	x, y := c.X, c.Y
	x = plotBit(img, x, y, sw, sh, colors[(c.Value>>6)&1])
	x = plotBit(img, x, y, sw, sh, colors[(c.Value>>5)&1])
	x = plotBit(img, x, y, sw, sh, colors[(c.Value>>4)&1])
	x = plotBit(img, x, y, sw, sh, colors[(c.Value>>3)&1])
	x = plotBit(img, x, y, sw, sh, colors[(c.Value>>2)&1])
	x = plotBit(img, x, y, sw, sh, colors[(c.Value>>1)&1])
	_ = plotBit(img, x, y, sw, sh, colors[(c.Value>>0)&1])
}

func plotBit(img *image.NRGBA, x, y, sw, sh int, clr color.NRGBA) int {
	for cy := 0; cy < sh; cy++ {
		for cx := 0; cx < sw; cx++ {
			img.SetNRGBA(x+cx, y+cy, clr)
		}
	}

	return x + sw
}
