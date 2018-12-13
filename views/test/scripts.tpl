<script type="text/javascript">

    var test = {
        $uploader:$('#uploaderId'),
        xsrf:$('input[name="_xsrf"]').val(),
        init:function(){
            this.render();
        },
        render: function () {
            var _this = this,options = {
                url:'/uploadToOSS',
                headers:{'X-Xsrftoken':_this.xsrf},
                chunk_size:0,
                onFileUploaded:_this.uploadCallback
            };
            _this.$uploader.uploader(options);
        },
        uploadCallback:function(file, responseObject){
            var $img = $('.img');
            console.log(responseObject);
            var resp = JSON.parse(responseObject.response)
            if(resp.code){
                layer.msg(resp.msg);
                return false;
            }

            layer.msg("上传成功");
            var image = '<div><img src="'+resp.imgUrl+'"><img src="data:image/img;base64,'+resp.qrCodebase64Code+'"></div>';
            $(image).appendTo($img);
        },
        handleEvent: function () {

        }
    };

    $(function(){
        test.init();

    });
</script>