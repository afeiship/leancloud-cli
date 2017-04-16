interface IPerson{
  name:string,
  age:number
}


class Person {
  constructor (public config:IPerson){}
}
