<div class="panel panel-default">
    <div class="panel-heading filter-container">
        <h3>{{.Title}}</h3>
    </div>
    <div class="panel-body">
        <div class="btn-group">
            <a href="{{.HtmlUriMiniProgramVersionCreate}}" class="btn btn-sm btn-create disabled" title="版本添加">
                <i class="icon icon-plus"></i> 添加
            </a>
            <a href="javascript:void(0)" class="btn btn-sm btn-edit disabled" title="版本编辑">
                编辑
            </a>
            <a href="javascript:void(0)" class="btn btn-sm btn-approving disabled" title="提交审核">
                提交审核
            </a>
            <a href="javascript:void(0)" class="btn btn-sm btn-online disabled" title="上线">
                上线
            </a>
            <a href="javascript:void(0)" class="btn btn-sm btn-offline disabled" title="下线">
                下线
            </a>
        </div>

        <div class="mt-10">
            <div id="datagrid" class="datagrid"></div>
        </div>
    </div>
    <div class="panel-footer">

    </div>
</div>