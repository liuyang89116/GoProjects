package main

func setup(task string) func() {
	println("do some setup for", task)
	return func() {
		println("do some teardown for", task)
	}
}

func main() {
	teardown := setup("demo")
	defer teardown()
	println("do some business stuff")
}

/*
	do some setup for demo
	do some business stuff
	do some teardown for demo
*/
