mysql 의 기본 스토리지 엔진인 InnoDB 는 프라이머리 키를 기준으로 클러스터링 인덱스를 사용한다.  
클러스터링 인덱스는 인덱스를 기준으로 정렬된 데이터를 그대로 디스크에 저장하는 방식이다.  
즉, 디스크에 저장된 데이터 자체가 인덱스를 기준으로 정렬되어 있기 때문에 프라이머리 키를 기준으로 한 검색이 굉장히 빠르다.
  
만약 프라이머리 키를 명시적으로 지정하지 않았다면  
1. Not null, unique 한 인덱스 중에서 첫번째를 사용하거나,
2. 자동으로 유니크하게 증가하는 칼럼을 내부적으로 추가해 사용한다.