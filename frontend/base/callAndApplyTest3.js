/**
 * 实现继承（单继承）
 */

/**
 * Person 类
 * @param {[type]} name [description]
 */
function Person(name){
	this.name = name;
	console.log(this);

	this.showName = function(){
		console.log(this.name);
	}
}

/**
 * Programer 类
 * @param {[type]} name [description]
 */
function Programer(name){
	//Person.call(this, name);
	Person.apply(this, [name]);
}

var pro = new Programer("Hello world."); // 输出：Programer {name: "Hello world."}
pro.showName();  // 输出：Hello world.
console.log(pro.name); // 输出： Hello world.

//=================================小结 start ========================================//
//	在Programer中使用call或apply方法将Person替换了当前的this对象，就有了Person中的
//	name属性和showName方法，Programer的实例就可以直接调用
//=================================小结 end ========================================//

