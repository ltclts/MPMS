<script type="text/javascript">

    let company = {
        $datagrid: $('#datagrid'),
        urlGetList:{{.UrlGetList}},
        init: function () {
            this.render();
        },
        render: function () {
            this.handleEvent();
            this.renderData();
        },
        handleEvent: function () {
            let _this = this;

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
                        {name: 'created_at', label: '创建时间'},
                    ],
                    array: [
                        {name: '两分钱', company_contact_user: '曹禺', company_contact_user_phone: '13221733659', status: '已启用', mp_count:2,creator:'曹禺',created_at:'2019.12.20 10:59'}
                    ]
                },
                checkable: true,
                sortable: true
            });
        }
    };

    $(function () {
        company.init();

    });
</script>