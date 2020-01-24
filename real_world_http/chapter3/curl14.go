package main

type Struct struct {
	Member string
}

func main() {
	a := Struct{
		Member: "Value",
	}
	println(a.Member)
	b := new(Struct)
	println(b.Member)
	c := make(map[string]string)
	println(c.Member)
	d := library.New()
	// println(d.Member)
}
