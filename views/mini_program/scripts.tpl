<script type="text/javascript">

    let mp = {
        operateType:{{.OperateType}},
        companyId:{{.CompanyId}},
        init: function () {
            this.render();
        },
        render: function () {
            this.handleEvent();
        },
        handleEvent: function () {
            let _this = this;
            console.log(_this.operateType);
        }
    };

    $(function () {
        mp.init();
    });
</script>