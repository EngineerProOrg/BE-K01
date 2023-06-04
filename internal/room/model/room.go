package roommodel


// create table ROOM (
//     ROOM_ID int primary key,
//     ROOM_LOC varchar(50),
//     ROOM_CAP varchar(50),
//     CLASS_ID int
// )
type Room struct {
	Room_id int `json:"room_id" gorm:"column:ROOM_ID;primaryKey"`
	Room_loc string `json:"room_loc" gorm:"column:ROOM_LOC"`
	Room_cap string `json:"room_cap" gorm:"column:ROOM_CAP"`
	Class_id int `json:"class_id" gorm:"column:CLASS_ID"`
}

func (Room) TableName() string {
	return "ROOM"
}