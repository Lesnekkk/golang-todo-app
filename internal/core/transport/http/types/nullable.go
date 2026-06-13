package core_http_types

import (
	"encoding/json"

	"github.com/Lesnekkk/golang-todo-app/internal/core/domain"
)

// Nullable — HTTP-версия доменного Nullable[T] с реализацией UnmarshalJSON.
// Различает три случая:
//   - Поле отсутствует в JSON → UnmarshalJSON не вызывается → Set==false
//   - Поле = null             → Set=true, Value=nil
//   - Поле = конкретное значение → Set=true, Value=&value
type Nullable[T any] struct {
	domain.Nullable[T]
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	n.Set = true

	if string(b) == "null" {
		n.Value = nil
		return nil
	}

	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	n.Value = &value
	return nil
}

func (n *Nullable[T]) ToDomain() domain.Nullable[T] {
	return domain.Nullable[T]{Value: n.Value, Set: n.Set}
}
