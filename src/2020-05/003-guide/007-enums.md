# enums
- https://www.tslang.cn/docs/handbook/enums.html

## auto enum
```ts
enum E2 {
  A = 1, B, C
}
```

## computed enum
```ts
enum FileAccess {
  // constant members
  None,
  Read    = 1 << 1,
  Write   = 1 << 2,
  ReadWrite  = Read | Write,
  // computed member
  G = "123".length
}
```

