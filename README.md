## server-sdk-go 版本说明：

### 使用说明
  * [FaaS SDK 使用指南](https://bytedance.feishu.cn/docx/XUIodmnkqojbFxxWNric2kQhnYf)
  * [Open SDK 使用指南](https://bytedance.feishu.cn/docx/YIc2dmy3ZozlHoxj0gEcBv1mnAc)
-------
### 版本：0.0.22 兼容升级
新功能：支持多分支开发能力

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
新功能：支持开发者获取飞书集成 app access token 及 tenant access token

### 版本：0.0.18 兼容升级
优化：废弃 GetContext().Flow.Variables 相关的 API

### 版本：0.0.17 兼容升级
新功能：支持模糊查询 FuzzySearch

-------
### 版本：0.0.15｜兼容升级
新功能：数据操作接口支持指定鉴权身份
- 不指定: 访问数据模型时, 权限 1.0 应用默认使用系统身份鉴权, 权限 2.0 应用默认使用用户身份鉴权
- useUserAuth(): 访问数据模型时, 使用用户身份鉴权
- useSystemAuth(): 访问数据模型时, 使用系统身份鉴权

-------
### 版本：0.0.14 兼容升级
新功能：支持流式查询 FindStream

优化：FindAll 标注为过时方法，用 FindStream 代替

-------

### 版本：0.0.12｜兼容升级
新功能：OQL 支持防注入用法

-------

### 版本：0.0.10 & 0.0.11｜兼容升级
优化：升级依赖包版本

-------
### 版本：0.0.9｜非兼容升级
优化：Find 仅查询 _id 字段时， limit 最大支持 10000

不兼容点：当 FindAll 传排序参数时，显式报错

-------
### 版本：0.0.8｜兼容升级
新功能：支持流程相关接口
* 获取流程实例中的人工任务信息
* 触发"点击按钮触发"类型的流程
* 获取流程实例信息
* 撤销包含人工任务的流程实例

-------
### 版本：0.0.7｜兼容升级
新功能：Tools 下支持 Error 处理及 Boe 泳道能力
* api.Tools.SetLaneNameToCtx
* api.Tools.ParseErr

优化：记录创建/更新/删除接口操作人由 **系统** 调整为 **实际触发人**

-------
### 版本：0.0.6｜兼容升级
修复：FindAll 查询的数据重复。

-------
### 版本：0.0.5｜兼容升级
修复：
* 事务批量更新时，修复了第1条记录失效的问题。
* RPC 请求时，如果 err 不为 nil，resp 为 nil 会导致 panic 的问题。

-------
### 版本：0.0.4｜兼容升级
修复：FaaS SDK 记录查询不支持两级及以上下钻。

-------
### 版本：0.0.3｜兼容升级
优化：记录创建接口和事务创建接口的返回结构体命名调整。
* RecordOnlyID => RecordID
* TransactionRecordOnlyID => TransactionRecordID

-------
### 版本：0.0.2｜兼容升级
修复：函数触发接口本地调试时，业务错误会被忽略。

-------
### 版本：0.0.1｜初始上线
新功能：FaaS SDK 和 Open SDK 初始功能上线。

