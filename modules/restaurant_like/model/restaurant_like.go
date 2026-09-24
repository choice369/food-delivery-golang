package restaurant_like_model

import (
	"fmt"
	"time"

	"food_delivery/common"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `gorm:"column:restaurant_id" json:"restaurant_id"`
	UserId       int                `gorm:"column:user_id" json:"user_id"`
	CreatedAt    time.Time          `gorm:"column:created_at" json:"created_at"`
	User         *common.SimpleUser `gorm:"preload:false;" json:"user"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("cannot like restaurant"),
		fmt.Sprintf("ErrCannotLikeRestaurant"),
	)
}

func ErrCannotUnLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("cannot unlike restaurant"),
		fmt.Sprintf("ErrCannotUnLikeRestaurant"),
	)
}
