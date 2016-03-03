<%@ page language="java" contentType="text/html; charset=UTF-8"
	pageEncoding="UTF-8"%>
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<title>webuploader上传插件测试</title>

<script type="text/javascript" src="./js/jquery-1.12.1.js"></script>

<!--引入CSS-->
<link rel="stylesheet" type="text/css" href="./js/webuploader/dist/webuploader.css">
<!--引入JS-->
<script type="text/javascript" src="./js/webuploader/dist/webuploader.js"></script>
<style type="text/css">

.tit_op {
    width: 98px;
    height: 32px;
    float: left;
    margin: 3px 5px;
    cursor: pointer;
    text-indent: 32px;
    color: #FFF;
    line-height: 32px;
    border-radius: 2px;
}
</style>
</head>
<body>

	<div id="uploader" class="wu-example">
		<!--用来存放文件信息-->
		<div id="thelist" class="uploader-list"></div>
		<div class="btns">
			<div id="picker" style="background-image: url(./img/ft_01.png); background-repeat: no-repeat;">选择文件</div>
			<div id="mediaFileBtn">多媒体文件</div>
			<button id="ctlBtn" class="btn btn-default">开始上传</button>
		</div>
	</div>
	
	<button id="changeBtn" class="btn btn-default">更改文件类型</button>

	<!--dom结构部分-->
	<div id="uploader-demo">
	    <!--用来存放item-->
	    <div id="fileList" class="uploader-list"></div>
	    <div id="filePicker">选择图片</div>
	</div>
	
