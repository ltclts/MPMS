<div class="panel panel-default">
    <div class="panel-heading filter-container">
        <h3>{{.Title}}</h3>
    </div>
    <div class="panel-body">
        <div class="btn-group admin-group hidden">
            <a href="{{.HtmlUriCompanyCreate}}" class="btn btn-sm btn-create" title="公司添加">
                <i class="icon icon-plus"></i> 添加
            </a>
            <a href="" class="btn btn-sm btn-edit disabled" title="公司编辑">
                编辑
            </a>
            <a href="javascript:void(0)" class="btn btn-sm btn-primary btn-in-use disabled" title="启用">
                启用
            </a>
            <a href="javascript:void(0)" class="btn btn-sm btn-primary btn-forbidden disabled" title="禁用">
                禁用
            </a>
        </div>
        <div class="btn-group">
            <a href="javascript:void(0)" class="btn btn-sm btn-mp-create disabled" title="小程序添加">
                <i class="icon icon-plus"></i> 小程序添加
            </a>
        </div>

        <div class="mt-10">
            <div id="datagrid" class="datagrid"></div>
        </div>
    </div>
    <div class="panel-footer">

    </div>
</div>