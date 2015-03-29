package gift

import (
	"image"
	"image/color"
	"testing"
)

func comparePix(p1, p2 []uint8) bool {
	if len(p1) != len(p2) {
		return false
	}
	for i := 0; i < len(p1); i++ {
		if p1[i] != p2[i] {
			return false
		}
	}
	return true
}

func TestRotate90(t *testing.T) {
	img0 := image.NewGray(image.Rect(-1, -1, 3, 1))
	img0.Pix = []uint8{
		1, 2, 3, 4,
		5, 6, 7, 8,
	}
	img1_exp := image.NewGray(image.Rect(0, 0, 2, 4))
	img1_exp.Pix = []uint8{
		4, 8,
		3, 7,
		2, 6,
		1, 5,
	}

	f := Rotate90()
	img1 := image.NewGray(f.Bounds(img0.Bounds()))
	f.Draw(img1, img0, nil)

	if img1.Bounds().Size() != img1_exp.Bounds().Size() {
		t.Errorf("expected %v got %v", img1_exp.Bounds().Size(), img1.Bounds().Size())
	}
	if !comparePix(img1_exp.Pix, img1.Pix) {
		t.Errorf("expected %v got %v", img1_exp.Pix, img1.Pix)
	}
}

func TestRotate180(t *testing.T) {
	img0 := image.NewGray(image.Rect(-1, -1, 3, 1))
	img0.Pix = []uint8{
		1, 2, 3, 4,
		5, 6, 7, 8,
	}
	img1_exp := image.NewGray(image.Rect(0, 0, 4, 2))
	img1_exp.Pix = []uint8{
		8, 7, 6, 5,
		4, 3, 2, 1,
	}

	f := Rotate180()
	img1 := image.NewGray(f.Bounds(img0.Bounds()))
	f.Draw(img1, img0, nil)

	if img1.Bounds().Size() != img1_exp.Bounds().Size() {
		t.Errorf("expected %v got %v", img1_exp.Bounds().Size(), img1.Bounds().Size())
	}
	if !comparePix(img1_exp.Pix, img1.Pix) {
		t.Errorf("expected %v got %v", img1_exp.Pix, img1.Pix)
	}
}

func TestRotate270(t *testing.T) {
	img0 := image.NewGray(image.Rect(-1, -1, 3, 1))
	img0.Pix = []uint8{
		1, 2, 3, 4,
		5, 6, 7, 8,
	}
	img1_exp := image.NewGray(image.Rect(0, 0, 2, 4))
	img1_exp.Pix = []uint8{
		5, 1,
		6, 2,
		7, 3,
		8, 4,
	}

	f := Rotate270()
	img1 := image.NewGray(f.Bounds(img0.Bounds()))
	f.Draw(img1, img0, nil)

	if img1.Bounds().Size() != img1_exp.Bounds().Size() {
		t.Errorf("expected %v got %v", img1_exp.Bounds().Size(), img1.Bounds().Size())
	}
	if !comparePix(img1_exp.Pix, img1.Pix) {
		t.Errorf("expected %v got %v", img1_exp.Pix, img1.Pix)
	}
}

func TestFlipHorizontal(t *testing.T) {
	img0 := image.NewGray(image.Rect(-1, -1, 3, 1))
	img0.Pix = []uint8{
		1, 2, 3, 4,
		5, 6, 7, 8,
	}
	img1_exp := image.NewGray(image.Rect(0, 0, 4, 2))
	img1_exp.Pix = []uint8{
		4, 3, 2, 1,
		8, 7, 6, 5,
	}

	f := FlipHorizontal()
	img1 := image.NewGray(f.Bounds(img0.Bounds()))
	f.Draw(img1, img0, nil)

	if img1.Bounds().Size() != img1_exp.Bounds().Size() {
		t.Errorf("expected %v got %v", img1_exp.Bounds().Size(), img1.Bounds().Size())
	}
	if !comparePix(img1_exp.Pix, img1.Pix) {
		t.Errorf("expected %v got %v", img1_exp.Pix, img1.Pix)
	}
}

func TestFlipVertical(t *testing.T) {
	img0 := image.NewGray(image.Rect(-1, -1, 3, 1))
	img0.Pix = []uint8{
		1, 2, 3, 4,
		5, 6, 7, 8,
	}
	img1_exp := image.NewGray(image.Rect(0, 0, 4, 2))
	img1_exp.Pix = []uint8{
		5, 6, 7, 8,
		1, 2, 3, 4,
	}

	f := FlipVertical()
	img1 := image.NewGray(f.Bounds(img0.Bounds()))
	f.Draw(img1, img0, nil)

	if img1.Bounds().Size() != img1_exp.Bounds().Size() {
		t.Errorf("expected %v got %v", img1_exp.Bounds().Size(), img1.Bounds().Size())
	}
	if !comparePix(img1_exp.Pix, img1.Pix) {
		t.Errorf("expected %v got %v", img1_exp.Pix, img1.Pix)
	}
}

