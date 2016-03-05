/**
 * 测试call方法第一个参数为基本类型时的执行。
 */

function outData(key, value){
	console.log(key + " , " + value);
	console.log(this);
	return key;
}

var data = {name: "zhangsan", age: 20};

for(var i in data){
	var ret = outData.call("hello", i, data[i]); // 执行outData方法
	console.log(ret);  // 返回值是outData执行后的返回值
}


/**
 *	输出结果为：
 		name , zhangsan
 		String {0: "h", 1: "e", 2: "l", 3: "l", 4: "o", length: 5, [[PrimitiveValue]]: "hello"}
 		age , 20
 		String {0: "h", 1: "e", 2: "l", 3: "l", 4: "o", length: 5, [[PrimitiveValue]]: "hello"}

 	call方法中的第一个参数"hello"输出为String的对象
 * 
 */

var arr = ["github", "googlecode"];
for(var i=0; i<arr.length; i++){
	var ret = outData.apply(i, [i, arr[i]]);
	console.log(ret);  //  返回值是outData执行后的返回值
} 

/**
 *  输出结果为：
		0 , github
		Number {[[PrimitiveValue]]: 0}
		1 , googlecode
		Number {[[PrimitiveValue]]: 1}

	apply方法中第一个参数i 输出为Number的对象
 */

console.log(undefined === false);