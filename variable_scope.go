package main

func main() {
	d := 1
	{
		// assignes to new instance of d inside  { }
		d := 2
		// assignes to global d outside { }
		//d = 2

		println(d)
	}
	println(d)
}
