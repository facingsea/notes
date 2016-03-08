/**
 *	反科里化测试
 * 
 */

// 方式一
Function.prototype.uncurryThis = function(){
	var f = this;
	console.log(this);  // this指调用uncurryThis的方法
	return function(){
		var a = arguments;
		// a[0]表示所要代替的函数，[].slice.call(a, 1)将第一个参数后的所有参数转换成一个新的数组
		f.apply(a[0], [].slice.call(a, 1));
	};
}

Array.forEach = Array.prototype.forEach.uncurryThis();
function printArgs(a, b, c){
	// 此处的Array.forEach相当于uncurryThis函数中返回的方法，而arguments和func是该返回函数的
	// 两个参数，uncurryThis里的f代表了Array.prototype.forEach，arguments是a[0]，func将会转变成
	// f.apply的第二个参数
	Array.forEach(arguments, function(elem, index){
		console.log(index + " ==> " + elem);
	});
	//	相当于：
	//	Array.prototype.forEach.apply(arguments, [function(...){...}]);
	//

}

printArgs("first", "second", "third");


// 方式二
Function.prototype.uncurryThis2 = function(){
	var f = this;
	console.log(this);
	return function(){
		// 如果针对的是有返回值的函数，那么这里需要return执行结果值，如果没有的话，就不需要return
		//return f.call.apply(f, arguments); 
		var c = f.call;
		console.log(arguments);				// [Arguments[3], 1]
		console.log(typeof arguments);		// object
		console.log(arguments instanceof Array);	// false
		var a = c.apply(f, arguments);
		return a;
	};
};

Array.slice = Array.prototype.slice.uncurryThis2();
function cutArgs(a, b, c){
	var ret = Array.slice(arguments, 1);
	console.log(ret);
}
cutArgs("first", "second", "third");

//
//	f.call.apply(f, arguments);  ==>   f.call(arguments)
//		
//	Array.prototype.slice.call(arguments);
//
//	Array.slice(args, 1){
//		var f = Array.prototype.slice;
//		var c = f.call;
//		var a = c.apply(f, arguments);
//	}
//

console.log("=====test apply===");
(function testApply(name, address, age){
	console.log(arguments);
	console.log(arguments instanceof Array);		// false
	var ret = Array.prototype.slice(arguments, 1);
	console.log(ret);
	console.log(typeof ret);
}("zhangsan", "Beijing", 30))

console.log("hello".this);   // undefined
console.log(typeof []);
console.log([] instanceof Array);  // true
