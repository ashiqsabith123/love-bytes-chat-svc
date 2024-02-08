package domain

import "time"

type ChatRooom struct {
	RoomID          uint `bson:"-"`
	User1ID         uint
	User2ID         uint
	LastMessageID   uint
	LastMessageTime time.Time
}
