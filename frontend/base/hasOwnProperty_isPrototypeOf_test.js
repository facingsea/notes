/**
 * 	
 *  hasOwnProperty：是用来判断一个对象是否有你给出名称的属性或对象。不过需要注意的是，
    				此方法无法检查该对象的原型链中是否具有该属性，该属性必须是对象本身的一个成员。
	isPrototypeOf:是用来判断要检查其原型链的对象是否存在于指定对象实例中，是则返回true，否则返回false。
 */

function siteAdmin(nickName, siteName){
	this.nickName = nickName;
	this.siteName = siteName;
}

siteAdmin.prototype.showAdmin = function(){
	console.log(this.nickName + " 是 " + this.siteName + " 的站长！");
}

siteAdmin.prototype.showSite = function(siteUrl){
	this.siteUrl = siteUrl;
	return this.siteName + " 的地址是：" + this.siteUrl;
}

var s1 = new siteAdmin("Github", "github.com");
var s2 = new siteAdmin("GoogleCode", "googlecode.com");
s1.age = 30;
s1.showAdmin();
var ret = s1.showSite("https://github.com");
console.log(ret);

console.log(s1.hasOwnProperty("nickName"));   	// true
console.log(s1.hasOwnProperty("age"));			// true
console.log(s1.hasOwnProperty("showAdmin"));	// false
console.log(s1.hasOwnProperty("showSite"));		// false

console.log(siteAdmin.prototype.hasOwnProperty("showAdmin"));	// true
console.log(siteAdmin.prototype.hasOwnProperty("showSite"));   	// true
console.log(siteAdmin.prototype.isPrototypeOf(s1));				// true
console.log(siteAdmin.prototype.isPrototypeOf(s2));				// true


//=================================

var obj = {};

var hasOwn = hasOwnProperty;
console.log(hasOwn);		// function hasOwnProperty() { [native code] }
obj.aaname = "zhangsan";
console.log(hasOwn);		// function hasOwnProperty() { [native code] }

var cc = {address: "Beijing", age: 20};
for(var i in cc){
	//console.log(obj.call(cc, i));  // Uncaught TypeError: obj.call is not a function
	//call应该由Function的实例调用
}

//var aa = hasOwn("aaname");
console.log(hasOwn.call(1, "aaname"));
console.log(window.hasOwnProperty("aaname")); 

//=================hasownProperty() 小结==============
//	
//	JavaScript中hasOwnProperty函数方法是返回一个布尔值，指出一个对象是否具有指定名称的属性。 使用方法：
// 		object.hasOwnProperty(proName)
// 			其中参数object是必选项。一个对象的实例。
// 			proName是必选项。一个属性名称的字符串值。
// 
// 如果 object 具有指定名称的属性，那么JavaScript中hasOwnProperty函数方法返回 true；反之则返回 false。
// 此方法无法检查该对象的原型链中是否具有该属性；该属性必须是对象本身的一个成员。
//
//	********在调用hasOwnProperty时必须有对象来调用，上例中的hasOwn返回的只是方法，但是执行的时候还是需要
//	找调用者的；在call或者apply中，hasOwn的调用者其实是第一个参数的类型。
//
//=================hasownProperty() 小结==============