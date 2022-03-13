자식 컴포넌트에서 부모로 이벤트를 전달하고 싶으면 this.$emit('onReceiveEvent', data)와 같이 넘기면 된다.  
받는 쪽에서는 @onReceiveEvent="onReceiveEvent" 와 같이 @를 통해 emit 된 메시지를 캡처하고  
컴포넌트 내에서 구현한 함수를 매핑하면 된다.