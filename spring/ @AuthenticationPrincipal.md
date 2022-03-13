@AuthenticationPrincipal 을 이용하면 자동으로 SecurityContextHolder 에서 현재 세션에 저장된 사용자의  
인증정보를 찾아 주입해준다. 선언된 타입에 따라 자동으로 principal 을 매핑해주므로 custom 한 userDetails 클래스를  
사용하는 경우 간편하게 쓸 수 있다.