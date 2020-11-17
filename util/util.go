package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"squish/config"
	"time"
)

func GetFileContentType(file *os.File) (string, error) {
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	ResetFile(file)

	return contentType, nil
}

func Startup() {
	d := config.SquishConfig.Destination
	if _, err := os.Stat("./" + d); os.IsNotExist(err) {
		if err := os.Mkdir(d, 0755); err != nil {
			panic(err)
		}
	}
}

func Cleanup() {
	d := config.SquishConfig.Destination
	squished, err := ioutil.ReadDir("./" + d)
	if err != nil {
		panic(err)
	}

	if len(squished) == 0 {
		if err := os.RemoveAll(d); err != nil {
			panic(err)
		}
	}
}

func Trim(fileName string) (trimmed string) {
	t := fileName[:len(fileName) - len(filepath.Ext(fileName))]
	return t
}

func ResetFile(file *os.File) {
	if _, err := file.Seek(0, 0); err != nil {
		panic(err)
	}
}

func LogDuration(start time.Time) {
	duration := time.Since(start).Seconds()
	fmt.Println(fmt.Sprintf("\nall image(s) squished in %f seconds\n", duration))
}