# GRPC-Study
grpc 相关的学习

### Protoc文件定义
- “包名.服务.proto”
- 语法
    - `syntax "proto3"`  定义proto版本号
    - `package ""` 定义package name
    - `service`    定义服务
    - `message`    定义消息
    - 字段序号
    - 字段操作选项
        - `optional` 可选参数
        - `repeated` 副数 list