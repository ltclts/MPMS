<script type="text/javascript">

    let mp = {
        $datagrid: $('#datagrid'),
        $adminGroup: $('.admin-group'),
        urlApiMiniProgramList:{{.ApiUriMiniProgramList}},
        urlHtmlMiniProgramEdit:{{.HtmlUriMiniProgramEdit}},
        urlHtmlMiniProgramVersion:{{.HtmlUriMiniProgramVersion}},
        userType:{{.UserType}},
        companyId:{{.CompanyId}},
        checkedItem: {},
        $versionManage: $('.btn-version-manage'),
        $edit: $('.btn-edit'),
        $create: $('.btn-create'),
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
                    _this.$versionManage.removeClass('disabled');
                    _this.$edit.removeClass('disabled');
                } else {
                    _this.checkedItem = {};
                    !_this.$versionManage.hasClass('disabled') && _this.$versionManage.addClass('disabled');
                    !_this.$edit.hasClass('disabled') && _this.$edit.addClass('disabled');
                }
                console.log(selectItems);
            }).on('click', '.btn-edit', function () {
                if (!_this.checkedItem) {
                    return;
                }

                location.href = _this.urlHtmlMiniProgramEdit + "?mp_id=" + _this.checkedItem.id;
            }).on('click', '.btn-version-manage', function () {
                //启用
                if (!_this.checkedItem) {
                    return;
                }

                location.href = _this.urlHtmlMiniProgramVersion + "?mp_id=" + _this.checkedItem.id;

                console.log('版本管理');
            });
        },
        renderHtml: function () {
            let _this = this;
            if (_this.userType === 0) {
                _this.$create.addClass('disabled');
            }

        },
        renderData() {
            let _this = this;
            layer.ajax({
                url: _this.urlApiMiniProgramList,
                type: 'post',
                data: {company_id: _this.companyId}
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
                        {name: 'created_at', label: '创建时间'},
                        {name: 'name', label: '名称'},
                        {name: 'appid', label: 'Appid'},
                        {name: 'version_count', label: '版本数'},
                    ],
                    array: info.list
                },
                checkable: true,
                sortable: true
            });
        }
    };

    $(function () {
        mp.init();
    });
</script>