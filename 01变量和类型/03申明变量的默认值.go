package main

func main() {
	var (
		a int     // 0
		b bool    // false
		c string  // ""
		d float64 // 0
		e *int    // nil
		f rune    // 0,用来表示char类型，用int32来表示
	)
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(&e)
	println(f)
}
