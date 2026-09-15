package restaurantmodel

import (
	"errors"
	"strings"

	"food_delivery/common"
)

const EntityName = "restaurant"

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Address         string         `json:"address" gorm:"column:address"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name    *string        `json:"name" gorm:"column:name"`
	Address *string        `json:"address" gorm:"column:address"`
	Logo    *common.Image  `json:"logo" gorm:"column:logo"`
	Cover   *common.Images `json:"cover" gorm:"column:cover"`
}

func (RestaurantUpdate) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	OwnerId         int            `json:"owner_id" gorm:"column:owner_id"`
	Name            string         `json:"name" gorm:"column:name"`
	Address         string         `json:"address" gorm:"column:address"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Cover           *common.Images `json:"cover" gorm:"column:cover"`
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

var ErrNameIsEmpty = errors.New("name is empty")
