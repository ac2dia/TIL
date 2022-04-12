# Java 8 new features and enhancements

## Table of Contents

```
1. Lambda Expressions
    1.1. Syntax
    1.2. Rules for writing lambda expressions

2. Functional Interfaces
    2.1. Functional Interface Example

3. Default Methods

4. Java 8 Streams
    4.1. Stream API Example

5. Java 8 Date/Time API Changes
    5.1. Date Classes
    5.2. Timestamp and Duration Classes
```

## 1. Lambda Expressions

### 1.1. Syntax

```java
(parameters) -> expression

(parameters) -> { statements; }

() -> expression

//This function takes two parameters and return their sum
(x, y) -> x + y
```

### 1.2. Rules for writing lambda expressions

- 람다 표기식은 0, 1 또는 이상의 파라미터를 가질 수 있습니다.
- 파라미터 타입은 명시적으로 선언되거나 문맥에 의해 추론될 수 있습니다.
- 파라미터들은 괄호 () 로 묶이며 , 로 구분되어집니다.
- 파라미터가 한 개이며 타입이 추론 가능한 경우 괄호 () 생략이 가능합니다.
- 본문에는 0, 1 또는 이상의 상태가 포함될 수 있습니다.

## 2. Functional Interfaces

```java
//Optional annotation
@FunctionalInterface
public interface MyFirstFunctionalInterface {
    public void firstWork();
}
```

- Functional Interface 는 Single Abstract Method interfaces (SAM Interfaces 라고도 불립니다.
- 오직 단 한개의 추상 메서드만을 가지고 있습니다.
- Java8 에서는 @FunctionalInterface 라는 어노테이션을 이용하여 표기합니다.

```java
@FunctionalInterface
public interface MyFirstFunctionalInterface
{
    public void firstWork();

    @Override
    public String toString();                //Overridden from Object class

    @Override
    public boolean equals(Object obj);        //Overridden from Object class
}
```

- Functional Interface 의 경우 java.lang.Object 의 public method 들을 override 해서 사용할 수 있습니다.
- 위 경우는 추상 메서드로 카운팅하지 않기 때문에 Functional Interface 인 상태로 유지됩니다.

## 3. Default Methods

```java
public interface Moveable {
    default void move(){
        System.out.println("I am moving");
    }
}

public class Animal implements Moveable{
    public static void main(String[] args){
        Animal tiger = new Animal();
        tiger.move();
    }
}

Output: I am moving
```

- interface 내부에 default 키워드가 붙은 메서드는 내부에 미리 구현을 할 수 있으며, 해당 interface 를 상속받는 class 는 @Override 해서 사용하거나 변경하지 않고 사용할 수 있습니다.

## 4. Java 8 Streams

```java
List<String> items = ...; //Initialize the list

String prefix = "str";

List<String> filteredList = items.stream()
          .filter(e -> (!e.startsWith(prefix)))
          .collect(Collectors.toList());
```

- filtering, transformation, sort 등 Collection 의 데이터를 다양한 방법으로 다룰 수 있습니다.

## 5. Java 8 Date/Time API Changes

### 5.1. Date Classes

- Date Class 를 3가지 형태로 변환할 수 있습니다.
- LocalDate = date
- LocalTime = time
- LocalDateTime = date and time

### 5.2. Timestamp and Duration Classes

```java
Instant instant = Instant.now();
Instant instant1 = instant.plus(Duration.ofMillis(5000));
Instant instant2 = instant.minus(Duration.ofMillis(5000));
Instant instant3 = instant.minusSeconds(10);
```

- 어느 순간의 timestamp 를 표현하기 위해 Instant 라는 클래스가 사용되어야 합니다.
- Instant Class 는 nano sec 단위로 매우 정확하게 시간을 표현할 수 있습니다.

```java
Duration duration = Duration.ofMillis(5000);
duration = Duration.ofSeconds(60);
duration = Duration.ofMinutes(10);
```

- Duration Class 는 milli sec, sec, min, hrs 와 같은 시간을 다루는데 사용됩니다.

```java
Period period = Period.ofDays(6);
period = Period.ofMonths(6);
period = Period.between(LocalDate.now(), LocalDate.now().plusDays(60));
```

- Period Class 는 days, month ... 와 같은 Duration 이상의 시간을 다루는데 사용됩니다.
