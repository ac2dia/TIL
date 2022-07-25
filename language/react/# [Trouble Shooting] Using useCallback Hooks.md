# [Trouble Shooting] Using useCallback Hooks - Click Event

``` javascript
    const onIncrease = useCallback(
    (user) => {
      setMembers(() => [...members, user]);
      setAvailableUsers(
        availableUsers.filter((availableUser) => availableUser.label !== user.label),
      );
    },
    [availableUsers, members], // 1번
  );

  const onDecrease = useCallback(
    (user) => {
      setAvailableUsers([...availableUsers, user]);
      setMembers(members.filter((member) => member.label !== user.label));
    },
    [availableUsers, members],
  );
```

- 1번
    - deps 의 경우 변경 대상을 지정해야 합니다! 저의 경우 set 함수를 지정하였기 때문에 이슈가 발생하였습니다