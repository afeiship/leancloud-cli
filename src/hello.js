var __extends = (this && this.__extends) || (function () {
    var extendStatics = Object.setPrototypeOf ||
        ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
        function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var Hello = (function () {
    function Hello() {
        this.state = {
            test: 1234
        };
    }
    Hello.say = function (inSomething) {
        console.log("say.." + inSomething);
    };
    Hello.prototype.method1 = function () {
        console.log('mehtod1');
    };
    return Hello;
}());
;
var World = (function (_super) {
    __extends(World, _super);
    function World() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    return World;
}(Hello));
World.say('tes');
