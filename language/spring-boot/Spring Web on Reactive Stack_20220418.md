# Spring Web on Reactive Stack

- Netty, Undertow ... 등 Non-blocking 서버 위에서 구동하는 Reactive Streams 기반 리액티브 스택 웹 어플리케이션
- spring-webflux 프로젝트

## Overview

- spring-webflux 는 적은 쓰레드로 동시 처리를 제어하고, 적은 하드웨어 리소스로 확장하기 위해 등장하였습니다.
- 이전에도 Non-blocking I/O를 위한 API가 제공되었으나 다른 동기 처리, Blocking 방식을 사용하는 API들로 인해 사용이 어려워 새로운 공통 API 를 만들게 되었습니다.
- 또 다른 이유는 Functional Programming 인데, lambda 표현식으로 인해 자바에서도 함수형 API를 작성할 수 있게 됨에 따라 대중화되었습니다.

### Reactive 란?

- "Reactive" 라는 용어는 변화에 반응하는 것을 중심에 두고 만든 프로그래밍 모델을 의미합니다.
- Non-blocking 은 작업을 기다리기보단 완료되거나 데이터를 사용할 수 있게 되면 반응합니다.
- Spring 에서 "Reactive" 와 관련한 중요 메커니즘 중 Non-blocking back pressure 라는 것이 있는데 동기식 명령 코드에서 Blocking 호출은 호출자를 강제로 기다리게 하는 일종의 back pressure 입니다.
- Reactive stack 은 back pressure 를 통한 비동기 컴포넌트 간의 상호작용을 정의한 간단한 스펙입니다. 이를 사용하는 주 목적은 subscriber 가 publisher 의 데이터 생산 속도를 제어하기 위함입니다.

### Reactive API

- Reactor 는 spring-webflux 가 선택한 Reactive 라이브러리입니다.
- Reactor 는 "Mono" (0~1), "Flux" (0~N) API 타입을 제공하며, 모든 연산자는 Non-blocking back pressure 를 지원합니다.

### Programming Models

- spring-webflux 는 두 가지 프로그래밍 모델을 지원합니다.
- Annotated Controllers
  - spring-mvc 와 동일하며 spring-web 모듈에 있는 같은 어노테이션을 사용합니다.
- Functional Endpoints
  - 경량화된 lambda 기반 Functional Programming 모델로써, 요청을 라우팅해주는 라이브러리, 유틸리티 모음이라고 볼 수 있습니다.
  - 어플리케이션의 시작부터 끝까지 모두 제어 가능합니다.

### Applicability

![MVC and WebFlux 다이어그램]](./spring-mvc-and-webflux.png)

### Servers

- Netty, Undertow 서버 환경에서도 동작하며, spring 을 통해 손 쉽게 어플리케이션을 구동할 수 있다.
- Tomcat, Jetty 는 spring-mvc, spring-webflux 모두에서 사용 가능하며, spring-mvc 의 경우 서블릿의 Blocking I/O 를 사용하며, 어플리케이션에서 필요하다면 서블릿 API 를 직접 사용할 수 있다. 하지만 spring-webflux 의 경우 Non-blocking I/O 로 동작하며, 서블릿 API 는 저수준 어댑터에서 서용하기 때문에 노출되어 있지 않습니다.

### Performance

- 성능상 이점보다는 "Reactive", "Non-blocking" 의 이점은 고정된 적은 Thread, Memory 로도 확장 가능하며, 부하 속에서 어플리케이션 복원 능력 또한 뛰어나다.

### Concurrency Model

- annotated controller 를 사용할 수 있다는 점은 동일해도, 동시성 모델과 Blocking/Thread 기본 전략은 상이하다.
- Invoking a Blocking API / Mutable State / Threading Model / Configuring
