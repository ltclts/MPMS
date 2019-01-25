<script type="text/javascript">

    let mpv = {
        id: +{{.Id}},
        mpId: +{{.MpId}},
        operateType: +{{.OperateType}},
        urlApiMiniProgramVersionGet:{{.ApiUriMiniProgramVersionGet}},
        urlApiMiniProgramVersionUpload:{{.ApiUriMiniProgramVersionUpload}},
        $btnEdit: $('.btn-edit'),
        $type: $('select[name="type"]'),
        $businessCardInfo: $('.business-card-info'),
        $carouselUploader: $('#carouselUploader'),
        $carouselImgTemplate: $('.carousel-img-template'),
        $carouselImgList: $('#carouselImgList'),
        $shareUploader: $('#shareUploader'),
        $shareImgTemplate: $('.share-img-template'),
        $shareImgList: $('#shareImgList'),
        $elegantDemeanorUploader: $('#elegantDemeanorUploader'),
        $elegantDemeanorImgTemplate: $('.elegant-demeanor-img-template'),
        $elegantDemeanorImgList: $('#elegantDemeanorImgList'),
        currentElegantDemeanorImgCount: 0,
        currentShareImgCount: 0,
        currentCarouseImgCount: 0,
        error:{{.Error}},
        notDeal: true,
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
            }).on('click', '.del-flag', function () {
                let $this = $(this), $parent = $this.parent();
                layer.dangerConfirm('确认删除？', function (index) {
                    $parent.addClass('hidden del');
                    layer.popupMsg("删除成功！");
                    layer.close(index);

                    if ($parent.hasClass('carousel-img')) {
                        _this.currentCarouseImgCount--;
                    } else if ($parent.hasClass('share-img')) {
                        _this.currentShareImgCount--;
                    } else if ($parent.hasClass('elegant-demeanor-img')) {
                        _this.currentElegantDemeanorImgCount--;
                    }
                    console.log(_this.currentShareImgCount);
                });
            });
        },
        renderUploader: function () {
            let _this = this, token = $('meta[name=_xsrf]').attr('content');

            //分享图
            _this.$shareUploader.uploader({//分享图上传插件初始化
                autoUpload: true,            // 当选择文件后立即自动进行上传操作
                url: _this.urlApiMiniProgramVersionUpload,
                chunk_size: 0,
                multipart: true,
                multipart_params: function (file, params) {
                    return {id: _this.id, refer_type: 1, current_count: _this.currentShareImgCount};
                },
                headers: {'X-Xsrftoken': token},
                responseHandler: function (responseObject, file) {
                    _this.$shareUploader.data('zui.uploader').removeFile(file.id);
                    let rsp = JSON.parse(responseObject.response);
                    console.log(rsp, file);
                    if (!+rsp.error) {
                        layer.popupMsg('上传成功！');
                        let $shareImgTemplate = _this.$shareImgTemplate.clone();
                        $shareImgTemplate.removeClass('hidden').removeClass('share-img-template').addClass('share-img new').attr('id', rsp.info.resource_id);
                        $shareImgTemplate.find('img').attr('src', rsp.info.url);
                        $shareImgTemplate.appendTo(_this.$shareImgList);
                        _this.currentShareImgCount++;
                        return;
                    }

                    layer.popupError('上传失败：' + rsp.msg);
                }
            });

            //轮播图
            _this.$carouselUploader.uploader({//轮播图上传插件初始化
                autoUpload: true,            // 当选择文件后立即自动进行上传操作
                url: _this.urlApiMiniProgramVersionUpload,
                chunk_size: 0,
                multipart: true,
                multipart_params: function (file, params) {
                    return {id: _this.id, refer_type: 2, current_count: _this.currentShareImgCount};
                },
                headers: {'X-Xsrftoken': token},
                responseHandler: function (responseObject, file) {
                    _this.$carouselUploader.data('zui.uploader').removeFile(file.id);
                    let rsp = JSON.parse(responseObject.response);
                    console.log(rsp, file);
                    if (!+rsp.error) {
                        layer.popupMsg('上传成功！');
                        let $carouselImgTemplate = _this.$carouselImgTemplate.clone();
                        $carouselImgTemplate.removeClass('hidden').removeClass('carousel-img-template').addClass('carousel-img new').attr('id', rsp.info.resource_id);
                        $carouselImgTemplate.find('img').attr('src', rsp.info.url);
                        $carouselImgTemplate.appendTo(_this.$carouselImgList);
                        _this.currentCarouseImgCount++;
                        return;
                    }

                    layer.popupError('上传失败：' + rsp.msg);
                }
            });
            _this.$carouselImgList.sortable();

            //风采图
            _this.$elegantDemeanorUploader.uploader({//风采图上传插件初始化
                autoUpload: true,            // 当选择文件后立即自动进行上传操作
                url: _this.urlApiMiniProgramVersionUpload,
                chunk_size: 0,
                multipart: true,
                multipart_params: function (file, params) {
                    return {id: _this.id, refer_type: 3, current_count: _this.currentShareImgCount};
                },
                headers: {'X-Xsrftoken': token},
                responseHandler: function (responseObject, file) {
                    _this.$elegantDemeanorUploader.data('zui.uploader').removeFile(file.id);
                    let rsp = JSON.parse(responseObject.response);
                    console.log(rsp, file);
                    if (!+rsp.error) {
                        layer.popupMsg('上传成功！');
                        let $elegantDemeanorImgTemplate = _this.$elegantDemeanorImgTemplate.clone();
                        $elegantDemeanorImgTemplate.removeClass('hidden').removeClass('elegant-demeanor-img-template').addClass('elegant-demeanor-img new').attr('id', rsp.info.resource_id);
                        $elegantDemeanorImgTemplate.find('img').attr('src', rsp.info.url);
                        $elegantDemeanorImgTemplate.appendTo(_this.$elegantDemeanorImgList);
                        _this.currentElegantDemeanorImgCount++;
                        return;
                    }

                    layer.popupError('上传失败：' + rsp.msg);
                }
            });
            _this.$elegantDemeanorImgList.sortable();

        },
        initHtml: function () {
            let _this = this;
            if (_this.error) {
                _this.$btnEdit.addClass('disabled');
                layer.popupError(_this.error);
                return
            }

            if (1 === +_this.$type.val()) {
                _this.$businessCardInfo.removeClass('hidden');
                _this.renderUploader();

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
                url: _this.urlApiMiniProgramVersionGet,
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
        mpv.init();
    });
</script>