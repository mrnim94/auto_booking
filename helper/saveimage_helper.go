package helper

import (
	"bytes"
	"image"
	"auto-booking/log"
	"image/png"
	"os"
)

func HelpSaveImage(photo []byte, name string) error {
	img, _, err := image.Decode(bytes.NewReader(photo))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	out, err := os.Create("./log_files/screenshots/" + name + ".png")
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
