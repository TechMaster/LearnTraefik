## Sử dụng traefik-forward-auth để xác thực bằng Google

Trước khi làm bài này cần bảo mật được Dashboard hãy xem bài [SecureDashBoard](../SecureDashBoard/ReadMe.md)

Mục tiêu sử dụng Google OAuth để xác thực và truy cập dịch vụ whoami ở địa chỉ http://localhost/whoami

Middleware sử dụng là [https://github.com/thomseddon/traefik-forward-auth](https://github.com/thomseddon/traefik-forward-auth)


Vào https://console.cloud.google.com/ để tạo `GOOGLE_CLIENT_ID` và `GOOGLE_CLIENT_SECRET`

![](GoogleAuth.jpg)

Triển khai thành công thì khi người dùng truy cập http://localhost/whoami, trình duyệt redirect Gmail, nếu người dùng đăng nhập thì được phép truy cập vào whoami

Còn nếu vào http://localhost:8080, bạn sẽ phải đăng nhập với userid: cuong, pass: xxxx để vào được dashboard. Bạn có thể đổi lại pass bằng cách vào https://hostingcanada.org/htpasswd-generator
rồi nhân đôi ký tự `$` lên.

