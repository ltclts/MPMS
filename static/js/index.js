;!function () {
    const layer = layui.layer, form = layui.form;
}();

!function () {
    const ajax = $.ajax;
    $.extend({
        ajax: function (url, options) {
            if (typeof url === 'object') {
                options = url;
                url = undefined;
            }
            options = options || {};
            url = options.url;
            let xsrftoken = $('meta[name=_xsrf]').attr('content');
            let headers = options.headers || {};
            let domain = document.domain.replace(/\./ig, '\\.');
            if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
                headers = $.extend(headers, {'X-Xsrftoken': xsrftoken});
            }
            options.headers = headers;
            return ajax(url, options);
        }
    });
}();

!(function (/*对layer的扩展*/) {
    layer.config({
        moveType: 1,
        shade: .6,
        shadeClose: true,
        shadeCloseAsCancel: true
    });

    if (!layer.hint_id_index) {
        layer.hint_id_index = new Date().getTime();
    }

    layer.popupMsg = function (msg, option) {
        new $.zui.Messager(msg, $.extend({type: 'default'}, option)).show();
    };
    layer.popupDanger = function (msg, option) {
        layer.popupMsg(msg, $.extend({type: 'danger'}, option))
    };
    layer.popupPrimary = function (msg, option) {
        layer.popupMsg(msg, $.extend({type: 'primary'}, option))
    };
    layer.popupInfo = function (msg, option) {
        layer.popupMsg(msg, $.extend({type: 'info'}, option))
    };
    layer.popupSuccess = function (msg, option) {
        layer.popupMsg(msg, $.extend({type: 'success'}, option))
    };
    layer.popupWarning = function (msg, option) {
        layer.popupMsg(msg, $.extend({type: 'warning'}, option))
    };
    layer.popupImportant = function (msg, option) {
        layer.popupMsg(msg, $.extend({type: 'important'}, option))
    };
    layer.popupSpecial = function (msg, option) {
        layer.popupMsg(msg, $.extend({type: 'special'}, option));
    };

    layer.popupError = layer.popupDanger;

    const ajaxFailCb = function (xhr, text, option) {
        if (+xhr.status === 422) {
            text = '';
            for (let i in xhr.responseJSON) {
                if (xhr.responseJSON.hasOwnProperty(i)) {
                    text = i + '：<br />' + xhr.responseJSON[i].join('<br />');
                }
            }
            layer.popupDanger('操作失败，请求参数不合法，请检查后重试！错误具体原因如下：<br />' + text, option);
        } else {
            layer.popupDanger(text || '操作失败，服务器或网络异常，请重试！status = ' + status, option);
        }
    };
    layer.select = function (options, yes) {
        let cache = layer.cache || {}, skin = function (type) {
            return (cache.skin ? (' ' + cache.skin + ' ' + cache.skin + '-' + type) : '');
        };
        options = options || {};
        if (typeof options === 'function') yes = options;
        let select, content = function () {
            let select_html = '<select  class="layui-layer-input layui-layer-select">';
            let items = options.items || [];
            items.forEach(function (e) {
                select_html += '<option value="' + e.value + '">' + e.text + '</option>';
            });
            select_html += '</select>';
            return select_html;
        }();
        return layer.open($.extend({
            btn: ['&#x786E;&#x5B9A;'],
            content: content,
            type: 1,
            moveType: 1,
            skin: 'layui-layer-select' + skin('select'),
            success: function (layero) {
                select = layero.find('.layui-layer-input');
                select.focus();
            }, yes: function (index) {
                let value = select.val();
                if (value === '') {
                    select.focus();
                } else if (value.length > (options.maxlength || 500)) {
                    layer.tips('&#x6700;&#x591A;&#x8F93;&#x5165;' + (options.maxlength || 500)
                        + '&#x4E2A;&#x5B57;&#x6570;', prompt, {tips: 1});
                } else {
                    yes && yes(value, index, select);
                }
            }
        }, options));
    };
    layer.hint = function (options, yes) {
        let cache = layer.cache || {}, skin = function (type) {
            return (cache.skin ? (' ' + cache.skin + ' ' + cache.skin + '-' + type) : '');
        };
        options = options || {};
        if (typeof options === 'function') yes = options;
        let content = function () {
            let html = '';
            let inputs = options.inputs || [];
            inputs.forEach(function (e) {
                if (e.type === 'input') {
                    html += '<div class="layui-layer-input-wrapper" >' +
                        '<span class="layui-layer-input-text">' + e.hint + '</span>' +
                        '<input type="text" class="layui-layer-input" id="layer_input_' + (e.id || (e.id = ++layer.hint_id_index)) + '" value="' + (e.value || '') + '" />' +
                        ('<span class="layui-layer-addon">' + (e.addon ? e.addon.html : '') + '</span>') +
                        '</div>'
                } else if (e.type === 'password') {
                    html += '<div class="layui-layer-input-wrapper" >' +
                        '<span class="layui-layer-input-text">' + e.hint + '</span>' +
                        '<input type="password" class="layui-layer-input" id="layer_input_' + (e.id || (e.id = ++layer.hint_id_index)) + '" value="' + (e.value || '') + '" />' +
                        ('<span class="layui-layer-addon">' + (e.addon ? e.addon.html : '') + '</span>') +
                        '</div>'
                } else if (e.type === 'textarea') {
                    html += '<div class="layui-layer-input-wrapper" >' +
                        '<span class="layui-layer-input-text">' + e.hint + '</span>' +
                        '<div class="layui-layer-textarea">' +
                        '<textarea id="layer_input_' + (e.id || (e.id = ++layer.hint_id_index)) + '">' + (e.value || '') + '</textarea>' +
                        '</div>' +
                        '</div>'
                } else if (e.type === 'date') {
                    html += '<div class="layui-layer-input-wrapper" >' +
                        '<span class="layui-layer-input-text">' + e.hint + '</span>' +
                        '<input type="text" class="layui-layer-input layui-layer-input-date" id="layer_input_' + (e.id || (e.id = ++layer.hint_id_index)) + '" value="' + (e.value || '') + '" />' +
                        ('<span class="layui-layer-addon">' + (e.addon ? e.addon.html : '') + '</span>') +
                        '</div>';
                } else if (e.type === 'datetime') {
                    html += '<div class="layui-layer-input-wrapper" >' +
                        '<span class="layui-layer-input-text">' + e.hint + '</span>' +
                        '<input type="text" class="layui-layer-input layui-layer-input-datetime" id="layer_input_' + (e.id || (e.id = ++layer.hint_id_index)) + '" value="' + (e.value || '') + '" />' +
                        ('<span class="layui-layer-addon">' + (e.addon ? e.addon.html : '') + '</span>') +
                        '</div>';
                } else if (e.type === 'select') {
                    html += '<div class="layui-layer-input-wrapper" >' +
                        '<span class="layui-layer-input-text">' + e.hint + '</span>';
                    html += '<select class="layui-layer-input" id="layer_input_' + (e.id || (e.id = ++layer.hint_id_index)) + '" >';
                    let items = e.items || [];
                    items.forEach(function (e) {
                        let str_attr = '';
                        (e.attrs || []).forEach(function (attr) {
                            str_attr += ' ' + attr.key + '=' + attr.value;
                        });
                        html += '<option ' + str_attr + ' value="' + e.value + '" ' + (e.selected ? 'selected' : '') + '>' + e.text + '</option>';
                    });
                    html += '</select>' +
                        ('<span class="layui-layer-addon">' + (e.addon ? e.addon.html : '') + '</span>') +
                        '</div>';
                } else if (e.type === 'multiple') {
                    e._values = [];
                    e._items = e.items;
                    e.id = e.id || (++layer.hint_id_index);
                    ('function' === typeof e._items) && (e._items = e._items());
                    html +=
                        '<div class="layui-layer-input-wrapper" >' +
                        '<span class="layui-layer-input-text">' + e.hint + '</span>' +
                        '<input type="hidden" class="layui-layer-input" id="layer_input_' + e.id + '" />' +
                        '<div class="layui-layer-multiple">' +
                        function (items) {
                            let html = '';
                            items.forEach(function (item) {
                                html += '<span value="' + (item.value || '') + '" title="' + (item.title || '') + '" class="layui-layer-checkbox ' + (item.selected ? (e._values.push(item.id), 'selected') : '') + '">' + (item.text || '') + '</span>'
                            });
                            return html;
                        }(e._items) +
                        '</div>' +
                        '</div>';
                }
            });

            return html;
        }();
        return layer.open($.extend({
            btn: [{text: '&#x786E;&#x5B9A;', className: 'btn btn-sm btn-primary'}],
            content: content,
            type: 1,
            moveType: 1,
            skin: 'layui-layer-hint' + skin('hint'),
            area: '290px',
            success: function (layer$) {
                layer$.find('.layui-layer-input').keyup(function () {
                    return false;
                });
                layer$.find('.layui-layer-input-date').datetimepicker({
                    format: 'yyyy-mm-dd',
                    autoclose: true,
                    language: 'zh-CN',
                    minView: 2
                });
                layer$.find('.layui-layer-input-datetime').datetimepicker({
                    format: 'yyyy-mm-dd hh:ii:ss',
                    autoclose: true,
                    language: 'zh-CN'
                });
                (options.inputs || []).forEach(function (e) {
                    e.elem = layer$.find('#layer_input_' + e.id);
                    e.attr && e.elem.attr(e.attr);
                    e.style && e.elem.css(e.style);
                    if (e.type === 'multiple') {
                        e.attr && e.elem.next().attr(e.attr);
                        e.style && e.elem.next().css(e.style);
                    }
                    if (e.addon && (e.addon.style || e.addon.className)) {
                        e.$addon = e.elem.siblings('.layui-layer-addon');
                        e.$addon && e.addon.style && e.$addon.css(e.addon.style);
                        e.$addon && e.addon.className && e.$addon.addClass(e.addon.className);
                    }
                });
                layer$.find('.layui-layer-multiple').each(function () {
                    let list = $(this);
                    let selected_list = [];
                    list.find('.selected').each(function (i, e) {
                        selected_list.push($(e).attr('value'));
                    });
                    list.prev().val(selected_list.join(','));
                });
                layer$.find('.layui-layer-checkbox').click(function () {
                    let list = $(this).toggleClass('selected').parent();
                    let selected_list = [];
                    list.find('.selected').each(function (i, e) {
                        selected_list.push($(e).attr('value'));
                    });
                    list.prev().val(selected_list.join(','));
                });
                options.hint_success && options.hint_success();
            }, yes: function (index) {
                yes && yes(index, options.inputs);
            }
        }, options));
    };
    layer.ajax = function (ajaxOption, option) {
        option = option || {};
        let loadingIndex = layer.msg(option.loadingText || '正在执行操作……', $.extend({
            icon: 16,
            time: -1,
            shade: .6
        }, option.loadingOption));
        let dtd = $.Deferred();
        //console.log(dtd.ajax);
        dtd.ajax = $.ajax(ajaxOption)
            .done(function (resp) {
                dtd.resolve(resp);
            })
            .fail(option.fail === false ? null : function (xhr) {
                ajaxFailCb(xhr, option.failText, option.failOption);
            })
            .complete(function () {
                layer.close(loadingIndex);
            });
        return dtd.promise();
    };
    layer.dangerConfirm = function (msg, option, yes, cancel) {
        let type = typeof option === 'function';
        if (type) {//表示不传入option
            cancel = yes;
            yes = option;
            option = null;
        }
        if (!cancel && option && option.cancelText || typeof(cancel) === 'string') {
            console.warn('cancel is set by cancelText');
            let cancelText = typeof(cancel) === 'string' ? cancel : option.cancelText;
            cancel = function () {
                layer.popupMsg(cancelText)
            }
        }
        option = $.extend(true, {
            icon: 7,
            btn: ['确定', '取消']
        }, option);
        layer.confirm(msg, option, yes, cancel);
    };
})();

