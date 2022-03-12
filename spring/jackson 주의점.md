spring 에서 요청 및 응답을 구조체와 바인딩할 때 기본적으로 jackson 을 사용하는데  
jackson 은 public 한 필드가 아니면 getter 를 이용해 필드에 접근하므로  
private 한 필드를 선언했다면 반드시 getter 를 선언해야 한다