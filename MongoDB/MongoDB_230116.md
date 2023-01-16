# MongoDB 개념

## NoSQL

- NoSQL은 Not Only SQL, SQL 뿐만 아니다라는 의미를 지니고있다. 즉, SQL을 사용하는 관계형 데이터베이스가 아닌 데이터베이스를 의미한다. 대표적인 관계형 데이터베이스로는 MySQL, Oracle, PostgreSQL이 있고, NoSQL 진영에는 이 포스트에서 다루는 MongoDB와 Redis, HBase 등이 있다. 더 다양한 NoSQL 제품을 보고 싶다면 링크를 들어가면 지금까지 출시된 NoSQL 데이터베이스 목록을 볼 수 있다.

## MongoDB 특징

- MongoDB는 앞서 설명한 것 처럼 NoSQL 데이터베이스고 다음 세 가지 특징을 가지고있다.

- Document
- BASE
- Open Source

![MongoDB 구조](./MongoDB_structure.jpeg)

## BASE

BASE는 ACID와 대립되는 개념으로 다음 세 가지로 이루어져있다.

- Basically Avaliable
  - 기본적으로 언제든지 사용할 수 있다는 의미를 가지고 있다.
    즉, 가용성이 필요하다는 뜻을 가진다.
- Soft state
  - 외부의 개입이 없어도 정보가 변경될 수 있다는 의미를 가지고 있다.
  - 네트워크 파티션 등 문제가 발생되어 일관성(Consistency)이 유지되지 않는 경우 일관성을 위해 데이터를 자동으로 수정한다.
- Eventually consistent
  - 일시적으로 일관적이지 않은 상태가 되어도 일정 시간 후 일관적인 상태가 되어야한다는 의미를 가지고 있다.
  - 장애 발생시 일관성을 유지하기 위한 이벤트를 발생시킨다.

## ACID

- Atomicity (원자성)
- Consistency (일관성)
- Isolation (고립성)
- Durability (지속성)

## 마무리!

- MongoDB와 RDBMS는 적합한 사용처가 다르다.
- NoSQL은 최대한 단순하면서 많은 데이터, RDBMS는 복잡하면서 무결성이 중요한 데이터에 적합하다고 생각한다.
- 물론 데이터를 단순화하는 것도 쉬운 일은 아니기 때문에 만약 당신이 MongoDB를 사용할 계획이 있다면 꼭 위 모델링 패턴을 참고하여 데이터 구조를 잡는 것을 추천한다.
