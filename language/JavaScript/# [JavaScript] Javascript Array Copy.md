# [JavaScript] Javascript Array Copy

- 배열의 얕은 복사: 배열 내에 배열이 있는 경우에 변경이 이뤄지면 같이 값이 변합니다.
- 배열의 깊은 복사: 어느 한 쪽의 수정으로 다른 쪽에 영향이 없습니다.

## Shallow Copy

### 1. Spread Operator (Shallow Copy)

```javascript
const array = [1, 2, [3]];
const array_copy = [...array];
console.log(array, array_copy);
// [1, 2, [3]]   [1, 2, [3]]

array.push(4);
console.log(array, array_copy);
// [1, 2, [3], 4]   [1, 2, [3]]

array[2].push(5);
console.log(array, array_copy);
// [1, 2, [3, 5], 4]   [1, 2, [3, 5]]
// 배열 내에 있는 배열을 수정 할 경우에 같이 값이 변경된다.
```

### 2. for문을 사용한 배열 복사 (Shallow Copy)

```javascript
const array = [1, 2, [3]];
const array_copy = [];
for (let i = 0; i < array.length; i++) {
  array_copy.push(array[i]);
}
console.log(array, array_copy);
// [1, 2, [3]]   [1, 2, [3]]

array.push(4);
console.log(array, array_copy);
// [1, 2, [3], 4]   [1, 2, [3]]

array[2].push(5);
console.log(array, array_copy);
// [1, 2, [3, 5], 4]   [1, 2, [3, 5]]
```

### 3. Array.filter()을 사용한 배열 복사 (Shallow Copy)

```javascript
const array = [1, 2, [3]];
const array_copy = array.filter(() => true);
console.log(array, array_copy);
// [1, 2, [3]]   [1, 2, [3]]

array.push(4);
console.log(array, array_copy);
// [1, 2, [3], 4]   [1, 2, [3]]

array[2].push(5);
console.log(array, array_copy);
// [1, 2, [3, 5], 4]   [1, 2, [3, 5]]
```

## Deep Copy

### 1. JSON.parse와 JSON stringify을 사용한 배열 복사 (Deep Copy)

```javascript
const array = [1, 2, [3]];
const array_copy = JSON.parse(JSON.stringify(array));
console.log(array, array_copy);
// [1, 2, [3]]   [1, 2, [3]]

array.push(4);
console.log(array, array_copy);
// [1, 2, [3], 4]   [1, 2, [3]]

array[2].push(5);
console.log(array, array_copy);
// [1, 2, [3, 5], 4]   [1, 2, [3]]
// 깊은 복사를 사용하여 배열 복사를 할 경우에는 한 쪽에 어떤 짓을 하더라도 다른 한 쪽에 영향을 끼치지 않는다.
profile;
```

## 정리

- 배열 형태의 자료 구조라면 배열 내 원소가 Value / Object 와는 관계없이 Shallow Copy 시에 동일하다.
