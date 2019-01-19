<script type="text/javascript">

    let miniProgramVersion = {
        $datagrid: $('#datagrid'),
        $adminGroup: $('.admin-group'),
        MpId:{{.MpId}},
        urlApiMiniProgramVersionList:{{.ApiUriMiniProgramVersionList}},
        urlHtmlMiniProgramVersionEdit:{{.HtmlUriMiniProgramVersionEdit}},
        checkedItem: {},
        $edit: $('.btn-edit'),
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
                } else {
                    _this.checkedItem = {};
                    !_this.$edit.hasClass('disabled') && _this.$edit.addClass('disabled');
                }
                console.log(selectItems);
            }).on('click', '.btn-edit', function () {
                if (!_this.checkedItem) {
                    return;
                }

                location.href = _this.urlHtmlMiniProgramVersionEdit + ("?mini_program_version_id=" + _this.checkedItem.id);
            });
        },
        renderHtml: function () {
            let _this = this;
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
                        {name: 'Code', label: '版本号'},
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