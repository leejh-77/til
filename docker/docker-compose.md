docker-compose 는 여러개의 도커 이미지들을 동시에 실행하기 위해 사용한다.  
docker-compose 를 통해 생성되는 컨테이너들은 기본적으로 같은 네트워크 안에서 실행된다.

```yml
version: "3"
services:
    app: myapp
    ///...
    
    app-mq: myapp-mq
    ///...

```

docker-compose up 
docker-compose up -d -> 데몬으로 실행
docker-compose down