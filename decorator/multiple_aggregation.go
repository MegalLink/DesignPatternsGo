package decorator

import "fmt"

//this patter is usefull when we want to add additional functionality to an object
// dont want to rewrite our existing code
// want to keep new functionality separate
// just decorate some existing structure
// solution : embed the decorated object and provide additional functionality

// problem combine some structures in single structure
// example a bird can fly, a lizard can crawl , and we want a dragon that can fly an crawl

type Bird struct {
	age int
}

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying")
	}
}

func (b *Bird) GetAge() int    { return b.age }
func (b *Bird) SetAge(age int) { b.age = age }

type Lizard struct {
	age int
}

func (l *Lizard) Crawl() {
	if l.age > 1 {
		fmt.Println("Crawling")
	}
}
func (l *Lizard) GetAge() int    { return l.age }
func (l *Lizard) SetAge(age int) { l.age = age }

/*create this way instead of directly
type Dragon struct{
	Bird
	Lizard
	}*/
type Dragon struct {
	bird   Bird
	lizard Lizard
}

// solution we cant do this like a proxy

// we need this interface to construc dragon, implement aged in lizard and bird
type Aged interface {
	GetAge() int
	SetAge(age int)
}

func NewDragon() *Dragon {
	return &Dragon{bird: Bird{}, lizard: Lizard{}}
}

func (d *Dragon) GetAge() int {
	return d.bird.GetAge()
}

func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}
