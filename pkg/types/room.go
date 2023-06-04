package types

type Room struct {
	RoomID  int    `gorm:"primaryKey;autoIncrement"`
	RoomLoc string `gorm:"column:room_loc"`
	RoomCap string `gorm:"column:room_cap"`
	ClassID int    `gorm:"column:class_id"`
}
