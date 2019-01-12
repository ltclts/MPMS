<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
    <link rel="bookmark" type="image/x-icon" href="/static/img/logo.ico"/>
    <link rel="shortcut icon" href="/static/img/logo.ico">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="/static/components/zui/css/zui.min.css">
    <link rel="stylesheet" href="/static/components/zui/lib/datagrid/zui.datagrid.min.css">
    <link rel="stylesheet" href="/static/components/layui/css/layui.css">
    <link rel="stylesheet" href="/static/components/zui/lib/chosen/chosen.css">
    <link rel="stylesheet" href="/static/components/zui/lib/datetimepicker/datetimepicker.min.css">
    <link rel="stylesheet" href="/static/components/zui/lib/uploader/zui.uploader.min.css">
    <link rel="stylesheet" href="/static/css/index.css">
    <meta name="_xsrf" content="{{.xsrfdata}}"/>
{{.HtmlCss}}
</head>
<body>

{{ template  "header.tpl" .}}
{{ template  "menu.tpl" .}}

<div id="global_container">
    <ol class="breadcrumb">
        <li class="active"><i class="icon icon-home"></i>首页</li>
        <li class="active">{{.MenuFirstName}}</li>
        <li class="active">{{.MenuSecondName}}</li>
    </ol>

{{.LayoutContent}}
</div>

<div>
{{.SideBar}}
</div>
<script src="/static/components/zui/lib/jquery/jquery.js"></script>
<script src="/static/components/vue/vue-2.5.16.min.js"></script>
<script src="/static/components/zui/js/zui.min.js"></script>
<script src="/static/components/layui/layui.all.js"></script>
<script src="/static/components/zui/lib/datagrid/zui.datagrid.min.js"></script>
<script src="/static/components/zui/lib/chosen/chosen.min.js"></script>
<script src="/static/components/zui/lib/datetimepicker/datetimepicker.min.js"></script>
<script src="/static/components/zui/lib/uploader/zui.uploader.min.js"></script>
<script src="/static/js/index.js"></script>
{{.Scripts}}
</body>
</html>
