# Lint Scope

[x] 1、表达是否概念的字段，必须使用boolean、bool、tinyint(1) 、bit 类型。命名方式应当为系表结构如：is_deleted, is_public，符合表达习惯。

[x] 2、表名、字段名必须使用小写字母、数字、下划线组成，禁止数字开头，禁止两个下划线中间只出现数字。

[] 3、表名应当是名词词性，表示一种资源。表名不使用复数名词。

[x] 4、禁止使用 MySQL 保留字，如 desc、range、match、delayed、limit 等。请参考 MySQL 官方文档中规定的保留字。

[x] 5、主键索引名为 pk_字段名 , 唯一索引名为 uk_字段名 , 普通索引名为 idx_字段名 。

[x] 6、小数类型为 decimal，禁止 float、double。

[x] 7、varchar 是可变长字符串，不预先分配存储空间，长度不要超过 5000。

[x] 8、表必备三字段：id, created_at, updated_at

[x] 9、字符串类型字段，应当是 NOT NULL 的。

[x] 10、参与计算的字段，应当是 NOT NULL 的。

[x] 11、所有字段必须有 comment。

[x] 12、提交的 Migration 不得具有破坏性。

[x] 13、禁止修改字段名称。

[x] 14、禁止删除字段。

[] 15、一般情况下禁止修改数据类型。

[x] 16、migration 工具保障每一个 data SQL (DML) 脚本都处于同一个事务，以保障数据一致性，防止执行中断造成未知后果。因此无须开发人员自行编写事务。

[x] 17、必须统一使用 utf8mb4 编码。

[x] 18、InnoDB 单列索引长度不能超过 767 bytes，联合索引长度不能超过 3072 bytes，utf8mb4 每字符占 4 bytes，创建索引时应当注意此限制。

[] 19、join 应当先筛选相关字段再连接。

[x] 20、不得使用外键与级联，一切外键概念必须在应用层解决。