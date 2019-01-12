<script type="text/javascript">

    let company = {
        $datagrid: $('#datagrid'),
        urlGetList:{{.UrlGetList}},
        urlMiniProgramCreate:{{.HtmlUriMiniProgramCreate}},
        checkedItem: {},
        $inUse: $('.btn-in-use'),
        $forbidden: $('.btn-forbidden'),
        $mpCreate: $('.btn-mp-create'),
        init: function () {
            this.render();
        },
        render: function () {
            this.handleEvent();
            this.renderData();
        },
        handleEvent: function () {
            let _this = this;

            $('body').on('click', '.datagrid-row', function () {
                let selectItems = _this.$datagrid.data('zui.datagrid').getCheckItems();
                if (selectItems.length === 1) {
                    _this.checkedItem = selectItems[0];
                    _this.$mpCreate.removeClass('disabled');
                    _this.$inUse.removeClass('disabled');
                    _this.$forbidden.removeClass('disabled');
                } else {
                    _this.checkedItem = {};
                    !_this.$mpCreate.hasClass('disabled') && _this.$mpCreate.addClass('disabled');
                    !_this.$inUse.hasClass('disabled') && _this.$inUse.addClass('disabled');
                    !_this.$forbidden.hasClass('disabled') && _this.$forbidden.addClass('disabled');
                }
                console.log(selectItems);
            }).on('click', '.btn-mp-create', function () {
                if (!_this.checkedItem) {
                    return;
                }
                location.href = _this.urlMiniProgramCreate + "?company_id=" + _this.checkedItem.id;
            });
        },
        renderData() {
            let _this = this;
            layer.ajax({
                url: _this.urlGetList,
                type: 'post',
                data: {}
            }, {loadingText: "数据加载中..."}).done(function (resp) {
                console.log(resp);
                if (0 !== +resp.error) {
                    layer.popupMsg("获取数据失败！" + resp.msg);
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
                        {name: 'name', label: '公司名称'},
                        {name: 'company_contact_user', label: '联系人'},
                        {name: 'company_contact_user_phone', label: '联系电话'},
                        {name: 'mp_count', label: '小程序数量'},
                        {name: 'status', label: '当前状态'},
                        {name: 'creator', label: '创建人'},
                        {name: 'expire_at', label: '过期时间', width: 160},
                    ],
                    array: info.list
                },
                checkable: true,
                sortable: true,
                configs: function (selector) {
                    let len = selector.length;
                    if (len > 2 && selector.substr(0, 2) !== 'R0' && selector.substr(len - 2, len) === 'C4') {
                        console.log(selector);
                        return {
                            html: true,
                            className: 'text-center',
                            style: {
                                color: '#00b8d4',
                                backgroundColor: '#e0f7fa'
                            }
                        }
                    }

                }
            });
        }
    };

    $(function () {
        company.init();

    });
</script>