# http-share
使用Golang构建http/1.0，http/1.1和http/2.0服务，比较客户端页面的加载速度，页面由1个html文件、1个js文件和625张图片构成。

## 开启http服务

需要注意的是：http服务不要和客户端处于同一局域网内，否则难以看出其差别。

开启服务的方法：进入指定的目录，例如`http/1.0`目录，然后执行`go run main.go`即可。监听端口可以任意更改。

## 附加

我在知乎和公众号上分享了一篇文章——《HTTP版本差异和底层TCP细节》，如有兴趣可以阅读此文章。

知乎文章：[HTTP版本差异和底层TCP细节](https://zhuanlan.zhihu.com/p/498225744)

公众号文章：[HTTP版本差异和底层TCP细节](https://mp.weixin.qq.com/s/3p0wHebpHAMMEsUIimoGmA)

知乎账户：https://www.zhihu.com/people/choubeihai

公众号：周北海



