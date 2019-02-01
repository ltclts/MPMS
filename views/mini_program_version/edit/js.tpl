<script charset="utf-8" src="{{.MapScriptSrc}}"></script>
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
        $lng: $('input[name="lng"]'),
        $lat: $('input[name="lat"]'),
        token: $('meta[name=_xsrf]').attr('content'),
        currentElegantDemeanorImgCount: 0,
        currentShareImgCount: 0,
        currentCarouseImgCount: 0,
        error:{{.Error}},
        notDeal: true,
        map: null,
        contentFields: ['name', 'flag', 'tel', 'address', 'lng', 'lat', 'map_key', 'column_other_name'],
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
        renderMap: function () {
            let _this = this;
            _this.map = new BMap.Map('map_container');
            let opts = {type: BMAP_NAVIGATION_CONTROL_SMALL, anchor: BMAP_ANCHOR_TOP_RIGHT};
            _this.map.addControl(new BMap.NavigationControl(opts));
            _this.map.enableScrollWheelZoom(true);     //开启鼠标滚轮缩放
            _this.map.setZoom(15);
            _this.mapTo(116.404, 39.915, true);
            _this.map.addEventListener('click', function (e) {
                _this.mapTo(e.point.lng, e.point.lat, false);
            });
        },
        mapTo: function (lng, lat, move) {
            console.log(lng, lat);
            let _this = this;
            //获取地图上所有的覆盖物
            let allOverlay = _this.map.getOverlays();
            for (let i = 0; i < allOverlay.length; i++) {
                _this.map.removeOverlay(allOverlay[i]);
            }

            let point = new BMap.Point(lng, lat);  // 创建点坐标
            if (move) {
                _this.map.centerAndZoom(point, _this.map.getZoom());
            }
            let marker = new BMap.Marker(point);        // 创建标注
            _this.map.addOverlay(marker); // 将标注添加到地图中

            _this.$lng.val(lng);
            _this.$lat.val(lat);
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
                _this.renderMap();

                if (1 === _this.operateType) {
                    //创建
                    _this.renderUploader();

                }

                if (2 === _this.operateType) {
                    _this.$btnEdit.attr("title", "保存");
                    _this.$btnEdit.html("保存");

                    //获取公司数据
                    _this.initHtmlInfo();
                }
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

                let item = resp.info.Version || {};
                _this.mpId = item.MpId || 0;
                _this.id = item.Id || 0;
                _this.$shareWords.val(item.ShareWords || "");
                let content = JSON.parse(item.Content);
                _this.mapTo(content.lng, content.lat, true);
                console.log(content);
                _this.contentFields.forEach(function (name) {
                    $('input[name="' + name + '"]').val(content[name]);
                });

                if (+item.Status !== 0) {
                    //非初始状态 将被禁用所有编辑
                    $.each($('input'), function () {
                        $(this).attr('disabled', true);
                    });
                    $.each($('button'), function () {
                        $(this).attr('disabled', true);
                    });
                } else {
                    _this.renderUploader();
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
            });
        },
        getBusinessCardEditInfo: function () {
            let _this = this;
            let content = {};
            _this.contentFields.forEach(function (name) {
                content[name] = $('input[name="' + name + '"]').val();
            });
            return {
                'carousel_info': _this.getCarouselInfo(),
                'elegant_demeanor_info': _this.getElegantDemeanorInfo(),
                'content': content
            };
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
        clearImg: function () {
            let _this = this;
            _this.$shareImgList.find('.share-img').each(function () {
                let $this = $(this);
                if ($this.hasClass('del')) {
                    $this.remove();
                    return true;
                }

                if ($this.hasClass('new')) {
                    $this.removeClass('new');
                }
            });
            _this.$carouselImgList.find('.carousel-img').each(function () {
                let $this = $(this);
                if ($this.hasClass('del')) {
                    $this.remove();
                    return true;
                }

                if ($this.hasClass('new')) {
                    $this.removeClass('new');
                }
            });
            _this.$elegantDemeanorImgList.find('.elegant-demeanor-img').each(function () {
                let $this = $(this);
                if ($this.hasClass('del')) {
                    $this.remove();
                    return true;
                }

                if ($this.hasClass('new')) {
                    $this.removeClass('new');
                }
            });
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
                    _this.clearImg();
                    return true;
                }
            });
        },
    };

    $(function () {
        mpv.init();
    });
</script>