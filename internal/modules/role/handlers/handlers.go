package handlers

import "microrole/internal/modules/role/models/dto"

type IRoleService interface {
	AddUserRole(userId string, dto dto.CreateRoleDTO) error
	RemoveUserRole(roleId string) error
	UpdateUserRole(roleId string, dto dto.CreateRoleDTO) error
}
