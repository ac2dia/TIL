# [Go] Go Reflect 

## type
- 속성들의 집합을 표현하는 struct
- 행위의 집합을 표현하는 interface

## interface

``` go
type Rectangle struct {
    width, height int
}

func (r *Rectangle) Area() int {
    return r.width * r.height
}
```
- r 의 경우 "method receiver" 역할을 함

``` go
type Shape interface {
    Area() int
    Perimeter() float64
}

// --------------------------------
type rectangle struct {
    W, H float64
}

type circle struct {
    R float64
}

// --------------------------------
func (r rectangle) Area() int {
    return r.W * r.H
}

func (r rectangle) Perimeter() floa64 {
    return 2 * (r.H + r.W)
}

func (c circle) Area() int {
    return c.R * c.R * math.Pi
}

func (c circle) Perimeter() floa64 {
    return 2 * c.R * math.Pi
}

// --------------------------------
func Calc(s shape) {
    switch s.(type) {
	case rectangle:
		fmt.Println("I'm a rectangle")
	case circle:
		fmt.Println("I'm a circle")
	}
	fmt.Printf("area: %d, perimeter: %.2f", s.Area(), s.Perimeter())
}

```
- 특정 instance 를 만들기 위한 interface 를 만드느 경우 반드시 공통으로 구현할 행위(method) 에 대해서만 선언한다.

## reflect
- 구조체에 대한 정보를 동적으로 알아내는데 사용되는 golang 기능
- 향후 새로운 type 이 추가되거나, 처리 대상 type 을 예측할 수 없는 상황에서 매우 유용합니다!
- 대표적인 예로 "fmt" 패키지가 있습니다.



# Reference
[1] "[golang] interface와 reflect, 그리고 OOP", https://incredible-larva.tistory.com/entry/golang-interface%EC%99%80-reflect-%EA%B7%B8%EB%A6%AC%EA%B3%A0-OOP
