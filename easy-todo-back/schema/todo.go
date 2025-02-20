package schema

import (
	"crypto/rand"
	"encoding/binary"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	Id        int64 `gorm:"primaryKey"`
	Value     string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	IsDone    bool
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}

func GetNewID() int64 {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}

	return int64(binary.BigEndian.Uint64(b[:])) & 0x7FFFFFFFFFFFFFFF
}
