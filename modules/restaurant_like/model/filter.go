package restaurant_like_model

type Filter struct {
	RestaurantId int `json:"-" form:"restaurant_id"`
	UserId       int `json:"-" form:"user_id"`
}
