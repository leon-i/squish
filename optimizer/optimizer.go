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
	util.Check(err)

	util.ResetFile(file)

	in := bytes.NewReader(i)

	var img image.Image

	if imageType == "image/jpeg" {
		img, err = jpeg.Decode(in)
		util.Check(err)
	} else {
		img, err = png.Decode(in)
		util.Check(err)
	}

	err = file.Close()
	util.Check(err)

	out := new(bytes.Buffer)

	err = mozjpegbin.Encode(out, img, &mozjpegbin.Options{
		Quality: conf.Quality,
		Optimize: true,
	})
	util.Check(err)

	path := fmt.Sprintf("%s/%s.jpg", conf.Destination, util.Trim(file.Name()))

	newFile, err := os.Create(path)
	util.Check(err)

	saved := (in.Size() - int64(out.Len())) * 100 / in.Size()

	_, err = io.Copy(newFile, out)
	util.Check(err)

	err = newFile.Close()
	util.Check(err)

	fmt.Println(fmt.Sprintf("%s squished - file size reduced by %d%%", file.Name(), saved))
}