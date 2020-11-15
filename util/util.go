package util

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"github.com/nickalie/go-mozjpegbin"
)

func GetFileContentType(file *os.File) (string, error) {
	// only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	resetFile(file)

	return contentType, nil
}

func OptimizeImage(file *os.File, imageType string, quality uint) {
	i, err := ioutil.ReadAll(file)
	Check(err)

	resetFile(file)

	in := bytes.NewReader(i)

	var img image.Image

	if imageType == "image/jpeg" {
		img, err = jpeg.Decode(in)
		Check(err)
	} else {
		img, err = png.Decode(in)
		Check(err)
	}

	err = file.Close()
	Check(err)

	out := new(bytes.Buffer)

	err = mozjpegbin.Encode(out, img, &mozjpegbin.Options{
		Quality: quality,
		Optimize: true,
	})
	Check(err)

	path := fmt.Sprintf("squished/%s.jpg", Trim(file.Name()))

	newFile, err := os.Create(path)
	Check(err)

	saved := (in.Size() - int64(out.Len())) * 100 / in.Size()

	_, err = io.Copy(newFile, out)
	Check(err)

	err = newFile.Close()
	Check(err)

	fmt.Println(fmt.Sprintf("%s squished - file size reduced by %d%%", file.Name(), saved))
}

func Startup() {
	if _, err := os.Stat("./squished"); os.IsNotExist(err) {
		if err := os.Mkdir("squished", 0755); err != nil {
			panic(err)
		}
	}
}

func Cleanup() {
	squished, err := ioutil.ReadDir("./squished")
	Check(err)

	if len(squished) == 0 {
		err = os.RemoveAll("squished")
		Check(err)
	}
}

func IsValidQuality(q uint) (valid bool) {
	if q < 0 || q > 100 {
		return false
	}

	return true
}

func Trim(fileName string) (trimmed string) {
	t := fileName[:len(fileName) - len(filepath.Ext(fileName))]
	return t
}

func resetFile(file *os.File) {
	if _, err := file.Seek(0, 0); err != nil {
		panic(err)
	}
}

func LogError(e error) {
	fmt.Println("ERROR - " + e.Error())
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}