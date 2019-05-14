# 论坛项目


1 go mod 配置方式 下载项目后 进入到项目所在目录 go mod init forum  go run main/main.go
2 配置文件分为正式版跟dev版
3 doc 文档为放置sql脚本以及一些需求地方
4 路由按照模块分组然后统一加进routers.go中
5 兼容以前版本。所以在controller中创建版本号。 在版本号中写代码
6 中间件统一放入到middleware中
7 菜逼。。各位大佬求轻喷。对还是个孩子的我要关爱
