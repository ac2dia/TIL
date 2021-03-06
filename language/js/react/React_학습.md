# React 기본 학습

## React의 특징

- SPA (Single Page Application)
  - Client Side Rendering
- 선언형 프로그래밍
  - 간결하게 목적만 명시하는 프로그래밍 방법
  - 불필요한 과정을 언급하지 않음
- 컴포넌트 기반
  - 컴포넌트란 독립적으로 존재할 수 있는 UI 요소
  - 페이지를 이루는 구성요소
  - 컴포넌트별 역할 분할로 인한 관심사 분리 (조립식 웹)
  - 컴포넌트와 산탄총 수술?
- Virtual DOM
  - DOM (Document Object Model)
  - 브라우저의 렌더링 과정
    - DOM 트리 생성 (HTML을 파싱하여 DOM객체로 이루어진 DOM트리 생성)
    - 스타일 규틱 생성 (CSS parser는 inline style과 CSS 코드를 파싱하여 CSSOM 트리 생성)
    - 렌더 트리 생성 (DOM과 CSSOM의 정보를 바탕으로 실제 브라우저의 화면에 노출되어야 하는 노드들에 대한 정보인 렌터 트리 생성)
    - 레이아웃 -> Reflow - 레이아웃이 변화할때 (새로운 요소 추가, 특정 요소 삭제, 요소 위치 또는 크기 변경)
    - 페인트 -> Repaint - CSS Style 이 바뀔 때
    - Reflow, Repaint overhead 발생 방지해야함!
- 높은 자유도
  - 프레임워크가 아닌 라이브러리!
  - 기능을 사용자가 원할 때 불러서 자유롭게 사용 가능, 사용자가 코드의 흐름을 통제함
- Learn Once Write Anywhere
  - React Native 를 이용하여 모바일 어플리케이션 제작

# Reference

[1] Hook의 개요, https://ko.reactjs.org/docs/hooks-intro.html
