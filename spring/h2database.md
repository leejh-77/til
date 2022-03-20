h2 database 를 사용하기 위한 패키지 및 datasource 설정  
```java
// gradle
testImplementation 'com.h2database:h2'

// application.properties
spring.datasource.url=jdbc:h2:~/test
spring.datasource.driver-class-name=org.h2.Driver
spring.datasource.username=sa
spring.datasource.password=sa
spring.jpa.database-platform=org.hibernate.dialect.H2Dialect
```