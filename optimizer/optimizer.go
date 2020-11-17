package optimizer

import (
	"bytes"
	"fmt"
	"github.com/nickalie/go-mozjpegbin"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"squish/config"
	"squish/util"
)

func ToMozillaJpeg(file *os.File, imageType string) {
	conf := config.SquishConfig

	i, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	util.ResetFile(file)

	in := bytes.NewReader(i)

	var img image.Image

	if imageType == "image/jpeg" {
		img, err = jpeg.Decode(in)
		if err != nil {
			panic(fmt.Errorf("ERROR - cannot decode image - %v", err))
		}
	} else {
		img, err = png.Decode(in)
		if err != nil {
			panic(fmt.Errorf("ERROR - cannot decode image - %v", err))
		}
	}

	if err := file.Close(); err != nil {
		panic(err)
	}

	out := new(bytes.Buffer)

	if err := mozjpegbin.Encode(out, img, &mozjpegbin.Options{
		Quality: conf.Quality,
		Optimize: true,
	}); err != nil {
		panic(fmt.Errorf("ERROR - cannot encode image, check mozjpeg config - %v", err))
	}

	path := fmt.Sprintf("%s/%s.jpg", conf.Destination, util.Trim(file.Name()))

	newFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	saved := (in.Size() - int64(out.Len())) * 100 / in.Size()

	if _, err = io.Copy(newFile, out); err != nil {
		panic(err)
	}

	if err := newFile.Close(); err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%s squished - file size reduced by %d%%", file.Name(), saved))
}