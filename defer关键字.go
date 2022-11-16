package main

/***
defer关键字实例
https://www.w3cschool.cn/go_internals/go_internals-drq7282o.html#:~:text=Go%E8%AF%AD%E8%A8%80%20%E5%85%B3%E9%94%AE%E5%AD%97%EF%BC%9Adefer.%20defer%E5%92%8Cgo%E4%B8%80%E6%A0%B7%E9%83%BD%E6%98%AFGo%E8%AF%AD%E8%A8%80%E6%8F%90%E4%BE%9B%E7%9A%84%E5%85%B3%E9%94%AE%E5%AD%97%E3%80%82.%20defer%E7%94%A8%E4%BA%8E%E8%B5%84%E6%BA%90%E7%9A%84%E9%87%8A%E6%94%BE%EF%BC%8C%E4%BC%9A%E5%9C%A8%E5%87%BD%E6%95%B0%E8%BF%94%E5%9B%9E%E4%B9%8B%E5%89%8D%E8%BF%9B%E8%A1%8C%E8%B0%83%E7%94%A8%E3%80%82.,%E4%B8%80%E8%88%AC%E9%87%87%E7%94%A8%E5%A6%82%E4%B8%8B%E6%A8%A1%E5%BC%8F%EF%BC%9A.%20%E5%A6%82%E6%9E%9C%E6%9C%89%E5%A4%9A%E4%B8%AAdefer%E8%A1%A8%E8%BE%BE%E5%BC%8F%EF%BC%8C%E8%B0%83%E7%94%A8%E9%A1%BA%E5%BA%8F%E7%B1%BB%E4%BC%BC%E4%BA%8E%E6%A0%88%EF%BC%8C%E8%B6%8A%E5%90%8E%E9%9D%A2%E7%9A%84defer%E8%A1%A8%E8%BE%BE%E5%BC%8F%E8%B6%8A%E5%85%88%E8%A2%AB%E8%B0%83%E7%94%A8%E3%80%82.%20%E4%B8%8D%E8%BF%87%E5%A6%82%E6%9E%9C%E5%AF%B9defer%E7%9A%84%E4%BA%86%E8%A7%A3%E4%B8%8D%E5%A4%9F%E6%B7%B1%E5%85%A5%EF%BC%8C%E4%BD%BF%E7%94%A8%E8%B5%B7%E6%9D%A5%E5%8F%AF%E8%83%BD%E4%BC%9A%E8%B8%A9%E5%88%B0%E4%B8%80%E4%BA%9B%E5%9D%91%EF%BC%8C%E5%B0%A4%E5%85%B6%E6%98%AF%E8%B7%9F%E5%B8%A6%E5%91%BD%E5%90%8D%E7%9A%84%E8%BF%94%E5%9B%9E%E5%8F%82%E6%95%B0%E4%B8%80%E8%B5%B7%E4%BD%BF%E7%94%A8%E6%97%B6%E3%80%82.%20%E5%9C%A8%E8%AE%B2%E8%A7%A3defer%E7%9A%84%E5%AE%9E%E7%8E%B0%E4%B9%8B%E5%89%8D%E5%85%88%E7%9C%8B%E4%B8%80%E7%9C%8B%E4%BD%BF%E7%94%A8defer%E5%AE%B9%E6%98%93%E9%81%87%E5%88%B0%E7%9A%84%E9%97%AE%E9%A2%98%E3%80%82.
返回值 = xxx
调用defer函数
空的return
*/
func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f11() (result int) {
	result = 0
	defer func() {
		result++
	}()
	return
}
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f22() (r int) {
	t := 5
	r = t
	defer func() {
		t = t + 5
	}()
	return
}

func f3() (r int) {
	r = 1
	defer func(r int) {
		r = r + 5
	}(r)
	return
}
