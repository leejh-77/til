go 의 error 타입은 아래의 인터페이스를 구현함으로써 정의할 수 있다.
```go
type error interface {
    Error() string
}
```
```go
errors.New("Error message")
```
로 간단히 만들 수 있다.
