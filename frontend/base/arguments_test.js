/**
 * arguments对象不能显式创建，arguments对象只有函数开始时才可用。
 * 函数的 arguments 对象并不是一个数组，访问单个参数的方式与访问数组元素的方式相同。
 * 索引 n 实际上是 arguments 对象的 0…n 属性的其中一个参数。
 *
 * 
 */

function testArguments(source){
	console.log(arguments);				// 输出 ["hello"]
	console.log(arguments.length);		// 输出 1
}

testArguments("hello");


/**
 * arguments对象的长度是由实参个数而不是形参个数决定的。
 * 形参是函数内部重新开辟内存空间存储的变量，但是其与arguments对象内存空间并不重叠。
 * 对于arguments和值都存在的情况下，两者值是同步的，但是针对其中一个无值的情况下，对于此无值的情形值不会得以同步。
 *
 * 如下例中，参数a传进来就有值，而c未传，在函数中对a的修改，无论是直接对a赋值（a = 100）还是使用arguments
 * 指定（arguments[0] = "got it"），两者都是同步的，但是对c赋值后，并没有同步到arguments。即使再对arguments[3]
 * 赋值，arguments[3]也是没关联的，由此可见，arguments的参数在调用函数时就与形参关联起来了。
 * 
 */

function testArguments2(a, b, c){
	console.log(arguments.length);	// 2
	a = 100;
	console.log(arguments[0]);		// 100
	arguments[0] = "got it";
	console.log(a);					// got it
	console.log(c);					// undefined
	c = 2016;
	console.log(c);					// 2016
	console.log(arguments[3]);		// undefined
	arguments[3] = 2017;			
	console.log(arguments[3]);		// 2017
	console.log(c);					// 2016
}

testArguments2(1, 2);


/**
 *	callee属性，返回正被执行的 Function 对象，也就是所指定的 Function 对象的正文。
 *	callee 属性是 arguments 对象的一个成员，仅当相关函数正在执行时才可用。
 *	callee 属性的初始值就是正被执行的 Function 对象，这允许匿名的递归函数。
 *
 * 
 */

function factorial(n){
	console.log(arguments.callee);   			// 输出 function factorial(n){....}
	if(n <= 0){
		return 1;
	}else{
		// callee属性指的是当前指定的函数对象，后面的括号表示传参，结构类似于(function xx(n){}())自执行
		return n * arguments.callee(n - 1);		
	}
}

var ret = factorial(3);
console.log(ret);	// 6
