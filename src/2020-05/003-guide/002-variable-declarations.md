# variable-declarations
- https://www.tslang.cn/docs/handbook/variable-declarations.html


## why not var
```ts
function 
```

## 作用域规则
- 变量 x是定义在*if语句里面*，但是我们却可以在语句的外面访问它。 
- 这是因为 var声明可以在包含它的函数，模块，命名空间或全局作用域内部任何位置被访问（我们后面会详细介绍），包含它的代码块对此没有什么影响。 
- 有些人称此为* var作用域或函数作用域*。 函数参数也使用函数作用域。

```ts
function f(shouldInitialize: boolean) {
  if (shouldInitialize) {
    var x = 10;
  }
  return x;
}

f(true);  // returns '10'
f(false); // returns 'undefined'


// 与下面的代码的区别
function f(){
  // ReferenceError: x is not defined
  return x;
}
```

## 其中之一就是，多次声明同一个变量并不会报错
- 注意代码中的 `i` 重复运用了。

```ts
function sumMatrix(matrix: number[][]) {
    var sum = 0;
    for (var i = 0; i < matrix.length; i++) {
        var currentRow = matrix[i];
        for (var i = 0; i < currentRow.length; i++) {
            sum += currentRow[i];
        }
    }
    return sum;
}
```

## 捕获变量怪异之处
> 出了名的闭包问题。
> 介绍一下，setTimeout会在若干毫秒的延时后执行一个函数（等待其它代码执行完毕）。

```ts
for (var i = 0; i < 10; i++) {
  setTimeout(function() { console.log(i); }, 100 * i);
}
// 10
// 10
// 10
// 10
// ...
```

## 期待的结果
```ts
for (var i = 0; i < 10; i++) {
    // capture the current state of 'i'
    // by invoking a function with its current value
    (function(i) {
        setTimeout(function() { console.log(i); }, 100 * i);
    })(i);
}
```

## let
> 当用let声明一个变量，它使用的是词法作用域或块作用域。 不同于使用 var声明的变量那样可以在包含它们的函数外访问，块作用域变量在包含它们的块或for循环之外是不能访问的。

```ts
function f(input: boolean) {
  let a = 100;
  if (input) {
    // Still okay to reference 'a'
    let b = a + 1;
    return b;
  }
  // Error: 'b' doesn't exist here
  return b;
}
```


## 块级作用域链
拥有块级作用域的变量的另一个特点是
- 它们不能在被声明之前读或写。 
- 虽然这些变量始终“存在”于它们的作用域里，但在直到声明它的代码之前的区域都属于 暂时性死区。 
- 它只是用来说明我们不能在 let语句之前访问它们，幸运的是TypeScript可以告诉我们这些信息。
