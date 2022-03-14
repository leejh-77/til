javascript 에서 외부로 export 할 때 default 키워드를 붙이면  
사용하는 쪽에서 {} 를 생략할 수 있다. 그냥 export 로 하면 {} 를 붙여야 함  

default 가 아닌 변수를 export 할 때는 export 와 동시에 초기화 시켜야한다.
default 일 때는 변수명 생략 가능

```javascript
const variable = new Class()
export default variable

export default new Class()

export const variable = new Class()
```