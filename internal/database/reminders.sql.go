// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: reminders.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addReminder = `-- name: AddReminder :one
INSERT INTO reminders(id, r_date, r_cc_id, investigator_id, note)
VALUES (gen_random_uuid(), $1, $2, $3, $4)
RETURNING id, r_date, r_cc_id, investigator_id, note
`

type AddReminderParams struct {
	RDate          time.Time
	RCcID          int32
	InvestigatorID uuid.UUID
	Note           string
}

func (q *Queries) AddReminder(ctx context.Context, arg AddReminderParams) (Reminder, error) {
	row := q.db.QueryRowContext(ctx, addReminder,
		arg.RDate,
		arg.RCcID,
		arg.InvestigatorID,
		arg.Note,
	)
	var i Reminder
	err := row.Scan(
		&i.ID,
		&i.RDate,
		&i.RCcID,
		&i.InvestigatorID,
		&i.Note,
	)
	return i, err
}

const deleteReminder = `-- name: DeleteReminder :exec
DELETE FROM reminders
WHERE $1 = id
`

func (q *Queries) DeleteReminder(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteReminder, id)
	return err
}

const getAllReminders = `-- name: GetAllReminders :many
SELECT id, r_date, r_cc_id, investigator_id, note FROM reminders
`

func (q *Queries) GetAllReminders(ctx context.Context) ([]Reminder, error) {
	rows, err := q.db.QueryContext(ctx, getAllReminders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reminder
	for rows.Next() {
		var i Reminder
		if err := rows.Scan(
			&i.ID,
			&i.RDate,
			&i.RCcID,
			&i.InvestigatorID,
			&i.Note,
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

const getAllTodayReminders = `-- name: GetAllTodayReminders :many
SELECT id, r_date, r_cc_id, investigator_id, note FROM reminders
WHERE r_date = $1
`

func (q *Queries) GetAllTodayReminders(ctx context.Context, rDate time.Time) ([]Reminder, error) {
	rows, err := q.db.QueryContext(ctx, getAllTodayReminders, rDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reminder
	for rows.Next() {
		var i Reminder
		if err := rows.Scan(
			&i.ID,
			&i.RDate,
			&i.RCcID,
			&i.InvestigatorID,
			&i.Note,
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

const getRemindersByCC = `-- name: GetRemindersByCC :many
SELECT id, r_date, r_cc_id, investigator_id, note FROM reminders
WHERE r_cc_id = $1
ORDER BY r_date
`

func (q *Queries) GetRemindersByCC(ctx context.Context, rCcID int32) ([]Reminder, error) {
	rows, err := q.db.QueryContext(ctx, getRemindersByCC, rCcID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reminder
	for rows.Next() {
		var i Reminder
		if err := rows.Scan(
			&i.ID,
			&i.RDate,
			&i.RCcID,
			&i.InvestigatorID,
			&i.Note,
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

const getRemindersDateRange = `-- name: GetRemindersDateRange :many
SELECT id, r_date, r_cc_id, investigator_id, note FROM reminders
WHERE (r_date BETWEEN $1 AND $2)
`

type GetRemindersDateRangeParams struct {
	RDate   time.Time
	RDate_2 time.Time
}

func (q *Queries) GetRemindersDateRange(ctx context.Context, arg GetRemindersDateRangeParams) ([]Reminder, error) {
	rows, err := q.db.QueryContext(ctx, getRemindersDateRange, arg.RDate, arg.RDate_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reminder
	for rows.Next() {
		var i Reminder
		if err := rows.Scan(
			&i.ID,
			&i.RDate,
			&i.RCcID,
			&i.InvestigatorID,
			&i.Note,
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

const getUserDayReminder = `-- name: GetUserDayReminder :many
SELECT id, r_date, r_cc_id, investigator_id, note FROM reminders
WHERE investigator_id = $1
AND r_date = $2
`

type GetUserDayReminderParams struct {
	InvestigatorID uuid.UUID
	RDate          time.Time
}

func (q *Queries) GetUserDayReminder(ctx context.Context, arg GetUserDayReminderParams) ([]Reminder, error) {
	rows, err := q.db.QueryContext(ctx, getUserDayReminder, arg.InvestigatorID, arg.RDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reminder
	for rows.Next() {
		var i Reminder
		if err := rows.Scan(
			&i.ID,
			&i.RDate,
			&i.RCcID,
			&i.InvestigatorID,
			&i.Note,
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

const getUserTodayReminders = `-- name: GetUserTodayReminders :many
SELECT id, r_date, r_cc_id, investigator_id, note FROM reminders
WHERE r_date = $1 AND investigator_id = $2
`

type GetUserTodayRemindersParams struct {
	RDate          time.Time
	InvestigatorID uuid.UUID
}

func (q *Queries) GetUserTodayReminders(ctx context.Context, arg GetUserTodayRemindersParams) ([]Reminder, error) {
	rows, err := q.db.QueryContext(ctx, getUserTodayReminders, arg.RDate, arg.InvestigatorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Reminder
	for rows.Next() {
		var i Reminder
		if err := rows.Scan(
			&i.ID,
			&i.RDate,
			&i.RCcID,
			&i.InvestigatorID,
			&i.Note,
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
