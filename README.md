# golang

以後台管理者為開發基礎進行開發
根據 RESTful API 協定
delete :: API 接口 - 刪除會員 (/v1/user/delete) 其請求方法以 DELETE 方式較佳而不是 POST
modifyPassword :: API 接口 - 修改會員密碼 (/v1/user/pwd/change) 其請求方法以 PATCH 方式較佳而不是 POST

因為以後台管理者為開發基礎進行開發
因此模擬後台管理者進行會員相關操作，且會有以下行為忽略
create :: API 接口 - 新增會員 (/v1/user/create) 沒有加入 confrim password 的作法
modifyPassword :: API 接口 - 修改會員密碼 (/v1/user/pwd/change) 沒有進行會員密碼驗證的作法
