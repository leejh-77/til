mysql 의 구조를 크게 둘로 나누면 mysql 엔진과 storage 엔진이다. 

## mysql 엔진
mysql 엔진은 쿼리 분석부터, 옵티마이징, 쿼리 실행기까지를 말한다.  

## storage 엔진
storage 엔진은 직접 파일로부터 데이터를 불러오거나 업데이트하는 역할을 담당한다.  
MyISAM 이나 InnoDB 가 storage 엔진에 해당하며 mysql 의 handler api 규격을 준수하면  
어떤 storage 엔진도 사용할 수 있다.