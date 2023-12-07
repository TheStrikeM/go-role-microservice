package repositories

import (
	"context"
	"log/slog"
	"microrole/internal/modules/role/models/dto"
	"microrole/internal/modules/role/models/models"
	"microrole/internal/storage/clients/postgresql"
)

type RoleRepository struct {
	client *postgresql.PSQLClient
	log    *slog.Logger
}

func New(client *postgresql.PSQLClient, log *slog.Logger) *RoleRepository {
	return &RoleRepository{
		client: client,
		log:    log,
	}
}

func (rr *RoleRepository) CreateRole(ctx context.Context, dto dto.CreateRoleDTO) (*models.Role, error) {
	query := `
			INSERT INTO public.role
				(userid, prefix, createnews, createposts, createevents)
			VALUES 
				($1, $2, $3, $4, $5)
			RETURNING (id, userid, prefix, createnews, createposts, createevents)
	`

	var role models.Role
	if err := rr.client.DbPool.QueryRow(
		ctx, query,
		dto.UserId,
		dto.Prefix,
		dto.Permissions.CreateNews,
		dto.Permissions.CreatePosts,
		dto.Permissions.CreateEvents,
	).Scan(&role); err != nil {
		// TODO: Разобраться с ошибками в postgres
		return nil, err
	}
	return &role, nil
}

func (rr *RoleRepository) DeleteRole(ctx context.Context, dto dto.DeleteRoleDTO) error {
	query := `
		DELETE FROM public.role
		WHERE id=$1
	`
	if err := rr.client.DbPool.QueryRow(ctx, query, dto.ID).Scan(); err != nil {
		return err
	}
	return nil
}

func (rr *RoleRepository) UpdateRole(ctx context.Context, dto dto.UpdateRoleDto) (*models.Role, error) {
	query := `
		UPDATE public.role
		SET prefix=$1, createnews=$1, createposts=$2, createevents=$3
		WHERE id=$4
		RETURNING (id, userid, prefix, createnews, createposts, createevents)
	`

	var role models.Role
	if err := rr.client.DbPool.QueryRow(ctx, query,
		dto.Prefix,
		dto.Permissions.CreateNews,
		dto.Permissions.CreatePosts,
		dto.Permissions.CreateEvents,
	).Scan(&role); err != nil {
		return nil, err
	}

	return &role, nil
}
