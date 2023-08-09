```
├── api
|   ├── v1 # v1版本接口服务
|       ├── system # 系统级服务
|       └── enter.go # 统一入口
├── config # 配置相关
├── core # 核心模块
├── dao # dao层
├── global # 全局变量
├── initialize # 配置启动初始化
├── middleware # 中间件
├── model # 数据库结构体
├── router # 路由
├── service # 业务层
├── utils # 工具函数
├── config.yaml # 配置文件
├── go.mod # 包管理
├── main.go # 项目启动文件
└── README.md # 项目README
```

* mkdir  base-showcase
* cd  base-showcase

# init go.mod
go mod init  base-showcase

#  安装依赖
go get -u github.com/gin-gonic/gin
go get -u github.com/spf13/viper 
go get -u go.uber.org/zap/zapcore 
go get -u github.com/lestrrat-go/file-rotatelogs 













