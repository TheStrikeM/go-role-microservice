package services

import (
	"context"
	"microrole/internal/modules/role/models/dto"
	"microrole/internal/modules/role/models/models"
)

type IRoleRepository interface {
	CreateRole(ctx context.Context, dto dto.CreateRoleDTO) (*models.Role, error)
	DeleteRole(ctx context.Context, dto dto.DeleteRoleDTO) error
	UpdateRole(ctx context.Context, dto dto.UpdateRoleDto) (*models.Role, error)
}
