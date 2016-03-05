/**
 * 模拟jQuery的each方法
 */

var class2type = {},
    hasOwn = class2type.hasOwnProperty;

function each( obj, iterator ) {
    var i;

    //fix error, add guard here
    if(!obj) {
        return;
    }

    // like array
    if ( typeof obj !== 'function' && typeof obj.length === 'number' ) {
        for ( i = 0; i < obj.length; i++ ) {
            // 执行iterator方法，传入i，obj[i]两个参数
            if ( iterator.call( obj[ i ], i, obj[ i ] ) === false ) {
                return obj;
            }
        }
    } else {
        for ( i in obj ) {
            //hasOwn.call方法实际执行的是string.hasOwnProperty(i);
            if ( hasOwn.call( obj, i ) && iterator.call( obj[ i ], i,
                   obj[ i ] ) === false ) {
               return obj;
            }
        }
    }

    return obj;
}

console.log("==============obj test===============");
var data = {name: "zhangsan", "address": "Beijing"};
var ret = each(data, function(key, value){
    console.log("key: " + key + " , value: " + value);
});
console.log(ret);

console.log("==============arr test===============");
var arr = ["hello", "world"];
var ret2 = each(arr, function(key, value){
    console.log("key: " + key + " , value: " + value);
});

console.log(window.hasOwnProperty);             // 输出 function hasOwnProperty() { [native code] } 
//console.log(hasOwnProperty("data"));            // Uncaught TypeError: Cannot convert undefined or null to object
console.log(window.hasOwnProperty("data"));     // true
console.log(class2type.hasOwnProperty("data")); // false