func TestTranspose(t *testing.T) {
	img0 := image.NewGray(image.Rect(-1, -1, 3, 1))
	img0.Pix = []uint8{
		1, 2, 3, 4,
		5, 6, 7, 8,
	}
	img1_exp := image.NewGray(image.Rect(0, 0, 2, 4))
	img1_exp.Pix = []uint8{
		1, 5,
		2, 6,
		3, 7,
		4, 8,
	}

	f := Transpose()
	img1 := image.NewGray(f.Bounds(img0.Bounds()))
	f.Draw(img1, img0, nil)

	if img1.Bounds().Size() != img1_exp.Bounds().Size() {
		t.Errorf("expected %v got %v", img1_exp.Bounds().Size(), img1.Bounds().Size())
	}
	if !comparePix(img1_exp.Pix, img1.Pix) {
		t.Errorf("expected %v got %v", img1_exp.Pix, img1.Pix)
	}
}

func TestTransverse(t *testing.T) {
	img0 := image.NewGray(image.Rect(-1, -1, 3, 1))
	img0.Pix = []uint8{
		1, 2, 3, 4,
		5, 6, 7, 8,
	}
	img1_exp := image.NewGray(image.Rect(0, 0, 2, 4))
	img1_exp.Pix = []uint8{
		8, 4,
		7, 3,
		6, 2,
		5, 1,
	}

	f := Transverse()
	img1 := image.NewGray(f.Bounds(img0.Bounds()))
	f.Draw(img1, img0, nil)

	if img1.Bounds().Size() != img1_exp.Bounds().Size() {
		t.Errorf("expected %v got %v", img1_exp.Bounds().Size(), img1.Bounds().Size())
	}
	if !comparePix(img1_exp.Pix, img1.Pix) {
		t.Errorf("expected %v got %v", img1_exp.Pix, img1.Pix)
	}
}

func TestCrop(t *testing.T) {
	testData := []struct {
		desc           string
		r              image.Rectangle
		srcb, dstb     image.Rectangle
		srcPix, dstPix []uint8
	}{
		{
			"crop (0, 0, 0, 0)",
			image.Rect(0, 0, 0, 0),
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 0, 0),
			[]uint8{
				0x00, 0x40, 0x00, 0x40, 0x00,
				0x60, 0xB0, 0xA0, 0xB0, 0x60,
				0x00, 0x80, 0x00, 0x80, 0x00,
			},
			[]uint8{},
		},
		{
			"crop (1, 1, -1, -1)",
			image.Rectangle{image.Pt(1, 1), image.Pt(-1, -1)},
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 0, 0),
			[]uint8{
				0x00, 0x40, 0x00, 0x40, 0x00,
				0x60, 0xB0, 0xA0, 0xB0, 0x60,
				0x00, 0x80, 0x00, 0x80, 0x00,
			},
			[]uint8{},
		},
		{
			"crop (-1, 0, 3, 2)",
			image.Rect(-1, 0, 3, 2),
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 4, 2),
			[]uint8{
				0x00, 0x40, 0x00, 0x40, 0x00,
				0x60, 0xB0, 0xA0, 0xB0, 0x60,
				0x00, 0x80, 0x00, 0x80, 0x00,
			},
			[]uint8{
				0x60, 0xB0, 0xA0, 0xB0,
				0x00, 0x80, 0x00, 0x80,
			},
		},
		{
			"crop (-100, -100, 2, 2)",
			image.Rect(-100, -100, 2, 2),
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x00, 0x40, 0x00, 0x40, 0x00,
				0x60, 0xB0, 0xA0, 0xB0, 0x60,
				0x00, 0x80, 0x00, 0x80, 0x00,
			},
			[]uint8{
				0x00, 0x40, 0x00,
				0x60, 0xB0, 0xA0,
				0x00, 0x80, 0x00,
			},
		},
		{
			"crop (-100, -100, 100, 100)",
			image.Rect(-100, -100, 100, 100),
			image.Rect(-1, -1, 4, 2),
			image.Rect(0, 0, 5, 3),
			[]uint8{
				0x00, 0x40, 0x00, 0x40, 0x00,
				0x60, 0xB0, 0xA0, 0xB0, 0x60,
				0x00, 0x80, 0x00, 0x80, 0x00,
			},
			[]uint8{
				0x00, 0x40, 0x00, 0x40, 0x00,
				0x60, 0xB0, 0xA0, 0xB0, 0x60,
				0x00, 0x80, 0x00, 0x80, 0x00,
			},
		},
	}

	for _, d := range testData {
		src := image.NewGray(d.srcb)
		src.Pix = d.srcPix

		f := Crop(d.r)
		dst := image.NewGray(f.Bounds(src.Bounds()))
		f.Draw(dst, src, nil)

		if !checkBoundsAndPix(dst.Bounds(), d.dstb, dst.Pix, d.dstPix) {
			t.Errorf("test [%s] failed: %#v, %#v", d.desc, dst.Bounds(), dst.Pix)
		}
	}
}

