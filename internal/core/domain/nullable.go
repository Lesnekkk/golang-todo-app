package domain

// Nullable позволяет различать три состояния в PATCH-запросах:
//   - Set=false           → поле не передано (ничего не меняем)
//   - Set=true, Value=nil → поле передано как null (очищаем в БД)
//   - Set=true, Value=&v  → поле передано с конкретным значением (обновляем)
type Nullable[T any] struct {
	Value *T
	Set   bool
}
