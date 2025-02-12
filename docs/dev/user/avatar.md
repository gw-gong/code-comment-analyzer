# 头像处理说明

## 上传头像的保存流程

1. **文件存储**：
    - 头像文件将被存储在以下位置：`file_storage/avatars/{{userID}}/avatarFileName`。
    - 注意：`avatarFileName` 使用 UUID 生成，确保文件名的唯一性。

2. **数据库更新**：
    - 数据库中存储的头像地址只需文件名即可，即 `avatarFileName`。

功能函数：`util.GetAvatarStoragePath()`

## 返回给前端的头像资源路径格式

- 格式为：`/file_storage/avatars/avatarFileName`。
- 解释：最前面的斜杠是为了保障前端路由不会再拼接当前路由。如果路径是 `file_storage/avatars/avatarFileName`，前端路由可能会拼接当前路由，导致路径错误，例如：
    - `/index/file_storage/avatars/avatarFileName`
    - `/file/file_storage/avatars/avatarFileName`
    - `/user/file_storage/avatars/avatarFileName`

功能函数：`util.TransformProfilePictureUrlToResourceUrl()`

## 通过资源路径访问头像

- 先从 URL 中获取文件名，然后拼接路径：`file_storage/avatars/{{userID}}/avatarFileName`。
- 注意：前面没有斜杠。

功能函数：`util.GetAvatarStoragePath()`
