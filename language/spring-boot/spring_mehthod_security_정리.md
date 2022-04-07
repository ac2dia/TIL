# Introduction to Spring Method Security

## Enabling Method Security

```java
@Configuration
@EnableGlobalMethodSecurity(
  prePostEnabled = true,
  securedEnabled = true,
  jsr250Enabled = true)
public class MethodSecurityConfig
  extends GlobalMethodSecurityConfiguration {
}
```

## Applying Method Security

### @Secured

```java
@Secured({ "ROLE_VIEWER", "ROLE_EDITOR" })
public boolean isValidUsername(String username) {
    return userRoleRepository.isValidUsername(username);
}
```

- @Secured 애노테이션의 경우 {} 내부에 정의된 ROLE 을 모두 가지고 있어야만합니다.

### @RoleAllowed

```java
@RolesAllowed({ "ROLE_VIEWER", "ROLE_EDITOR" })
public boolean isValidUsername2(String username) {
    //...
}
```

- @RoleAllowed 애노테이션의 경우 {} 내부에 정의된 ROLE 중 하나 이상 가지고 있으면됩니다.

### @PreFilter and @PostFilter

```java
@PostFilter("hasRole('MANAGER') or filterObject.assignee == authentication.name")
List<Task> findAll() {
    // ...
}

@PreFilter("hasRole('MANAGER') or filterObject.assignee == authentication.name")
Iterable<Task> save(Iterable<Task> entities) {
    // ...
}
```

- @PreFilter, @PostFilter 애노테이션을 사용하면 SpEL을 사용하여 세분화된 보안 규칙을 정의할 수 있습니다.
- @PostFilter 의 경우 작업 이 후 반환되는 데이터 Object 를 검사하여 조건 일치하는 경우 정상적으로 리턴
- @PreFilter 의 경우 작업 이 전 들어온 데이터 Object 를 검사하여 조건 일치하는 경우 정상적으로 메서드 실행

### Meta Annotation

```java
@Target(ElementType.METHOD)
@Retention(RetentionPolicy.RUNTIME)
@PreAuthorize("hasRole('VIEWER')")
public @interface IsViewer {
}

@IsViewer
public String getUsername4() {
    //...
}
```

- 용도에 맞게 커스텀하게 애노테이션을 만들어 사용할 수 있습니다.

### Security Annotation at the Class Level

```java
@Service
@PreAuthorize("hasRole('ROLE_ADMIN')")
public class SystemService {

    public String getSystemYear(){
        //...
    }

    public String getSystemDate(){
        //...
    }
}
```

- method 뿐만 아니라 class 에도 Security 관련 애노테이션을 사용할 수 있습니다.

## Important Considerations

1. 기본적으로 Spring AOP proxing 은 method security 를 적용하는데 사용됩니다.
2. Spring SecurityContext 는 thread-bound 합니다. 기본적으로 SecurityContext 는 child threads 를 전파하지 않습니다.

# Reference

[1] [https://www.baeldung.com/spring-security-method-security](https://www.baeldung.com/spring-security-method-security)
