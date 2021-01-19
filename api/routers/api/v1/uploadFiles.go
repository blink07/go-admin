package v1

import (
	"github.com/gin-gonic/gin"
	"go-admin/api/service"
	"go-admin/api/utils/app"
	"go-admin/api/utils/e"
	"net/http"
)


// 图片统一上传
func ImageUpload(c *gin.Context) {
	appG := app.Gin{c}
	data := make(map[string]string)
	file, image, err := c.Request.FormFile("image")

	if err != nil {
		appG.Response(http.StatusOK, e.FILE_PARAM_GET_FAIL, err.Error())
		return
	}

	if image == nil {
		appG.Response(http.StatusOK, e.FILE_IS_EMPTY, nil)
		return
	}

	imageName := service.GetImageName(image.Filename)
	savePath := service.GetImagePath()
	src := savePath + imageName

	if !service.CheckImageExt(imageName) || !service.CheckImageSize(file) {
		appG.Response(http.StatusOK, e.FILE_NOT_STANDARD, nil)
		return
	}

	err = service.CheckImage(savePath)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, err.Error())
	}

	if err := c.SaveUploadedFile(image, src); err!=nil {
		appG.Response(http.StatusOK, e.FILE_UPLOAD_FAIL, err.Error())
	}
	data["image_url"] = service.FullImagePath(imageName)
	data["image_save_url"] =savePath + imageName

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

