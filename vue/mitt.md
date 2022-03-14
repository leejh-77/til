vue2 에서는 Vue 컴포넌트 객체를 따로 생성해 event bus 로 쓰는 패턴이 많이 쓰였지만  
vue3 에 와서는 mitt 라이브러리를 통해 이벤트를 전달하도록 권장하고 있다  

```javascript
const app = createApp(App)
export const emitter = mitt()
app.config.globalProperties.$emitter = emitter

this.$emitter.emit('event', data)

this.$emitter.on('event', (data) => {
    // do something
})
```