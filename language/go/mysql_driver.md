# Go Mysql Driver

- golang 에서 Mysql 을 사용하기 위해서는 "database/sql", "github.com/go-sql-driver/mysql" 패키지를 import 해야합니다.

## Mysql Query - Single Row

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 하나의 Row를 갖는 SQL 쿼리
    var name string
    err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name)
```

- sql 패키지의 Open 메서드를 사용하여 db 에 연결을 시도합니다.
- defer db.Close() 를 이용하여 문제가 생기는 경우 항상 close 될 수 있도록 합니다.
- QueryRow 메서드의 파라미터로 실제 sql query 문을 작성하여 결과를 조회합니다.

## Mysql Query - Multi Row

```go
func main() {
    // sql.DB 객체 생성
    db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 복수 Row를 갖는 SQL 쿼리
    var id int
    var name string
    rows, err := db.Query("SELECT id, name FROM test1 where id >= ?", 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close() //반드시 닫는다 (지연하여 닫기)

    for rows.Next() {
        err := rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name)
    }
}

```

- 복수개의 row 를 리턴하는 경우 리스트로 담기며 Next 메서드를 이용하여 순서대로 데이터를 조회합니다.
- Query 메서드에서 ? (placeholder) 를 사용하여 parameterized query 를 사용하고 있으며, SQL Injection 을 방지할 수 있습니다.
- 또한 placeholder ? 에는 1이 대입되며, 데이터베이스의 종류에 따라 다른 표기를 사용합니다.
  - Mysql => ?
  - Oracle => :val1, :val2
  - PostgreSQL => $1, $2
