package controller

import (
	"hot-chat/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AttachController struct{}

func NewAttachController() *AttachController {
	return &AttachController{}
}

type UploadResult struct {
	Path   string `json:"path"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

func (c *AttachController) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	// 按照时间创建文件夹
	now := time.Now()
	dstDir := filepath.Join("./static", strconv.Itoa(now.Year()), strconv.Itoa(int(now.Month())), strconv.Itoa(now.Day()))
	err = os.MkdirAll(dstDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		utils.ResFailed(ctx, err)
		return
	}
	// 保存原始文件
	dstFile := filepath.Join(dstDir, file.Filename)
	err = ctx.SaveUploadedFile(file, dstFile)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	// 解析图片宽高信息
	fileBuf, err := ioutil.ReadFile(dstFile)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	imgConfig, err := utils.DecodeImageConfig(fileBuf)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}
	// 转为 webp
	webpPath, err := utils.ConvertToWebp(dstFile)
	if err != nil {
		utils.ResFailed(ctx, err)
		return
	}

	utils.ResOk(ctx, UploadResult{
		Path:   webpPath,
		Width:  imgConfig.Width,
		Height: imgConfig.Height,
	})
}

func (c *AttachController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/attach", c.Upload)
}