!(function (undefined/*页面初始化*/) {
    let  $treeMenu = $('#treeMenu');
    $treeMenu && $treeMenu.on('click', 'a', function() {
        $('#treeMenu li.active').removeClass('active');
        $(this).closest('li').addClass('active');
    });

    $.fn.popover && $('[data-toggle=popover]').popover();
    $.fn.chosen && $('select.chosen-select').each(function (i, e) {
        let $this = $(e),
            search = $this.attr('data-search-input');
        $this.chosen({
            no_results_text: '没有找到',
            disable_search_threshold: search === undefined ? 10 : 0,
            search_contains: true,
            allow_single_deselect: true
        })
    });
    $.fn.datatable && $('table.datatable').each(function (i, e) {
        let $e = $(e);
        $e.datatable({
            sortable: $e.hasClass('sortable'),
            checkable: $e.hasClass('checkable')
        }).on('beforeSort.zui.datatable', function (event) {
            layer.popupMsg('注意：目前仅支持页面内的条目进行排序！');
            console.log(event);
        });
    });

    let $formdate = $("input.form-date");
    $formdate && $formdate.datetimepicker(
        {
            language: "zh-CN",
            weekStart: 1,
            todayBtn: 1,
            autoclose: 1,
            todayHighlight: 1,
            startView: 2,
            minView: 2,
            forceParse: 0,
            format: "yyyy-mm-dd"
        });

    let $formtime = $("input.form-time");
    $formtime && $formtime.datetimepicker({
        language: "zh-CN",
        weekStart: 1,
        todayBtn: 1,
        autoclose: 1,
        todayHighlight: 1,
        startView: 1,
        minView: 0,
        maxView: 1,
        forceParse: 0,
        format: 'hh:ii'
    });
    // 选择时间和日期
    let $formdatetime = $("input.form-datetime");
    $formdatetime && $formdatetime.datetimepicker({
        weekStart: 1,
        todayBtn: 1,
        autoclose: 1,
        todayHighlight: 1,
        startView: 2,
        forceParse: 0,
        showMeridian: 1,
        format: "yyyy-mm-dd hh:ii"
    });

    $.fn.tooltip && $('[data-toggle=tooltip]').tooltip();

    $('body').on('click', '.logout', function () {
        layer.dangerConfirm("确认注销？", function () {
            layer.ajax({
                url: "/api/user/logout",
                type: 'post',
                data: {}
            }, {loadingText: "正在注销..."}).done(function (resp) {
                console.log(resp);
                if (0 !== +resp.error) {
                    layer.alert("注销失败！");
                    return false;
                }
                location.href = resp.info.uri;
            });
        });
    });

})();
