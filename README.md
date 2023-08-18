## server-sdk-go 版本说明：

-------
### 版本：0.0.20 兼容升级
新功能: 调整 Faas 相关 SDK 的挂载空间
- task 从 server-sdk 迁移到 baas-sdk
    - baas.task
- 新增包 faas-sdk，tool 和 function 从 server-sdk 迁移到 faas-sdk
    - faas.tool
    - faas.function
- 将原先server-sdk的其他相关能力（如api.Data、api.MetaData等）从api移动到application下
- 新增application.App用于获取应用信息
- 新增application.Event用于获取触发事件信息

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
