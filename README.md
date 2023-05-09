# go-admin
基于golang的后台管理系统，力求适中简化后台开发过程，满足中小项目开发需要。
使用到的主要技术框架：
Gin、Vue2.0等
![启动](/doc/1.png)

两个账户：
账户：admin 密码：123456  顶级账户
账户：test  密码：123456  菜单权限受限的账户

## 1. 项目地址
[github go-admin](https://github.com/jason-wj/go-admin)  
[gitee go-admin](https://gitee.com/jason-wj/go-admin)

## 2. 当前主要功能
基础功能：   
-- 登录身份鉴权，支持session和jwt两种方式、支持单点登录   
-- 公司、岗位、角色、管理员等基本通用功能  
-- 数据权限管理   
-- 根据模板快速生成前端页面以及后端相应业务代码，满足常用的增删改   
-- 插件化管理，尽可能独立化每个功能   
-- 支持mysql
-- 接口支持多语言
-- 其余常用的功能，这里就不一一列举了

扩展功能（一般项目都能用到的）   
-- app用户管理，包括对应配置管理里   
-- 等级管理   
-- 账变记录
-- 国际区号
-- 插件（基本的CMS内容管理、App安装包管理）

## 3. 遵守规则
1. 数据库表
   1. 系统基础表，以`sys_`为前缀，后面衔接模块业务名
   2. 插件表，以`plugins_`为前缀，后面衔接模块业务名
   3. 主业务项目，以`app_`为前缀，后面衔接模块业务名
2. 项目结构
   1. 参考现有结构，时间不充裕，没法详细解释
3. 别的规则暂时记不起来了，后面想到了我再补充

## 4. 部署
后台登录默认账户：`admin`
默认密码：`123456`

### 4.1 数据库部署
1. 项目根目录的`app.sql`是当前的最新库，自行导入数据库即可

### 4.2 后台接口部署与启动
1. 编译项目，项目根目录执行如下：
```shell
go mod tidy

# 我本人是mac环境，直接如下编译
go build -a -o go-admin-api main.go

# 可交叉编译为linux的：
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o go-admin-api main.go
```
2. 项目根目录，找到`conf/`，可以直接再`settings.yml`中配置数据库，或者可以拷贝一份，重新命名，比如：`settings.dev.yml`，具体的配置信息，请自行去文件中参考
3. 启动项目，参考如下两种方式：
```shell
# 将会默认读取conf/目录中settings.yml配置，并启动接口
./admin server

# 读取自定义配置
./admin server -c=config/settings.dev.yml
```
4. 启动并录入所有接口到表`sys_api`中，需要在启动的时候加参数`-a`，这样就能方便菜单接口权限管理了：
```shell
# 这是使用默认配置
./admin server -a
```

### 4.3 web页面编译和部署
1. 依赖安装：
```shell
cd ./web
npm install --force
```
2. 本地启动：
```shell
npm run dev
```
3. 编译：
```shell
npm run build:prod
```
4. 检测package.json中哪些依赖是否有用，可以精简无效依赖，我已经精简了，这里记录下：
```shell
npm i -g npm-check
cd web
npm-check
```

## 5. 其余
这里只是大概讲了下基本情况，具体更多内容，还是需要在使用中一点点去发现。

## 感谢
[go-admin-team](https://github.com/go-admin-team)  
[flipped-aurora](https://github.com/flipped-aurora)
