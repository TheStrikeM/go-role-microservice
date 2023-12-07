package models

import "microrole/internal/modules/role/models"

type Role struct {
	ID          string             `json:"id"`
	UserID      string             `json:"user_id"`
	Prefix      string             `json:"prefix"`
	Permissions models.Permissions `json:"permissions"`
}
