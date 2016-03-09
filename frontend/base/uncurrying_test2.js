/**
 * 测试call.apply连用
 */

Function.prototype.uncurryThis3 = function(){
	var f = this;
	return function(){
		console.log(arguments);
		console.log(arguments instanceof Array);  // false
		return f.call.apply(f, arguments);
	};
}

Array.forEach = Array.prototype.forEach.uncurryThis3();

Array.forEach(["first", "second", "third"], function(ele, index){
	console.log(index + " ==> " + ele);
});

//============================================================================
//
//	uncurryThis3中，f指代调用uncurryThis3的函数，本例中为Array.prototype.forEach
//	那么f.call.apply(f, arguments) 就可以换成：
//	
//		Array.prototype.forEach.call.apply(Array.prototype.forEach, arguments);
//  ==> (Array.prototype.forEach.call).apply(Array.prototype.forEach, arguments);
//  
//  通过Apply方法，Array.prototype.forEach替换了(Array.prototype.forEach.call)中的this，
//  而(Array.prototype.forEach.call)中的this指的是Array.prototype.forEach，
//  所以实际执行的是Array.prototype.forEach.call(arguments);
//  即：
//  	Array.prototype.forEach.call(["first", "second", "third"], function(){....}); 
//
//
//============================================================================



//==针对js的forEach测试
//
var data = ["beijing", "shanghai", "hangzhou"];
data.forEach(function(elem, index){
	console.log(index + "---->" + elem);   // 执行成功，输出 0---->beijing ....
});

// 测试call中传入arguments的执行过程
// 
function testCallArguments(a, b){
	console.log(arguments);
	console.log(a.concat(b));
	var f = function(){
		console.log(arguments);
		console.log(arguments instanceof Array);  // false
		var aa = Array.prototype.concat.call(arguments);
		console.log(aa);
		console.log(aa instanceof Array);  // true
	}
	f("hello", arguments);
}
testCallArguments("github", "googlecode", "SourceForge");


function testForEachFunc(){
	console.log(arguments);
	Array.prototype.forEach.call(arguments);
}

testForEachFunc(["first", "second", "third"], function(elem, index){
	console.log(index + "---->" + elem); 
});