<script type="text/javascript">
	// 文件上传
	jQuery(function() {
	    var $ = jQuery,
	        $list = $('#thelist'),
	        $btn = $('#ctlBtn'),
	        $changeBtn = $('#changeBtn'),
	        state = 'pending',
	        uploader;

	    uploader = WebUploader.create({
	    	// swf文件路径
		    swf: 'http://cdn.staticfile.org/webuploader/0.1.5/Uploader.swf'
		    // 文件接收服务端。
		    , server: 'http://localhost:8080/webuploader/FileAction'
		    // 选择文件的按钮。可选。
		    // 内部根据当前运行是创建，可能是input元素，也可能是flash.
		    //pick: '#picker',
		    , pick: {
		    	id: "#picker",  // 指定选择文件的按钮容器，不指定则不创建按钮
		    	//style: "",
		    	multiple: true // 是否开起同时选择多个文件能力
		    }
		    , accept: {  //  {Arroy} [可选] [默认值：null] 指定接受哪些类型的文件。 由于目前还有ext转mimeType表，所以这里需要分开指定。
		    	title: "Images" // 文字描述
		    	,extensions: "gif,jpg,jpeg,bmp,png" // 允许的文件后缀，不带点，多个用逗号分割。
		    	//,mimeTypes: "images/*" // 多个用逗号分割
		    }
		    , thumb: { // {Object} [可选] 配置生成缩略图的选项。
		    	width: 110
    			, height: 110
    			, quality: 70  // // 图片质量，只有type为`image/jpeg`的时候才有效。
    			, allowMagnify: true   // 是否允许放大，如果想要生成小图的时候不失真，此选项应该设置为false.
    			, crop: true  // 是否允许裁剪。
    			//, type: "image/jpeg"
		    }
		    , threads: 1  // [可选] [默认值：3] 上传并发数。允许同时最大上传进程数。
		    , formData: {  //[可选] [默认值：{}] 文件上传请求的参数表，每次发送都会发送此对象中的参数。
		    	username: "zhangsan"
		    	, address: "Beijing"
		    }
		    , fileNumLimit: 2  // [可选] [默认值：undefined] 验证文件总数量, 超出则不允许加入队列。
		    , fileSizeLimit: 5*1024*1000 // (字节) [可选] [默认值：undefined] 验证文件总大小是否超出限制, 超出则不允许加入队列。
		    , fileSingleSizeLimit: 1024*1024*10 // [可选] [默认值：undefined] 验证单个文件大小是否超出限制, 超出则不允许加入队列。
		    , duplicate: false  // {Boolean} true：可以重复，false：不允许重复 [可选] [默认值：undefined] 去重， 根据文件名字、文件大小和最后修改时间来生成hash Key.
		    
		    // 不压缩image, 默认如果是jpeg，文件上传前会压缩一把再上传！
		    , resize: false
	    });
	
	    uploader.addButton({
	    	//id: "#mediaFileBtn"
	    	//, style: ''
	    });
	    
	    //当文件被加入队列之前触发，此事件的handler返回值为false，则此文件不会被添加进入队列。
	    uploader.on('beforeFileQueued', function(file){
	    	var exist = false;
	    	//校验文件名是否已经存在
	    	$.ajax({
	    		data: {"fileName": file.name}, 
	    		url: 'http://localhost:8080/webuploader/ValidatorAction', 
	    		method: 'post',
	    		async: false,  // 同步
	    		success: function(data){
	    			console.log(data);
	    			var ret = JSON.parse(data);
	    			//alert("ret" + ret.success);
	    			if(!ret.success){
		    			exist = true;
	    			}
	    		}, 
	    		error: function(){
	    			
	    		}
	    	});
	    	
	    	if(exist){
	    		alert("文件已存在");
	    		return false;
	    	}else{
	    		alert("文件不存在");
	    		return true;
	    	}
	    
	    	//if(file.name.indexOf('.png') > -1){
	    	//	console.log('file not allowed.');
	    	//	return false;
	    	//}else {
	    	//	return true;
	    	//}
	    });
	    
	    // 当一批文件添加进队列以后触发。
	    uploader.on('filesQueued', function(file){
	    	console.log(file.length);
	    });
	    
	    // 当有文件添加进来的时候
	    uploader.on( 'fileQueued', function( file ) {
	        $list.append( '<div id="' + file.id + '" class="item">' +
	            '<h4 class="info">' + file.name + '</h4>' +
	            '<p class="state">等待上传...</p>' +
	        '</div>' );
	    });
		
	    // 当文件被移除队列后触发
	    uploader.on('fileDequeued', function(file){
	    	console.log('file '+file.name+' is removed.');
	    });
	    
	    // 文件上传过程中创建进度条实时显示。
	    uploader.on( 'uploadProgress', function( file, percentage ) {
	        var $li = $( '#'+file.id ),
	            $percent = $li.find('.progress .progress-bar');

	        // 避免重复创建
	        if ( !$percent.length ) {
	            $percent = $('<div class="progress progress-striped active">' +
	              '<div class="progress-bar" role="progressbar" style="width: 0%">' +
	              '</div>' +
	            '</div>').appendTo( $li ).find('.progress-bar');
	        }

	        $li.find('p.state').text('上传中');

	        $percent.css( 'width', percentage * 100 + '%' );
	    });

	    uploader.on( 'uploadSuccess', function( file ) {
	        $( '#'+file.id ).find('p.state').text('已上传');
	        console.log('uploadSuccess: file id is: ' + file.id);
	    });


		/**
		默认如果啥也不处理，只要有返回数据就认为是成功，就算返回的是错误信息，也认为是成功了。
		但是，```在认为成功前会派送一个事件uploadAccept，这个事件是用来询问是否上传成功的```。
		在这个事件中你可以拿到上传的是哪个文件，以及对应的服务端返回reponse。
		
		response是个对象，如果服务器返回是json格式，那么正和你意，都已经解析好了，
		如果不是json格式，response._raw里面可以拿到原始数据。所以，webuploader对于后端返回的数据格式是没有要求的。
		
		*/
		uploader.on('uploadAccept', function(file, response){
			console.log('uploadAccept:');
			console.log(file.name);
			console.log(response);
			if(response.error != undefined ){
				if(response.error){
					alert("上传失败，信息为：" + response.msg);
					return false;
				}else{
					return true;
				}
			}else{
				alert("未知错误");
				return false;
			}
		});
		
	    uploader.on( 'uploadError', function( file, reason  ) {
	        $( '#'+file.id ).find('p.state').text('上传出错');
	        console.log("uploadError:");
	        console.log(reason);
	        console.log(file.error);
	    });

	    uploader.on( 'uploadComplete', function( file ) {
	        $( '#'+file.id ).find('.progress').fadeOut();
	    });

	    uploader.on( 'all', function( type ) {
	        if ( type === 'startUpload' ) {
	            state = 'uploading';
	        } else if ( type === 'stopUpload' ) {
	            state = 'paused';
	        } else if ( type === 'uploadFinished' ) {
	            state = 'done';
	        }

	        if ( state === 'uploading' ) {
	            $btn.text('暂停上传');
	        } else {
	            $btn.text('开始上传');
	        }
	    });
	    
	    /**
	    	当validate不通过时，会以派送错误事件的形式通知调用者。
	    	通过upload.on('error', handler)可以捕获到此类错误，目前有以下错误会在特定的情况下派送错来。
	    		1. Q_EXCEED_NUM_LIMIT 在设置了fileNumLimit且尝试给uploader添加的文件数量超出这个值时派送。
	    		2. Q_EXCEED_SIZE_LIMIT 在设置了Q_EXCEED_SIZE_LIMIT且尝试给uploader添加的文件总大小超出这个值时派送。
	    		3. Q_TYPE_DENIED 当文件类型不满足时触发。
	    */
	    uploader.on('error', function(type){
	    	console.log("error");
	    	console.log(type);
	    	if(type == 'Q_EXCEED_NUM_LIMIT'){
	    		alert('文件超出大小');
	    	}else if (type == 'Q_EXCEED_SIZE_LIMIT'){
	    		alert('总文件超出设置大小');
	    	}else if (type == 'Q_TYPE_DENIED'){
	    		alert('非法文件类型');
	    	}else if (type == 'F_DUPLICATE'){
	    		alert('列表中文件已存在');
	    	}
	    });

	    $btn.on( 'click', function() {
	        if ( state === 'uploading' ) {
	            uploader.stop();
	        } else {
	            uploader.upload();
	        }
	    });
	    
	    $changeBtn.on('click', function(){
			alert("change the accept types");
			console.log(uploader);
			//重新设置accept不好使
	    	uploader.option('accept', {  //  {Arroy} [可选] [默认值：null] 指定接受哪些类型的文件。 由于目前还有ext转mimeType表，所以这里需要分开指定。
		    	title: "Csv" // 文字描述
		    	,extensions: "csv" // 允许的文件后缀，不带点，多个用逗号分割。
		    	//,mimeTypes: "images/*" // 多个用逗号分割
		    }
		    );
			console.log(uploader);
	    })
	});


	// 图片上传demo
	jQuery(function() {
	    var $ = jQuery,
	        $list = $('#fileList'),
	        // 优化retina, 在retina下这个值是2
	        ratio = window.devicePixelRatio || 1,

	        // 缩略图大小
	        thumbnailWidth = 100 * ratio,
	        thumbnailHeight = 100 * ratio,

	        // Web Uploader实例
	        uploader;

	    // 初始化Web Uploader
	    uploader = WebUploader.create({

	        // 自动上传。
	        auto: true,

	     	// swf文件路径
		    swf: 'http://cdn.staticfile.org/webuploader/0.1.5/Uploader.swf',
		    // 文件接收服务端。
		    server: 'http://localhost:8080/webuploader/FileAction',

	        // 选择文件的按钮。可选。
	        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
	        pick: '#filePicker',

	        // 只允许选择文件，可选。
	        accept: {
	            title: 'Images',
	            extensions: 'gif,jpg,jpeg,bmp,png',
	            mimeTypes: 'image/*'
	        }
	    });

	    // 当有文件添加进来的时候
	    uploader.on( 'fileQueued', function( file ) {
	        var $li = $(
	                '<div id="' + file.id + '" class="file-item thumbnail">' +
	                    '<img>' +
	                    '<div class="info">' + file.name + '</div>' +
	                '</div>'
	                ),
	            $img = $li.find('img');

	        $list.append( $li );

	        // 创建缩略图
	        uploader.makeThumb( file, function( error, src ) {
	            if ( error ) {
	                $img.replaceWith('<span>不能预览</span>');
	                return;
	            }

	            $img.attr( 'src', src );
	        }, thumbnailWidth, thumbnailHeight );
	    });

	    // 文件上传过程中创建进度条实时显示。
	    uploader.on( 'uploadProgress', function( file, percentage ) {
	        var $li = $( '#'+file.id ),
	            $percent = $li.find('.progress span');

	        // 避免重复创建
	        if ( !$percent.length ) {
	            $percent = $('<p class="progress"><span></span></p>')
	                    .appendTo( $li )
	                    .find('span');
	        }

	        $percent.css( 'width', percentage * 100 + '%' );
	    });

	    // 文件上传成功，给item添加成功class, 用样式标记上传成功。
	    uploader.on( 'uploadSuccess', function( file ) {
	        $( '#'+file.id ).addClass('upload-state-done');
	    });

	    // 文件上传失败，显示上传出错。
	    uploader.on( 'uploadError', function( file ) {
	        var $li = $( '#'+file.id ),
	            $error = $li.find('div.error');

	        // 避免重复创建
	        if ( !$error.length ) {
	            $error = $('<div class="error"></div>').appendTo( $li );
	        }

	        $error.text('上传失败');
	    });

	    // 完成上传完了，成功或者失败，先删除进度条。
	    uploader.on( 'uploadComplete', function( file ) {
	        $( '#'+file.id ).find('.progress').remove();
	    });
	    
	});
</script>
</body>
</html>