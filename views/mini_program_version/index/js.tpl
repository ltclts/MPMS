<script type="text/javascript">

    let miniProgramVersion = {
        $datagrid: $('#datagrid'),
        $adminGroup: $('.admin-group'),
        MpId:{{.MpId}},
        urlApiMiniProgramVersionList:{{.ApiUriMiniProgramVersionList}},
        urlHtmlMiniProgramVersionEdit:{{.HtmlUriMiniProgramVersionEdit}},
        urlApiMiniProgramVersionUpdateStatus:{{.ApiUriMiniProgramVersionUpdateStatus}},
        checkedItem: {},
        $edit: $('.btn-edit'),
        $create: $('.btn-create'),
        $online: $('.btn-online'),
        $offline: $('.btn-offline'),
        $approving: $('.btn-approving'),
        init: function () {
            this.render();
        },
        render: function () {
            this.handleEvent();
            this.renderHtml();
            this.renderData();
        },
        handleEvent: function () {
            let _this = this;

            $('body').on('click', '.datagrid-row', function () {
                let selectItems = _this.$datagrid.data('zui.datagrid').getCheckItems();
                if (selectItems.length === 1) {
                    _this.checkedItem = selectItems[0];
                    _this.$edit.removeClass('disabled');
                    (+_this.checkedItem.Status === 0) && _this.$approving.removeClass('disabled'); //初始状态才能提交审核
                    ((+_this.checkedItem.Status === 2) || (+_this.checkedItem.Status === 4)) &&
                    _this.$online.removeClass('disabled'); //审核后或者已下线才能上线
                    (+_this.checkedItem.Status === 3) && _this.$offline.removeClass('disabled'); //上线后才能下线
                } else {
                    _this.checkedItem = {};
                    !_this.$edit.hasClass('disabled') && _this.$edit.addClass('disabled');
                    !_this.$online.hasClass('disabled') && _this.$online.addClass('disabled');
                    !_this.$offline.hasClass('disabled') && _this.$offline.addClass('disabled');
                    !_this.$approving.hasClass('disabled') && _this.$approving.addClass('disabled');
                }
                console.log(selectItems);
            }).on('click', '.btn-edit', function () {
                if (!_this.checkedItem) {
                    return;
                }

                location.href = _this.urlHtmlMiniProgramVersionEdit + ("?mini_program_version_id=" + _this.checkedItem.Id);
            }).on('click', '.btn-online', function () {
                _this.updateStatus(_this.checkedItem.Status, 3, _this.checkedItem.Id)
            }).on('click', '.btn-offline', function () {
                _this.updateStatus(3, 4, _this.checkedItem.Id)
            }).on('click', '.btn-approving', function () {
                _this.updateStatus(0, 1, _this.checkedItem.Id)
            });
        },
        updateStatus: function (from_status, to_status, id) {
            let _this = this;
            let statusToActNamesMap = {
                1: '确认提交审核？提交审核后的版本将禁止更新！',
                3: '确认上线？若有其他上线版本，将被强制下线！',
                4: '确认将此版本下线？'
            };
            if (!statusToActNamesMap[to_status]) {
                return;
            }
            layer.dangerConfirm(statusToActNamesMap[to_status], function () {
                layer.ajax({
                    url: _this.urlApiMiniProgramVersionUpdateStatus,
                    type: 'post',
                    data: {id: id, to_status: to_status, from_status: from_status}
                }, {loadingText: "操作处理中..."}).done(function (resp) {
                    if (0 !== +resp.error) {
                        layer.popupError("操作处理失败！" + resp.msg);
                        return false;
                    }
                    layer.popupMsg("操作成功！");
                    setTimeout(function(){
                        location.reload();
                    },500)
                });
            })
        },
        renderHtml: function () {
            let _this = this;
            if (_this.MpId) {
                _this.$create.removeClass('disabled');
            }
        },
        renderData() {
            let _this = this;
            layer.ajax({
                url: _this.urlApiMiniProgramVersionList,
                type: 'post',
                data: {mini_program_id: _this.MpId}
            }, {loadingText: "数据加载中..."}).done(function (resp) {
                console.log(resp);
                if (0 !== +resp.error) {
                    layer.popupError("获取数据失败！" + resp.msg);
                    return false;
                }
                _this.renderDataTable(resp.info);
            });
        },
        renderDataTable: function (info) {
            let _this = this;
            _this.$datagrid.datagrid({
                dataSource: {
                    cols: [
                        {name: 'CreatedAt', label: '创建时间'},
                        {name: 'StatusName', label: '状态'},
                        {name: 'MpName', label: '小程序'},
                        {name: 'CShortName', label: '所属'},
                    ],
                    array: info.list
                },
                checkable: true,
                sortable: true
            });
        }
    };

    $(function () {
        miniProgramVersion.init();

    });
</script>