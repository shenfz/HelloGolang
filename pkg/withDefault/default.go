package withDefault

/**
 * @Author shenfz
 * @Date 2024/5/19 16:41
 * @Email 1328919715@qq.com
 * @Description:
 **/

// Default returns defaultValue if value is zero, otherwise returns defaultValue.
//
//	Default("","xxx") // "xxx"
//	Default("yyy","xxx") // "yyy"
func Default[T comparable](value T, defaultValue T) T {
	var zero T
	if value == zero {
		return defaultValue
	}
	return value
}

// DefaultWithFunc returns defaultValue if value is zero, otherwise value.
//
//	DefaultWithFunc("", func() string { return "foo" }) // "foo"
//	DefaultWithFunc("bar", func() string { return "foo" }) // "bar"
func DefaultWithFunc[T comparable](value T, defaultValue func() T) T {
	var zero T
	if value == zero {
		return defaultValue()
	}
	return value
}
