<div class="panel panel-default">
    <div class="panel-heading filter-container">
        <h3>{{.Title}}</h3>
    </div>
    <div class="panel-body">
        <div class="btn-group admin-group">
            <a href="{{.HtmlUriMiniProgramCreate}}" class="btn btn-sm btn-create" title="添加">
                <i class="icon icon-plus"></i> 添加
            </a>
            <a href="javascript:void(0)" class="btn btn-sm btn-edit disabled" title="编辑">
                编辑
            </a>
            <a href="javascript:void(0)" class="btn btn-sm btn-version-manage disabled" title="版本管理">
                版本管理
            </a>
        </div>

        <div class="mt-10">
            <div id="datagrid" class="datagrid"></div>
        </div>
    </div>
    <div class="panel-footer">

    </div>
</div>