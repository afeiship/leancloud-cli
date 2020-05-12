# interfaces
> 接口的作用就是为这些类型命名和为你的代码或第三方代码定义契约。


## normal/optional/readonly
```ts
interface Person{
  name: string;
  age: number;
}

interface SquareConfig {
  color?: string;
  width?: number;
}

interface Point {
    readonly x: number;
    readonly y: number;
}
```


## misunderstand
- 可索引的类型
- implements(少见使用，暂时不了解)
