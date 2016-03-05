
/**
 * RegExp.$1是RegExp的一个属性,指的是与正则表达式匹配的第一个 子匹配(以括号为标志)字符串，
 * 以此类推，RegExp.$2，RegExp.$3，..RegExp.$99总共可以有99个匹配
 */

function testRegExp(){


	var reg = /^#(.*)$/;

	var ss = "#hello.world#javascript";

	reg.test(ss);
	console.log(RegExp.$1);		// hello.world#javascript
	console.log(RegExp.$2);		// "" （空字符串）
	reg.test(ss);
	console.log(RegExp.$1);		// hello.world#javascript
	console.log(RegExp.$2);		// "" （空字符串）

	reg.exec(ss);
	console.log(RegExp.$1);		// hello.world#javascript
	console.log(RegExp.$2);		// "" （空字符串）
	reg.exec(ss);
	console.log(RegExp.$1);		// hello.world#javascript
	console.log(RegExp.$2);		// "" （空字符串）
}

testRegExp();
