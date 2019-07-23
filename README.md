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
| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 池名称 | pool | string | | required |
| 镜像名称 | image | string | | required |
| 镜像大小 | size | int | 单位GB | required |
#### 2.DELETE 删除image
##### URL
`DELETE /api/v1/ceph/rbd/:pool/image/:name`
##### Path Param
| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 池名称 | pool | string | | required |
| 镜像名称 | image | string | | required |
#### 3.UPDATE 修改image size
##### URL
`PUT /api/v1/ceph/rbd/:pool/image/:name/:size`
##### Path Param
| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 池名称 | pool | string | | required |
| 镜像名称 | image | string | | required |
| 镜像大小 | size | int | 单位GB | required |
#### 4.GET USAGE 获取已使用大小
##### URL
`PUT /api/v1/ceph/rbd/:pool/image/:name/usage`
##### Path Param
| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 池名称 | pool | string | | required |
| 镜像名称 | image | string | | required |
##### Response Body
| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 池名称 | pool | string | | required |
| 镜像名称 | image | string | | required |
| 返回值 | data | string | 带单位B MB | required |
```json
{
    "code": 1200,
    "msg": "获取已使用大小",
    "data": "0B"
}
```
#### 5.Batch Create Images批量创建image
##### URL
`POST /api/v1/ceph/rbd/:pool/images`
##### Path Param
| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 池名称 | pool | string | | required |
##### Request Body 
| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 镜像列表 | images | Array(CreateImage) | | required |

`CreateImage` object:

| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 镜像名称 | name | string | | required |
| 镜像大小 | size | int | 单位GB | required |
```json
{
  "images": [
      {
          "name": "my-disk-1",
          "size": 10
      },
      {
          "name": "my-disk-2",
          "size": 50
      }
  ]
}
```

#### 5.Batch Delete Images批量删除image
##### URL
`DELETE /api/v1/ceph/rbd/:pool/images`
##### Path Param
| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 池名称 | pool | string | | required |
##### Request Body

| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 镜像列表 | name | Array(DeleteImage) | | required |

`DeleteImage` object:

| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 镜像名称 | name | string | | required |

```json
{
  "images": [
      {
          "name": "my-disk-1"
      },
      {
          "name": "my-disk-2"
      }
  ]
}
```

#### 6.Batch Get Images Usage批量获取image使用大小
##### URL
`GET /api/v1/ceph/rbd/:pool/images/usages`
##### Path Param
| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 池名称 | pool | string | | required |
##### Request Body

| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 镜像列表 | name | Array(GetImage) | | required |

`GetImage` object:

| 名称 | 变量名称 | 类型 | 备注 | Optional |
| :---: | :---: | :---: | :---: | :---: |
| 镜像名称 | name | string | | required |

```json
{
  "images": [
      {
          "name": "my-disk-1"
      },
      {
          "name": "my-disk-2"
      }
  ]
}
```
##### Response Body
```json
{
    "code": 1200,
    "msg": "查询成功",
    "data": [
        {
            "name": "my-disk-1",
            "used": "0B"
        },
        {
            "name": "my-disk-2",
            "used": "472MiB"
        }
    ]
}
```