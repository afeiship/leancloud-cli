class Hello {

  static say(inSomething?: string){
    console.log(`say..${inSomething}`);
  }

  state = {
    test:1234
  };

  method1(){
    console.log('mehtod1');
  }
}

interface User{
  name:string,
  age:number
};

class World extends Hello{

}



//Call say method:
World.say('tes');
