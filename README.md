# go-ceph - Go bindings for Ceph APIs

[![Build Status](https://travis-ci.org/ceph/go-ceph.svg)](https://travis-ci.org/ceph/go-ceph) [![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ceph/go-ceph) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/ceph/go-ceph/master/LICENSE)

## Installation
在这个项目中我直接把github.com/ceph/go-ceph的代码copy到了goceph目录下，为了以后方便扩展
一些对ceph的操作。

在go-ceph中调用了librados c的库所以在开发调试的机器上需要安装下面的库


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
