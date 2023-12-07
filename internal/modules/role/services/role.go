package services

import (
	"context"
	"fmt"
	"microrole/internal/modules/role"
	"microrole/internal/modules/role/models/dto"
	"microrole/internal/modules/role/models/models"
	"unicode"
)

type IRoleRepository interface {
	CreateRole(ctx context.Context, dto dto.CreateRoleDTO) (*models.Role, error)
	DeleteRole(ctx context.Context, dto dto.DeleteRoleDTO) error
	UpdateRole(ctx context.Context, dto dto.UpdateRoleDto) (*models.Role, error)
}

type RoleService struct {
	repo IRoleRepository
}

func New(repo IRoleRepository) *RoleService {
	return &RoleService{
		repo: repo,
	}
}

func (rs *RoleService) AddUserRole(userId string, dto dto.CreateRoleDTO) error {
	if !validateRolePrefix(dto.Prefix) {
		return fmt.Errorf("AddUserRole-error %s", role.ErrPrefixNotCorrect)
	}
}

func validateRolePrefix(prefix string) bool {
	for _, char := range prefix {
		switch {
		case unicode.IsNumber(char):
			return false
		case unicode.IsPunct(char):
			return false
		}
	}

	return true
}
