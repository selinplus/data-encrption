package api

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/selinplus/data-encrption/pkg/app"
	"github.com/selinplus/data-encrption/pkg/e"
	"github.com/selinplus/data-encrption/pkg/logging"
	"github.com/selinplus/data-encrption/pkg/upload"
)

func UploadFile(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, err := c.Request.FormFile("file")
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	//savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_FILE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_FILE_FAIL, nil)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_FILE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"name": image.Filename,
		"url":  upload.GetImageFullUrl(imageName),
	})
}
