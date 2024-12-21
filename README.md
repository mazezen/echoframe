### <p align="center">EchoFrame</p>
#### <p align="center">基于echo框架,搭建一个便捷的脚手架</p>
#### <p align="center"><a href="https://github.com/mazezen/cgncode/releases"><img src="https://img.shields.io/github/release/cgncode/releases.svg" alt="GitHub release"></a><a href="https://github.com/mazezen/cgncode/blob/master/LICENSE"><img src="https://img.shields.io/github/license/mashape/apistatus.svg" alt="license"></a><p>
#### <p align="center"><a href="./README.md" target="_blank">简体中文</a> | <a href="./README_en.md" target="_blank">English</a> </p>

# EchoFrame
> 本项目是一个基于 Golang 和 Echo 框架的脚手架。
> 
> 旨在为开发者提供一个高效、灵活的开发基础。该脚手架集成了多种强大功能。


## 功能点

* 日志系统：内置日志功能，支持不同级别的日志记录，方便调试和监控。
* XORM：使用 XORM 作为 ORM 库，简化 MySQL 数据库操作，支持复杂查询。
* MySQL：集成 MySQL 数据库，提供可靠的数据存储解决方案。
* Redis：支持 Redis，提供快速的缓存和消息队列功能。
* LevelDB：集成 LevelDB，适合需要高效键值存储的场景。
* MongoDB：支持 MongoDB，适合处理非结构化数据和高并发场景。
* JWT：实现 JSON Web Token 身份验证，确保安全的用户管理。
* 代码生成器：提供代码生成工具，快速创建项目结构和模板，提升开发效率。
* gtools 工具库：内置 gtools 库，提供常用工具函数，进一步简化开发过程。 

这个脚手架适合需要快速搭建应用的开发者，减少了重复工作，让你专注于业务逻辑的实现，同时提升代码质量和开发效率

## 下载使用

### 安装docker
linux
```shell
make linux
```

macos
```shell
make macos
```

### 搭建mysql、redis、mongodb、naocs
```shell
make env
```

### 安装依赖
```shell
make mod
```

### 运行
```shell
make run
```

### 代码生成器 auto
<a href="./cmd/auto/README.md"  target="_blank">简体中文</a>
|
<a href="./cmd/auto/README_en.md"  target="_blank">English</a>

### 贡献
欢迎提交问题和贡献代码!

### 支持一杯咖啡
![wx](./cmd/auto/images/wx.jpg)

### License
该项目遵循 MIT 许可证。有关详细信息，请查看 LICENSE 文件。