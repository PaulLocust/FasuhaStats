package models

import (
	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model        // Включает ID, CreatedAt, UpdatedAt, DeletedAt
	ReviewerID uint   `gorm:"not null"`  // ID игрока, который оставил отзыв (ссылка на Player)
	TargetID   uint   `gorm:"not null"`  // ID игрока, которому оставлен отзыв (ссылка на Player)
	MatchID    uint   `gorm:"not null"`  // ID матча, связанного с отзывом (ссылка на Match)
	Score      int    `gorm:"not null"`  // Оценка (1-5)
	Comment    string `gorm:"type:text"` // Комментарий (опционально)

	// Связи
	Reviewer Player `gorm:"foreignKey:ReviewerID"` // Связь с игроком, который оставил отзыв
	Target   Player `gorm:"foreignKey:TargetID"`   // Связь с игроком, которому оставили отзыв
	Match    Match  `gorm:"foreignKey:MatchID"`    // Связь с матчем
}
