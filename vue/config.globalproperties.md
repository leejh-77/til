vue2 에서는 컴포넌트에서 글로벌하게 접근 가능한 변수를 protoype 을 이용해 선언해주었지만  
vue3 에 와서는 app.config.globalproperties 에 선언하도록 하고 있다  

```javascript
const app = createApp(App)
app.config.globalProperties.$variable = 1

this.$variable // 1
```