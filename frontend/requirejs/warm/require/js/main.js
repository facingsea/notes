/**
 * Created by wangzhf on 2016/2/27.
 */

require.config({
    paths: {
        "math": "lib/math",
        "substring": "lib/substring",
        "concat": "lib/concat"
    }
});

require(['math'], function(math){
    alert(math.add(2, 3));
});

require(['concat'], function(concat){
    alert(concat.concat('well', 'world'));
});

