<script type="text/javascript">

    let company = {
        id:{{.Id}},
        operateType: +{{.OperateType}},
        urlApiCompanyGetEditInfo:{{.ApiUriCompanyGetEditInfo}},
        urlHtmlCompanyEdit:{{.HtmlUriCompanyEdit}},
        urlApiCompanyEdit:{{.ApiUriCompanyEdit}},
        urlApiUserGetCheckCode:{{.ApiUriUserGetCheckCode}},
        urlApiMiniProgramVersionCarouselUpload:{{.ApiUriMiniProgramVersionCarouselUpload}},
        $btnEdit: $('.btn-edit'),
        $type: $('select[name="type"]'),
        $businessCardInfo: $('.business-card-info'),
        $carouselUploader: $('#carouselUploader'),
        error:{{.Error}},
        init: function () {
            this.render();
        },
        render: function () {
            this.handleEvent();
            this.initHtml();
        },
        handleEvent: function () {
            let _this = this;
            $('body').on('click', '.btn-edit', function () {
                _this.$btnEdit.addClass("disabled");
                _this.edit();
            });
        },
        initHtml: function () {
            let _this = this;

            console.log(_this.$type.val());
            if (_this.error) {
                _this.$btnEdit.addClass('disabled');
                layer.popupError(_this.error);
                return
            }

            if (1 === +_this.$type.val()) {
                _this.$businessCardInfo.removeClass('hidden')
                _this.$carouselUploader.uploader({//轮播图上传插件初始化
                    autoUpload: true,            // 当选择文件后立即自动进行上传操作
                    url: _this.urlApiMiniProgramVersionCarouselUpload,
                    chunk_size: 0,
                    headers: {'X-Xsrftoken': $('meta[name=_xsrf]').attr('content')},
                    responseHandler: function (responseObject, file) {
                        _this.$carouselUploader.data('zui.uploader').removeFile(file.id);
                        let rsp = JSON.parse(responseObject.response);
                        console.log(rsp, file);
                        if (!+rsp.error) {
                            layer.popupMsg('上传成功！');
                            //todo 已上传图片展示到前端

                            return;
                        }

                        layer.popupError('上传失败：' + rsp.msg);
                    }
                });


            }

            if (1 === _this.operateType) {


            }

            if (2 === _this.operateType) {
                _this.$btnEdit.attr("title", "保存");
                _this.$btnEdit.html("保存");

                //获取公司数据
                _this.initHtmlInfo();
            }
        },
        initHtmlInfo: function () {
            let _this = this;
            layer.ajax({
                url: _this.urlApiCompanyGetEditInfo,
                type: 'get',
                data: {
                    id: _this.id
                }
            }, {loadingText: "数据加载中，请稍后..."}).done(function (resp) {
                console.log(resp);
                if (0 !== +resp.error) {
                    layer.popupError("数据失败：" + resp.msg);
                    return false;
                }

                //下划线转大驼峰
                if (resp.info.company_info) {
                    $.each(_this.companyFieldToInputNameMap, function (k, v) {
                        $('input[name="' + v + '"]').val(resp.info.company_info[k.toLargeHump()]);
                    })
                }

                if (resp.info.user_info) {
                    $.each(_this.userFieldToInputNameMap, function (k, v) {
                        $('input[name="' + v + '"]').val(resp.info.user_info[k.toLargeHump()]);
                    });
                    _this.email = resp.info.user_info['Email'];
                }

            });
        },
        edit: function () {
            let _this = this;
            let companyInfo = {}, userInfo = {};
            let $needValFields = [];


            $.each(_this.companyFieldToInputNameMap, function (k, v) {
                let $item = $('input[name="' + v + '"]');
                if (!$item.val()) {
                    $needValFields.push($item);
                } else {
                    companyInfo[k] = $item.val().trim();
                }
            });
            $.each(_this.userFieldToInputNameMap, function (k, v) {
                let $item = $('input[name="' + v + '"]');
                if (!$item.val()) {
                    $needValFields.push($item);
                } else {
                    userInfo[k] = $item.val().trim();
                }
            });

            if ($needValFields.length) {
                $needValFields[0].focus();
                _this.$btnEdit.removeClass("disabled");
                layer.popupDanger('请将信息填写完整！');
                return;
            }

            if (_this.operateType === 2) {
                if (!_this.id) {
                    layer.popupDanger('参数错误，请刷新页面重试！');
                    return;
                }
                companyInfo['id'] = _this.id;
            }


            layer.ajax({
                url: _this.urlApiCompanyEdit,
                type: 'post',
                data: {
                    operate_type: _this.operateType,
                    company_info: companyInfo,
                    user_info: userInfo
                }
            }, {loadingText: "操作中，请稍后..."}).done(function (resp) {
                console.log(resp);
                if (0 !== +resp.error) {
                    layer.popupError("操作失败：" + resp.msg);
                    _this.$btnEdit.removeClass("disabled");
                    return false;
                }
                if (_this.operateType === 1) {
                    layer.popupMsg("创建成功！");
                    setTimeout(function () {
                        location.href = _this.urlHtmlCompanyEdit + '?company_id=' + resp.info.id;
                    }, 1000)
                } else {
                    layer.popupMsg("编辑成功！");
                    _this.$btnEdit.removeClass("disabled");
                    return true;
                }
            });
        },
    };

    $(function () {
        company.init();
    });
</script>