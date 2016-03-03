/**
 *  Created by wangzhf on 2016/3/3
 */

define(function(){
	var noWord = "no word out!",
		outTip = "out word is : ";

	var outStr = function(str){
		if(str != undefined && str != null){
			console.log(outTip + str);
		}else{
			console.log(noWord);
		}
	}


	return {
		outStr: outStr
	}
});
