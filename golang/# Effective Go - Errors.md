# Effective Go - Errors

## Error

- Go Library 들은 대게 error 를 리턴하는 형식으로 되어 있습니다.
- Go multi value return 으로 인해서 일반적인 caller 의 반환 값 뿐만 아니라 상세한 에러 문구를 같이 반환하는 것도 쉬워졌습니다.

```go
type error interface {
    Error() string
}
```

- error interface 의 경우 다음과 같은 형태로 되어 있습니다.

```go
for try := 0; try < 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    }
    if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
        deleteTempFiles()  // Recover some space.
        continue
    }
    return
}
```

- 해당 코드와 같이 파일을 생성하였을 때 err 반환 값이 존재하는 경우 에러에 대한 처리를 할 수 있도록 합니다.

## Panic

- error 에 대해 리포팅하는 보편적인 방법은 error 를 extra return value 로써 반환하는 것입니다.
  - extra return value <- 아마 코드의 실행이 끝난 마지막 라인을 의미?

````go
// A toy implementation of cube root using Newton's method.
func CubeRoot(x float64) float64 {
    z := x/3   // Arbitrary initial value
    for i := 0; i < 1e6; i++ {
        prevz := z
        z -= (z*z*z-x) / (3*z*z)
        if veryClose(z, prevz) {
            return z
        }
    }
    // A million iterations has not converged; something is wrong.
    panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}

``` go
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
````

## Recover

- panic() 이 호출되었을 때, 즉시 실행중인 코드를 멈춥니다.

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

# Reference

[1] Effective_go - errors, https://go.dev/doc/effective_go#errors
