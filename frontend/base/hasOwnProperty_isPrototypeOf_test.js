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



