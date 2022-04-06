inverted index 는 elastic search 에서 도큐먼트 데이터를 검색하기 위해 사용하는 인덱스다.  

inverted index 는 고유한 단어들을 키로, 이 단어가 포함된 문서의 id 리스트를 매핑한다.  
때문에 특정 단어를 찾기 위해 모든 row 를 뒤져야하는 RDB 에 비해 full text search 를 훨씬 빠르게 할 수 있다.  

inverted index 에 단어들의 목록을 저장할 때는 book/Book, book/books 과 같이 유사한 단어는 같은 단어로 보고  
정규화한 뒤 검색어에 해당하는 단어도 정규화해 검색을 수행한다.
