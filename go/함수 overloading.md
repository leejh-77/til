go 는 함수 오버로딩을 지원하지 않는다..\
따라서 아래와 같이 함수를 정의할 시 컴파일 에러가 난다.
```go
func Function(p1 string) {
}

func Function(p1 string, p2 string) {
}
```