func TestCropToSize(t *testing.T) {
	testData := []struct {
		desc           string
		w, h           int
		anchor         Anchor
		srcb, dstb     image.Rectangle
		srcPix, dstPix []uint8
	}{
		{
			"crop to size (0, 0, center)",
			0, 0, CenterAnchor,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 0, 0),
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
			[]uint8{},
		},
		{
			"crop to size (3, 3, center)",
			3, 3, CenterAnchor,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
			[]uint8{
				0x06, 0x07, 0x08,
				0x0b, 0x0c, 0x0d,
				0x10, 0x11, 0x12,
			},
		},
		{
			"crop to size (3, 3, top-left)",
			3, 3, TopLeftAnchor,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
			[]uint8{
				0x00, 0x01, 0x02,
				0x05, 0x06, 0x07,
				0x0a, 0x0b, 0x0c,
			},
		},
		{
			"crop to size (3, 3, top)",
			3, 3, TopAnchor,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
			[]uint8{
				0x01, 0x02, 0x03,
				0x06, 0x07, 0x08,
				0x0b, 0x0c, 0x0d,
			},
		},
		{
			"crop to size (3, 3,, top-right)",
			3, 3, TopRightAnchor,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
			[]uint8{
				0x02, 0x03, 0x04,
				0x07, 0x08, 0x09,
				0x0c, 0x0d, 0x0e,
			},
		},
		{
			"crop to size (3, 3, left)",
			3, 3, LeftAnchor,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
			[]uint8{
				0x05, 0x06, 0x07,
				0x0a, 0x0b, 0x0c,
				0x0f, 0x10, 0x11,
			},
		},
		{
			"crop to size (3, 3, right)",
			3, 3, RightAnchor,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
			[]uint8{
				0x07, 0x08, 0x09,
				0x0c, 0x0d, 0x0e,
				0x11, 0x12, 0x13,
			},
		},
		{
			"crop to size (3, 3, bottom-left)",
			3, 3, BottomLeftAnchor,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
			[]uint8{
				0x0a, 0x0b, 0x0c,
				0x0f, 0x10, 0x11,
				0x14, 0x15, 0x16,
			},
		},
		{
			"crop to size (3, 3, bottom)",
			3, 3, BottomAnchor,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
			[]uint8{
				0x0b, 0x0c, 0x0d,
				0x10, 0x11, 0x12,
				0x15, 0x16, 0x17,
			},
		},
		{
			"crop to size (3, 3, bottom-right)",
			3, 3, BottomRightAnchor,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
			[]uint8{
				0x0c, 0x0d, 0x0e,
				0x11, 0x12, 0x13,
				0x16, 0x17, 0x18,
			},
		},
		{
			"crop to size (100, 100, center)",
			100, 100, CenterAnchor,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 5, 5),
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
			[]uint8{
				0x00, 0x01, 0x02, 0x03, 0x04,
				0x05, 0x06, 0x07, 0x08, 0x09,
				0x0a, 0x0b, 0x0c, 0x0d, 0x0e,
				0x0f, 0x10, 0x11, 0x12, 0x13,
				0x14, 0x15, 0x16, 0x17, 0x18,
			},
		},
	}

	for _, d := range testData {
		src := image.NewGray(d.srcb)
		src.Pix = d.srcPix

		f := CropToSize(d.w, d.h, d.anchor)
		dst := image.NewGray(f.Bounds(src.Bounds()))
		f.Draw(dst, src, nil)

		if !checkBoundsAndPix(dst.Bounds(), d.dstb, dst.Pix, d.dstPix) {
			t.Errorf("test [%s] failed: %#v, %#v", d.desc, dst.Bounds(), dst.Pix)
		}
	}
}

