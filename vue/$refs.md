컴포넌트에 ref="name" 를 선언하면  
javascript 코드에서 this.$refs.name 으로 컴포넌트에 직접 접근할 수 있다  
name 에 - 가 들어가는 경우에는 this.$ref['name-slash'] 로 접근

```javascript
<p ref="name">
<p ref="name-slash">

func() {
    const p1 = this.$refs.name
    const p2 = this.$refs['name-slash']
}
```
