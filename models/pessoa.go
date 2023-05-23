package models

import "gorm.io/gorm"

type Pessoa struct {
	gorm.Model `swaggerignore:"true"`
	Nome       string `json:"nome"`
	Idade      int    `json:"idade"`
}
