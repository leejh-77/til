SSE(Server-Sent Event) 는 클라이언트와 서버가 연결을 맺은 뒤 서버가 클라이언트로 메시지를 지속적으로 보낼 수 있는 단방향 전송 방식이다.  
웹소켓이랑 비교해서 서버를 따로 띄울 필요없이 http 스펙 안에서 지속적인 메시징을 할 수 있다는 장점이 있다.  

Content type 은 text/event-stream 로 설정한다