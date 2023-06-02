## server-sdk-go 版本说明：

-------
### 版本：0.0.19 兼容升级
优化：废弃 GetContext().Flow.Variables 相关的 API

-------
### 版本：0.0.18 兼容升级
修复：规范 FindOne 查询不到记录时的错误信息

-------
### 版本：0.0.16 兼容升级
新功能：支持模糊查询 FuzzySearch

-------
### 版本：0.0.14 兼容升级
新功能：支持流式查询 FindStream

优化：FindAll 标注为过时方法，用 FindStream 代替
