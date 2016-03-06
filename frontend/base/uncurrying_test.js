/**
 *	反科里化测试
 * 
 */

function uncurrying(fn){
	return function(){
		return Function.call.apply(fn, arguments);
	}
}

var obj = {};
var push = Array.prototype.push.uncurrying();
push(obj, "first");
console.log(obj.length);
console.log(obj[0]);


