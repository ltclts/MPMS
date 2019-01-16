<div class="panel panel-default">
    <div class="panel-heading filter-container">
        <h3>{{.Title}}</h3>
    </div>
    <div class="panel-body">
        <div class="input-group btn-group">
            <a class="btn btn-info btn-edit" title="创建">
                <i class="icon icon-plus"></i>创建
            </a>
        </div>

        <div class="list common-content company-info">
            <header>
                <h2>主体信息</h2>
            </header>
            <section class="items mt-10">
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">名&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;称</button></span>
                    <input type="text" class="form-control bg-danger" name="company-name" placeholder="公司全称"
                           maxlength="50">
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">简&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;称</button></span>
                    <input type="text" class="form-control" name="company-short-name" placeholder="公司简称" maxlength="50">
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">到期时间</button></span>
                    <input type="text" class="form-control form-datetime" name="company-expire-at"
                           placeholder="过期时间 YYYY-MM-DD H:i:s">
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">备&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;注</button></span>
                    <input type="text" class="form-control" name="company-remark" placeholder="备注(公司简要介绍)"
                           maxlength="80">
                </div>
            </section>
        </div>
        <div class="list common-content contact-info">
            <header>
                <h2>联系人信息</h2>
            </header>
            <section class="items mt-10">
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">姓&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;名</button></span>
                    <input type="text" class="form-control" name="contact-user-name" placeholder="联系人姓名" maxlength="50">
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">邮&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;箱</button></span>
                    <input type="email" class="form-control" name="contact-user-email" placeholder="登陆邮箱"
                           maxlength="50">
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3 hidden">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">邮箱验证</button></span>
                    <input type="text" class="form-control" name="contact-user-check-code" placeholder="验证码"
                           maxlength="50">
                    <span class="input-group-btn"><button class="btn btn-primary check-code"
                                                          type="button">获取验证码</button></span>
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">联系电话</button></span>
                    <input type="tel" class="form-control" name="contact-user-phone" placeholder="手机号码" maxlength="20">
                </div>
            </section>
        </div>
    </div>
</div>