# [React] useEffect 깔끔하게 마스터하기

## useEffect

- Mount Update, Unmount
- Callback 함수를 통해 호출

```javascript
// 렌더링 될 때마다 실행
useEffect(() => {
  // 작업 ...
});

// 초기 렌더링
useEffect(() => {}, []);

// 초기 렌더링 + 배열 내부 요소가 변경될 때
useEffect(() => {}, [요소]);

// mount + unmount
useEffect(() => {
  // 구독 ...

  return () => {
    // 구독 해제
    // unmount or 다음 effect 이전
  };
}, []);
```

# Reference

[1] React Hooks에 취한다 - useEffect 깔끔하게 마스터하기 | 리액트 훅스 시리즈, https://youtu.be/kyodvzc5GHU
