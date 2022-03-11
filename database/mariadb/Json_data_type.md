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

## 유용한 사용 방법?
- key:value 형식의 데이터를 저장하려는 경우 사용
- 애플리케이션 설정 정보를 담고 있는 테이블?
- 자원 정보를 담고 있는 테이블?

# Reference
[1] JSON Data Type, https://mariadb.com/kb/en/json-data-type/