/**
 *call方法: 
	语法：call([thisObj[,arg1[, arg2[,   [,.argN]]]]]) 
	定义：调用一个对象的一个方法，以另一个对象替换当前对象。 
	说明： 
	call 方法可以用来代替另一个对象调用一个方法。call 方法可将一个函数的对象上下文从初始的上下文改变为由 thisObj 指定的新对象。 
	如果没有提供 thisObj 参数，那么 Global 对象被用作 thisObj。 

  apply方法： 
	语法：apply([thisObj[,argArray]]) 
	定义：应用某一对象的一个方法，用另一个对象替换当前对象。 
	说明： 
	如果 argArray 不是一个有效的数组或者不是 arguments 对象，那么将导致一个 TypeError。 
	如果没有提供 argArray 和 thisObj 任何一个参数，那么 Global 对象将被用作 thisObj， 并且无法被传递任何参数。
 *
 */

function add(a, b){
	console.log(this);
	console.log(a + b);
}

function sub(a, b){
	console.log(this);
	console.log(a - b);
}

// 这个例子中的意思就是用 add 来替换 sub（执行sub函数时用add函数来替换），add.call(sub,3,1) == add(3,1) ，
// 所以运行结果为：alert(4); 
// 注意：js 中的函数其实是对象，函数名是对 Function 对象的引用。
add.call(sub, 3, 1); 		// 输出 4，this输出sub函数定义

//var a1 = new add(); 		// this输出 add{}即实例a1 ，a+b输出 NaN
// a1.sub(2, 5);  			// Uncaught TypeError: a1.sub is not a function
//var s1 = new sub(); 		// this输出 sub{}即实例s1 , a-b输出NaN
//s1.add(2, 5); 			// Uncaught TypeError: s1.add is not a function

//add.apply(sub, 3, 1);		// Uncaught TypeError: Function.prototype.apply: Arguments list has wrong type

add.apply(sub, [3, 1]);		// 输出4， this输出sub函数定义

//var a2 = new add();		// this输出 add{}即实例a2 , a+b输出 NaN
//a2.sub(3, 6);				// Uncaught TypeError: a2.sub is not a function
//var s2 = new sub();			// this输出 sub{}即实例s2 ，a-b输出 NaN
//s2.add(3, 6);				// Uncaught TypeError: s2.add is not a function

//=================================小结 start ========================================//
//	call和apply函数都是调用者来替换参数中函数的执行，参数列表中，第一位都是要替换的
//	目标函数，call里后面的参数对应实际调用者的参数列表，apply后面的参数是一个数组，
//	里面的值对应调用者的实际参数列表，call适合已知参数列表且不可变的情况，
//	apply适合参数列表可变的情况， 具体参数可以push到数组中。
//=================================小结 end ========================================//


