InnoDB 에서 외래키를 지정하면 자동으로 해당 칼럼에 인덱스가 지정되므로 기본적으로 인덱스가 가지는 특징을 가진다.  
외래키는 참조 관계를 보장해야하기 때문에 업데이트 시 경우에 따라 쓰기 잠금이 걸릴 수 있다.