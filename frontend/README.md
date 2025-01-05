```
user  nobody;
worker_processes  10;
worker_rlimit_nofile 20000;

#error_log  logs/error.log;
#error_log  logs/error.log  notice;
#error_log  logs/error.log  info;
#error_log logs/error.log crit;  # 只记录严重错误

#pid        logs/nginx.pid;


events {
    worker_connections  4096;
}


http {
    include       mime.types;
    tcp_nodelay on;  # 确保小数据包不会被延迟发送

    sendfile        on;
    tcp_nopush       on; #需要在sendfile开启模式才有效，防止网路阻塞，积极的减少网络报文段的数量。将响应头和正文的开始部分一起发送，而不一个接一个的发送。

    keepalive_timeout  65;
    keepalive_requests 20000;  # 每个连接可以处理的最大请求数

    #gzip  on;

    server {
        listen       80;
        server_name  localhost;

        # 配置前端静态文件处理
        location /static/ {
            root /Users/guowei.gong/Documents/workspace/projects/Go_Entry_Task/frontend; # 静态文件存放的根目录
            expires 30d; # 设置静态文件的 HTTP 缓存过期时间
            access_log off; # 不记录访问静态文件的日志
        }

        # 配置页面路由
        # 指定根路由 '/' 直接指向 index.html
        location = / {
            root /Users/guowei.gong/Documents/workspace/projects/code-comment-analyzer/frontend;
            try_files /index.html =404;
            types { }
            default_type text/html; 
        }
        location = /index {
            root /Users/guowei.gong/Documents/workspace/projects/code-comment-analyzer/frontend;
            try_files /index.html =404;
            types { }
            default_type text/html; 
        }

        location / {
            proxy_pass http://127.0.0.1:9999; # 转发到后端服务器
            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection 'upgrade';
            proxy_set_header Host $host;
            proxy_cache_bypass $http_upgrade;
            proxy_set_header X-Real-IP $remote_addr; # 传递真实 IP
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
        }

        error_page  404              /404.html;
        location = /404.html {
            root   /Users/guowei.gong/Documents/workspace/projects/code-comment-analyzer/frontend/error_pages;
        }
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   /Users/guowei.gong/Documents/workspace/projects/code-comment-analyzer/frontend/error_pages;
        }
    }
    include servers/*;
}

```