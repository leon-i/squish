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

func OptimizeImage(file *os.File, imageType string) {
	i, err := ioutil.ReadAll(file)
	check(err)

	resetFile(file)

	in := bytes.NewReader(i)

	var img image.Image

	if imageType == "image/jpeg" {
		img, err = jpeg.Decode(in)
		check(err)
	} else {
		img, err = png.Decode(in)
		check(err)
	}

	file.Close()

	out := new(bytes.Buffer)

	err = mozjpegbin.Encode(out, img, &mozjpegbin.Options{
		Quality: 75,
		Optimize: true,
	})
	check(err)

	path := fmt.Sprintf("squished/%s.jpg", Trim(file.Name()))

	newFile, err := os.Create(path)
	check(err)

	saved := (in.Size() - int64(out.Len())) * 100 / in.Size()
	io.Copy(newFile, out)
	newFile.Close()

	fmt.Println(fmt.Sprintf("%s squished - file size reduced by %d%%", file.Name(), saved))
}

func Cleanup() {
	squished, err := ioutil.ReadDir("./squished")
	check(err)

	if len(squished) == 0 {
		os.RemoveAll("squished")
	}
}

func Trim(fileName string) (trimmed string) {
	t := fileName[:len(fileName) - len(filepath.Ext(fileName))]
	return t
}

func resetFile(file *os.File) {
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}