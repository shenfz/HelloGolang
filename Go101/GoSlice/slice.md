# slice

> *  数组是一组同类型数据的集合，它是值类型，在初始化后长度是固定 
> *  切片是地址传递；切片可以通过数组来初始化

```go
	var (
		arrays = [...]int{1, 2, 3, 4, 5, 6}
		silces = []int{1, 2, 3, 4, 5}
		s2     = arrays[:]
	)
```

## [sliceAchieve切片实现](./sliceAchieve)
> 自定义结构实现切片

## [sliceCopy切片拷贝](./sliceCopy)
> unsafe 切片拷贝 

## [sliceAndArray](./sliceAndArray)
> 切片和数组比较： 地址、元素更新特性

## [sliceSort](./sliceSort)
> 排序