# This file is subject to a BSD license.
# Its contents can be found in the enclosed LICENSE file.

IN = sample.txt
BS = binaryscarf

all: $(IN)
	$(BS) -spacing 1 -border 1 -columns 3 \
		-stitch-width 5 -stitch-height 5 \
		-repeat 1 -out 1.png $<
	$(BS) -spacing 2 -border 2 -columns 3 \
		-stitch-width 5 -stitch-height 5 \
		-repeat 1 -out 2.png $<
	$(BS) -spacing 2 -border 1 -columns 2 \
		-stitch-width 5 -stitch-height 5 \
		-repeat 1 -out 3.png $<
	$(BS) -spacing 2 -border 1 -columns 1 \
		-stitch-width 5 -stitch-height 5 \
		-repeat 1 -out 4.png $<
	$(BS) -spacing 2 -border 0 -columns 3 \
		-stitch-width 2 -stitch-height 3 \
		-repeat 10 -out 5.png $<

clean:
	rm *.png

