<script type="text/javascript">

    let company = {
        id:{{.Id}},
        operateType: +{{.OperateType}},
        urlApiCompanyGetEditInfo:{{.ApiUriCompanyGetEditInfo}},
        urlHtmlCompanyEdit:{{.HtmlUriCompanyEdit}},
        urlApiCompanyEdit:{{.ApiUriCompanyEdit}},
        $btnEdit: $('.btn-edit'),
        $checkCode: $('.check-code'),
        companyFieldToInputNameMap: {
            'name': 'company-name',
            'short_name': 'company-short-name',
            'expire_at': 'company-expire-at',
            'remark': 'company-remark'
        },
        userFieldToInputNameMap: {
            'name': 'contact-user-name',
            'email': 'contact-user-email',
            'phone': 'contact-user-phone'
        },
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
            }).on('click', '.check-code', function () {
                _this.timeWait(60);
                _this.$checkCode.addClass('disabled');
            });
        },
        timeWait(wait) {
            let _this = this;
            if (wait === 0) {
                _this.$checkCode.removeClass('disabled');
                _this.$checkCode.html('获取验证码');
            } else {
                _this.$checkCode.addClass('disabled');
                _this.$checkCode.html('重新发送(' + wait + 's)');
                setTimeout(function () {
                    _this.timeWait(--wait)
                }, 1000);
            }
        },
        initHtml: function () {
            let _this = this;

            if (1 === _this.operateType) {
                _this.$checkCode.show();
                //增加校验码必传校验
                _this.userFieldToInputNameMap['check_code'] = 'contact-user-check-code';
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
                    })
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
                    companyInfo[k] = $item.val();
                }
            });
            $.each(_this.userFieldToInputNameMap, function (k, v) {
                let $item = $('input[name="' + v + '"]');
                if (!$item.val()) {
                    $needValFields.push($item);
                } else {
                    userInfo[k] = $item.val();
                }
            });

            if ($needValFields.length) {
                $needValFields[0].focus();
                _this.$btnEdit.removeClass("disabled");
                layer.popupDanger('请将信息填写完整！');
                return;
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
                    location.href = _this.urlHtmlCompanyEdit + '?company_id=' + resp.info.id;
                } else {
                    layer.popupMsg("编辑成功！")
                }

            });
        },
    };

    $(function () {
        company.init();
    });
</script>