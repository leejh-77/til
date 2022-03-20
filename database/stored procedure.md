stored procedure 는 반복적으로 사용하는 쿼리를 DB 서버에 저장해놓고  
클라이언트에서는 필요한 데이터만 보내 쿼리에 사용되는 리소스를 줄이기 위해 사용된다.  

장점  
- 데이터만 보내면 되기 때문에 통신에 사용되는 데이터량을 줄일 수 있다.  
- 데이터뿐만 아니라 네트워킹에 사용되는 비용도 줄일 수 있다.  
- stored procedure 는 DB 서버 내부에서 컴파일된 상태로 있기 때문에 구문 분석에 사용되는 비용도 줄일 수 있다.  

단점
- procedure 내부에서 데이터를 가공하기 위한 연산을 수행할 경우 일반적으로 어플리케이션에서 하는 것보다 성능이 떨어진다.  
- 디버깅하기 어렵다.