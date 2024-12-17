package models

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model        // Включает ID, CreatedAt, UpdatedAt, DeletedAt
	FaceitID   string `gorm:"unique;not null"` // Уникальный ID от Faceit API
	Nickname   string `gorm:"unique;not null"` // Никнейм игрока
	Avatar     string `gorm:"default:null"`    // URL аватара
	Country    string `gorm:"default:null"`    // Страна игрока
	SkillLevel int    `gorm:"default:0"`       // Уровень навыка (1-10)
}
