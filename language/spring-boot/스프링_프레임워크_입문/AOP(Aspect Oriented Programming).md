# AOP(Aspect Oriented Programming)
- 흩어진 코드를 한 곳으로 모으고, 다른 기타 클래스는 자신이 해야하는 일만 하도록 (SRP)!
- 객체지향원칙에 맞게 코딩할 수 있도록 도와주는 코딩 기법

## 흩어진 AAAA BBBB
``` java
class A {
    method a() {
        AAAA
        ...
        BBBB
    }

    method b() {
        AAAA
        ...
        BBBB
    }
}

class B {
    method c() {
        AAAA
        ...
        BBBB
    }
}

## 모아놓은 AAAA BBBB
class A {
    method a() {
        ...
    }

    method b() {
        ...
    }
}

class B {
    method c() {
        ...
    }
}

class AAAABBBBB {
    method aaaabbbbb(JoinPoint point) {
        AAAA
        point.execute()
        BBBB
    }
}
```
## AOP 주요 애노테이션
- @Aspect : AOP를 정의하는 클래스에 할당
- @Pointcut : AOP를 적용 시킬 지점 설정
- @Before : 메소드 실행하기 이전
- @After : 메소드가 성공적으로 실행 후 예외가 발생되더라도 실행
- @AfterReturing : 메소드 호출 성공 실행 시
- @AfterThrowing : 메소드 호출 실패 예외 발생
- @Around : Before/After 모두 제어

## AOP 적용하는 방법
1. Byte 코드 조작
2. Proxy 패턴 적용
- Spring AOP 는 해당 방법을 사용합니다.

## AOP 적용 예제
- 애노테이션은 스스로 동작하는 것이 아니라 AOP 등을 통해 실제로 동작하는 코드들이 있습니다.

``` java
@Target(ElementType.METHOD) // 애노테이션 적용 범위
@Retention(RetentionPolicy.RUNTIME) // 애노테이션 생명 주기
public @interface LogExecutionTime {

}

@Aspect
@Component
public class LogAspect {

    Logger logger = LoggerFactory.getLogger(LogAspect.class);

    @Around("@annotation(LogExecutionTime)")
    public object logExecutionTime(ProceedingJoinPoint joinPoint) throws ThreadException {
        stopWatch stopWatch = new StopWatch();
        stopWatch.start();

        Object ret = jointPoint.proceed();

        stopWatch.stop();
        logger.info(stopWatch.prettyPoint());

        return ret;
    }
}
```

# Reference
[1] AOP  소개, https://www.inflearn.com/course/spring/lecture/15543?tab=curriculum&volume=1.00<br>
[2] AOP 적용 예제, https://www.inflearn.com/course/spring/lecture/15544?tab=curriculum&volume=1.00<br>
[3] [Spring Boot] AOP, Object Mapper, 어노테이션 정리, https://velog.io/@malgum/Spring-Boot-AOP