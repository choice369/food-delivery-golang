package ginupload

import (
	"net/http"

	"food_delivery/common"
	"food_delivery/component/appctx"
	uploadbiz "food_delivery/modules/upload/biz"

	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
		//	panic(err)
		//}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err = file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(img))
		//c.JSON(http.StatusOK, common.SimpleSuccessRes(common.Image{
		//	Id:        0,
		//	Url:       "http://localhost:8080/static/" + fileHeader.Filename,
		//	Width:     0,
		//	Height:    0,
		//	CloudName: "Local",
		//	Extension: "jpeg",
		//}))
	}
}
