@Schduled 어노테이션을 이용하면 일정 주기로 반복적인 작업을 할 수 있다.  
주기 설정은 스프링에서 제공하는 방식 외에 crontab 의 표현 방식도 지원한다.  

@Schduled 를 사용하려면 어플리케이션이나 Configuration 클래스에 @EnableScheduling 을 
달아주어야 한다.