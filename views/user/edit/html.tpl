<div class="panel panel-default">
    <div class="panel-heading filter-container">
        <h3>{{.Title}}</h3>
    </div>
    <div class="panel-body">
        <div class="input-group btn-group">
            <a class="btn btn-info btn-edit" title="保存">
                保存
            </a>
        </div>
        <div class="list common-content">
            <section class="items mt-10">
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">登陆邮箱</button></span>
                    <input type="email" class="form-control" name="email" placeholder="邮箱" disabled>
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">联系电话</button></span>
                    <input type="tel" class="form-control" name="phone" placeholder="手机号码" maxlength="20">
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">新密码&nbsp;&nbsp;&nbsp;</button></span>
                    <input type="password" class="form-control" name="new-password" placeholder="新密码"
                           maxlength="20">
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">确认密码</button></span>
                    <input type="password" class="form-control" name="new-password-confirm" placeholder="确认密码"
                           maxlength="20">
                </div>
            </section>
        </div>
    </div>
</div>