# go-admin
基于golang的后台管理系统，力求适中简化后台开发过程，满足中小项目开发需要。
使用到的主要技术框架： Gin、TypeScript、React18、Ant-design5等

`关于前端技术我解释下`：  
1. 由于业务和环境的变化，原本vue2.0前端的维护压力越来越大。再三考虑后，使用React对前端进行了重构，借此也对golang接口端的各种问题进行了集中处理。  
2. vue2.0版的最终完整版本，我单独设置了一个分支，该分支仅做永久备份，不再进行更新和维护。后续将只会迭代新的react版
3. react版分为内部版和开源版，内部版源码不对外开放，但为了开源版的完整性，我借用了第三方开源项目框架[Hooks-Admin](https://github.com/HalseySpicy/Hooks-Admin) 来实现。
   1. 所有涉及业务接口的页面，都是从内部版完整迁移到这个框架，这些页面一行代码也没改
   2. 该框架原本使用的是ant-design4.x，为了兼容业务页面，我变更为5.x，框架本身针对性的做了一些调整
   3. 最终目的是让迁移的页面能在不改代码的情况下正常运行，而框架本身的基础功能非必要我不做调整
   4. 该开源框架本身其实已经很老，很多细节实现上也需要完善或者重构，由于个人精力有限，这个只能是使用者根据自身情况自行迭代完善了。
  
`初衷`：  
有很多开源的go-admin，但其实都是在别的开源项目的基础上各种迭代的，代码冗余，功能冗余，结构上也不尽人意。  
为了满足个人日常需要（统计数据、中小单体项目开发等），我经过多个项目的磨合，对该项目做了大量的精简和调整，也对很多功能做了重新规划和设计。  
对很多人而言，仅仅需要最基础的功能，剩下的根据需要自行迭代，而这套框架也刚好满足你。  

我的开发准则：  
`大道至简`

![启动](/logo.png)

在线体验：  
[go-admin react开源版入口](http://120.48.157.250:22222/)  
[go-admin vue2.0版入口(不再更新)](http://120.48.157.250:60000/)  

内置两个默认账户：  
账户：admin 密码：123456  顶级账户    
账户：test  密码：123456  菜单权限受限的账户  
`注：`：该项目没有任何商业行为，只作为公开分享和交流，请友善体验，不要非法攻击该在线地址。所有功能对外开放，期间可能有小伙伴无意间操作导致页面显示异常，我发现后会及时恢复数据库，或者可以发Issues提醒我处理。  

## 1. 项目结构
前后端都放在一个项目了，对于中小项目，这个管理反而会很方便。  
其实它本质还是前后端分离。  
`web`目录是前端目录，完全独立，可以独立开发和部署。除了该目录，剩下的都属于后台接口的  
  
//-------------以下待更新--------------  

## 2. 当前主要功能
基础功能：   
-- 登录身份鉴权，支持session和jwt两种方式、支持单点登录   
-- 公司、岗位、角色、管理员等基本通用功能  
-- 数据权限管理   
-- 根据模板快速生成前端页面以及后端相应业务代码，满足常用的增删改   
-- 插件化管理，尽可能独立化每个功能   
-- 支持mysql（20231022-打算支持postgres，先等我熟悉了，冒失新增，怕造成不稳定）  
-- 接口支持多语言  
-- 其余常用的功能，这里就不一一列举了  

扩展功能（一般项目都能用到的）   
-- app用户管理，包括对应配置管理里   
-- 等级管理   
-- 账变记录  
-- 国际区号  
-- 插件（基本的CMS内容管理、App安装包管理、消息管理）  

`备注`：截止目前，我也不打算提供定时任务和swagger功能，对我而言，我认为完全没必要。首先，定时任务和后台耦合，每次后台发布，很容易引起定时任务的异常，定时任务完全可以搞成一个独立的脚本去运行，彻底解耦。而swagger，我只感觉鸡肋，除了给自己增加工作量，眼前还没发现任何优势。这两个功能谁如果需要，完全可以自行在现有框架中加进去。

## 3. 遵守规则
1. 数据库表
   1. 系统基础表，以`sys_`为前缀，后面衔接模块业务名  
   2. 插件表，以`plugins_`为前缀，后面衔接模块业务名  
   3. 主业务项目，以`app_`为前缀，后面衔接模块业务名  
2. 项目结构
   1. 参考现有结构，时间不充裕，没法详细解释  
3. 别的规则暂时记不起来了，后面想到了我再补充  

## 4. 部署
内置两个默认账户：  
账户：admin 密码：123456  顶级账户    
账户：test  密码：123456  菜单权限受限的账户  

### 4.1 数据库部署
1. 项目根目录的`app.sql`是当前的最新库，自行导入数据库即可  

### 4.2 后台接口部署与启动
1. 编译项目，项目根目录执行如下：  
```shell
go mod tidy

# 直接编译
go build -a -o go-admin-api main.go

# 交叉编译为linux amd64：
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
浏览器打开：http://127.0.0.1:1688

3. 补充：编译：
```shell
npm run build:prod
```
4. 补充：检测package.json中哪些依赖是否有用，可以精简无效依赖，我已经精简了，这里记录下：
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
