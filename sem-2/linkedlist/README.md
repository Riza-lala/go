# linkedlist

Дан узел односвязного списка:

```go
type Node struct {
    Value int
    Next  *Node
}
```

Указатель `nil` обозначает пустой список. Реализуйте функции:

```go
func Prepend(head *Node, value int) *Node
func Length(head *Node) int
func Find(head *Node, value int) *Node
func Reverse(head *Node) *Node
```

- `Prepend` создаёт новую голову со значением `value`; прежний список идёт
  следом.
- `Length` возвращает число узлов.
- `Find` возвращает указатель на первый узел с нужным значением или `nil`.
- `Reverse` разворачивает связи между существующими узлами и возвращает новую
  голову. Создавать новые узлы в `Reverse` не нужно.

Обратите внимание: сам `head` передаётся по значению. Поэтому операции, которые
меняют голову списка, возвращают новый указатель.

## Проверка

```shell
go test ./linkedlist/...
```
