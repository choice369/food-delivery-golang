package common

type SimpleUser struct {
	SQLModel  `bson:",inline"`
	LastName  string  `bson:"last_name" json:"last_name" bson:"last_name"`
	FirstName string  `bson:"first_name" json:"first_name" bson:"first_name"`
	Role      string  `bson:"role" json:"role" bson:"role"`
	Avatar    *string `bson:"avatar" json:"avatar" bson:"avatar"`
}

func (u *SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}
