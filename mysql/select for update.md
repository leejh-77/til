select ... for update 구문은 트랜잭션 안에서 select 의 결과로 반환된 row 들이  
다른 세션에서 접근하지 못하도록 막는다. 접근을 시도하는 세션은 락이 걸린다.