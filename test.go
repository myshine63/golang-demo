package main

type People struct {
	name string
}

func main() {

}

func (user *People) changeName(name string) {
	user.name = name
}
