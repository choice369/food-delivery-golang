package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"

	uploadmodel "food_delivery/modules/upload/model"

	"food_delivery/component/upload_provider"

	"food_delivery/common"
)

type CreateImageStore interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider upload_provider.UploadProvider
	imgStore CreateImageStore
}

func NewUploadBiz(provider upload_provider.UploadProvider, imgStore CreateImageStore) *uploadBiz {
	return &uploadBiz{provider, imgStore}
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}

func (u *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)
	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage()
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName) // img.ipg => .jpg
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := u.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s%s", folder, fileName))
	if err != nil {
		return nil, uploadmodel.ErrFileSaveFailed(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}
