<script type="text/javascript">

    let mp = {
        id:{{.Id}},
        operateType: +{{.OperateType}},
        companyId: +{{.CompanyId}},
        urlMiniProgramEdit:{{.ApiUriMiniProgramEdit}},
        urlHtmlMiniProgramEdit:{{.HtmlUriMiniProgramEdit}},
        urlApiMiniProgramList:{{.ApiUriMiniProgramList}},
        $btnEdit: $('.btn-edit'),
        $mpName: $('input[name="mp-name"]'),
        $mpAppid: $('input[name="mp-appid"]'),
        $mpRemark: $('input[name="mp-remark"]'),
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


            console.log(_this.operateType);
        },
        initHtml: function () {
            let _this = this;

            if (1 === _this.operateType) {
                _this.$mpAppid.removeClass('disabled');
            }

            if (2 === _this.operateType) {
                _this.$btnEdit.attr("title", "保存");
                _this.$btnEdit.html("保存");
                _this.initHtmlInfo();
            }
        },
        initHtmlInfo: function () {
            let _this = this;
            layer.ajax({
                url: _this.urlApiMiniProgramList,
                type: 'post',
                data: {
                    id: _this.id
                }
            }, {loadingText: "数据加载中，请稍后..."}).done(function (resp) {
                console.log(resp);
                if (0 !== +resp.error || !resp.info || !resp.info.list) {
                    layer.popupError("获取数据失败!");
                    return false;
                }

                if (resp.info.list[0]) {
                    _this.$mpName.val(resp.info.list[0].name);
                    _this.$mpAppid.val(resp.info.list[0].appid);
                    _this.$mpRemark.val(resp.info.list[0].remark);
                }

            });
        },
        edit: function () {
            let _this = this;
            let mpInfo = {};
            mpInfo['name'] = _this.$mpName.val();
            if (!mpInfo['name']) {
                layer.popupImportant('请填写小程序名称！');
                _this.$btnEdit.removeClass("disabled");
                return false;
            }
            mpInfo['appid'] = _this.$mpAppid.val();
            if (!mpInfo['appid']) {
                layer.popupImportant('请填写小程序Appid！');
                _this.$btnEdit.removeClass("disabled");
                return false;
            }
            mpInfo['remark'] = _this.$mpRemark.val();
            mpInfo['company_id'] = _this.companyId;

            if (2 === _this.operateType) {
                if (!_this.id) {
                    layer.popupImportant('参数缺失，请刷新重试！');
                    _this.$btnEdit.removeClass("disabled");
                    return false;
                }
                mpInfo['id'] = _this.id;
            }

            layer.ajax({
                url: _this.urlMiniProgramEdit,
                type: 'post',
                data: {
                    operate_type: _this.operateType,
                    mp_info: mpInfo,
                }
            }, {loadingText: "操作中，请稍后..."}).done(function (resp) {
                console.log(resp);
                if (0 !== +resp.error) {
                    layer.popupError("操作失败：" + resp.msg);
                    _this.$btnEdit.removeClass("disabled");
                    return false;
                }
                if (_this.operateType === 1) {
                    location.href = _this.urlHtmlMiniProgramEdit + '?mp_id=' + resp.info.id;
                } else {
                    layer.popupMsg("编辑成功！");
                    _this.$btnEdit.removeClass("disabled");
                }

            });
        },
    };

    $(function () {
        mp.init();
    });
</script>