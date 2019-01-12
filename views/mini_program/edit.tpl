<div class="panel panel-default">
    <div class="panel-heading filter-container">
        <h3>{{.Title}}</h3>
    </div>
    <div class="panel-body">
        <div class="input-group btn-group">
            <a class="btn btn-info" title=" 创建" onclick="$('.form').submit()">
                <i class="icon icon-plus"></i>创建
            </a>
        </div>
        <div class="search-area">
            <form class="form form-horizontal" method="post" action="">
                <fieldset>
                    <div class="form-group">
                        <label for="serial" class="col-sm-2 col-xs-12">名称</label>
                        <div class="col-md-6 col-sm-8 col-xs-12">
                            <input type="text" class="form-control" name="name" id="name"
                                   value="" placeholder="名称">
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="serial" class="col-sm-2 col-xs-12">类型</label>
                        <div class="col-md-6 col-sm-8 col-xs-12">
                            <select id="finance_type_id" name="finance_type_id" class="form-control">

                                <option value="" selected >1</option>
                                <option value="" selected >2</option>
                            </select>
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="serial" class="col-sm-2 col-xs-12">发生时间</label>
                        <div class="col-md-6 col-sm-8 col-xs-12">
                                <input type="text" class="form-control form-datetime" name="deal_at" id="deal_at"
                                   value="" placeholder="请选择发生时间" readonly>
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="serial" class="col-sm-2 col-xs-12">涉及金额</label>
                        <div class="col-md-6 col-sm-8 col-xs-12">
                            <input type="number" class="form-control" name="amount" id="amount"
                                   value="" placeholder="涉及金额">
                        </div>
                    </div>
                    <div class="form-group">
                        <label for="serial" class="col-sm-2 col-xs-12">备注</label>
                        <div class="col-md-6 col-sm-8 col-xs-12">
                            <input type="text" class="form-control" name="remark" id="remark"
                                   value="" placeholder="备注">
                        </div>
                    </div>
                </fieldset>
            </form>
        </div>
    </div>
</div>