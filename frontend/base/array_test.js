/**
 *	js Array测试
 * 
 */

/**
 * forEach 测试
 * refer: https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/forEach
 */

var a1 = ["first", "second", "third"];

// a1.forEach(function(value, index, currentArray){
// 	console.log(value);
// 	console.log(index);
// 	console.log(currentArray);
// });

a1.forEach(function(value, index){
	console.log(value);
	console.log(index);
});

console.log(a1.shift());
console.log(a1.shift());
console.log(a1.shift());
console.log(a1.shift());

