# [JavaScript] The JavaScript Handbook

## 7. Variables

- const
  - 초기화 이 후 변경이 불가능한 변수
- let

  - 초기화 이 후에도 변경이 가능한 변수

- var
  - ES2015 이전에 사용된 변수 선언 방식

## 8. Types

### 8.1. Primitive types

- numbers
- strings
- booleans
- symbols

### 8.2. Object types

- properties
- methods

## 14. Arrays

- const a = []
- const a = Array()

- a.length
- a.concat
- a.find((element, index, array) => {
  // return true or false
  })
- a.find((x) => x.id === my_id)
- a.includes(value)

## 16. Loops

- while loops
- for loops
- for..of loops

## 17. Functions

function getData() {
return ['Flavio', 37]
}

let [name, age] = getData()

const getData = () => {
const dosomething = () => {}
dosomething()
return 'test'
}

## 18. Arrow functions

let getData = function () {
//...
}

let getData = () => {
//...
}

# Reference

[1] The JavaScript Handbook, https://thevalleyofcode.com/js/#1-introduction-to-javascript
