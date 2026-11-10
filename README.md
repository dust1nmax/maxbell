# Maxbell

一个基于 Go + Gin 的社区论坛后端服务，实现了用户注册登录、社区浏览、发帖、投票以及按时间/分数排序的帖子列表等功能，是一个类 Reddit / 掘金的社区项目。

## 技术栈

| 类别 | 技术 |
| --- | --- |
| Web 框架 | [Gin](https://github.com/gin-gonic/gin) |
| 数据库 | MySQL（[sqlx](https://github.com/jmoiron/sqlx)） |
| 缓存 / 投票计数 | Redis（[go-redis](https://github.com/go-redis/redis)，ZSet / Pipeline 事务） |
| 认证 | JWT（HS256，有效期 1 年） |
| ID 生成 | Snowflake 雪花算法 |
| 配置管理 | Viper（支持热加载） |
| 日志 | Zap + Lumberjack（日志切割），集成 Gin 日志与 Recovery 中间件 |
| 参数校验 | validator/v10 + 中文翻译器 |
| 部署 | Docker 多阶段构建、docker-compose、Air 热重载 |

## 功能特性

- **用户模块**：注册、登录，登录成功签发 JWT，密码加密存储
- **社区模块**：社区列表、社区详情
- **帖子模块**：创建帖子（雪花算法生成 ID）、帖子详情、分页帖子列表
- **投票模块**：赞成 / 反对 / 取消投票，基于 Redis 实现
  - 每票 432 分（86400 / 200，即 200 张赞成票可为帖子"续命"一天）
  - 帖子发布一周内允许投票，超时拒绝
  - 防止重复投票，支持改票与取消
- **帖子排序**：基于 Redis ZSet，支持按创建时间或投票分数排序 + 分页
- **优雅关机**：监听 SIGINT / SIGTERM，5 秒内处理完剩余请求后退出

## 项目结构

```
maxbell/
├── main.go             # 入口：初始化配置/日志/MySQL/Redis/雪花算法，启动服务
├── conf/               # 配置文件 config.yaml
├── settings/           # Viper 配置加载与热更新
├── routes/             # 路由注册
├── controllers/        # 请求处理：参数校验、响应封装、业务码、校验器翻译
├── logic/              # 业务逻辑层
├── dao/
│   ├── mysql/          # MySQL 数据访问（user / community / post）
│   └── redis/          # Redis 数据访问（投票、帖子分数、社区帖子集合）
├── models/             # 数据模型、请求参数、建表 SQL
├── middlewares/        # JWT 认证中间件
├── logger/             # Zap 日志初始化、Gin 日志中间件
├── pkg/
│   ├── jwt/            # JWT 生成与解析
│   └── snowflake/      # 雪花算法 ID 生成
├── static/ templates/  # 前端静态资源（Vue 构建产物）
├── Dockerfile          # 多阶段构建
├── docker-compose.yaml # MySQL + Redis + 应用编排
├── Makefile
└── .air.toml           # Air 热重载配置
```

## 快速开始

### 环境要求

- Go 1.25+
- MySQL 8.x
- Redis

### 1. 初始化数据库

创建数据库并导入表结构（含 4 个示例社区）：

```bash
mysql -u root -p < models/create_table.sql
```

### 2. 修改配置

编辑 `conf/config.yaml`，按实际环境修改 MySQL / Redis 连接信息：

```yaml
name: "maxbell"
mode: "dev"        # dev / release
port: 8081
start_time: "2026-10-09"  # 雪花算法起始时间
mechine_id: 1             # 机器 ID

mysql:
  host: "127.0.0.1"
  port: 13306
  user: "root"
  password: "1111"
  dbname: "maxbell"
  # ...

redis:
  host: "127.0.0.1"
  port: 16379
  db: 0
  # ...
```

> **注意**：程序当前从**工作目录下的 `config.yaml`** 读取配置（见 `settings/settings.go`），因此本地运行前需要把配置文件复制到项目根目录：
>
> ```bash
> cp conf/config.yaml ./config.yaml
> ```

### 3. 运行

```bash
# 直接运行
go run .

# 或使用 Makefile
make run      # go run ./
make build    # 交叉编译 Linux amd64 二进制
make all      # go fmt + go vet + build

# 开发模式（热重载，需先安装 air）
air
```

服务默认监听 `:8081`。

## API 接口

基础路径：`/api/v1`。除注册 / 登录外，其余接口需要在请求头携带 JWT：

```
Authorization: Bearer <token>
```

| 方法 | 路径 | 说明 | 认证 |
| --- | --- | --- | --- |
| POST | `/signup` | 注册（username / password / re_password） | 否 |
| POST | `/login` | 登录，返回 token | 否 |
| GET | `/community` | 社区列表 | 是 |
| GET | `/community/:id` | 社区详情 | 是 |
| POST | `/post` | 创建帖子（title / content / community_id） | 是 |
| GET | `/post/:id` | 帖子详情 | 是 |
| GET | `/posts/` | 帖子列表（page / size 分页） | 是 |
| GET | `/posts2/` | 帖子列表升级版（page / size / order=time\|score） | 是 |
| POST | `/vote` | 投票（post_id / direction：1 赞成、-1 反对、0 取消） | 是 |

### 响应格式

```json
{
  "code": 1000,
  "msg": "success",
  "data": {}
}
```

| code | 含义 |
| --- | --- |
| 1000 | 成功 |
| 1001 | 请求参数错误 |
| 1002 | 用户已存在 |
| 1003 | 用户不存在 |
| 1004 | 用户名或密码错误 |
| 1005 | 服务繁忙 |
| 1006 | 请求头中 auth 为空 |
| 1007 | 无效的 Token |

### 示例

```bash
# 注册
curl -X POST http://localhost:8081/api/v1/signup \
  -H "Content-Type: application/json" \
  -d '{"username":"tom","password":"123456","re_password":"123456"}'

# 登录
curl -X POST http://localhost:8081/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username":"tom","password":"123456"}'

# 按分数获取帖子列表
curl "http://localhost:8081/api/v1/posts2/?page=1&size=10&order=score" \
  -H "Authorization: Bearer <token>"

# 投赞成票
curl -X POST http://localhost:8081/api/v1/vote \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"post_id":"123456789","direction":"1"}'
```

## Docker 部署

项目提供了 `Dockerfile`（多阶段构建，最终为 scratch 小镜像）和 `docker-compose.yaml`（MySQL + Redis + 应用）。

> **注意**：当前的 Docker 编排文件尚不能直接运行，存在以下问题需要修正：
> - `docker-compose.yaml` 中 `depends_on: mysql8019` 引用了不存在的服务（MySQL 服务名为 `maxbell`）
> - `Dockerfile` 中 `COPY --from=builder /build/bubble /` 的二进制名与实际构建产物 `maxbell` 不一致
> - compose 中引用的 `wait-for.sh` 脚本不存在于仓库中
> - 端口映射不一致：应用配置为 8081，compose 映射 8888；MySQL 配置 13306 vs compose 33061；Redis 配置 16379 vs compose 26379
>
> 建议先本地运行验证，或修正上述问题后再使用 Docker 部署。

## 已知问题 / 待完善

- 按社区查询帖子列表的 handler（`GetCommunityPostListHandler2`）已实现，但尚未在 `routes/routes.go` 中注册路由
- `static/`、`templates/` 中已包含前端构建产物，但后端未配置对应的静态资源 / 页面路由
- 投票到期后将分数持久化到 MySQL 的逻辑尚未实现（见 `dao/redis/vote.go` 注释）

## License

本项目仅供学习交流使用。