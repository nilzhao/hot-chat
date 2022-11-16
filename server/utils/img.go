package utils

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func GetFileMimeType(imgBuf []byte) (string, error) {
	// 前 512 字节就可以
	buf := imgBuf[0:512]
	mimeType := http.DetectContentType(buf)
	return mimeType, nil
}

func ConvertToWebp(imgPath string) (string, error) {
	file, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return "", err
	}
	img, err := DecodeImage(file)
	if err != nil {
		return "", err
	}
	newImgPath := strings.Replace(imgPath, filepath.Ext(imgPath), ".webp", 1)
	output, err := os.Create(newImgPath)
	if err != nil {
		return "", err
	}
	defer output.Close()
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		return "", err
	}

	if err := webp.Encode(output, img, options); err != nil {
		return "", err
	}
	return newImgPath, err
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
