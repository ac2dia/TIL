# [Cheat Sheet] Updating Collection in React State 

## Arrays
``` javascript
const [todos, setTodos] = useState([]);
```

``` javascript
// Add to array
// 1번 방법
const handleAdd = (todo) => {
  const newTodos = todos.slice();
  newTodos.push(todo);
  setTodos(newTodos);
}

// 2번 방법
const handleAdd = (todo) => {
    setTodos([...todos, todo]);
}
```
- 2번 방법의 경우 "syntactic sugar" = "보다 코드를 쉽게 읽을 수 있도록 하는 행위" 에 해당합니다.

``` javascript
// Remove from array
// 1번 방법
const handleRemove = (todo) => {
  const newTodos = todos.filter((t) => t !== todo);
  setTodos(newTodos);
}

// 2번 방법
const handleRemove = (todo) => {
  setTodos(todos.filter((t) => t !== todo));
}

// Update array
const handleUpdate = (index, todo) => {
  const newTodos = [...todos];
  newTodos[index] = todo;
  setTodos(newTodos);
}
```

- 배열뿐만 아니라 객체도 동일한 방법을 사용합니다.


# Reference
[1] Grepper.com, https://www.codegrepper.com/code-examples/javascript/how+to+add+object+list+into+array+in+react
[2] https://dev.to/andyrewlee/cheat-sheet-for-updating-objects-and-arrays-in-react-state-48np