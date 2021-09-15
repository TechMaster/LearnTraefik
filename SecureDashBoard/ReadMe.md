# Bảo mật DashBoard Traefik sử dụng Basic Auth

Vào [https://hostingcanada.org/htpasswd-generator/](https://hostingcanada.org/htpasswd-generator/) để tạo Password. Nên chọn BCrypt

Chuỗi trả về như sau

```
cuong:$2y$10$A6vuYpsxe.NPH2wHtPdflOgDjHGScxSbrq0YqKgmJ3E8HmS7kzWVC"
```

Cần nhân đôi ký tự `$` lên thành `$$`
```
cuong:$$2y$$10$$A6vuYpsxe.NPH2wHtPdflOgDjHGScxSbrq0YqKgmJ3E8HmS7kzWVC"
```

Cấu hình Traefik - Docker compose như sau
```yaml
version: '3.8'

services:
  traefik:
    image: traefik:v2.5
    command:      
      - --providers.docker
      - --api.insecure=false
      - --api.dashboard=true
      - --entryPoints.web.address=:80
      - --entryPoints.api.address=:8080
    ports:
      - "80:80"
      - "8080:8080"
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    labels:          
      - "traefik.enable=true"                             
      - "traefik.http.routers.dashboard.rule=Host(`localhost`)"
      - "traefik.http.routers.dashboard.service=api@internal"   
      - "traefik.http.routers.dashboard.entrypoints=api"      
      - "traefik.http.routers.dashboard.middlewares=auth"
      - "traefik.http.middlewares.auth.basicauth.users=cuong:$$2y$$10$$A6vuYpsxe.NPH2wHtPdflOgDjHGScxSbrq0YqKgmJ3E8HmS7kzWVC"
```

