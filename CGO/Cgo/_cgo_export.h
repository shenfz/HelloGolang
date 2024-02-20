extern void helloString(_GoString_ p0);
extern void helloSlice(_GoSlice_ p0);\


/*
  使用了预定义 _GoString_ ，避免循环引用头文件依赖导致的问题
   更严谨的做法是为C语言函数接口定义严格的头文件，然后基于稳定的头文件实现代码


  用于获取字符串结构中的长度和指针信息
  size_t _GoStringLen(_GoString_ s);
  const char *_GoStringPtr(_GoString_ s);
*/