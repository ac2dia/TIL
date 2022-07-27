# [JavaScript] Useful String methods

## String Property

```javascript
const name = 'alice';
console.log(name.length); // 5
```

- String.length
  - 문자열 내의 문자 갯수를 반환합니다.

## String Method

- 문자열을 immutable 한 원시 값이기 때문에 String Method 는 언제나 새로운 문자열을 반환합니다.

- indexOf()

  - String 인스턴스에서 특정 문자나 문자열이 처음으로 등장하는 위치의 인덱스를 반환함.

- lastIndexOf()

  - String 인스턴스에서 특정 문자나 문자열이 마지막으로 등장하는 위치의 인덱스를 반환함.

- charAt()

  - String 인스턴스에서 전달받은 인덱스에 위치한 문자를 반환함.

- charCodeAt()

  - String 인스턴스에서 전달받은 인덱스에 위치한 문자의 UTF-16 코드를 반환함. (0 ~ 65535)

- charPointAt()

  - String 인스턴스에서 전달받은 인덱스에 위치한 문자의 유니코드 코드 포인트(unicode code point)를 반환함.

- slice()

  - String 인스턴스에서 전달받은 시작 인덱스부터 종료 인덱스 바로 앞까지의 문자열을 추출한 새 문자열을 반환함.

- substring()

  - String 인스턴스에서 전달받은 시작 인덱스부터 종료 인덱스 바로 앞까지의 문자열을 추출한 새 문자열을 반환함.

- substr()

  - String 인스턴스에서 전달받은 시작 인덱스부터 길이만큼의 문자열을 추출한 새로운 문자열을 반환함.

- split()

  - String 인스턴스에서 구분자(separator)를 기준으로 나눈 후, 나뉜 문자열을 하나의 배열로 반환함.

- concat()

  - String 인스턴스에 전달받은 문자열을 결합한 새로운 문자열을 반환함.

- toUpperCase()

  - String 인스턴스의 모든 문자를 대문자로 변환한 새로운 문자열을 반환함.

- toLowerCase()

  - String 인스턴스의 모든 문자를 소문자로 변환한 새로운 문자열을 반환함.

- trim()

  - String 인스턴스의 양 끝에 존재하는 공백과 모든 줄 바꿈 문자(LF, CR 등)를 제거한 새로운 문자열을 반환함.

- search()

  - 인수로 전달받은 정규 표현식에 맞는 문자나 문자열이 처음으로 등장하는 위치의 인덱스를 반환함.

- replace()

  - 인수로 전달받은 패턴에 맞는 문자열을 대체 문자열로 변환한 새 문자열을 반환함.

- match()

  - 인수로 전달받은 정규 표현식에 맞는 문자열을 찾아서 하나의 배열로 반환함.

- includes()

  - 인수로 전달받은 문자나 문자열이 포함되어 있는지를 검사한 후 그 결과를 불리언 값으로 반환함.

- startsWith()

  - 인수로 전달받은 문자나 문자열로 시작되는지를 검사한 후 그 결과를 불리언 값으로 반환함.

- endsWith()

  - 인수로 전달받은 문자나 문자열로 끝나는지를 검사한 후 그 결과를 불리언 값으로 반환함.

- toLocaleUpperCase()

  - 영문자뿐만 아니라 모든 언어의 문자를 대문자로 변환한 새로운 문자열을 반환함.

- toLocaleLowerCase()

  - 영문자뿐만 아니라 모든 언어의 문자를 소문자로 변환한 새로운 문자열을 반환함.

- localeCompare()

  - 인수로 전달받은 문자열과 정렬 순서로 비교하여 그 결과를 정수 값으로 반환함.

- normalize()

  - 해당 문자열의 유니코드 표준화 양식(Unicode Normalization Form)을 반환함.

- repeat()

  - 해당 문자열을 인수로 전달받은 횟수만큼 반복하여 결합한 새로운 문자열을 반환함.

- toString()

  - String 인스턴스의 값을 문자열로 반환함.

- valueOf()
  - String 인스턴스의 값을 문자열로 반환함.

# Reference

[1] 자바스크립트 String 메소드, https://inpa.tistory.com/entry/JS-📚-String-메소드-✏️-정리
