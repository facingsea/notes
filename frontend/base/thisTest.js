// 情况一：纯粹的函数调用
// 当前指定的是全局的window对象
function test(){
	this.x = 1;
	console.log(this); // windows
}

test();

//==================================

var x = 1;

function test2(){
	this.x = 0;
}

test2();
console.log(x); // 0

//==================================
//情况二：作为对象方法的调用

function test3(){
	console.log(this.x);
}
var o = {};
o.x = 3;
o.f = test3;
o.f(); // this指向o，输出为 3


//==================================
// 情况三 作为构造函数调用

function test4(){
	this.x = 4;
}
var t4 = new test4();
console.log(t4.x); // 当前this指向t4，输出 4


//==================================
// 情况四 apply调用

var x5 = 5;
function test5(){
	console.log(this.x5);
}
var t5 = {};
t5.x5 = 55;
t5.f = test5;
t5.f.apply(); // apply()的参数为空时，（会执行方法本身）默认调用全局对象。因此，这时的运行结果为 5 ，证明this指的是全局对象。

t5.f.apply(t5); // 运行结果就变成了 55 ，证明了这时this代表的是对象t5。


//==================================
// 情况五 call调用

var x6 = 6;
function test6(){
	console.log(this.x6);
}
var t6 = {};
t6.x6 = 66;
t6.f = test6;
t6.f.call(); // 输出结果为 6，证明此时this指向全局

t6.f.call(t6); // 输出 66， this指向了t6
