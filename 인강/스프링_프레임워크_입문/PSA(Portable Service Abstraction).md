# PSA(Portable Service Abstraction)
- 잘 만든 인터페이스
- 나의 코드, 확장성이 좋지 못한 코드, 기술에 특화되어 있는 코드는 문제가 있다!

## Resource 추상화
- Resource 어떻게 읽어들일 것이냐?

## i18n
- message 를 어떻게 읽어들일 것이냐?

## Spring Transaction
- @Transactional
- Platform TransactionManager
  - JpaTransactionManager | Datasource--- | Hibernate---

## Cache
- @EnableCaching
- @Cacheable | @CacheEvict | ...
- CacheManager
  - JCacheManager | ConcurrentMapCacheManager | EhCacheCacheManager | ...

## Web MVC
- @PostMapping | @GetMapping | ...
- 해당 애노테이션을 통해서는 Servlet 기반인지, Reactive 기반인지 알 수 없다!

## 결론
- 잘 만들어진 추상화를 이용한다면 나의 코드가 변하더라도 소스에 영향을 끼치지 않는다!
- 또한 Platform, 기술 등이 변경되어도 소스에 영향을 끼치지 않는다!
