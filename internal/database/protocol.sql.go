// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: protocol.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addBalance = `-- name: AddBalance :exec
UPDATE protocols SET balance = (balance + $2)
WHERE $1 = id
`

type AddBalanceParams struct {
	ID      uuid.UUID
	Balance int32
}

func (q *Queries) AddBalance(ctx context.Context, arg AddBalanceParams) error {
	_, err := q.db.ExecContext(ctx, addBalance, arg.ID, arg.Balance)
	return err
}

const createProtocol = `-- name: CreateProtocol :one
INSERT INTO protocols(id, p_number, primary_investigator, title, allocated, balance, expiration_date, is_active, previous_protocol)
VALUES(gen_random_uuid(), $1, $2, $3, $4, $5, $6, true, $7)
RETURNING id, p_number, primary_investigator, title, allocated, balance, expiration_date, is_active, previous_protocol
`

type CreateProtocolParams struct {
	PNumber             string
	PrimaryInvestigator uuid.UUID
	Title               string
	Allocated           int32
	Balance             int32
	ExpirationDate      time.Time
	PreviousProtocol    uuid.NullUUID
}

func (q *Queries) CreateProtocol(ctx context.Context, arg CreateProtocolParams) (Protocol, error) {
	row := q.db.QueryRowContext(ctx, createProtocol,
		arg.PNumber,
		arg.PrimaryInvestigator,
		arg.Title,
		arg.Allocated,
		arg.Balance,
		arg.ExpirationDate,
		arg.PreviousProtocol,
	)
	var i Protocol
	err := row.Scan(
		&i.ID,
		&i.PNumber,
		&i.PrimaryInvestigator,
		&i.Title,
		&i.Allocated,
		&i.Balance,
		&i.ExpirationDate,
		&i.IsActive,
		&i.PreviousProtocol,
	)
	return i, err
}

const getProtocolByID = `-- name: GetProtocolByID :one
SELECT id, p_number, primary_investigator, title, allocated, balance, expiration_date, is_active, previous_protocol FROM protocols
WHERE $1 = id
`

func (q *Queries) GetProtocolByID(ctx context.Context, id uuid.UUID) (Protocol, error) {
	row := q.db.QueryRowContext(ctx, getProtocolByID, id)
	var i Protocol
	err := row.Scan(
		&i.ID,
		&i.PNumber,
		&i.PrimaryInvestigator,
		&i.Title,
		&i.Allocated,
		&i.Balance,
		&i.ExpirationDate,
		&i.IsActive,
		&i.PreviousProtocol,
	)
	return i, err
}

const getProtocolByNumber = `-- name: GetProtocolByNumber :one
SELECT id, p_number, primary_investigator, title, allocated, balance, expiration_date, is_active, previous_protocol FROM protocols
WHERE $1 = p_number
`

func (q *Queries) GetProtocolByNumber(ctx context.Context, pNumber string) (Protocol, error) {
	row := q.db.QueryRowContext(ctx, getProtocolByNumber, pNumber)
	var i Protocol
	err := row.Scan(
		&i.ID,
		&i.PNumber,
		&i.PrimaryInvestigator,
		&i.Title,
		&i.Allocated,
		&i.Balance,
		&i.ExpirationDate,
		&i.IsActive,
		&i.PreviousProtocol,
	)
	return i, err
}

const getProtocols = `-- name: GetProtocols :many
SELECT id, p_number, primary_investigator, title, allocated, balance, expiration_date, is_active, previous_protocol FROM protocols
ORDER BY p_number DESC
`

func (q *Queries) GetProtocols(ctx context.Context) ([]Protocol, error) {
	rows, err := q.db.QueryContext(ctx, getProtocols)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Protocol
	for rows.Next() {
		var i Protocol
		if err := rows.Scan(
			&i.ID,
			&i.PNumber,
			&i.PrimaryInvestigator,
			&i.Title,
			&i.Allocated,
			&i.Balance,
			&i.ExpirationDate,
			&i.IsActive,
			&i.PreviousProtocol,
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

const updateAllocated = `-- name: UpdateAllocated :exec
UPDATE protocols SET allocated = $2
WHERE $1 = p_number
`

type UpdateAllocatedParams struct {
	PNumber   string
	Allocated int32
}

func (q *Queries) UpdateAllocated(ctx context.Context, arg UpdateAllocatedParams) error {
	_, err := q.db.ExecContext(ctx, updateAllocated, arg.PNumber, arg.Allocated)
	return err
}

const updateProtocol = `-- name: UpdateProtocol :exec
UPDATE protocols
SET p_number = $2,
    primary_investigator = $3,
    title = $4,
    allocated = $5,
    balance = $6,
    expiration_date = $7
WHERE $1 = id
`

type UpdateProtocolParams struct {
	ID                  uuid.UUID
	PNumber             string
	PrimaryInvestigator uuid.UUID
	Title               string
	Allocated           int32
	Balance             int32
	ExpirationDate      time.Time
}

func (q *Queries) UpdateProtocol(ctx context.Context, arg UpdateProtocolParams) error {
	_, err := q.db.ExecContext(ctx, updateProtocol,
		arg.ID,
		arg.PNumber,
		arg.PrimaryInvestigator,
		arg.Title,
		arg.Allocated,
		arg.Balance,
		arg.ExpirationDate,
	)
	return err
}
