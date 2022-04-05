# Innogrid IaaS 1팀 - FE 파트 면접

## 1. 브라우저 렌더링 과정에 대해서 설명해주세요.

- DOM 트리 생성 (HTML을 파싱하여 DOM객체로 이루어진 DOM트리 생성)
- 스타일 규틱 생성 (CSS parser는 inline style과 CSS 코드를 파싱하여 CSSOM 트리 생성)
- 렌더 트리 생성 (DOM과 CSSOM의 정보를 바탕으로 실제 브라우저의 화면에 노출되어야 하는 노드들에 대한 정보인 렌터 트리 생성)
- 레이아웃 -> Reflow - 레이아웃이 변화할때 (새로운 요소 추가, 특정 요소 삭제, 요소 위치 또는 크기 변경)
- 페인트 -> Repaint - CSS Style 이 바뀔 때

## 2. 브라우저에서 성능 저하가 발생하는 원인

- 주요 원인은 DOM을 수정할 때 발생하는 Reflow, Repair 과정

## 3. VirtualDOM

- HTML 문서를 파싱하여 "문서의 구성요소들을 객체로 구조화하여 나타낸 것"
- DOM을 수정하는 일은 수반되는 비용이 크기 때문에, 성능저하를 최소화하기 위해서는 결국 DOM을 최소한으로 수정해야한다.
- 이러한 문제점을 해결하기 위해 Virtual DOM이 등장하게 되었다.

## 4. React 의 useState, useEffect Hook 에 대해서 설명해주세요.

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
