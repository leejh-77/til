javascript 는 동적으로 클래스에 프로퍼티를 추가할 수 있기 때문에 아래와 같이  
prototype 을 이용하면 Vue 컴포넌트 내에서 접근할 수 있는 변수를 추가할 수 있다.

```javascript
Vue.prototype.$bus = eventBus

this.$bus
```