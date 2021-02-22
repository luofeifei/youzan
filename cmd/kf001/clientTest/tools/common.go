package tools

// 判断来源是否合法
func IsPlatform(plat string) (err bool) {
	return plat == "wx" || plat == "h5" || plat == "app" || plat == "web" || plat == "ali"
}