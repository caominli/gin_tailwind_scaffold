# Gin+tailwindcss脚手架

实现了基本导航，页脚，用户注册，注销，修改密码，通过jwt，token保存在cookie因为我是SSR+SPA双模式

### 安装

`npm install -D tailwindcss`  #安装
`npx tailwindcss init` #初始化
`npm install @headlessui/vue @heroicons/vue` #如果使用vue则可选安

#构建css
`npx tailwindcss -i ./static/tailwind.css -o ./static/index.css`

## 结构

controllers
- controller 视图控制器
- user.go 用户的视图控制器
models
- init.go 数据库初始化
- user.go 用户模型相关，验证器，验证码，数据模型

routers
- router.go 路由和gin初始化

services
- text.go 文本处理相关，密码加密，验证码，转大写
- jwt.go JWT的实现
- email.go 邮件发送实现

templates 模板文件

static 静态文件，js css
- theme.js 主题和 消息闪现


### 编译

在生产环境需要修改 models>init.go的Domain和Port注释
router.go 中正式模式反注释
`$env:GOOS='linux'; $env:GOARCH='amd64'; go build -o btcmai`


Linux正式编译
`$env:GOOS='linux'; $env:GOARCH='amd64'; go build -ldflags="-s -w" -o btcmai`

换回win
`$env:GOOS='windows'; $env:GOARCH='amd64';`

### 部署命令
#启动
sudo systemctl start btcmai

#查看运行
sudo systemctl status btcmai

#停止
sudo systemctl stop btcmai

#重启
sudo systemctl restart btcmai

#查看错误
sudo journalctl -u btcmai.service -r


### tailwindcss构建
`npx tailwindcss -i ./static/tailwind.css -o ./static/index.css`



