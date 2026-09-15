# divmod

Реализуйте функцию:

```go
func DivMod(dividend, divisor int) (quotient, remainder int, err error)
```

Если `divisor` не равен нулю, функция должна вернуть те же частное и остаток,
что выражения `dividend / divisor` и `dividend % divisor` в Go. Ошибка в этом
случае равна `nil`.

Если `divisor == 0`, функция должна вернуть нулевые частное и остаток, а в
качестве ошибки — `ErrDivisionByZero`.

Паниковать при ожидаемой ошибке нельзя.

## Проверка

```shell
go test ./divmod/...
```
