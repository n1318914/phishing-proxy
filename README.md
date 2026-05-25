域名准备：
直接配置cf， 然后设置通配符域名
服务器申请通配符域名

证书放到evilginx下面
~/.evilginx/certs/sites/xxx.com
fullchain.pem  privkey.pem(名字保持不变)



evilginx部署
1. install go  nodejs
2. git clone
3. make
4. setcap CAP_NET_BIND_SERVICE=+eip ./build/evilginx
5. cp phishlets  redirectors
6. ./evilginx 

## nginx配置
    server {
        listen 443 ssl http2;
    
        server_name chemicalguys.top *.chemicalguys.top;
    
        # SSL 证书（自签或正式证书）
        ssl_certificate     /root/.evilginx/crt/sites/chemicalguys.top/fullchain.pem;
        ssl_certificate_key /root/.evilginx/crt/sites/chemicalguys.top/privkey.pem;
    
        # TLS 配置
        ssl_protocols TLSv1.2 TLSv1.3;
        ssl_ciphers 'ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305';
        ssl_prefer_server_ciphers on;
        ssl_session_cache shared:SSL:10m;
        ssl_session_timeout 10m;
    
        # HSTS 强制 HTTPS
        add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
    
        # 日志
        access_log /var/log/nginx/evilginx.access.log;
        error_log  /var/log/nginx/evilginx.error.log warn;
    
        # 反代到 evilginx
        location / {
            proxy_pass https://127.0.0.1:18443;
    
            # -------------------------------
            # Host / Origin / Referer 完全保留浏览器访问域名
            # -------------------------------
            #proxy_set_header Host $host;
            #proxy_set_header Origin $scheme://$host;
            #proxy_set_header Referer $http_referer;
            proxy_set_header Host $http_host;
            proxy_set_header Origin $http_origin;
            proxy_set_header Referer $http_referer;
    
            # -------------------------------
            # 客户端真实 IP
            # -------------------------------
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
    
            proxy_pass_header Authorization;
            proxy_pass_header Set-Cookie;
            # -------------------------------
            # Cookie 保持原始 domain / path / SameSite
            # -------------------------------
            # 不做 proxy_cookie_domain 改写
            # proxy_cookie_path / /;
    
            # -------------------------------
            # WebSocket / HTTP/2 支持
            # -------------------------------
            proxy_http_version 1.1;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection "upgrade";
    
            # -------------------------------
            # SSL 反代设置
            # -------------------------------
            proxy_ssl_verify off;       # evilginx 自签或 TLS
            proxy_ssl_server_name on;
            proxy_ssl_name $host;
            proxy_ssl_protocols TLSv1.2 TLSv1.3;
    
            # -------------------------------
            # 超时 & 缓冲
            # -------------------------------
            proxy_connect_timeout 30s;
            proxy_read_timeout 90s;
            proxy_send_timeout 90s;
            client_max_body_size 20M;
    
        }
    }







<p align="center">
  <img alt="Evilginx2 Logo" src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/evilginx2-logo-512.png" height="160" />
  <p align="center">
    <img alt="Evilginx2 Title" src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/evilginx2-title-black-512.png" height="60" />
  </p>
</p>

# Evilginx 3.0

**Evilginx** is a man-in-the-middle attack framework used for phishing login credentials along with session cookies, which in turn allows to bypass 2-factor authentication protection.

