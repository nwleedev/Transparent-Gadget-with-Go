package pkg

import (
	"bytes"
	"image/png"

	"github.com/lxn/walk"
)

func LoadImage(buf []byte, _err error) *walk.Bitmap {
	HandleError(_err)
	imageDecoded, err := png.Decode(bytes.NewReader(buf))
	HandleError(err)
	imageBitmap, err := walk.NewBitmapFromImageForDPI(imageDecoded, 96)
	HandleError(err)
	return imageBitmap
}
