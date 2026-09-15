# permissions

Права доступа представлены отдельными битами типа `Permission`:

```go
const (
    Read Permission = 1 << iota
    Write
    Execute
)
```

Реализуйте три функции:

```go
func Grant(current, added Permission) Permission
func Revoke(current, removed Permission) Permission
func Has(current, required Permission) bool
```

- `Grant` возвращает права из `current` вместе со всеми правами из `added`.
- `Revoke` удаляет из `current` все права из `removed`.
- `Has` возвращает `true`, только если в `current` присутствуют все биты из
  `required`.

Функции не должны менять аргументы: значения типа `Permission` передаются по
значению.

## Проверка

```shell
go test ./permissions/...
```
