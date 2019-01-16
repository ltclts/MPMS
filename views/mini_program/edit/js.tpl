<script type="text/javascript">

    let mp = {
        id:{{.Id}},
        operateType: +{{.OperateType}},
        companyId: +{{.CompanyId}},
        urlMiniProgramEdit:{{.ApiUriMiniProgramEdit}},
        urlHtmlMiniProgramEdit:{{.HtmlUriMiniProgramEdit}},
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
            if (2 === _this.operateType) {
                _this.$btnEdit.attr("title", "保存");
                _this.$btnEdit.html("保存");
            }
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
                    layer.popupMsg("编辑成功！")
                }

            });
        },
    };

    $(function () {
        mp.init();
    });
</script>