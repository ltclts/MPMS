<!DOCTYPE html>
<html lang="en_US">
<head>
    <link rel="bookmark" type="image/x-icon" href="/static/img/logo.icon"/>
    <link rel="shortcut icon" href="/static/img/logo.icon">

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="chrome=1,IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <link rel="stylesheet" href="/static/components/zui/css/zui.min.css">
    <link rel="stylesheet" href="/static/components/layui/css/layui.css">
    <link rel="stylesheet" href="/static/components/zui/lib/chosen/chosen.css">
    <link rel="stylesheet" href="/static/css/index.css">

    <link rel="stylesheet" type="text/css" href="/static/css/auth/normalize.css"/>
    <link rel="stylesheet" type="text/css" href="/static/css/auth/demo.css"/>
    <link rel="stylesheet" type="text/css" href="/static/css/auth/component.css"/>
    <meta name="_xsrf" content="{{.xsrfdata}}"/>
    <title>{{.Title}}</title>
</head>
<body>

<div class="auth-container">
    <div class="content">
        <div id="large-header" class="large-header">
            <div style="text-align: center" class="alert hide">
                <span class="alert-info"></span>
            </div>

            <canvas id="demo-canvas"></canvas>
            <div class="logo-box">
                <h3>{{.Title}}</h3>
                <div class="input_outer">
                    <span class="u_email"></span>
                    <input name="email" class="text" type="text" value="" placeholder="请输入邮箱">
                </div>
                <div class="input_outer">
                    <span class="us_uer"></span>
                    <input name="password" class="text" type="password" placeholder="请输入密码">
                </div>
                <div class="mb2">
                    <a class="act-but submit">登录</a>
                </div>
            </div>
        </div>
    </div>
</div>
<script src="/static/components/zui/lib/jquery/jquery.js"></script>
<script src="/static/js/auth/html5.js"></script>
<script src="/static/js/auth/TweenLite.min.js"></script>
<script src="/static/js/auth/EasePack.min.js"></script>
<script src="/static/js/auth/rAF.js"></script>
<script src="/static/js/auth/demo-1.js"></script>
<script>
    let login = {
        urlLogin: "/api/user/login",
        _xsrf: $('meta[name="_xsrf"]').attr('content'),
        $email: $('input[name="email"]'),
        $password: $('input[name="password"]'),
        $alert: $('.alert'),
        $alertInfo: $('.alert-info'),
        init: function () {
            this.render();
        },
        render: function () {
            let _this = this;
            console.log(_this._xsrf);
            $('body').on('click', '.submit', function () {
                _this.login();
            });
        },
        showErrMsg: function (msg) {
            this.$alert.removeClass('hide');
            this.$alertInfo.html(msg);
        },
        login: function () {
            const _this = this;
            let email = _this.$email.val(), pwd = _this.$password.val();

            if (!email) {
                _this.showErrMsg('请输入用户名！');
                return;
            }

            if (!pwd) {
                _this.showErrMsg('请输入密码！');
                return;
            }

            $.ajax({
                url: _this.urlLogin,
                type: 'post',
                headers: {'X-Xsrftoken': _this._xsrf},
                data: {email: email, password: pwd}
            }).done(function (resp) {
                console.log(resp);
                if (0 !== +resp.error) {
                    _this.showErrMsg(resp.msg);
                    return false;
                }
                location.href = resp.info.uri;
            });
        }

    };

    $(function () {
        login.init();
    });
</script>
</body>
</html>