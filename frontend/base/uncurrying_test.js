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
	return function(){
		// 如果针对的是有返回值的函数，那么这里需要return执行结果值，如果没有的话，就不需要return
		return f.call.apply(f, arguments); 
	};
};

Array.slice = Array.prototype.slice.uncurryThis2();
function cutArgs(a, b, c){
	var ret = Array.slice(arguments, 1);
	console.log(ret);
}
cutArgs("first", "second", "third");

//
//	Array.prototype.slice.call.apply(Array.prototype.slice, arguments);
//
