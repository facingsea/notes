/**
 *  实现多继承
 */

function class01(){
	this.showSub = function(a, b){
		console.log(a - b);
	}
}

function class02(){
	this.showAdd = function(a, b){
		console.log(a + b);
	}
}

function class03(){
	class01.call(this);
	class02.apply(this);
	console.log(this); // 输出 class03{}，里面包含showSub和showAdd方法
}

var cls = new class03();
cls.showSub(3, 1);    // 输出 2
cls.showAdd(3, 1);	  // 输出 4

