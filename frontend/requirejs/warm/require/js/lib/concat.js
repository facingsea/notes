/**
 * Created by wangzhf on 2016/2/27.
 */

//如果这个模块还依赖其他模块，那么define()函数的第一个参数，必须是一个数组，指明该模块的依赖性。
define(['substring'], function(substring){
    var concat = function(s1, s2){
        if(substring.substring() != null){
            return substring.substring() + s2.toString();
        }
        return s1.toString() + s2.toString();
    };

    return {
        concat: concat
    };
});
