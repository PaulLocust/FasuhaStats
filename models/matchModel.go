package models

import "gorm.io/gorm"

type Match struct {
	gorm.Model        // Включает ID, CreatedAt, UpdatedAt, DeletedAt
	MatchID    string `gorm:"unique;not null"` // Уникальный ID матча от Faceit API
	Game       string `gorm:"not null"`        // Название игры (например, CS:GO)

	Players []Player `gorm:"many2many:match_players;"` // Связь многие-ко-многим
}
