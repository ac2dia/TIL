# 01. JPA 소개

# 목차

- SQL 중심적인 개발의 문제점
- JPA 소개

## SQL 중심적인 개발의 문제점

- CRUD 무한 반복
- 객체 CRUD, 필드 추가
- 객체와 관계형 데이터베이스 사이에서 패러다임의 불일치 발생

## JPA?

- Java Persistance API
- 자바 진영의 ORM 기술 표준
- 인터페이스의 모음
- 구현체로 Hibernate, EclipseLink, DataNucleus 등 이 있음

## ORM?

- Object-relational mapping (객체 관계 매핑)
- 객체는 객체대로 설계

## JPA를 왜 사용해야 하는가?

- SQL 중심적인 개발에서 객체 중심으로 개발 - 생산성
- 유지보수
- 패러다임의 불일치 해결
- 성능
- 데이터 접근 추상화와 벤더 독립성 - 표준

## JPA와 패러다임의 불일치 해결

- 1.JPA와 상속
- 2.JPA와 연관관계
- 3.JPA와 객체 그래프 탐색
- 4.JPA와 비교하기

## JPA의 성능 최적화 기능

- 1. 1차 캐시와 동일성(identity) 보장
- 2. 트랜잭션을 지원하는 쓰기 지연(transactional write-behind)
- 3. 지연 로딩(Lazy Loading)
