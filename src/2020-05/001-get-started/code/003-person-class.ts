// https://www.tslang.cn/docs/handbook/typescript-in-5-minutes.html

interface Person {
  firstName: string,
  lastName: string
}


class Student {
  fullName: string;
  constructor(public firstName, public middleInitial, public lastName) {
    this.fullName = firstName + " " + middleInitial + " " + lastName;
  }
}


function greeter(person: Person) {
  // return "Hello, " + person.firstName + " " + person.lastName;
  return "Hello, " + person.firstName + ' ' + person.lastName;
}

let user = new Student("Jane", "M.", "User");

document.body.innerHTML = greeter(user);
