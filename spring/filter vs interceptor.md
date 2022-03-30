### Filter
J2EE 표준 스펙으로 Dispatcher Servlet 전에 실행된다.

### Interceptor
스프링의 기능으로 Servlet 을 거친 뒤, 컨트롤러로 요청이 들어오기 전에 실행된다.

기본적으로 Filter 는 Servlet Container 에 의해 관리되기 때문에  
스프링 컨텍스트 안에서 빈으로 Filter 를 만들어 등록하는 것은 불가능했으나, 이후 DelegatingFilterProxy 를 통해 주입할 수 있게 되었다.  
스프링 부트의 경우에는 내장 웹서버를 지원하기 때문에 서블릿 컨테이너까지 관리할 수 있게 되어 DelegateFilterProxy 가 필요없다.