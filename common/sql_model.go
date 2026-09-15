package common

import "time"

type SQLModel struct {
	Id        int       `json:"id" gorm:"column:id"`
	OwnerId   int       `json:"owner_id" gorm:"column:owner_id"`
	Status    int       `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
	FakeId    *UID      `json:"fake_id gorm:"column:"-"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeId = uid
}