func TestRotate(t *testing.T) {
	testData := []struct {
		desc           string
		a              float32
		bg             color.Color
		interp         Interpolation
		srcb, dstb     image.Rectangle
		srcPix, dstPix []uint8
	}{
		{
			"rotate 0x0 90 white nearest",
			90, color.White, NearestNeighborInterpolation,
			image.Rect(0, 0, 0, 0),
			image.Rect(0, 0, 0, 0),
			[]uint8{},
			[]uint8{},
		},
		{
			"rotate 1x1 90 white nearest",
			90, color.White, NearestNeighborInterpolation,
			image.Rect(-1, -1, 0, 0),
			image.Rect(0, 0, 1, 1),
			[]uint8{0x80},
			[]uint8{0x80},
		},
		{
			"rotate 3x3 -90 white nearest",
			-90, color.White, NearestNeighborInterpolation,
			image.Rect(-1, -1, 2, 2),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x10, 0x20, 0x30,
				0x40, 0x50, 0x60,
				0x70, 0x80, 0x90,
			},
			[]uint8{
				0x70, 0x40, 0x10,
				0x80, 0x50, 0x20,
				0x90, 0x60, 0x30,
			},
		},
		{
			"rotate 3x3 -90 white linear",
			-90, color.White, LinearInterpolation,
			image.Rect(-1, -1, 2, 2),
			image.Rect(0, 0, 3, 3),
			[]uint8{
				0x10, 0x20, 0x30,
				0x40, 0x50, 0x60,
				0x70, 0x80, 0x90,
			},
			[]uint8{
				0x70, 0x40, 0x10,
				0x80, 0x50, 0x20,
				0x90, 0x60, 0x30,
			},
		},
		{
			"rotate 3x3 45 black nearest",
			45, color.Black, NearestNeighborInterpolation,
			image.Rect(-1, -1, 2, 2),
			image.Rect(0, 0, 5, 5),
			[]uint8{
				0x10, 0x20, 0x30,
				0x40, 0x50, 0x60,
				0x70, 0x80, 0x90,
			},
			[]uint8{
				0x00, 0x00, 0x30, 0x00, 0x00,
				0x00, 0x20, 0x30, 0x60, 0x00,
				0x10, 0x10, 0x50, 0x90, 0x90,
				0x00, 0x40, 0x70, 0x80, 0x00,
				0x00, 0x00, 0x70, 0x00, 0x00,
			},
		},
		{
			"rotate 5x5 45 black linear",
			45, color.Black, LinearInterpolation,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 8, 8),
			[]uint8{
				0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff,
			},
			[]uint8{
				0x00, 0x00, 0x00, 0x26, 0x26, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x2c, 0xe0, 0xe0, 0x2c, 0x00, 0x00,
				0x00, 0x2c, 0xe0, 0xff, 0xff, 0xe0, 0x2c, 0x00,
				0x26, 0xe0, 0xff, 0xff, 0xff, 0xff, 0xe0, 0x26,
				0x26, 0xe0, 0xff, 0xff, 0xff, 0xff, 0xe0, 0x26,
				0x00, 0x2c, 0xe0, 0xff, 0xff, 0xe0, 0x2c, 0x00,
				0x00, 0x00, 0x2c, 0xe0, 0xe0, 0x2c, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x26, 0x26, 0x00, 0x00, 0x00,
			},
		},
		{
			"rotate 5x5 45 black cubic",
			45, color.Black, CubicInterpolation,
			image.Rect(-1, -1, 4, 4),
			image.Rect(0, 0, 8, 8),
			[]uint8{
				0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff,
				0xff, 0xff, 0xff, 0xff, 0xff,
			},
			[]uint8{
				0x00, 0x00, 0x00, 0x23, 0x23, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x28, 0xf1, 0xf1, 0x28, 0x00, 0x00,
				0x00, 0x28, 0xe3, 0xff, 0xff, 0xe3, 0x28, 0x00,
				0x23, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xf1, 0x23,
				0x23, 0xf1, 0xff, 0xff, 0xff, 0xff, 0xf1, 0x23,
				0x00, 0x28, 0xe3, 0xff, 0xff, 0xe3, 0x28, 0x00,
				0x00, 0x00, 0x28, 0xf1, 0xf1, 0x28, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x23, 0x23, 0x00, 0x00, 0x00,
			},
		},
	}

	for _, d := range testData {
		src := image.NewGray(d.srcb)
		src.Pix = d.srcPix

		f := Rotate(d.a, d.bg, d.interp)
		dst := image.NewGray(f.Bounds(src.Bounds()))
		f.Draw(dst, src, nil)

		if !checkBoundsAndPix(dst.Bounds(), d.dstb, dst.Pix, d.dstPix) {
			t.Errorf("test [%s] failed: %#v, %#v", d.desc, dst.Bounds(), dst.Pix)
		}
	}

}
