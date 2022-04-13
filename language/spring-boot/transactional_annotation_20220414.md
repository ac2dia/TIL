# @Transcational 정리

```java
@Target(value={TYPE,METHOD})
 @Retention(value=RUNTIME)
 @Inherited
 @Documented
public @interface Transactional
```

- 대상 범위의 경우 method, class, interface 에 사용 가능
- 동시에 사용되는 경우 method 우선
- Proxy Mode (default), AspectJ Mode 가 있다.

## @Transcational 설정

```
value(String) : 사용할 트랜잭션 관리자
propagation(enum: Propagation) : 사용할 전파 설정
isolation(enum: isolation) : 선택적 격리 수준
readOnly(boolean) : 읽기/쓰기 vs 읽기 전용 트랜잭션
timeout(int(초)) : 트랜잭션 타임 아웃
rollbackFor(Throwable 로부터 얻을 수 있는 Class 객체 배열) : 롤백이 수행되어야하는 선택적인 예외 클래스의 배열
rollbackForClassName(Throwable 로부터 얻을 수 있는 클래스 이름 배열) : 롤백이 수행되어야 하는, 선택적인 예외 클래스 이름의 배열
noRollbackFor(Throwable 로부터 얻을 수 있는 Class 객체 배열) : 롤백이 수행되지 않아야 하는, 선택적인 예외 클래스의 배열
noRollbackForClassName(Throwable 로부터 얻을 수 있는 클래스 이름 배열) : 롤백이 수행되지 않아야 하는, 선택적인 예외 클래스 이름의 배열
```

## 주의점

1. private 접근 제한자를 사용하는 경우 @Transactional 적용되지 않습니다.
2. 같은 클래스 내에서 여러 @Transactional method 를 호출하는 경우 @Transactional 기능이 동작하지 않습니다.
3. @Transcational 의 경우 AOP 를 사용하여 Proxy 형태로 동작합니다.

## 사용

- 현재 repository interface 에 대해서만 @Repository 어노테이션을 통해 @Transactional 이 적용된 상태입니다.
- 그래서 추가로 한 서비스에서 여러 트랜잭션이 일어나거나 트랜잭션 이 후 추가 로직이 있는 경우 repository 의 메서드를 호출하는 service method 에도 @Transactional 설정이 필요합니다.

# Reference

[1] Annotation Type Transactional
, https://docs.spring.io/spring-framework/docs/current/javadoc-api/org/springframework/transaction/annotation/Transactional.html
[2] Proxy형태로 동작하는 JPA @Transactional
, https://cobbybb.tistory.com/17
