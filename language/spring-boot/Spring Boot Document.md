# Spring Boot

# Developing with Spring Boot

## 6.6. Using the @SpringBootApplication Annotation

@SpringBootApplication 어노테이션을 이용하여 다음과 같은 설정들을 활성화할 수 있습니다.

- @EnabledAutoConfiguration
  - 스프링 부트의 자동 설정 메커니즘을 활성화
- @ComponentScan
  - 애플리케이션이 위치한 패키지에 @Component 스캔을 활성화
- @SpringBootConfiguration
  - 컨텍스트에 별도로 빈 등록이나 추가 설정 클래스 임포넌트를 활성화

## 6.8. Developer Tools

- 스프링 부트는 애플리케이션 개발을 쉽게할 수 있도록 개발 툴 셋을 제공하고 있습니다.

```java
// Maven
<dependencies>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-devtools</artifactId>
        <optional>true</optional>
    </dependency>
</dependencies>

// Gradle
dependencies {
    developmentOnly("org.springframework.boot:spring-boot-devtools")
}
```

6.8.2. Automatic Restart

- sprin-boot-devtools 를 사용하는 애플리케이션의 경우 class path 에 있는 파일이 변경될 때마다 자동으로 재시작됩니다. 이로 인해 static asset 이나 view template 같은 특정 리소스들은 애플리케이션을 재시작할 필요가 없습니다

Restart vs Reload

- 스프링 부트에서 애플리케이션 재시작시 두 가지 클래스 로더를 이용하는 기술을 사용합니다. 변경되지 않는 클래스는 base class loader 에서 로드하고, 현재 개발 중인 클래스는 restart class loader 에 로드합니다.
- 애플리케이션 재시작시 기존 restart class loader 는 폐기되고, 새로운 class loader 를 생성합니다. 이렇게 되면 base class loader 는 이미 로드한 상태로 준비되어 있기 때문에, 애플리케이션 재시작은 보통 “clod start” 보단 훨씬 빠릅니다.
- 하지만 재시작이 빠르지 않거나 class loading 이슈를 만나게 된다면 JRebel 과 같은 리로딩 기술을 고려하는 경우도 있습니다.

  6.8.3. LiveReload

- spring-boot-devtools 모듈에는 리소스가 변경되면 브라우저 새로 고침을 트리거할 수 있는 임베디드 LiveReload 서버가 포함되어 있습니다.

# Reference

[1] 토리맘의 한글라이즈 프로젝트 - Spring Boot, [https://godekdls.github.io/Spring Boot/contents/](https://godekdls.github.io/Spring%20Boot/contents/)
