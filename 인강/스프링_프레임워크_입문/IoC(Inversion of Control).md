# IoC(Inversion of Control)
- 잘 만든 인터페이스
- 나의 코드, 확장성이 좋지 못한 코드, 기술에 특화되어 있는 코드는 문제가 있다!


## IoC 소개
- 스스로 객체를 제어하는 것이 아니라 프레임워크 등에서 객체에 의존성을 주입하도록 하는 방법

## IoC 컨테이너
- ApplicationContext
- Bean 을 만들어주고 엮어주며 제공해준다.

## Bean
- 스프링 IoC 컨테이너가 관리하는 객체

- @Component
  - @Repository
  - @Service
  - @Controller
  - ...

## DI(Dependency Injection)
- 필요한 의존성을 어떻게 받을 수 있을까?

- @Autowired | @Inject ...

- DI 권장하는 우선 순위
  - 생성자 (1순위) 
  - 필드 (3순위)
  - Setter (2순위)

## 생성자 의존성 주입 예제
``` java
@Controller
public class OwnerController {
    private final OwnerRepository owners;

    public OwnerController(ownerRepository clinicService) {
            this.owners = clinicService;
    }
}
```