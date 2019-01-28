<script type="text/javascript">

    let mpv = {
        id: +{{.Id}},
        mpId: +{{.MpId}},
        operateType: +{{.OperateType}},
        urlHtmlMiniProgramVersionEdit:{{.HtmlUriMiniProgramVersionEdit}},
        urlApiMiniProgramVersionGet:{{.ApiUriMiniProgramVersionGet}},
        urlApiMiniProgramVersionUpload:{{.ApiUriMiniProgramVersionUpload}},
        urlApiMiniProgramVersionEdit:{{.ApiUriMiniProgramVersionEdit}},
        $btnEdit: $('.btn-edit'),
        $type: $('select[name="type"]'),
        $shareWords: $('input[name="share-words"]'),
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
        token: $('meta[name=_xsrf]').attr('content'),
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
            let _this = this, token = _this.token;

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
                        _this.renderShareImg(rsp.info.resource_id, rsp.info.url, true);
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
                        _this.renderCarouselImg(rsp.info.resource_id, rsp.info.url, true);
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
                        _this.renderElegantDemeanorImg(rsp.info.resource_id, rsp.info.url, true);
                        return;
                    }

                    layer.popupError('上传失败：' + rsp.msg);
                }
            });
            _this.$elegantDemeanorImgList.sortable();

        },
        renderShareImg: function (id, src, add) {
            let _this = this, addClass = 'share-img';

            if (add) {
                addClass += ' new';
            }

            let $shareImgTemplate = _this.$shareImgTemplate.clone();
            $shareImgTemplate.removeClass('hidden').removeClass('share-img-template').addClass(addClass).attr('id', id);
            $shareImgTemplate.find('img').attr('src', src);
            $shareImgTemplate.appendTo(_this.$shareImgList);
            _this.currentShareImgCount++;
        },
        renderCarouselImg: function (id, src, add) {
            let _this = this, addClass = 'carousel-img';

            if (add) {
                addClass += ' new';
            }
            let $carouselImgTemplate = _this.$carouselImgTemplate.clone();
            $carouselImgTemplate.removeClass('hidden').removeClass('carousel-img-template').addClass(addClass).attr('id', id);
            $carouselImgTemplate.find('img').attr('src', src);
            $carouselImgTemplate.appendTo(_this.$carouselImgList);
            _this.currentCarouseImgCount++;
        },
        renderElegantDemeanorImg: function (id, src, add) {
            let _this = this, addClass = 'elegant-demeanor-img';

            if (add) {
                addClass += ' new';
            }
            let $elegantDemeanorImgTemplate = _this.$elegantDemeanorImgTemplate.clone();
            $elegantDemeanorImgTemplate.removeClass('hidden').removeClass('elegant-demeanor-img-template').addClass(addClass).attr('id', id);
            $elegantDemeanorImgTemplate.find('img').attr('src', src);
            $elegantDemeanorImgTemplate.appendTo(_this.$elegantDemeanorImgList);
            _this.currentElegantDemeanorImgCount++;
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

                let carouselImgList = resp.info.CarouselImgList || [];
                carouselImgList.forEach(function (v) {
                    _this.renderCarouselImg(v.Id, v.Path, false);
                });

                let shareImgList = resp.info.ShareImgList || [];
                shareImgList.forEach(function (v) {
                    _this.renderShareImg(v.Id, v.Path, false);
                });

                let elegantDemeanorImgList = resp.info.ElegantDemeanorImgList || [];
                elegantDemeanorImgList.forEach(function (v) {
                    _this.renderElegantDemeanorImg(v.Id, v.Path, false);
                });

                let item = resp.info.Version || {};
                _this.mpId = item.MpId || 0;
                _this.id = item.Id || 0;
                _this.$shareWords.val(item.ShareWords || "");
            });
        },
        getBusinessCardEditInfo: function () {
            let _this = this;
            return {'carousel_info': _this.getCarouselInfo(), 'elegant_demeanor_info': _this.getElegantDemeanorInfo()};
        },
        getShareInfo: function () {
            //共有属性获取
            //分享图片与寄语
            let _this = this, shareInfo = {}, imgToAdd = [], imgToDel = [];
            shareInfo['share_words'] = _this.$shareWords.val();
            _this.$shareImgList.find('.share-img').each(function () {
                let $this = $(this), id = +$this.attr('id');
                if ($this.hasClass('del')) {
                    imgToDel.push(id);
                    return true;
                }

                if ($this.hasClass('new')) {
                    imgToAdd.push(id);
                }
            });
            shareInfo['img_to_add'] = imgToAdd;
            shareInfo['img_to_del'] = imgToDel;
            return shareInfo;
        },
        getCarouselInfo: function () {
            //共有属性获取
            //分享图片与寄语
            let _this = this, imgToAdd = [], imgToDel = [], imgToSort = [];
            _this.$carouselImgList.find('.carousel-img').each(function () {
                let $this = $(this), id = +$this.attr('id');
                imgToSort.push(id); //所有的id
                if ($this.hasClass('del')) {
                    imgToDel.push(id);
                    return true;
                }

                if ($this.hasClass('new')) {
                    imgToAdd.push(id);
                }
            });
            return {'img_to_sort': imgToSort, 'img_to_add': imgToAdd, 'img_to_del': imgToDel};
        },
        getElegantDemeanorInfo: function () {
            //共有属性获取
            //分享图片与寄语
            let _this = this, imgToAdd = [], imgToDel = [], imgToSort = [];
            _this.$elegantDemeanorImgList.find('.elegant-demeanor-img').each(function () {
                let $this = $(this), id = +$this.attr('id');
                imgToSort.push(id); //所有的id
                if ($this.hasClass('del')) {
                    imgToDel.push(id);
                    return true;
                }

                if ($this.hasClass('new')) {
                    imgToAdd.push(id);
                }
            });
            return {'img_to_sort': imgToSort, 'img_to_add': imgToAdd, 'img_to_del': imgToDel};
        },
        edit: function () {
            let _this = this,
                type = +_this.$type.val(),
                param = {};

            if (1 === type) { //名片展示
                param = _this.getBusinessCardEditInfo();
            }

            if (_this.operateType === 1) {
                if (!_this.mpId) {
                    layer.popupDanger('参数错误，请刷新页面重试！');
                    return;
                }
            }

            if (_this.operateType === 2) {
                if (!_this.id) {
                    layer.popupDanger('参数错误，请刷新页面重试！');
                    return;
                }
            }

            layer.ajax({
                url: _this.urlApiMiniProgramVersionEdit,
                type: 'post',
                data: {
                    operate_type: _this.operateType,
                    'id': _this.id,
                    'mp_id': _this.mpId,
                    'type': type,
                    'share_info': _this.getShareInfo(),
                    param: param
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
                        location.href = _this.urlHtmlMiniProgramVersionEdit + '?mini_program_version_id=' + resp.info.id;
                    }, 1000)
                } else {
                    layer.popupMsg("编辑成功！");
                    _this.$btnEdit.removeClass("disabled");
                    setTimeout(function () {
                        location.reload();
                    }, 1000);
                    return true;
                }
            });
        },
    };

    $(function () {
        mpv.init();
    });
</script>