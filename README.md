# panda

#### 介绍
go语言gofiber脚手架

#### 相关技术
| 序号 | 名称                          | 描述 |
|----|-----------------------------|----|
| 1  |github.com/gofiber/fiber/v2|基于Fasthttp的Go web框架，Fasthttp是Go最快的HTTP引擎。它旨在简化快速开发，同时考虑到零内存分配和性能|
| 2  |github.com/google/uuid|基于github.com/pborman/uuid包（以前称为code.google.com/p/go-uuid）。它与这些早期包的不同之处在于，UUID是一个16字节的数组，而不是一个字节切片。此更改造成的一个损失是表示无效UUID的能力（与NIL UUID相比）|
| 3  |go.uber.org/zap|zap提供快速、结构化、分级的日志记录。对于登录热路径的应用程序，基于反射的序列化和字符串格式化的成本非常高——它们占用大量CPU，并且会进行许多小的分配。换句话说，使用json。元帅和fmt。Fprintf记录大量接口{}会使您的应用程序变慢。Zap采取了不同的方法。它包括一个无反射、零分配的JSON编码器，基本Logger尽可能避免序列化开销和分配。通过在此基础上构建高级SugaredLogger，zap允许用户选择何时需要计算每个分配，以及何时更喜欢更熟悉、松散类型的API|
| 4  |gopkg.in/yaml.v3|yaml包使Go程序能够轻松地编码和解码yaml值。它是作为juju项目的一部分在Canonical内部开发的，基于著名的libyaml C库的纯Go端口，可以快速可靠地解析和生成YAML数据|
| 5  |gorm.io/gorm|1. 全功能 ORM。2. 关联 (拥有一个，拥有多个，属于，多对多，多态，单表继承)。3. Create，Save，Update，Delete，Find 中钩子方法。4. 支持 Preload、Joins 的预加载。5. 事务，嵌套事务，Save Point，Rollback To to Saved Point。6. Context、预编译模式、DryRun 模式。7. 批量插入，FindInBatches，Find/Create with Map，使用 SQL 表达式、Context Valuer 进行 CRUD。8. SQL 构建器，Upsert，锁，Optimizer/Index/Comment Hint，命名参数，子查询。9. 复合主键，索引，约束。10. 自动迁移。11. 自定义 Logger。12. 灵活的可扩展插件 API：Database Resolver（多数据库，读写分离）、Prometheus…。13. 每个特性都经过了测试的重重考验。14. 开发者友好。|

#### gofiber卓越的性能
![输入图片说明](https://gofiber.io/assets/images/benchmark-pipeline.png)

#### 软件架构
软件架构说明


#### 安装教程

1.  切换 GO 镜像

    go env -w GOPROXY=https://goproxy.cn,direct

    清除缓存 go clean -modcache

    验证镜像切换 go env | grep GOPROXY

    找到 GOPROXY=https://goproxy.cn,direct

2.  xxxx
3.  xxxx

#### 使用说明
[接口文档](https://doc.apipost.net/docs/2ef94399bc66000)

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
