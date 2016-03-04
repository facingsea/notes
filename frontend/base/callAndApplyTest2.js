// Animal类
function Animal(){
	this.name = "Animal";
	this.showName = function(){
		console.log(this.name);
	}
}

// Cat类
function Cat(){
	this.name = "Cat";
}

var animal = new Animal();
var cat = new Cat();

// 相当于把animal.showName中的this换成了cat，this.name ==> cat.name
animal.showName.call(cat);   //输出 Cat
// 调用者必须是function
animal.name.call(cat); // Uncaught TypeError: animal.name.call is not a function

animal.showName.apply(cat);  //输出 Cat


