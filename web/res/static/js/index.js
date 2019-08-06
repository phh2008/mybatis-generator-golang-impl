layui.use(['form', 'layer', 'table', 'util'], function () {
    const form = layui.form;
    const layer = layui.layer;
    const table = layui.table;
    const util = layui.util;
    const $ = layui.jquery;
    const ctxPath = $("#ctxPath").val();

    form.on('submit(query)', function (data) {
        console.log(data.field);
        table.reload('gen', {
            where: data.field
        });
        return false;
    });

    table.render({
        elem: '#gen',
        id: 'gen',
        height: 'full-280',
        url: ctxPath + '/tables',
        method: 'get',
        limit: 10000,
        loading: true,
        even: true,
        page: false,
        cols: [[
            {checkbox: true, fixed: 'left'},
            {field: 'Name', title: '表名', width: 160, fixed: 'left'},
            {
                field: 'CreateTime', title: '创建时间', width: 160, templet: function (d) {
                    return util.toDateString(d.CreateTime, "yyyy-MM-dd HH:mm:ss");
                }
            },
            {field: 'Comment', title: '注释'},
        ]]
    });

});