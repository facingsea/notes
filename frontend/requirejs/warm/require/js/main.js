/**
 * Created by wangzhf on 2016/2/27.
 */

/**
 *	require方式参见AMD：https://github.com/amdjs/amdjs-api/wiki/AMD-%28%E4%B8%AD%E6%96%87%E7%89%88%29
 * 
 */

require.config({
    paths: {
        "math": "lib/math",
        "substring": "lib/substring",
        "concat": "lib/concat", 
        "str": "lib/str"
    }
});

//require(['math'], function(math){
//    alert(math.add(2, 3));
//});

// require(['concat'], function(concat){
//     alert(concat.concat('well', 'world'));
// });


require(["str"], function(str){
	var a = null,
		b = "zhangsan";
	str.outStr(a);
	str.outStr(b);
})
