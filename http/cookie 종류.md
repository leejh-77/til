Session Cookie - 메모리에만 저장되며, 브라우저 종료 시 삭제된다  
Persistent Cookie - 파일로 저장되어 브라우저 종료와 상관없이 유지된다  
Secure Cookie - HTTPS 에서 사용. 쿠키 정보가 암회화되어 전송된다  
Third-Party Cookie - 방문한 도메인과 다른 도메인의 쿠키

서버에서 Set-Cookie 로 쿠키를 설정할 때 만료 기간을 따로 정해주지 않으면 Session Cookie 로 간주하고,  
만료 기한이 있으면 그 시간까지는 Persistent 로 저장된다.