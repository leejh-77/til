vue.config.js 에 아래와 같은 옵션을 주면 프록시 서버를 실행한다.  
개발 단계에서 cors 를 우회하기 위한 방법으로 사용될 수 있다.

```javascript
devServer: {
    port: 3000,
    proxy: {
        '/api/*': {
        target: 'http://localhost:8080'
        }
    }
}
```