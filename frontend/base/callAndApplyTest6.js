/**
 * 测试apply
 * apply的第一个参数是运行函数的作用域，第二个参数可以是Array的实例，也可以是arguments对象
 */

function sum(a, b){
	return a + b;
}

function calcSum(a, b){
	return sum.apply(this, arguments);
}
function calcSum2(a, b){
	return sum.apply(this, [a, b]);
}

var ret = calcSum(3, 5);
console.log(ret);
var ret2 = calcSum(5, 7);
console.log(ret2);
