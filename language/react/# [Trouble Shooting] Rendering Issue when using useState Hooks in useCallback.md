# [Trouble Shooting] Rendering Issue when using useState Hooks in useCallback

## 문제 상황
``` javascript
const onIncrease = useCallback(
    (users) => {asdasd
      const newMembers = members;

      // eslint-disable-next-line array-callback-return
      users.map((user) => {
        newMembers.push(user);
      });
sd
      setMembers(newMembers);
    },
    [members],
  );
```
- useState 를 통해 생성된 변수에 ... (spread operator) 를 사용하지 않고 한번에 setState 를 한 경우는 문제가 없습니다.


``` javascript
const onIncrease = useCallback(
    (users) => {
      const newMembers = [];

      // eslint-disable-next-line array-callback-return
      users.map((user) => {
        newMembers.push(user);
      });

      setMembers({...members, newMembers});

      handleAllCheck(false, true);
      handleAllCheck(false, false);
    },
    [members],
  );
```

- useState 를 통해 생성된 변수에 ... (spread operator) 를 사용한 경우 useCallback 에 의한 렌더링이 정상적으로 이루어지지 않았습니다.

## 해결 방안
- spread operator 내부 동작에 대해 확인해보았지만 원인을 확인할 수 없었습니다.
- !! 우선 1번 방법을 통해 해결이 되어 추후 재확인이 필요합니다.