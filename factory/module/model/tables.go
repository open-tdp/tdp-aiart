package model

// 图片

type Artwork struct {
	Id             uint `gorm:"primaryKey"`
	UserId         uint `gorm:"index"`
	Username       string
	Subject        string
	Prompt         string
	NegativePrompt string
	Styles         string
	Strength       float64
	InputFile      string
	OutputFile     string
	Favorites      uint
	Recommends     uint
	Status         string
	CreatedAt      int64
	UpdatedAt      int64
}

// 系统配置

type Config struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"uniqueIndex:idx_config"`
	Value       string
	Type        string `gorm:"default:string"`
	Module      string `gorm:"uniqueIndex:idx_config"`
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// 迁移记录

type Migration struct {
	Id          uint   `gorm:"primaryKey"`
	Version     string `gorm:"uniqueIndex"`
	Description string
	CreatedAt   int64
	UpdatedAt   int64
}

// 用户

type User struct {
	Id           uint   `gorm:"primaryKey"`
	Username     string `gorm:"uniqueIndex"`
	Password     string `json:"-"`
	Level        uint   `gorm:"default:5"`
	AppId        string `gorm:"uniqueIndex"`
	AppKey       string `json:"-"`
	Email        string `gorm:"uniqueIndex,default:null"`
	Description  string `gorm:"default:挥一挥手"`
	ArtworkQuota uint   `gorm:"default:10"`
	CreatedAt    int64
	UpdatedAt    int64
}