<div class="panel panel-default">
    <div class="panel-heading filter-container">
        <h3>{{.Title}}</h3>
    </div>
    <div class="panel-body">
        <div id="uploaderId" class="uploader">
            <input type="hidden" name="_xsrf" value="{{ .xsrfdata }}">
            <div class="uploader-message text-center">
                <div class="content"></div>
                <button type="button" class="close">×</button>
            </div>
            <div class="uploader-files file-list file-list-lg" data-drag-placeholder="请拖拽文件到此处"></div>
            <div class="uploader-actions">
                <div class="uploader-status pull-right text-muted"></div>
                <button type="button" class="btn btn-link uploader-btn-browse"><i class="icon icon-plus"></i> 选择文件</button>
                <button type="button" class="btn btn-link uploader-btn-start"><i class="icon icon-cloud-upload"></i> 开始上传</button>
            </div>
        </div>
        <div class="img">
        </div>
    </div>
    <div class="panel-footer">
    </div>
</div>