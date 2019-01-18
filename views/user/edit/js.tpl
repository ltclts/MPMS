<script type="text/javascript">

    let user = {
        urlApiUserGetUserInfo:{{.ApiUriUserGetUserInfo}},
        urlApiUserInfoChange:{{.ApiUriUserInfoChange}},
        urlHtmlUserLogin:{{.HtmlUriUserLogin}},
        $btnEdit: $('.btn-edit'),
        $email: $('input[name="email"]'),
        $phone: $('input[name="phone"]'),
        $newPwd: $('input[name="new-password"]'),
        $newPwdConfirm: $('input[name="new-password-confirm"]'),
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
            }).on('change', 'input[name="new-password-confirm"]', function () {
                if (_this.$newPwdConfirm.val() !== _this.$newPwd.val()) {
                    layer.popupDanger('确认密码与新密码不一致！');
                    _this.$newPwdConfirm.val('');
                    _this.$newPwdConfirm.focus();
                }
            });
        },
        initHtml: function () {
            let _this = this;
            layer.ajax({
                url: _this.urlApiUserGetUserInfo,
                type: 'get',
            }, {loadingText: "数据加载中，请稍后..."}).done(function (resp) {
                console.log(resp);
                if (0 !== +resp.error) {
                    layer.popupError("获取数据失败：" + resp.msg);
                    return false;
                }

                _this.$email.val(resp.info.user['Email']);
                _this.$phone.val(resp.info.user['Phone']);
            });
        },
        edit: function () {
            let _this = this,
                phone = _this.$phone.val(),
                newPwd = _this.$newPwd.val(),
                newPwdConfirm = _this.$newPwdConfirm.val();

            if (!phone) {
                _this.$phone.focus();
                _this.$btnEdit.removeClass("disabled");
                layer.popupDanger('请将信息补充完整！');
                return;
            }
            if (!newPwd) {
                _this.$newPwd.focus();
                _this.$btnEdit.removeClass("disabled");
                layer.popupDanger('请将信息补充完整！');
                return;
            }
            if (!newPwdConfirm) {
                _this.$newPwdConfirm.focus();
                _this.$btnEdit.removeClass("disabled");
                layer.popupDanger('请将信息补充完整！');
                return;
            }

            if (newPwd !== newPwdConfirm) {
                layer.popupDanger('确认密码与新密码不一致！');
                _this.$newPwdConfirm.val('');
                _this.$newPwdConfirm.focus();
                _this.$btnEdit.removeClass("disabled");
                return;
            }

            layer.ajax({
                url: _this.urlApiUserInfoChange,
                type: 'post',
                data: {phone: phone, new_password: newPwd}
            }, {loadingText: "操作中，请稍后..."}).done(function (resp) {
                console.log(resp);
                if (0 !== +resp.error) {
                    layer.popupError("操作失败：" + resp.msg);
                    _this.$btnEdit.removeClass("disabled");
                    return false;
                }
                layer.popupMsg("信息修改成功！");
                setTimeout(function () {
                    location.href = _this.urlHtmlUserLogin;
                }, 1000)
            });
        },
    };

    $(function () {
        user.init();
    });
</script>