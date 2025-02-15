// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: orders.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createNewOrder = `-- name: CreateNewOrder :one
INSERT INTO orders(id, order_number, expected_date, protocol_id, investigator_id, strain_id, note, received)
VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, $6, false)
RETURNING id, order_number, expected_date, protocol_id, investigator_id, strain_id, note, received
`

type CreateNewOrderParams struct {
	OrderNumber    string
	ExpectedDate   time.Time
	ProtocolID     uuid.UUID
	InvestigatorID uuid.UUID
	StrainID       uuid.UUID
	Note           sql.NullString
}

func (q *Queries) CreateNewOrder(ctx context.Context, arg CreateNewOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createNewOrder,
		arg.OrderNumber,
		arg.ExpectedDate,
		arg.ProtocolID,
		arg.InvestigatorID,
		arg.StrainID,
		arg.Note,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.OrderNumber,
		&i.ExpectedDate,
		&i.ProtocolID,
		&i.InvestigatorID,
		&i.StrainID,
		&i.Note,
		&i.Received,
	)
	return i, err
}

const getAllExpectedOrders = `-- name: GetAllExpectedOrders :many
SELECT id, order_number, expected_date, protocol_id, investigator_id, strain_id, note, received FROM orders
WHERE received = false
`

func (q *Queries) GetAllExpectedOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getAllExpectedOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.OrderNumber,
			&i.ExpectedDate,
			&i.ProtocolID,
			&i.InvestigatorID,
			&i.StrainID,
			&i.Note,
			&i.Received,
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

const getAllOrders = `-- name: GetAllOrders :many
SELECT id, order_number, expected_date, protocol_id, investigator_id, strain_id, note, received FROM orders
`

func (q *Queries) GetAllOrders(ctx context.Context) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getAllOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.OrderNumber,
			&i.ExpectedDate,
			&i.ProtocolID,
			&i.InvestigatorID,
			&i.StrainID,
			&i.Note,
			&i.Received,
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

const getOrderByID = `-- name: GetOrderByID :one
SELECT id, order_number, expected_date, protocol_id, investigator_id, strain_id, note, received FROM orders
WHERE id = $1
`

func (q *Queries) GetOrderByID(ctx context.Context, id uuid.UUID) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderByID, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.OrderNumber,
		&i.ExpectedDate,
		&i.ProtocolID,
		&i.InvestigatorID,
		&i.StrainID,
		&i.Note,
		&i.Received,
	)
	return i, err
}

const getOrderByNumber = `-- name: GetOrderByNumber :one
SELECT id, order_number, expected_date, protocol_id, investigator_id, strain_id, note, received FROM orders
WHERE order_number = $1
`

func (q *Queries) GetOrderByNumber(ctx context.Context, orderNumber string) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderByNumber, orderNumber)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.OrderNumber,
		&i.ExpectedDate,
		&i.ProtocolID,
		&i.InvestigatorID,
		&i.StrainID,
		&i.Note,
		&i.Received,
	)
	return i, err
}

const getOrderDateRange = `-- name: GetOrderDateRange :many
SELECT id, order_number, expected_date, protocol_id, investigator_id, strain_id, note, received FROM orders
WHERE (expected_date BETWEEN $1 AND $2)
`

type GetOrderDateRangeParams struct {
	ExpectedDate   time.Time
	ExpectedDate_2 time.Time
}

func (q *Queries) GetOrderDateRange(ctx context.Context, arg GetOrderDateRangeParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getOrderDateRange, arg.ExpectedDate, arg.ExpectedDate_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.OrderNumber,
			&i.ExpectedDate,
			&i.ProtocolID,
			&i.InvestigatorID,
			&i.StrainID,
			&i.Note,
			&i.Received,
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

const getOrderExpectedToday = `-- name: GetOrderExpectedToday :many
SELECT id, order_number, expected_date, protocol_id, investigator_id, strain_id, note, received FROM orders
WHERE (expected_date = $1) AND received = false
`

func (q *Queries) GetOrderExpectedToday(ctx context.Context, expectedDate time.Time) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getOrderExpectedToday, expectedDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.OrderNumber,
			&i.ExpectedDate,
			&i.ProtocolID,
			&i.InvestigatorID,
			&i.StrainID,
			&i.Note,
			&i.Received,
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

const getUserExpectedOrders = `-- name: GetUserExpectedOrders :many
SELECT id, order_number, expected_date, protocol_id, investigator_id, strain_id, note, received FROM orders
WHERE received = false
AND expected_date <= $1
AND investigator_id = $2
`

type GetUserExpectedOrdersParams struct {
	ExpectedDate   time.Time
	InvestigatorID uuid.UUID
}

func (q *Queries) GetUserExpectedOrders(ctx context.Context, arg GetUserExpectedOrdersParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getUserExpectedOrders, arg.ExpectedDate, arg.InvestigatorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.OrderNumber,
			&i.ExpectedDate,
			&i.ProtocolID,
			&i.InvestigatorID,
			&i.StrainID,
			&i.Note,
			&i.Received,
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

const markOrderReceived = `-- name: MarkOrderReceived :one
UPDATE orders
SET received = true
WHERE id = $1
RETURNING id, order_number, expected_date, protocol_id, investigator_id, strain_id, note, received
`

func (q *Queries) MarkOrderReceived(ctx context.Context, id uuid.UUID) (Order, error) {
	row := q.db.QueryRowContext(ctx, markOrderReceived, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.OrderNumber,
		&i.ExpectedDate,
		&i.ProtocolID,
		&i.InvestigatorID,
		&i.StrainID,
		&i.Note,
		&i.Received,
	)
	return i, err
}

const updateOrder = `-- name: UpdateOrder :exec
UPDATE orders
SET expected_date = $2,
    protocol_id = $3,
    investigator_id = $4,
    strain_id = $5,
    note = $6
WHERE $1 = id
`

type UpdateOrderParams struct {
	ID             uuid.UUID
	ExpectedDate   time.Time
	ProtocolID     uuid.UUID
	InvestigatorID uuid.UUID
	StrainID       uuid.UUID
	Note           sql.NullString
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) error {
	_, err := q.db.ExecContext(ctx, updateOrder,
		arg.ID,
		arg.ExpectedDate,
		arg.ProtocolID,
		arg.InvestigatorID,
		arg.StrainID,
		arg.Note,
	)
	return err
}
