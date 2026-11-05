package main

// ---

type Singleton struct {
}

var singleInstance = &Singleton{}

// ---

func getInstance() *Singleton {
	return singleInstance
}

// ---

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}
}
