<!DOCTYPE html>
<html class="x-admin-sm">
    <head>
        <meta charset="UTF-8">
        <title>图书销售管理系统</title>
        <meta name="renderer" content="webkit">
        <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
        <meta name="viewport" content="width=device-width,user-scalable=yes, minimum-scale=0.4, initial-scale=0.8,target-densitydpi=low-dpi" />
        <link rel="stylesheet" href="./static/css/font.css">
        <link rel="stylesheet" href="./static/css/xadmin.css">
        <script src="./static/lib/layui/layui.js" charset="utf-8"></script>
        <script type="text/javascript" src="./static/js/xadmin.js"></script>
        <!--[if lt IE 9]>
          <script src="./static/js/html5.min.js"></script>
          <script src="./static/js/respond.min.js"></script>
        <![endif]-->
    </head>
    <body>
        <div class="layui-fluid">
            <div class="layui-row layui-col-space15">
                <div class="layui-col-md12">
                    <div class="layui-card">
                        <div class="container">
                            <div class="logo">
                                <a href="./index.html">图书展示</a></div>
                            <ul class="layui-nav right" lay-filter="">
                                <li class="layui-nav-item">
                                    <a href="javascript:;">tanxiaowei</a>
                                    <dl class="layui-nav-child">
                                        <!-- 二级菜单 -->
                                        <dd><a href="/">后台管理</a></dd>
                                        <dd><a href="/register.html">注册用户</a></dd>
                                        <dd><a href="login.html">退出登录</a></dd>
                                    </dl>
                                </li>
                            </ul>
                        </div>
                        <div class="layui-card-body ">
                            <table class="layui-table layui-form">
                              <tbody>
                                <tr>
                                  <td><img src="{{ .Pic_Path }}" height="200" width="200" />
                                      <div>
                                          <span>库存:</span>
                                          <span>购买:</span>
                                      </div>
                                  </td>
                                  <td>admin</td>
                                  <td>18925139194</td>
                                  <td>113664000@qq.com</td>
                                  <td>超级管理员</td>
                                </tr>
                              </tbody>
                            </table>
                        </div>
                        <div class="layui-card-body ">
                            <div class="page">
                                <div>
                                  <a class="prev" href="">&lt;&lt;</a>
                                  <a class="num" href="">1</a>
                                  <span class="current">2</span>
                                  <a class="num" href="">3</a>
                                  <a class="num" href="">489</a>
                                  <a class="next" href="">&gt;&gt;</a>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div> 
    </body>
    <script>
      layui.use(['laydate','form'], function(){
        var laydate = layui.laydate;
        var form = layui.form;
        
        //执行一个laydate实例
        laydate.render({
          elem: '#start' //指定元素
        });

        //执行一个laydate实例
        laydate.render({
          elem: '#end' //指定元素
        });
      });

       /*用户-停用*/
      function member_stop(obj,id){
          layer.confirm('确认要停用吗？',function(index){

              if($(obj).attr('title')=='启用'){

                //发异步把用户状态进行更改
                $(obj).attr('title','停用')
                $(obj).find('i').html('&#xe62f;');

                $(obj).parents("tr").find(".td-status").find('span').addClass('layui-btn-disabled').html('已停用');
                layer.msg('已停用!',{icon: 5,time:1000});

              }else{
                $(obj).attr('title','启用')
                $(obj).find('i').html('&#xe601;');

                $(obj).parents("tr").find(".td-status").find('span').removeClass('layui-btn-disabled').html('已启用');
                layer.msg('已启用!',{icon: 5,time:1000});
              }
              
          });
      }

      /*用户-删除*/
      function member_del(obj,id){
          layer.confirm('确认要删除吗？',function(index){
              //发异步删除数据
              $(obj).parents("tr").remove();
              layer.msg('已删除!',{icon:1,time:1000});
          });
      }



      function delAll (argument) {

        var data = tableCheck.getData();
  
        layer.confirm('确认要删除吗？'+data,function(index){
            //捉到所有被选中的，发异步进行删除
            layer.msg('删除成功', {icon: 1});
            $(".layui-form-checked").not('.header').parents('tr').remove();
        });
      }
    </script>
    <script>var _hmt = _hmt || []; (function() {
        var hm = document.createElement("script");
        hm.src = "https://hm.baidu.com/hm.js?b393d153aeb26b46e9431fabaf0f6190";
        var s = document.getElementsByTagName("script")[0];
        s.parentNode.insertBefore(hm, s);
      })();</script>
</html>