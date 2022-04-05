# React Hook 학습

## Hook

- Hook은 함수 컴포넌트에서 React state와 생명주기 기능(lifecycle features)을 “연동(hook into)“할 수 있게 해주는 함수입니다.
- Hook은 class 안에서는 동작하지 않습니다. 대신 class 없이 React를 사용할 수 있게 해주는 것입니다.
- 하지만 이미 짜놓은 컴포넌트를 모조리 재작성하는 것은 권장하지 않습니다. 대신 새로 작성하는 컴포넌트부터는 Hook을 이용하시면 됩니다.

## useState (State Hook)

- 가장 기본적인 hook
- 해당 기능 사용시 컴포넌트가 다시 렌더링되어도 state 값이 그대로 유지된다.
- useState는 현재의 값을 저장하는 기능과 업데이트하는 함수를 쌍으로 제공한다.
- 기본적으로 하나의 상태 값만 관리할 수 있다.

## useEffect (Effect Hook)

- Effect Hook을 사용하면 함수 컴포넌트에서 side effect를 수행할 수 있습니다.
- React 는 매 렌덩링 이후에 effects 를 실행하기 때문에, DOM을 변경할 시 effect 함수를 실행합니다.
- Effects 는 컴포넌트 안에 선언되어있기 때문에 props와 state에 접근할 수 있습니다.
- useState 와 마찬가지로 여러 개의 effect 를 사용할 수 있습니다.

## Hook 사용 규칙

- Hook은 그냥 JavaScript 함수이지만, 두 가지 규칙을 준수해야 합니다.

- 최상위(at the top level)에서만 Hook을 호출해야 합니다. 반복문, 조건문, 중첩된 함수 내에서 Hook을 실행하지 마세요.
- React 함수 컴포넌트 내에서만 Hook을 호출해야 합니다. 일반 JavaScript 함수에서는 Hook을 호출해서는 안 됩니다. (Hook을 호출할 수 있는 곳이 딱 한 군데 더 있습니다. 바로 직접 작성한 custom Hook 내입니다. 이것에 대해서는 나중에 알아보겠습니다.)

# Reference

[1] Hook의 개요, https://ko.reactjs.org/docs/hooks-intro.html
