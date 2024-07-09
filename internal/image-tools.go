package imageTools

import (
	"github.com/h2non/bimg"
)

func Resize(buffer []byte, width int, height int) (resizedBuffer []byte, err error) {
	resizedBuffer, err = bimg.NewImage(buffer).Resize(width, height)
	if err != nil {
		return nil, err
	}

	// TODO: Check the new image? `bimg.NewImage(resizedBuffer).Size() === {Width, height}`

	return resizedBuffer, nil
}

// TODO: ForcedResize (breaks aspect ratio?)

// TODO: Crop
