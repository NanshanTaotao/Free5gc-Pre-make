# Free5gc-pre-make
This project is clone from free5gc GitHub store and already make all file in NFs, you can clone this repository when you have troubles to clone the official free5gc repository.

**tips:**

1. clone this repository
2. unzip gtp5g.zip
3. cd gtp5g
4. sudo make
5. sudo make install
6. lsmod | grep gtp # check gtp5g installed status

**test command:**

1. cd ~/free5gc 
2. ./test.sh TestRegistration # success if show 'Pass'



这个项目来自于free5gc的官方仓库，已经安装官方文档中的安装教程make了所有NFs的文件，如果你在下载官方仓库代码或者make amf之类时出现问题并无法解决，可以尝试使用本仓库文件。

1. 下载源码
2. unzip gtp5g.zip # 解压gtp5g
3. cd gtp5g
4. sudo make
5. sudo make install
6. lsmod | grep gtp # 查看安装状态

测试安装状态：

1. cd ~/free5gc 
2. ./test.sh TestRegistration # 如果出现pass字段则说明安装成功



参考博客：[free5GC安装、运行、测试及注意事项_poplar-master的博客-CSDN博客_free5gc](https://blog.csdn.net/weixin_41973619/article/details/109121184)