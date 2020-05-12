/**
 *
 * cd /Users/feizheng/github/typescript-notes/src/2020-05/001-get-started/code
 * tsc 001-greeter.ts
 */
function greeter(person) {
    return "Hello, " + person;
}
var user = "Jane User";
document.body.innerHTML = greeter(user);
