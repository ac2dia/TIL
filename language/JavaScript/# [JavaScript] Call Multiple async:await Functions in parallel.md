# [JavaScript] Call Multiple async/await Functions in parallel

## use async/await

```javascript
// 1번 방식: async/await
function getTallestPerson() {
  var promise = new Promise((resolve, reject) => {
    setTimeout(() => resolve(123), 1000);
  });
  return promise;
}

function getUserDetails(studentId) {
  var promise = new Promise((resolve, reject) => {
    setTimeout(
      () =>
        resolve({
          id: studentId,
          name: 'John',
          company: 'clairvoyant',
        }),
      1000
    );
  });
  return promise;
}

async function getUser() {
  let studentIdPromise = getTallestPerson();
  let studentId = await studentIdPromise;

  let userDetailsPromise = getUserDetails(studentId);
  let userDetails = await userDetailsPromise;
  console.log('userDetails', userDetails);
}

getUser();
```

- await 키워드는 혼자서는 아무런 의미가 없는 코드이지만, async 키워드 이 후 호출되는 경우 block 상태로 선언되어 response 가 오기까지 기다립니다.

```javascript
// 2번 방식: promise chain
function getTallestPerson() {
  var promise = new Promise((resolve, reject) => {
    setTimeout(() => resolve(123), 1000);
  });
  return promise;
}

function getUserDetails(studentId) {
  var promise = new Promise((resolve, reject) => {
    setTimeout(
      () =>
        resolve({
          id: studentId,
          name: 'John',
          company: 'clairvoyant',
        }),
      1000
    );
  });
  return promise;
}

function getUser() {
  let studentIdPromise = getTallestPerson();
  studentIdPromise
    .then((studentId) => {
      return getUserDetails(studentId);
    })
    .then((userDetails) => {
      console.log('userDetails', userDetails);
    });
}
```

- 비동기 기법으로 async/await 대신에 promise 를 사용할 수 있습니다.

## Call Multiple async/await

```javascript
function process(val, sec, rej = false) {
  return new Promise((resolve, reject) => {
    setTimeout(() => {
      if (!rej) {
        resolve(val);
      } else {
        reject('Error');
      }
    }, sec);
  });
}

async function process_one() {
  return process(10, 500);
}

async function process_two() {
  return process(40, 100);
}

async function process_three() {
  return process(50, 1000);
}

async function process_four() {
  return process(20, 800);
}

async function run_processes() {
  let res = null;
  try {
    res = await Promise.all([
      process_one(),
      process_two(),
      process_three(),
      process_four(),
    ]);
    console.log('Success >>', res);
  } catch (err) {
    console.log('Fail >>', res, err);
  }
}

run_processes();
```

- async/await 키워드를 통해 async 함수가 완료될 때까지 block 상태가 됩니다.
- 또한 Promise.all([func]) 를 통해 여러 개의 async 함수에 대하여 결과를 반환할때 까지 대기합니다.

# Reference

[1] How to use Async/Await in JavaScript, https://blog.clairvoyantsoft.com/how-to-use-async-await-in-javascript-9f640ec96aad
[2] Call Multiple async/await Functions Parallel or Sequential, https://ankurpatel.in/blog/call-multiple-async-await-functions-parallel-or-sequential/
[3] Call async/await functions in parallel, https://stackoverflow.com/questions/35612428/call-async-await-functions-in-parallel
