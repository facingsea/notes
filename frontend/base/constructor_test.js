/**
 * 测试constructor
 *
 * 	constructor属性是专门为function而设计的，它存在于每一个function的prototype属性中。
 * 	这个constructor保存了指向function的一个引用。
 */


var testConstructor = function(){
	var name = "zhangsan";
	var sayHello = function(){
		console.log("hello.");
	}
}

console.log(testConstructor);
console.log(testConstructor.prototype);
console.log(testConstructor.constructor);


