# extends

## 成员权限
- default(public)
- public
- private
- protected


## readonly
```ts
class Octopus {
  readonly name: string;
  readonly numberOfLegs: number = 8;
  constructor (theName: string) {
    this.name = theName;
  }
}
```

## abstract class
> 抽象类做为其它派生类的基类使用。 它们一般不会直接被实例化。 
> 不同于接口，抽象类可以包含成员的实现细节。 
> abstract关键字是用于定义抽象类和在抽象类内部定义抽象方法。

```ts
abstract class Animal {
  abstract makeSound(): void;
  move(): void {
    console.log('roaming the earch...');
  }
}
```

## 把类当做接口使用
> 类定义会创建两个东西：类的实例类型和一个构造函数。 因为类可以创建出类型，所以你能够在允许使用接口的地方使用类。

```ts
class Point {
  x: number;
  y: number;
}

interface Point3d extends Point {
  z: number;
}

let point3d: Point3d = {
  x: 1, 
  y: 2, 
  z: 3
};
```
