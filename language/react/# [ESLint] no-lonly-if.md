# [ESLint] no-lonly-if

``` javascript
// If an if statement is the only statement in the else block, it is often clearer to use an else if form.
if (foo) {
    // ...
} else {
    if (bar) {
        // ...
    }
}


// should be rewritten as
if (foo) {
    // ...
} else if (bar) {
    // ...
}
```

- ESLint 규칙에 의해 수정을 하긴 했지만, boolean 변수의 경우 참, 거짓 두 가지뿐일텐데 굳이 해야할까?

# Reference
[1] no-lonly-if, https://eslint.org/docs/latest/rules/no-lonely-if