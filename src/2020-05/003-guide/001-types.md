# types


## boolean
```ts
let isDone: boolean = false;
```

## numbers
```ts
let decLiteral: number = 6;
let hexLiteral: number = 0xf00d;
let binaryLiteral: number = 0b1010;
let octalLiteral: number = 0o744;
```

## string
```ts
let name: string = `Gene`;
let age: number = 37;
let sentence: string = `Hello, my name is ${ name }.

I'll be ${ age + 1 } years old next month.`;
```

## array
```ts
let list1: number[] = [1, 2, 3];
let list2: Array<number> = [1, 2, 3];
```

--- 

## Tuple
> 元组类型允许表示一个已知元素数量和类型的数组，各元素的类型不必相同。 比如，你可以定义一对值分别为 string和number类型的元组。

```ts
// Declare a tuple type
let x: [string, number];
// Initialize it
x = ['hello', 10]; // OK
// Initialize it incorrectly
x = [10, 'hello']; // Error
```

## Enum
> 枚举

```ts
// 无默认值(0,1,2)
enum Color { Red, Green, Blue };

// 翻译成 js 如下
"use strict";
var Color;
(function (Color) {
    Color[Color["Red"] = 0] = "Red";
    Color[Color["Green"] = 1] = "Green";
    Color[Color["Blue"] = 2] = "Blue";
})(Color || (Color = {}));
console.log(Color.Green);


// 有默认值
enum Color { Red: 1, Green: 2, Blue: 4};
```

## Any
> 有时候，我们会想要为那些在编程阶段还不清楚类型的变量指定一个类型。
> 这些值可能来自于动态的内容，比如来自用户输入或第三方代码库。
> 这种情况下，我们不希望类型检查器对这些值进行检查而是直接让它们通过编译阶段的检查。 那么我们可以使用 any类型来标记这些变量：

```ts
let notSure: any = 4;
notSure = "maybe a string instead";
notSure = false; // okay, definitely a boolean

let list: any[] = [1, true, "free"];
```

## Void
> 某种程度上来说，void类型像是与any类型相反，它表示没有任何类型。

```ts
function warnUser(): void {
  console.log("This is my warning message");
}

// 声明一个void类型的变量没有什么大用，因为你只能为它赋予 undefined 和 null ：
let unusable: void = undefined;
```

## Null 和 Undefined
```ts
// Not much else we can assign to these variables!
let u: undefined = undefined;
let n: null = null;
```

## Never
> never类型表示的是那些永不存在的值的类型。

```ts
// 返回never的函数必须存在无法达到的终点
function error(message: string): never {
    throw new Error(message);
}

// 推断的返回值类型为never
function fail() {
    return error("Something failed");
}

// 返回never的函数必须存在无法达到的终点
function infiniteLoop(): never {
    while (true) {
    }
}
```

## Object
> object表示非原始类型，也就是除number，string，boolean，symbol，null或undefined之外的类型。

```ts
// object表示非原始类型，也就是除number，string，boolean，symbol，null或undefined之外的类型。

declare function create(o: object | null): void;

create({ prop: 0 }); // OK
create(null); // OK

create(42); // Error
create("string"); // Error
create(false); // Error
create(undefined); // Error
```

## 类型断言
> 个人感觉，强制类型转换。

```ts
// 类型断言有两种形式。 其一是“尖括号”语法：

let someValue: any = "this is a string";
let strLength: number = (<string>someValue).length;

// 另一个为as语法：

let someValue: any = "this is a string";
let strLength: number = (someValue as string).length;
```
