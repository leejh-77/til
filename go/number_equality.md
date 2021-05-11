go 에서 숫자는 int, int8, int16.., uint, uint8... 의 타입을 갖는데\
숫자값이 같더라도 타입이 다르면 동등성 체크에서 false 가 나온다.
```go
i := 1
i2 := int64(1)

if i == i2 { // compile 에러.
    
}
```
기본적으로 == 로 타입이 다른 숫자를 비교할 때는 컴파일 에러가 난다.\

아래와 같은 상황에서는 주의해야한다.
```go
func TestInt(t *testing.T) {
	i := 1
	i2 := int64(1)

	log.Println(isEqual(i, i2)) // false
}

func isEqual(i1 interface{}, i2 interface{}) bool {
	return i1 == i2
}
```