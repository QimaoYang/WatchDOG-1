# WatchDOG

## API Design
* POST powercubicle/v1/manage/passwd
（内/外）更新密钥

输入：密钥

返回：success/failed

* POST powercubicle/v1/seat
（内）生成座位对应二维码

输入：座位号

返回：二维码图片

* POST powercubicle/v1/seat/set
（外）扫描二维码占座

输入：二维码图片

返回：success/failed

* POST powercubicle/v1/userinfo
（外）用于注册

输入：用户名，密码，手机号

返回：success/failed

* GET powercubicle/v1/userinfo
（外）

返回：座位号/None