package util

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

func OptimizeImage(file *os.File, size int64) {
	i, err := ioutil.ReadAll(file)
	check(err)

	resetFile(file)

	img, err := jpeg.Decode(bytes.NewReader(i))
	check(err)

	file.Close()

	out := new(bytes.Buffer)

	err = mozjpegbin.Encode(out, img, &mozjpegbin.Options{
		Quality: 75,
		Optimize: true,
	})
	check(err)

	path := fmt.Sprintf("squished/%s", file.Name())

	newFile, err := os.Create(path)
	check(err)

	io.Copy(newFile, out)
	newFile.Close()
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