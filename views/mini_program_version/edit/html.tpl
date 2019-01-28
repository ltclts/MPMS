<div class="panel panel-default">
    <div class="panel-heading filter-container">
        <h3>{{.Title}}</h3>
    </div>
    <div class="panel-body">
        <div class="list common-content col-sm-12 col-xs-12">
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
        <div class="list common-content col-sm-12 col-xs-12 share-info">
            <header class="col-sm-offset-1">
                <h2>分享设置</h2>
            </header>
            <section class="items col-sm-12 col-xs-12 mt-10">
                <header class="col-sm-offset-2">
                    <h3>分享图上传</h3>
                </header>
                <div class="col-sm-6 col-xs-12 share-img-template hidden" id="">
                    <p class="del-flag" title="点击可进行删除">删除</p>
                    <img src="#">
                </div>
                <div class="share-img-list list-group col-sm-6 col-xs-12 col-sm-offset-3 mt-10" id="shareImgList">

                </div>
                <div id="shareUploader" class="col-sm-6 col-xs-12 col-sm-offset-3 uploader">
                    <div class="file-list" data-drag-placeholder="请拖拽文件到此处"></div>
                    <button type="button" class="btn btn-primary uploader-btn-browse"><i
                                class="icon icon-cloud-upload"></i> 选择文件
                    </button>
                </div>
                <div class="input-group col-sm-6 col-xs-12 col-sm-offset-3">
                    <span class="input-group-btn"><button class="btn btn-default"
                                                          type="button">分享寄语</button></span>
                    <input type="text" class="form-control" name="share-words" placeholder="分享寄语" maxlength="50">
                </div>
            </section>
        </div>
        <div class="list common-content col-sm-12 col-xs-12 business-card-info hidden">
            <header class="col-sm-offset-1">
                <h2>名片信息设置</h2>
            </header>
            <section class="items col-sm-12 col-xs-12 mt-10">
                <header class="col-sm-offset-2">
                    <h3>轮播图上传（最多可上传4张）</h3>
                </header>
                <div class="col-sm-6 col-xs-6 carousel-img-template hidden" id="">
                    <p class="del-flag" title="点击可进行删除">删除</p>
                    <img src="#">
                </div>
                <div class="carousel-img-list list-group col-sm-6 col-xs-12 col-sm-offset-3 mt-10"
                     id="carouselImgList"
                     title="拖动可进行排序">
                </div>
                <div id="carouselUploader" class="col-sm-6 col-xs-12 col-sm-offset-3 uploader">
                    <div class="file-list" data-drag-placeholder="请拖拽文件到此处"></div>
                    <button type="button" class="btn btn-primary uploader-btn-browse"><i
                                class="icon icon-cloud-upload"></i> 选择文件
                    </button>
                </div>
            </section>
            <section class="items col-sm-12 col-xs-12 mt-10">
                <header class="col-sm-offset-2">
                    <h3>风采图上传（最多可上传4张）</h3>
                </header>
                <div class="col-sm-12 col-xs-12 elegant-demeanor-img-template hidden" id="">
                    <p class="del-flag" title="点击可进行删除">删除</p>
                    <img src="#">
                </div>
                <div class="elegant-demeanor-img-list list-group col-sm-6 col-xs-12 col-sm-offset-3 mt-10"
                     id="elegantDemeanorImgList"
                     title="拖动可进行排序">
                </div>
                <div id="elegantDemeanorUploader" class="col-sm-6 col-xs-12 col-sm-offset-3 uploader">
                    <div class="file-list" data-drag-placeholder="请拖拽文件到此处"></div>
                    <button type="button" class="btn btn-primary uploader-btn-browse"><i
                                class="icon icon-cloud-upload"></i> 选择文件
                    </button>
                </div>
            </section>
        </div>
        <div class="list common-content col-sm-12 col-xs-12">
            <div class="input-group btn-group col-sm-6 col-xs-12  col-sm-offset-1">
                <a class="btn btn-info btn-edit" title="创建">
                    <i class="icon icon-plus"></i>创建
                </a>
            </div>
        </div>
    </div>
</div>