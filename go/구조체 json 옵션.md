go 에서는 기본적으로 구조체에 \`json\` 태그를 붙이면 해당 구조체를 json 형식으로 직렬화할 수 있다.
```go
type User struct {
    EmailAddress string `json:"emailAddress"`
    Password string `json:"password"`
}

json.Marshal(/* 구조체 포인터 */)
```
위에서 User 의 EmailAddress, Password 필드는 태그로 설정한 emailAddress, password 의 이름으로 직렬화된다.\
이때 옵션을 줄 수가 있는데 아래와 같다.
```go
type User struct {
	// omitempty 시 필드의 데이터가 없을 경우 json 에서 빼버린다.
    EmailAddress string `json:"emailAddress,omitempty"`
    // - 를 붙일 시 해당 필드는 json 에서 무조건 빠진다.
    // json 태그를 아예 붙이지 않으면 직렬화시 에러남.
    Password string `json:"-"`
}
```