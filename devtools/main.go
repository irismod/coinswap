package main

func main() {
	ss := []rune("卧槽")
	j := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[j] = ss[j], ss[i]
		j--
	}
	println(string(ss))
}
