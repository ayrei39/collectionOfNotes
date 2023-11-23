# Nginx

解决C10K问题（同时处理10000个并发连接）

Nginx可以解决5w个并发连接

## 安装
- 包管理安装
  - mac `brew install nginx`
  
  - linux 
    
    ```
    sudo apt update
    sudo install nginx
    ```
    
  - windows
    
    ```
    scoop install nginx
    choco install nginx
    ```
  
- 源码编译安装

- Docker安装 `docker pull nginx`

## 服务启停

### 启用

terminal输入`nginx` → 没有消息代表启动成功

浏览器localhost:8080`

启动后作为后台进程常驻`ps -ef|grep nginx`检查进程 

进程：

- master：主进程（唯一），读取&验证配置文件，管理worker进程
- worker：工作进程（可存在多个），处理实际请求，数量通过配置文件调整

> pid：进程ID
>
> `lsof -i:端口号`：查看端口占用

### 停止

`nginx -s signal`

signal:

- `quit`: 优雅停止
- `stop`:立即停止
- `reload`:重载配置文件
- `reopen`:重新打开日志文件

### 常用命令

```shell
nginx               # 启动Nginx
nginx -c filename   # 指定配置⽂件
nginx -V            # 查看Nginx的版本和编译参数等信息
nginx -t            # 检查配置⽂件是否正确，也可⽤来定位配置⽂件的位置
nginx -s quit       # 优雅停⽌Nginx
nginx -s stop       # 快速停⽌Nginx
nginx -s reload     # 重新加载配置⽂件
nginx -s reopen     # 重新打开⽇志⽂件
```

## 静态站点部署

> `nginx -V`:查看nginx版本、安装目录、编译参数、配置文件、日志文件的目录……

### 配置文件

`nginx.conf`

`server.location`

```
location / {
            root   html;
            index  index.html index.htm;
        }
```

> 1. `location / {`： 这行指定了一个location块，该块用于匹配URL路径为`/`的请求。`/`通常表示网站的根目录。
> 2. `root html;`： 这行指定了Nginx在文件系统中查找文件的根目录。在这个例子中，Nginx将在名为`html`的目录中寻找文件。这个目录通常是Nginx安装目录下的`html`子目录。
> 3. `index index.html index.htm;`： 这行指定了默认的文件索引顺序。当客户端请求的URL以`/`结尾时，Nginx会尝试在指定的根目录下查找文件，按照指定的顺序依次尝试文件名为`index.html`和`index.htm`的文件。第一个找到的文件将作为响应发送给客户端。
>
> 举例说明，如果有一个请求访问`http://yourdomain/`，Nginx将尝试在`html`目录下找到并返回`index.html`文件。如果找不到，它会继续尝试`index.htm`文件。这样的配置使得当用户访问站点的根目录时，可以自动返回站点的默认首页。
>
> 需要注意的是，Nginx的配置文件使用了块结构，用花括号 `{}` 括起来的一组配置指令被认为是一个块。在这个例子中，`location / {...}` 就是一个块，用于配置处理根路径请求的行为。

### 利用Hexo部署静态页面

> Hexo：
>
> - 安装：`npm install hexo-cli -g`
> - 初始化：`hexo init blog`
> - 安装依赖：`cd blog;npm install`
> - 本地运行：`hexo server / hexo s`
> - `hexo g`:把markdown文章转换成静态页面放到public目录下

#### 把blog部署到nginx

 ```shell
 cd public
 cp -rf * /opt/homebrew/var/www/
 ```

> 一键部署：`hexo d`

## 配置文件

### 结构

```
# 全局块
worker_processes  1;
events {
   # events块
}
http {
   # http块
   server {
       # server块
       location / {
           # location块
           }
   }
}
```

