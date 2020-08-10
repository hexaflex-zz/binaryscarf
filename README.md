## binaryscarf

binaryscarf generates [double-knit][1] patterns for some corpus of input
text. The layout follows the same style as described [here][2]. Output
is saved as a PNG image.

The characters are encoded as 7-bit binary (ASCII), where each binary digit
represents the foreground or background color. The scarf will use a
configurable number of columns.

The length of the pattern will be dependent on the number of selected
columns, plus the length of the input text. The width depends on the number
of columns and the configurable spacing between colums.

The words in the input text are layed out top-to-bottom, left-to-right.
For instance, the phrase "Machines take me by surprise with great frequency.",
with 2 colums will be drawn as follows:

	===            ===============
	M i            1001101 1101001
	a s            1100001 1110011
	c e            1100011 1100101
	h              1101000
	i w            1101001 1110111
	n i            1101110 1101001
	e t            1100101 1110100
	s h            1110011 1101000
	              
	t g            1110100 1100111
	a r            1100001 1110010
	k e            1101011 1100101
	e a     =>     1100101 1100001
	  t                    1110100
	m              1101101
	e f            1100101 1100110
	  r                    1110010
	b e            1100010 1100101
	y q            1111001 1110001
	  u                    1110101
	s e            1110011 1100101
	u n            1110101 1101110
	r c            1110010 1100011
	p y            1110000 1111001
	r .            1110010 0101110
	===            ===============


The pattern does not assume any specific cast-on, bind-off or bordering.
This will be up to the knitter to choose. However, for best results, I
recommend the following:

* [Two-colour alternating, invisible cast-on][3]
* [Double-sided knitting with slip-stitch edges][4]
* [One-needle Kitchener bind-off][5]


[1]: http://tutorials.knitpicks.com/double-knitting/
[2]: http://smeech.co.uk/binary-scarves/
[3]: https://www.youtube.com/watch?v=Y5oi2fknnH0
[4]: https://www.youtube.com/watch?v=TLZQEXQl4Yw
[5]: https://www.youtube.com/watch?v=upoeo5hwDFo


## Install

    $ go get github.com/hexaflex/binaryscarf


## Usage

	usage: binaryscarf [options] <textfile>
	   or: cat <textfile> | binaryscarf [options]
	
	  -border int
	    	Optional decorative n-row border at top and bottom of work: [0..n] (default 2)
	  -color-a string
	    	Foreground color in 24-bit hexadecimal notation. (default "0xffffff")
	  -color-b string
	    	Background color in 24-bit hexadecimal notation. (default "0x647384")
	  -columns int
	    	Number of 7-bit columns to generate: [1..n] (default 3)
	  -out string
	    	Filename for resulting PNG image. (default "out.png")
	  -repeat int
	    	Number of times to repeat the input text: [1..n] (default 1)
	  -spacing int
	    	Number of stitches to leave blank between columns: [0..n] (default 2)
	  -stitch-height int
	    	Height of a single stitch, in pixels: [1..n] (default 3)
	  -stitch-width int
	    	Width of a single stitch, in pixels: [1..n] (default 2)
	  -version
	    	Display version information.


Refer to `testdata/Makefile` for some invocation examples. And the images
in `testdata/` for output samples.

The `stitch-width` and `stitch-height` values are provided, so the resulting
pattern may reflect the way real knit stitches are shaped. They are never
perfectly square. Desiging a pattern with square stitches, will end up
looking squashed in either width or height, when knitted. Adjusting the
pattern to account for this property, will ensure the desired result is
achieved. A stitch-width of 2 and stitch-height of 3 usually gives a
reasonable approximation of the real thing.


## license

Unless otherwise noted, the contents of this project are subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.
