package utils

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"os/exec"
)

func GetFileMimeType(imgBuf []byte) (string, error) {
	// 前 512 字节就可以
	buf := imgBuf[0:512]
	mimeType := http.DetectContentType(buf)
	return mimeType, nil
}

func ConvertToWebp(imgPath string) (webpPath string, err error) {
	imgBuf, err := os.ReadFile(imgPath)
	if err != nil {
		return
	}
	mimeType, err := GetFileMimeType(imgBuf)
	if err != nil {
		return
	}
	webpLib := "cwebp"
	if mimeType == "image/gif" {
		webpLib = "gif2webp"
	}
	webpPath = strings.Replace(imgPath, filepath.Ext(imgPath), ".webp", 1)
	cmd := exec.Command(webpLib, imgPath, "-o", webpPath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return
	}
	return
}

func DecodeImage(imgBuf []byte) (img image.Image, err error) {
	imgReader := bytes.NewReader(imgBuf)
	mimeType, err := GetFileMimeType(imgBuf)
	if err != nil {
		return nil, err
	}
	switch mimeType {
	case "image/jpeg":
		return jpeg.Decode(imgReader)
	case "image/png":
		return png.Decode(imgReader)
	case "image/gif":
		return gif.Decode(imgReader)
	}
	return nil, errors.New("不支持的图片格式")
}

func DecodeImageConfig(imgBuf []byte) (config image.Config, err error) {
	imgReader := bytes.NewReader(imgBuf)
	mimeType, err := GetFileMimeType(imgBuf)
	if err != nil {
		return
	}
	switch mimeType {
	case "image/jpeg":
		return jpeg.DecodeConfig(imgReader)
	case "image/png":
		return png.DecodeConfig(imgReader)
	case "image/gif":
		return gif.DecodeConfig(imgReader)
	}
	err = errors.New("不支持的图片格式")
	return
}
