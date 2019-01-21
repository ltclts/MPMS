<div class="panel panel-default">
    <div class="panel-heading filter-container">
        <h3>{{.Title}}</h3>
    </div>
    <div class="panel-body">
        <div class="list common-content">
            <header class="col-sm-offset-1">
                <h2>类型选择</h2>
            </header>
            <section class="items mt-10">
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">类&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;型</button></span>
                    <select class="form-control" name="type" disabled>
                        {{range $i, $elem := .MiniProgramVersionTypeToNameMap}}
                            <option value="{{$i}}">{{$elem}}</option>
                        {{end}}
                    </select>
                </div>
            </section>
        </div>
        <div class="list common-content business-card-info hidden">
            <header class="col-sm-offset-1">
                <h2>名片信息设置</h2>
            </header>
            <section class="items mt-10">
                <div>
                    <header class="col-sm-offset-2">
                        <h3>轮播图上传（最多可上传4张）</h3>
                    </header>
                    <div class="carousel-img-list col-sm-6 col-xs-12 col-sm-offset-3">
                        <div class="col-sm-3 col-xs-3 carousel-img-template hidden">
                            <img src="" id="">
                        </div>
                    </div>
                    <div id="carouselUploader" class="col-sm-6 col-xs-12 col-sm-offset-3 uploader mt-10">
                        <div class="file-list" data-drag-placeholder="请拖拽文件到此处"></div>
                        <button type="button" class="btn btn-primary uploader-btn-browse"><i
                                    class="icon icon-cloud-upload"></i> 选择文件
                        </button>
                    </div>

                </div>
            </section>
        </div>
        <div class="list common-content">
            <div class="input-group btn-group col-sm-6 col-xs-12  col-sm-offset-1">
                <a class="btn btn-info btn-edit" title="创建">
                    <i class="icon icon-plus"></i>创建
                </a>
            </div>
        </div>
    </div>
</div>