### From swift
build setting 에서 bridging-header 파일을 지정해주면 objective-c 클래스를 사용할 수 있는데,\
기본적으로 swift 와 objective-c 가 혼용된 프로젝트에서 파일을 추가하면 briding-header 를 추가할건지 묻는다.\
추가한다고 선택하면 "프로젝트 이름 + -Bridging-Header.h" 로 헤더파일이 생성되는데\
이곳에 swift 에서 사용할 objective-c 헤더들을 import 해주면 된다.

### From objective-c
@obj 를 붙여준 swift 클래스와 프로퍼티 및 함수들은 컴파일할 때 objective-c 에서 사용할 수 있도록\
"프로젝트 이름 + -Swift.h" 파일에 자동 생성된다. 사용하는 곳에서는 이 헤더파일을 import 해서
사용하면 된다.