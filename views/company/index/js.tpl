<script type="text/javascript">

    let company = {
        $datagrid: $('#datagrid'),
        $adminGroup: $('.admin-group'),
        urlGetList:{{.UrlGetList}},
        urlHtmlCompanyEdit:{{.HtmlUriCompanyEdit}},
        urlMiniProgramCreate:{{.HtmlUriMiniProgramCreate}},
        urlApiCompanyUpdateStatus:{{.ApiUriCompanyUpdateStatus}},
        urlHtmlMiniProgram:{{.HtmlUriMiniProgram}},
        companyId:{{.CompanyId}},
        checkedItem: {},
        $inUse: $('.btn-in-use'),
        $edit: $('.btn-edit'),
        $forbidden: $('.btn-forbidden'),
        $mpCreate: $('.btn-mp-create'),
        $mpList: $('.btn-mp-list'),
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
                    _this.$mpCreate.removeClass('disabled');
                    _this.$inUse.removeClass('disabled');
                    _this.$forbidden.removeClass('disabled');
                    _this.$edit.removeClass('disabled');
                    _this.$mpList.removeClass('disabled');
                } else {
                    _this.checkedItem = {};
                    !_this.$mpCreate.hasClass('disabled') && _this.$mpCreate.addClass('disabled');
                    !_this.$inUse.hasClass('disabled') && _this.$inUse.addClass('disabled');
                    !_this.$forbidden.hasClass('disabled') && _this.$forbidden.addClass('disabled');
                    !_this.$edit.hasClass('disabled') && _this.$edit.addClass('disabled');
                    !_this.$mpList.hasClass('disabled') && _this.$mpList.addClass('disabled');
                }
                console.log(selectItems);
            }).on('click', '.btn-mp-create', function () {
                if (!_this.checkedItem) {
                    return;
                }
                //如果有companyId则是用户登陆 那么是管理员登陆
                location.href = _this.urlMiniProgramCreate + (_this.companyId ? "" : ("?company_id=" + _this.checkedItem.id));
            }).on('click', '.btn-edit', function () {
                if (!_this.checkedItem) {
                    return;
                }
                //如果有companyId则是用户登陆 那么是管理员登陆
                location.href = _this.urlHtmlCompanyEdit + (_this.companyId ? "" : ("?company_id=" + _this.checkedItem.id));
            }).on('click', '.btn-in-use', function () {
                //启用
                if (!_this.checkedItem) {
                    return;
                }
                _this.updateStatus(1);
            }).on('click', '.btn-forbidden', function () {
                //禁用
                if (!_this.checkedItem) {
                    return;
                }
                _this.updateStatus(2);
            }).on('click', '.btn-mp-list', function () {
                if (!_this.checkedItem) {
                    return;
                }
                //如果有companyId则是用户登陆 那么是管理员登陆
                location.href = _this.urlHtmlMiniProgram + (_this.companyId ? "" : ("?company_id=" + _this.checkedItem.id));
            });
        },
        updateStatus: function (status) {
            let _this = this;
            console.log(_this.checkedItem);
            let actName = status === 1 ? '启用' : (status === 2 ? '禁用' : '该');
            layer.dangerConfirm('确认进行' + actName + '操作？', function () {
                layer.ajax({
                    url: _this.urlApiCompanyUpdateStatus,
                    type: 'post',
                    data: {id: _this.checkedItem.id, to_status: status, from_status: _this.checkedItem.status}
                }, {loadingText: "操作处理中..."}).done(function (resp) {
                    if (0 !== +resp.error) {
                        layer.popupError("操作处理失败！" + resp.msg);
                        return false;
                    }
                    layer.popupMsg(actName + "成功！");
                    location.reload();
                });
            })
        },
        renderHtml: function () {
            let _this = this;
            if (!_this.companyId) {
                _this.$adminGroup.removeClass('hidden');
            }
        },
        renderData() {
            let _this = this;
            layer.ajax({
                url: _this.urlGetList,
                type: 'post',
                data: {id: _this.companyId ? _this.companyId : 0}
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
                        {name: 'name', label: '公司名称'},
                        {name: 'company_contact_user', label: '联系人'},
                        {name: 'company_contact_user_phone', label: '联系电话'},
                        {name: 'mp_count', label: '小程序数量'},
                        {name: 'status_name', label: '当前状态'},
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