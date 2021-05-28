package main

import (
	_ "image"
	_ "image/jpeg"
	_ "image/png"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	println("x")
}

//
//def remove_transparency_from_png(img_path, background_color):
//	with Image.open(img_path).convert('RGBA') as png:
//	with Image.new('RGBA', png.size, background_color) as background:
//	# draw logo on background with given color
//	alpha_composite = Image.alpha_composite(background, png)
//	with alpha_composite.convert('RGB') as converted:
//	converted.save(img_path)
