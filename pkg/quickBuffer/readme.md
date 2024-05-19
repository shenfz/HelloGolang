# bytebufferpool

>  bytebufferpool 实现了自己的Buffer类型，并使用一个简单的算法降低扩容带来的性能损失。
bytebufferpool已经在大名鼎鼎的 Web 框架fasthttp和灵活的 Go 模块库quicktemplate得到了应用
> 按容量大小分20个区间 2^6 ---- 2^25
## 小细节：
> 容量最小值取 2^6 = 64，因为这就是 64 位计算机上 CPU 缓存行的大小。这个大小的数据可以一次性被加载到 CPU 缓存行中，再小就无意义了。
代码中多次使用atomic原子操作，避免加锁导致性能损失
## 缺点：
> 浪费部分内存
 
You can find results [here](https://omgnull.github.io/go-benchmark/buffer/).

# bytebufferpool users

* [fasthttp](https://github.com/valyala/fasthttp)
* [quicktemplate](https://github.com/valyala/quicktemplate)