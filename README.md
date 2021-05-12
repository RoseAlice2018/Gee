# Gee


## 动态路由
- 实现的动态路由具备以下两个功能
### 参数匹配
- 例如 
### 通配*
- 例如 /static/*filepath,可以匹配/static/fav.ico,也可以
匹配/static/js/jQuery.js,这种模式常用于静态服务器，能够递归
  匹配子路径。
  
## 分组控制
- 分组控制(Group Control)是 Web 框架应提供的基础功能之一。所谓分组，是指路由的分组。如果没有路由分组，我们需要针对每一个路由进行控制。但是真实的业务场景中，往往某一组路由需要相似的处理。例如：
  - 以/post开头的路由匿名可访问。
  - 以/admin开头的路由需要鉴权
  - 以/api开头的路由是 RESTful 接口，可以对接第三方平台，需要三方平台鉴权。
  