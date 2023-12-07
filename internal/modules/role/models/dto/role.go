package dto

import "microrole/internal/modules/role/models"

type CreateRoleDTO struct {
	UserId      string             `json:"user_id"`
	Prefix      string             `json:"prefix"`
	Permissions models.Permissions `json:"permissions"`
}

type DeleteRoleDTO struct {
	ID string `json:"id"`
}

type UpdateRoleDto struct {
	ID          string             `json:"id"`
	Prefix      string             `json:"prefix"`
	Permissions models.Permissions `json:"permissions"`
}
