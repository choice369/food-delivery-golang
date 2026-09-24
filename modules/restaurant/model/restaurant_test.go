package restaurantmodel

import "testing"

type testData struct {
	Input  RestaurantCreate
	Expect error
}

func TestRestaurantCreate_Validate(t *testing.T) {
	dataTable := []testData{
		{
			Input: RestaurantCreate{Name: ""}, Expect: ErrNameIsEmpty,
		},
		{
			Input: RestaurantCreate{Name: "test"}, Expect: nil,
		},
	}

	for _, item := range dataTable {
		err := item.Input.Validate()
		if err != item.Expect {
			t.Errorf("Validate err expected %v, got %v", item.Expect, err)
		}
	}
}
