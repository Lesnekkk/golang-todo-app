package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	core_errors "github.com/Lesnekkk/golang-todo-app/internal/core/errors"
)

type Task struct {
	ID      uuid.UUID
	Version int

	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorUserID uuid.UUID
}

func NewTask(
	id uuid.UUID,
	version int,
	title string,
	description *string,
	completed bool,
	createdAt time.Time,
	completedAt *time.Time,
	authorUserID uuid.UUID,
) Task {
	return Task{
		ID:           id,
		Version:      version,
		Title:        title,
		Description:  description,
		Completed:    completed,
		CreatedAt:    createdAt,
		CompletedAt:  completedAt,
		AuthorUserID: authorUserID,
	}
}

func CreateTask(title string, description *string, authorUserID uuid.UUID) Task {
	return NewTask(uuid.New(), 1, title, description, false, time.Now(), nil, authorUserID)
}

func (t *Task) CompletionDuration() *time.Duration {
	if !t.Completed || t.CompletedAt == nil {
		return nil
	}
	d := t.CompletedAt.Sub(t.CreatedAt)
	return &d
}

func (t *Task) Validate() error {
	titleLen := len([]rune(t.Title))
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf("invalid `Title` len: %d: %w", titleLen, core_errors.ErrInvalidArgument)
	}

	if t.Description != nil {
		descLen := len([]rune(*t.Description))
		if descLen < 1 || descLen > 1000 {
			return fmt.Errorf("invalid `Description` len: %d: %w", descLen, core_errors.ErrInvalidArgument)
		}
	}

	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf("`CompletedAt` can't be nil if Completed==true: %w", core_errors.ErrInvalidArgument)
		}
		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf("`CompletedAt` can't be before `CreatedAt`: %w", core_errors.ErrInvalidArgument)
		}
	} else {
		if t.CompletedAt != nil {
			return fmt.Errorf("`CompletedAt` must be nil if Completed==false: %w", core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

type TaskPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func NewTaskPatch(title Nullable[string], description Nullable[string], completed Nullable[bool]) TaskPatch {
	return TaskPatch{Title: title, Description: description, Completed: completed}
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
	if patch.Title.Set && patch.Title.Value == nil {
		return fmt.Errorf("`Title` can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}
	if patch.Completed.Set && patch.Completed.Value == nil {
		return fmt.Errorf("`Completed` can't be patched to NULL: %w", core_errors.ErrInvalidArgument)
	}

	tmp := *t

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}
	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}
	if patch.Completed.Set {
		tmp.Completed = *patch.Completed.Value
		if tmp.Completed {
			now := time.Now()
			tmp.CompletedAt = &now
		} else {
			tmp.CompletedAt = nil
		}
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched task: %w", err)
	}

	*t = tmp
	return nil
}
