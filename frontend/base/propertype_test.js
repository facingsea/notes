/**
 * 测试propertype原型
 */

/**
 * 创建的每一个函数都有一个prototype（原型）属性，这个属性是一个指针，指向一个对象，
 * 而这个对象的用途可以由特定类型的所有实例共享的属性和方法。如果按照字面意思来理解，
 * 那么  prototype就是通过调用构造函数而创建的那个对象实例的原型对象。  使用原型对象的好处
 * 是可以让所有对象实例共享他所包含的属性和方法。换句话说，不必在构造函数中定义对象实例
 * 的信息，而是可以将信息直接添加到原型对象中，如下例：
 */

function Person(){

}

Person.prototype.name = "Zhangsan";
Person.prototype.age = 30;
Person.prototype.job = "Software Engineer";
Person.prototype.sayName = function(){
	console.log(this.name);
};

var p1 = new Person();
p1.sayName();		// zhangsan

var p2 = new Person();
p2.sayName();		// zhangsan

console.log(p1.sayName == p2.sayName);		// true
console.log(p1.sayName === p2.sayName);		// true
