package restaurantbiz

import (
	"context"
	"errors"
	"testing"

	"food_delivery/common"

	restaurantmodel "food_delivery/modules/restaurant/model"
)

type mokeCreateStore struct{}

func (mokeCreateStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "Phuc" {
		return common.ErrDB(errors.New("test"))
	}

	data.Id = 200

	return nil
}

func TestNewCreateRestaurantBiz(t *testing.T) {
	biz := NewCreateRestaurantBiz(mokeCreateStore{})

	dataTest := restaurantmodel.RestaurantCreate{Name: ""}
	err := biz.CreateRestaurant(context.Background(), &dataTest)

	if err != nil && err.Error() != "test" {
		t.Errorf("failed")
	}

	dataTest = restaurantmodel.RestaurantCreate{Name: "Phuc"}
	err = biz.CreateRestaurant(context.Background(), &dataTest)

	if err == nil && err.Error() != "test" {
		t.Errorf("failed")
	}

	dataTest = restaurantmodel.RestaurantCreate{Name: "Phuz"}
	err = biz.CreateRestaurant(context.Background(), &dataTest)
	if err != nil {
		t.Errorf("failed")
	}

	t.Log("TestNewCreateRestaurantBiz")
}
