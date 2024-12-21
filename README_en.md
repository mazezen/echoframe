### <p align="center">EchoFrame</p>
#### <p align="center">Based on the echo framework, build a convenient scaffolding</p>
#### <p align="center"><a href="https://github.com/mazezen/cgncode/releases"><img src="https://img.shields.io/github/release/cgncode/releases.svg" alt="GitHub release"></a><a href="https://github.com/mazezen/cgncode/blob/master/LICENSE"><img src="https://img.shields.io/github/license/mashape/apistatus.svg" alt="license"></a><p>
#### <p align="center"><a href="./README.md" target="_blank">Simplified Chinese</a> | <a href="./README_en.md" target="_blank">English</a> </p>

# EchoFrame
> This project is a scaffold based on Golang and Echo framework.
>
> It aims to provide developers with an efficient and flexible development foundation. 
> 
> The scaffold integrates a variety of powerful functions.

## Function Points

* Log system: built-in log function, support different levels of logging, convenient debugging and monitoring.
* XORM: use XORM as ORM library to simplify MySQL database operations and support complex queries.
* MySQL: integrated MySQL database, providing reliable data storage solutions.
* Redis: support Redis, provide fast cache and message queue functions.
* LevelDB: integrated LevelDB, suitable for scenarios requiring efficient key-value storage.
* MongoDB: support MongoDB, suitable for processing unstructured data and high concurrency scenarios.
* JWT: implement JSON Web Token authentication to ensure secure user management.
* Code generator: provide code generation tools, quickly create project structure and templates, and improve development efficiency.
* gtools tool library: built-in gtools library, provide common tool functions, further simplify the development process.

This scaffold is suitable for developers who need to build applications quickly. It reduces repetitive work and allows you to focus on the implementation of business logic while improving code quality and development efficiency.
## Download and use function points

### Install Docker
linux
```shell
make linux
```

macos
```shell
make macos
```

### mysql、redis、mongodb、naocs
```shell
make env
```

### Install Dependencies
```shell
make mod
```

### run
```shell
make run
```

### Code Generator
<a href="./cmd/auto/README.md"  target="_blank">简体中文</a>
|
<a href="./cmd/auto/README_en.md"  target="_blank">English</a>

### Contributions
Welcome to submit issues and contribute code!

### Support a cup of coffee
![wx](./cmd/auto/images/wx.jpg)

### License
This project is licensed under the MIT License. See the LICENSE file for details.