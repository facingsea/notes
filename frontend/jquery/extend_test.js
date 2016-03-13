/**
 * 测试jQuery的 $.extend()方法
 */


/**
 *	Jquery的扩展方法原型是:
 *		extend(dest,src1,src2,src3...);
 *	它的含义是将src1,src2,src3...合并到dest中,返回值为合并后的dest,
 *	由此可以看出该方法合并后，是修改了dest的结构的。如果想要得到合并的结果却又不想修改dest的结构，可以如下使用：
 * 		var newSrc=$.extend({},src1,src2,src3...) //也就是将"{}"作为dest参数。
 */
var data1 = {name: "zhangsan", age: 30 };
var data2 = {name: "lisi", address: "Beijing"};
var data3 = {work: "Programer", address: "Shanghai"};

var newData = $.extend(data1, data2, data3);
console.log("newData: " + JSON.stringify(newData)); // 输出： {"name":"lisi","age":30,"address":"Shanghai","work":"Programer"}
console.log("data1: " + JSON.stringify(data1)); // 输出： {"name":"lisi","age":30,"address":"Shanghai","work":"Programer"}
console.log("data2: " + JSON.stringify(data2)); // 输出： {"name":"lisi","address":"Beijing"}
console.log("data3: " + JSON.stringify(data3)); // 输出： {"work":"Programer", address: "Shanghai"}

//=========================分析=========================
//	由data1和newData的值可以看出，extend方法将data2、data3的值添加到了
//	data1的里面，直接改变了原值。并且，data1中的同名属性会被后来的同名
//	属性依次替换
//
//======================================================

//如果不想改变data1 的结构，可以传递一个空的对象
var data4 = {name: "wangwu"};
var data5 = {name: "zhaoliu", age: 20};
var newData2 = $.extend({}, data5);
console.log(newData2);
console.log(data4);
console.log(data5);


/**
 * 省略dest参数
 * 		 上述的extend方法原型中的dest参数是可以省略的，如果省略了，
 * 		 则该方法就只能有一个src参数，而且是将该src合并到调用extend方法的对象中去，如：
 * 		 	$.extend(src)	// 1
 * 		 该方法就是将src合并到jquery的全局对象中去
 */

$.extend({
	testExtend: function(){
		console.log("This is testExtend function.");
	}
});

$.testExtend();  // 调用通过 $.extend添加的方法，输出：This is testExtend function.

// 2. $.fn.extend(src)  // 该方法将src合并到jquery的实例对象中去
$.fn.extend({
	testFnExtend: function(){
		console.log("This is testFnExtend function.");
	}
});

//$.testFnExtend();				// 该方法下不可调用
$("#firstId").testFnExtend();	// 可以调用，输出：This is testFnExtend function.



/**
 *	Jquery的extend方法还有一个重载原型：
 *		extend(boolean,dest,src1,src2,src3...)
 *	第一个参数boolean代表是否进行深度拷贝，其余参数和前面介绍的一致
 *
 * 
 */
// 深层复制
var newData3 = $.extend(true, {}, 
				{name: "zhangsan", location: {city: "Beijing", country: "China"}}, 
				{nickName: "xiaowang", location: {state: "MA", country: "America"}});
console.log(newData3);
/* 输出结果为：
	{
		name: "zhangsan", 
		nickName: "xiaowang", 
		location: {
			city: "Beijing", 
			state: "MA", 
			country: "America"
		}
	}
 */


// 不深层复制
var newData4 = $.extend(false, {}, 
				{name: "zhangsan", location: {city: "Beijing", country: "China"}}, 
				{nickName: "xiaowang", location: {state: "MA", country: "America"}});
console.log(newData4);
/* 输出结果为：
	{
		name: "zhangsan", 
		nickName: "xiaowang", 
		location: {
			state: "MA", 
			country: "America"
		}
	}
 */
