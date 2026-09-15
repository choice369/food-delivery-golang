package uploadmodel

import (
	"errors"

	"food_delivery/common"
)

const EntityName = "Upload"

type Upload struct {
	common.SQLModel `bson:",inline"`
	common.Image    `bson:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("file too large"),
		"file too large",
		"ErrFileTooLarge",
	)

	ErrFileIsNotImage = common.NewCustomError(
		errors.New("file is not image"),
		"file is not image",
		"ErrFileIsNotImage",
	)

	ErrFileSaveFailed = common.NewCustomError(
		errors.New("file save failed"),
		"file save failed",
		"ErrFileSaveFailed",
	)
)
