# Innogrid IaaS 1팀 - FE 파트 면접

## 1. SPA (Single Page Application) 의 장점에 대해서 설명해주세요

- Client Side Rendering 로 서버에 부담이 없고, 속도가 좀 더 빠르다.
- ~

## 2. VirtualDOM 에 대해서 설명해주세요.

- DOM은 HTML 문서를 파싱하여 "문서의 구성요소들을 객체로 구조화하여 나타낸 것"
- DOM을 수정하는 일은 수반되는 비용이 크기 때문에, 성능저하를 최소화하기 위해서는 결국 DOM을 최소한으로 수정해야한다.
  - DOM을 추상화한 후, 변경 사항에 대해서만 실제 DOM에 업데이트를 통해 최적화
- 이러한 문제점을 해결하기 위해 Virtual DOM이 등장하게 되었다.

## 3. 클래스형 컴포넌트와 함수형 컴포넌트의 차이점에 대해서 설명해주세요.

- 기존에는 함수형 컴포넌트에서는 불가능하고, 클래스형 컴포넌트에서만 라이프 사이클 및 상태를 관리하였습니다.
- react Hook 의 도입으로 함수형 컴포넌트에서도 라이프 사이클(useEffect) 및 상태 관리(useState) 를 할 수 있습니다.

### Hook

- Hook은 함수 컴포넌트에서 React state와 생명주기 기능(lifecycle features)을 “연동(hook into)“할 수 있게 해주는 함수입니다.
- Hook은 class 안에서는 동작하지 않습니다. 대신 class 없이 React를 사용할 수 있게 해주는 것입니다.
- 하지만 이미 짜놓은 컴포넌트를 모조리 재작성하는 것은 권장하지 않습니다. 대신 새로 작성하는 컴포넌트부터는 Hook을 이용하시면 됩니다.

### useState (State Hook)

- 가장 기본적인 hook
- 해당 기능 사용시 컴포넌트가 다시 렌더링되어도 state 값이 그대로 유지된다.
- useState는 현재의 값을 저장하는 기능과 업데이트하는 함수를 쌍으로 제공한다.
- 기본적으로 하나의 상태 값만 관리할 수 있다.

### useEffect (Effect Hook)

- Effect Hook을 사용하면 함수 컴포넌트에서 side effect를 수행할 수 있습니다.
- React 는 매 렌덩링 이후에 effects 를 실행하기 때문에, DOM을 변경할 시 effect 함수를 실행합니다.
- Effects 는 컴포넌트 안에 선언되어있기 때문에 props와 state에 접근할 수 있습니다.
- useState 와 마찬가지로 여러 개의 effect 를 사용할 수 있습니다.

## 4. react 의 state 와 props 의 특징 및 차이점에 대해서 설명해주세요.

- props는 부모 컴포넌트에서 자식 컴포넌트로 전달되는 데이터입니다.
- props는 수정될 수 없으며 표시되거나 다른 값을 계산하는데만 사용됩니다.

- state는 컴포넌트의 생명 주기 동안 수정될 수 있는 내부 데이터로, 다시 렌더링해도 유지됩니다.
- 또한, setState() 함수를 사용하여 값을 변경할 수 있으며, 직접 state 의 값을 변경하는 경우 react 에서 감지할 수 없습니다.

- 상태(state)란? 동적인 데이터를 다루기 위한 객체로 컴포넌트 내부에서 변경될 수 있는 값.
- props는 (함수 매개변수처럼) 컴포넌트에 전달되는 반면, state는 (함수 내 선언된 변수처럼) 컴포넌트 안에서 관리된다는 차이가 있다.

## 5. prop drilling 이란 무엇이고, 이로 인한 문제점과 해결 방법에 대해서 설명해주세요.

- prop drilling은 부모 컴포넌트에서 하위 컴포넌트(자식 컴포넌트의 자식 컴포넌트 등으)로 데이터를 전달할 때 발생하는 것으로, props를 전달하는 것 외에는 props를 필요로 하지 않는 다른 컴포넌트를 통해 “drilling”(내리꽂기) 됩니다.

- 문제점으로는 컴포넌트 depth가 깊어질 수록 해당 prop을 추적하기 힘들어져서 유지 보수가 어려워집니다.

- 해결방법
  - 전역 상태관리 라이브러리 사용 (redux, MobX, recoil, context API 등)
  - children 활용

## 5. 자바스크립트를 이용한 비동기 처리 방법에 대해서 설명해주세요.

- 콜백 함수 사용
- Promise
  - 대기(pending), 이행(fulfilled), 거부(rejected) 중 하나의 상태를 가진다.
  - new Promise(function(resolve, reject) {});
  - resolve 는 결과가 성공인 Promise 객체, reject 는 결과가 실패인 Promise 객체
- Promise를 활용한 async/await
  - ES2017 등장

## 6. 프로젝트를 진행하면서 마주한 이슈와 어떤 방식으로 해결하였는지 경험을 공유해주세요.

## 7. 소속되었던 팀 내부적 혹은 외부적으로 지식 공유나 지식 전파 같은 활동들을 해본 적이 있는지?
