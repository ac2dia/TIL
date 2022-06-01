# Singleton in Go

- 생성 패턴 중의 하나
- 오직 하나의 객체만 존재하는 것을 보장

## Conceptual Example

- 일반적으로 getInstance 메서드를 이용해서 정의합니다.
- 멀티 goroutine 속에서도 singleton struct 는 동일한 인스턴스를 반환합니다.
- go SingleInstance struct 는 lock 과 함께 생성됩니다.

```go
package main

import (
    "fmt"
    "sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creating single instance now.")
            singleInstance = &single{}
        } else {
            fmt.Println("Single instance already created.")
        }
    } else {
        fmt.Println("Single instance already created.")
    }

    return singleInstance
}
```

- sync Mutext 를 활용하여 동시 접근을 제어합니다.
- nil 검사를 통해 현재 인스턴스가 초기화되었는지를 검사하고, 초기화되어있지 않은 경우 초기화 후 반환합니다.

```go
package main

import (
    "fmt"
)

func main() {

    for i := 0; i < 30; i++ {
        go getInstance()
    }

    // Scanln is similar to Scan, but stops scanning at a newline and
    // after the final item there must be a newline or EOF.
    fmt.Scanln()
}
```

- goroutine 을 이용하여 동시에 singleton 으로 구현된 getInstance 메서드를 호출합니다.
- goroutine 을 사용하더라도 Mutex 에 의해 임계 영역에 대한 동시 접근이 제한됩니다.
- 가장 처음 호출시에만 새로 초기화를 하며 이 후는 초기화된 인스턴스를 반환합니다.
