// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: positions.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createAdminPosition = `-- name: CreateAdminPosition :one
INSERT INTO positions(id, title, can_activate, can_deactivate, can_add_orders, can_receive_orders, can_query, can_change_protocol, can_add_staff, can_add_reminders, is_admin)
VALUES(gen_random_uuid(), 'admin', true, true, true, true, true, true, true, true, true)
RETURNING id, title, can_activate, can_deactivate, can_add_orders, can_receive_orders, can_query, can_change_protocol, can_add_staff, can_add_reminders, is_admin
`

func (q *Queries) CreateAdminPosition(ctx context.Context) (Position, error) {
	row := q.db.QueryRowContext(ctx, createAdminPosition)
	var i Position
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.CanActivate,
		&i.CanDeactivate,
		&i.CanAddOrders,
		&i.CanReceiveOrders,
		&i.CanQuery,
		&i.CanChangeProtocol,
		&i.CanAddStaff,
		&i.CanAddReminders,
		&i.IsAdmin,
	)
	return i, err
}

const createPosition = `-- name: CreatePosition :one
INSERT INTO positions(id, title, can_activate, can_deactivate, can_add_orders, can_receive_orders, can_query, can_change_protocol, can_add_staff, can_add_reminders)
VALUES(gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, title, can_activate, can_deactivate, can_add_orders, can_receive_orders, can_query, can_change_protocol, can_add_staff, can_add_reminders, is_admin
`

type CreatePositionParams struct {
	Title             string
	CanActivate       bool
	CanDeactivate     bool
	CanAddOrders      bool
	CanReceiveOrders  bool
	CanQuery          bool
	CanChangeProtocol bool
	CanAddStaff       bool
	CanAddReminders   bool
}

func (q *Queries) CreatePosition(ctx context.Context, arg CreatePositionParams) (Position, error) {
	row := q.db.QueryRowContext(ctx, createPosition,
		arg.Title,
		arg.CanActivate,
		arg.CanDeactivate,
		arg.CanAddOrders,
		arg.CanReceiveOrders,
		arg.CanQuery,
		arg.CanChangeProtocol,
		arg.CanAddStaff,
		arg.CanAddReminders,
	)
	var i Position
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.CanActivate,
		&i.CanDeactivate,
		&i.CanAddOrders,
		&i.CanReceiveOrders,
		&i.CanQuery,
		&i.CanChangeProtocol,
		&i.CanAddStaff,
		&i.CanAddReminders,
		&i.IsAdmin,
	)
	return i, err
}

const getPositionByTitle = `-- name: GetPositionByTitle :one
SELECT id, title, can_activate, can_deactivate, can_add_orders, can_receive_orders, can_query, can_change_protocol, can_add_staff, can_add_reminders, is_admin FROM positions
WHERE $1 = title
`

func (q *Queries) GetPositionByTitle(ctx context.Context, title string) (Position, error) {
	row := q.db.QueryRowContext(ctx, getPositionByTitle, title)
	var i Position
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.CanActivate,
		&i.CanDeactivate,
		&i.CanAddOrders,
		&i.CanReceiveOrders,
		&i.CanQuery,
		&i.CanChangeProtocol,
		&i.CanAddStaff,
		&i.CanAddReminders,
		&i.IsAdmin,
	)
	return i, err
}

const getPositions = `-- name: GetPositions :many
SELECT id, title, can_activate, can_deactivate, can_add_orders, can_receive_orders, can_query, can_change_protocol, can_add_staff, can_add_reminders, is_admin FROM positions
`

func (q *Queries) GetPositions(ctx context.Context) ([]Position, error) {
	rows, err := q.db.QueryContext(ctx, getPositions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Position
	for rows.Next() {
		var i Position
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.CanActivate,
			&i.CanDeactivate,
			&i.CanAddOrders,
			&i.CanReceiveOrders,
			&i.CanQuery,
			&i.CanChangeProtocol,
			&i.CanAddStaff,
			&i.CanAddReminders,
			&i.IsAdmin,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserPosition = `-- name: GetUserPosition :one
SELECT id, title, can_activate, can_deactivate, can_add_orders, can_receive_orders, can_query, can_change_protocol, can_add_staff, can_add_reminders, is_admin FROM positions
WHERE $1 = id
`

func (q *Queries) GetUserPosition(ctx context.Context, id uuid.UUID) (Position, error) {
	row := q.db.QueryRowContext(ctx, getUserPosition, id)
	var i Position
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.CanActivate,
		&i.CanDeactivate,
		&i.CanAddOrders,
		&i.CanReceiveOrders,
		&i.CanQuery,
		&i.CanChangeProtocol,
		&i.CanAddStaff,
		&i.CanAddReminders,
		&i.IsAdmin,
	)
	return i, err
}

const updatePosition = `-- name: UpdatePosition :exec
UPDATE positions
SET title = $1,
    can_activate = $2,
    can_deactivate = $3,
    can_add_orders = $4,
    can_receive_orders = $5,
    can_query = $6,
    can_change_protocol = $7,
    can_add_staff = $8,
    can_add_reminders = $9
WHERE $10 = id
`

type UpdatePositionParams struct {
	Title             string
	CanActivate       bool
	CanDeactivate     bool
	CanAddOrders      bool
	CanReceiveOrders  bool
	CanQuery          bool
	CanChangeProtocol bool
	CanAddStaff       bool
	CanAddReminders   bool
	ID                uuid.UUID
}

func (q *Queries) UpdatePosition(ctx context.Context, arg UpdatePositionParams) error {
	_, err := q.db.ExecContext(ctx, updatePosition,
		arg.Title,
		arg.CanActivate,
		arg.CanDeactivate,
		arg.CanAddOrders,
		arg.CanReceiveOrders,
		arg.CanQuery,
		arg.CanChangeProtocol,
		arg.CanAddStaff,
		arg.CanAddReminders,
		arg.ID,
	)
	return err
}
