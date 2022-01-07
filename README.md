# journal

#### 介绍

beego日志库的一键调用  

国内开源平台 gitee 地址为：https://gitee.com/bingcqust/journal

#### 软件说明

一行代码即可使用beego自带的logs日志库，将logs输出转到文件上，生成独立的log文件夹，并按照日期（年月日）进行分割，调用简单，可视化强

#### 安装教程

1.  import 导入```gitee.com/bingcqust/journal```或```github.com/Shishengbing/journal```
2.  在main函数中写入```journal.Start```
3.  在代码任意地方使用beego自带的logs.Error(),logs.Warnning()等方法

#### 使用说明

1.  源代码中配置的Warnning及以上级别的日志会被写入日志库，符合行业习惯
2.  日志保存在引入文件的第一级目录```./log/日志时间/info.log```
3.  日志目录会自动创建，后期考虑引入conf配置文件，使得日志导出级别、保存位置、日志名称等能够自定义

#### 参与贡献

1. BingCQUST

   

#### 其他

1.  欢迎大家与我交流

