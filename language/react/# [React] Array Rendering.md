# [React] Array Rendering

## Array Rendering

```javascript
import React from 'react';

function User({ user }) {
  return (
    <div>
      <b>{user.username}</b> <span>({user.email})</span>
    </div>
  );
}

function UserList() {
  const users = [
    {
      id: 1,
      username: 'velopert',
      email: 'public.velopert@gmail.com',
    },
    {
      id: 2,
      username: 'tester',
      email: 'tester@example.com',
    },
    {
      id: 3,
      username: 'liz',
      email: 'liz@example.com',
    },
  ];

  return (
    <div>
      {users.map((user) => (
        <User user={user} key={user.id} />
      ))}
    </div>
  );
}

export default UserList;
```

- 각 고유 원소에 key 가 있어야만 배열이 업데이트 될 때 효율적으로 렌더링 될 수 있기 때문입니다.

  - 언제? 업데이트 될 때! / 왜? 인덱스에 효율적으로 접근할 수 있기 때문!

- 배열에 key 가 없다면 특정 원소를 찾기 위해 순차적으로 탐색하여 O(n) 의 시간 복잡도를 가지지만, key 로 인해서 해당 인덱스를 탐색하는 경우 O(1) 의 시간 복잡도를 가집니다.

- 또한 가장 중요한 Point!! 배열 내에 중복된 key 값을 가지는 경우는 렌더링시에 오류가 발생하게 됩니다.

# Reference

[1] 11. 배열 렌더링하기, https://react.vlpt.us/basic/11-render-array.html
