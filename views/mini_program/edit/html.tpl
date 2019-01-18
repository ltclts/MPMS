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

        <div class="list common-content">
            <header>
                <h2>主体信息</h2>
            </header>
            <section class="items mt-10">
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">名&nbsp;&nbsp;&nbsp;称</button></span>
                    <input type="text" class="form-control" name="mp-name" placeholder="小程序名称" maxlength="50">
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default" type="button">Appid</button></span>
                    <input type="text" class="form-control" name="mp-appid" placeholder="微信小程序的appid" maxlength="50" disabled>
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">备&nbsp;&nbsp;&nbsp;注</button></span>
                    <input type="text" class="form-control" name="mp-remark" placeholder="该小程序用途" maxlength="50">
                </div>
            </section>
        </div>
    </div>
</div>