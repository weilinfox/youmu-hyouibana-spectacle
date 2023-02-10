# hyouibana spectacle

从 0 开始的抓包和研判；协议的解释见 [zlib/main.go](zlib/main.go) 最下面的注释。

凭依华联机协议及观战服务器 demo

实现的功能：

+ 判断对战双方的状态
+ 存储对战数据
+ 独立于对战双方的观战服务器

对战方连接 ``127.0.0.1:4646`` ，观战方连接 ``127.0.0.1:4647``

所有的 zlib 貌似都不能用 golang 的 zlib 库来压缩， th155 不认，需要用 ``zlib.h`` 。

已经应用在 [thlink](https://github.com/weilinfox/youmu-thlink) 联机器中。
