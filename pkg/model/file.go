package model

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model

	Name string `gorm:"index;unique"`

	Url string

	Cnt int
}

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
func CreateFile(db *gorm.DB, url string) (string, error) {
	name := stringWithCharset(16, charset)
	err := db.Create(&File{Name: name, Url: url}).Error

	return name, err
}

func RecordCnt(db *gorm.DB, name string) error {
	result := db.Model(&File{}).Where("name = ?", name).Update("cnt", gorm.Expr("cnt + 1"))

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetFile(db *gorm.DB, name string) (*File, error) {
	var file File
	result := db.Where("name = ?", name).First(&file)

	if result.Error != nil {
		return nil, result.Error
	}

	return &file, nil
}
