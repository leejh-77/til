RDB 마다 문법이 조금씩 달라서 사용할 때마다 헷갈린다. (그래봤자 sqlite, mysql 만 사용하지만..)\
Mysql 은 auto_increment 를 사용한다. (sqlite 는 autoincrement)
```sql
CREATE TABLE `table` (
    `id` int primary key auto_increment,
    ...
)
```