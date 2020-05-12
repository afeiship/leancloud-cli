/**
 *
 * cd /Users/feizheng/github/typescript-notes/src/2020-05/001-get-started
 * tsc 001-greeter.ts
 */


function greeter(person: string) {
  return "Hello, " + person;
}

let user = "Jane User";

document.body.innerHTML = greeter(user);



