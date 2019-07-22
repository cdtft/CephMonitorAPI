# CephMonitorAPI - Go bindings for Ceph RESTful APIs
## Ceph as a service
[![Build Status](https://travis-ci.org/ceph/go-ceph.svg)](https://travis-ci.org/ceph/go-ceph) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/ceph/go-ceph/master/LICENSE)

在kubernetes中使用CEPH作为持久化存储卷，当然可以使用StorageClass动态的创建PV和CEPH中的image但是这对持久化的存储卷的管理
带来了麻烦。例如：获取持久化的卷的使用率等，所以建立一套REST API完成持久化卷的管理无论是rbd还是cephfs存储方式都可以更灵活的使用并监管。

## Installation
在这个项目中我直接把github.com/ceph/go-ceph的代码copy到了goceph目录下，为了以后方便扩展
一些对ceph的操作。

在go-ceph中调用了librados c的库所以在开发调试的机器上需要安装下面的依赖库


The native RADOS library and development headers are expected to be installed.

On debian systems (apt):
```sh
libcephfs-dev librbd-dev librados-dev
```

On rpm based systems (dnf, yum, etc):
```sh
libcephfs-devel librbd-devel librados-devel
```

就是因为上面所依赖的库可能对开发调试带来难度，应为我使用的mac没有找到如何安装依赖的C库还好我找了一个方法。
VScode有一款插件可以远程编程，也就是可以使用服务器的环境在本地开发。

## API 文档
### 一、IMAGE
#### 1.CREAT 创建image
##### URL
`POST /api/v1/ceph/rbd/:pool/image/:name/:size`
##### Path Param
| 名称 | 变量名称 | 类型 | 备注 |
| :---: | :---: | :---: | :---: |
| 池名称 | pool | string | |
| 镜像名称 | image | string | |
| 镜像大小 | size | int | 单位GB |
#### 2.DELETE 删除image
##### URL
`DELETE /api/v1/ceph/rbd/:pool/image/:name`
##### Path Param
| 名称 | 变量名称 | 类型 | 备注 |
| :---: | :---: | :---: | :---: |
| 池名称 | pool | string | |
| 镜像名称 | image | string | |
#### 3.UPDATE 修改image size
##### URL
`PUT /api/v1/ceph/rbd/:pool/image/:name/:size`
##### Path Param
| 名称 | 变量名称 | 类型 | 备注 |
| :---: | :---: | :---: | :---: |
| 池名称 | pool | string | |
| 镜像名称 | image | string | |
| 镜像大小 | size | int | 单位GB |
#### 4.GET USAGE 获取已使用大小
##### URL
`PUT /api/v1/ceph/rbd/:pool/image/:name/usage`
##### Path Param
| 名称 | 变量名称 | 类型 | 备注 |
| :---: | :---: | :---: | :---: |
| 池名称 | pool | string | |
| 镜像名称 | image | string | |
##### Response Body