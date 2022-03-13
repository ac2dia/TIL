# JSON Data Type

- MariaDB 에서 JSON 타입의 경우 10.2.7 버전 이후로 추가되었습니다.

## json data type 을 가지는 table 생성
``` sql
CREATE TABLE t (j JSON);

DESC t;
+-------+----------+------+-----+---------+-------+
| Field | Type     | Null | Key | Default | Extra |
+-------+----------+------+-----+---------+-------+
| j     | longtext | YES  |     | NULL    |       |
+-------+----------+------+-----+---------+-------+
```

- mariadb 에서 JSON 의 경우 longtext 타입으로 생성됩니다.


## json data type 
``` sql
CREATE TABLE t2 (
  j JSON 
  CHECK (JSON_VALID(j))
);

INSERT INTO t2 VALUES ('invalid');
ERROR 4025 (23000): CONSTRAINT `j` failed for `test`.`t2`

INSERT INTO t2 VALUES ('{"id": 1, "name": "Monty"}');
Query OK, 1 row affected (0.13 sec)
```

- CHECK (JSON_VALID($COULMN_NAME)) 을 통해 json format 의 데이터가 아닌 경우 ERROR 가 발생합니다.

## 언제 사용하면 좋을까?
- 빠른 조회보다 데이터 저장에 중점을 두는 데이터
- 변화가 적고, 정형화되어 있지 않으며 다양한 형식의 변화가 요구되는 데이터
- 지속적인 데이터의 형식에 변화가 가능한 데이터

## 의문점
- JSON과 같은 비정형 데이터를 다루는 경우 NoSQL을 이용하는 경우가 많은데 RDB를 사용해야 얻는 장점이 있을까?

# Reference
[1] JSON Data Type, https://mariadb.com/kb/en/json-data-type/
[2] (꿀팁) Mysql에서 JSON 저장과 조회하기, https://couplewith.tistory.com/entry/%EA%BF%80%ED%8C%81-Mysql%EC%97%90%EC%84%9C-Json-%EC%A0%80%EC%9E%A5%EA%B3%BC-%EC%A1%B0%ED%9A%8C%ED%95%98%EA%B8%B0
[3] JSON + MariaDB: How Hybrid Data Models Allow You to Have Your Cake and Eat It Too, https://dev.to/probablyrealrob/json-mariadb-how-hybrid-data-models-allow-you-to-have-your-cake-and-eat-it-too-282i