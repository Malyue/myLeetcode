package main

// 字符串的最大公因子
// 对于字符串S和T，只有在S = T + ... + T（T自身连接 1 次或多次）时，我们才认定“T 能除尽 S”
// 返回最长字符串X，要求满足X能除尽str1且X能除尽str2

func gcdOfStrings(str1 string, str2 string) string {
	// 辗转相除法
	// 保持str1是长度较长的
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	var remainder = mod(str1, str2)
	// 如果返回的是“”，说明刚好整除str2
	if remainder == "" {
		return str2
	} else if remainder == str1 {
		return ""
	}
	return gcdOfStrings(str2, remainder)
}

// 字符串取余数
// pwdpwdpwd pwd => "",其实取余数就是看是否前缀和str2一致
func mod(str1, str2 string) string {
	length := len(str2)
	var remainder = str1
	for {
		// 如果str1长度比str2小，或者str1的前面一段不是str2，则说明无法取余
		if len(remainder) < length || remainder[:length] != str2 {
			break
		}
		remainder = remainder[length:]
	}
	return remainder
}