This tool is a successor to [Evilginx](https://github.com/kgretzky/evilginx), released in 2017, which used a custom version of nginx HTTP server to provide man-in-the-middle functionality to act as a proxy between a browser and phished website.
Present version is fully written in GO as a standalone application, which implements its own HTTP and DNS server, making it extremely easy to set up and use.

<p align="center">
  <img alt="Screenshot" src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/screen.png" height="320" />
</p>

## Disclaimer

I am very much aware that Evilginx can be used for nefarious purposes. This work is merely a demonstration of what adept attackers can do. It is the defender's responsibility to take such attacks into consideration and find ways to protect their users against this type of phishing attacks. Evilginx should be used only in legitimate penetration testing assignments with written permission from to-be-phished parties.

## Evilginx Pro is now available!

This is it! After over two years of development, countless delays, and hundreds of manual company verifications, concluded with multiple hurdles related to export regulations, [Evilginx Pro is finally live!](https://evilginx.com)

<p align="center">
  <a href="https://evilginx.com"><img alt="Evilginx Mastery" src="https://breakdev.org/content/images/size/w2000/2025/03/evilginx_pro_release_cover.png" height="320" /></a>
</p>

Evilginx Pro is the fruit of a passion I've had for a long time in developing offensive security tools for cybersecurity enthusiasts. The journey has just begun, and now that the product is officially released, I can focus on making it even better by implementing all the ideas I've planned for it.

### Key features:

- Out-of-the-box **phishing detection evasion** (including Chrome's Enchanced Browser Protection)
- Tested and maintained **official phishlets database**
- **Botguard** to **prevent bot traffic** by default (same concept as Cloudflare Turnstile)
- **Evilpuppet** for advanced phishing capability (Google)
- External **DNS providers** with multi-domain support
- **Website spoofing** for unauthorized requests
- **JavaScript** & **HTML obfuscation**
- **Wildcard TLS certificates**
- **Automated** server deployment
- **SQLite** database support

Find out more on the [official release blog post](https://breakdev.org/evilginx-pro-release/).

## Evilginx Mastery Training Course

If you want everything about reverse proxy phishing with **Evilginx** - check out my [Evilginx Mastery](https://academy.breakdev.org/evilginx-mastery) course!

<p align="center">
  <a href="https://academy.breakdev.org/evilginx-mastery"><img alt="Evilginx Mastery" src="https://raw.githubusercontent.com/kgretzky/evilginx2/master/media/img/evilginx_mastery.jpg" height="320" /></a>
</p>

Learn everything about the latest methods of phishing, using reverse proxying to bypass Multi-Factor Authentication. Learn to think like an attacker, during your red team engagements, and become the master of phishing with Evilginx.

Grab it here:
https://academy.breakdev.org/evilginx-mastery

## Official Gophish integration

If you'd like to use Gophish to send out phishing links compatible with Evilginx, please use the official Gophish integration with Evilginx 3.3.
You can find the custom version here in the forked repository: [Gophish with Evilginx integration](https://github.com/kgretzky/gophish/)

If you want to learn more about how to set it up, please follow the instructions in [this blog post](https://breakdev.org/evilginx-3-3-go-phish/)

## Write-ups

If you want to learn more about reverse proxy phishing, I've published extensive blog posts about **Evilginx** here:

[Evilginx 2.0 - Release](https://breakdev.org/evilginx-2-next-generation-of-phishing-2fa-tokens)

[Evilginx 2.1 - First Update](https://breakdev.org/evilginx-2-1-the-first-post-release-update/)

[Evilginx 2.2 - Jolly Winter Update](https://breakdev.org/evilginx-2-2-jolly-winter-update/)

[Evilginx 2.3 - Phisherman's Dream](https://breakdev.org/evilginx-2-3-phishermans-dream/)

[Evilginx 2.4 - Gone Phishing](https://breakdev.org/evilginx-2-4-gone-phishing/)

[Evilginx 3.0](https://breakdev.org/evilginx-3-0-evilginx-mastery/)

[Evilginx 3.2](https://breakdev.org/evilginx-3-2/)

[Evilginx 3.3](https://breakdev.org/evilginx-3-3-go-phish/)

## Help

In case you want to learn how to install and use **Evilginx**, please refer to online documentation available at:

https://help.evilginx.com

## Support

I DO NOT offer support for providing or creating phishlets. I will also NOT help you with creation of your own phishlets. Please look for ready-to-use phishlets, provided by other people.

## License

**evilginx2** is made by Kuba Gretzky ([@mrgretzky](https://twitter.com/mrgretzky)) and it's released under BSD-3 license.
