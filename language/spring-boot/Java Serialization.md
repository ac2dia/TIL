# Java Serialization

# 직렬화(Serialize)란?

- 자바 시스템 내부에서 사용되는 Object 또는 Data를 외부의 자바 시스템에서도 사용할 수 있도록 byte 형태로 데이터를 변화하는 기술
- JVM의 메모리에 상주되어 있는 객체 데이터를 바이트 형태로 변환하는 기술

# 역직렬화(Deserialize)란?

- byte로 변환된 Data를 원래대로 Object 또는 Data로 변환하는 기술
- 직렬화된 바이트 형태의 데이터를 객체로 변환해서 JVM으로 상주시키는 형태

# 어디에 사용되는가?

- 서블릿 세션 (Servlet Session)
- 캐시 (Cache)
- 자바 RMI (Remote Method Inv

# serialVersionUID

- Java Class 멤버 변수가 추가, 변경, 삭제가 일어나는 경우 serialVersionUID 가 달라지는데 이러한 경우 역직렬화시 serialVersionUID 가 다른 경우 java.io.InvalidClassException 예외가 발생합니다.
- 그러므로 Class 별로 serialVersionUID 를 직접 관리해야 클래스가 변경되어도 문제 없이 직렬화/역직렬화를 할 수 있습니다.

# 실무 사용 조언

- serialVersionUID는 개발 시 직접 관리
- 역직렬화 대상 클래스의 멤버 변수 타입변경 지양
- 외부(DB, 캐시 서버, NoSQL 서버 등)에 장기간 저장될 정보는 Java 직렬화 사용 지양 (클래스 변경을 예측할 수 없으므로)
- 개발자가 직접 컨트롤할 수 없는 클래스(프레임워크 또는 라이브러리에서 제공하는 클래스)는 직렬화 지양
- 자주 변경되는 클래스는 Java 직렬화를 사용하지 않는 것이 좋다.
- 역직렬화에 실패하는 상황에 대한 예외처리 필수
- 직렬화된 데이터는 타입 정보등의 클래스 메타정보를 포함하기 때문에 JSON 포맷에 비해 약 2~10배 더 사이즈가 크다. 특히 직렬화된 데이터를 메모리 서버(Redis, Memcached)에 저장하는 환경에서 트래픽에 따라 네트워크 비용과 캐시 서버 비용이 급증할 수 있으므로, JSON 포맷으로의 변경을 고려해야 한다.

# Reference

[1] Java의 직렬화(Serialize)란?, https://nesoy.github.io/articles/2018-04/Java-Serialize
[2] Java Serialization 개념 정리, https://ryan-han.com/post/java/serialization/
[3] Java Docs - Interface Serializable, https://docs.oracle.com/javase/8/docs/api/java/io/Serializable.html
