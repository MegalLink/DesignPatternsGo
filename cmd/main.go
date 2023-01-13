package main

import (
	"fmt"

	"github.com/MegalLink/design-patterns/adapter"
	"github.com/MegalLink/design-patterns/bridge"
	"github.com/MegalLink/design-patterns/builder"
	"github.com/MegalLink/design-patterns/command"
	"github.com/MegalLink/design-patterns/composite"
	"github.com/MegalLink/design-patterns/decorator"
	"github.com/MegalLink/design-patterns/facade"
	"github.com/MegalLink/design-patterns/factory"
	"github.com/MegalLink/design-patterns/flyweight"
	"github.com/MegalLink/design-patterns/iterator"
	"github.com/MegalLink/design-patterns/logger"
	"github.com/MegalLink/design-patterns/mediator"
	"github.com/MegalLink/design-patterns/prototype"
	"github.com/MegalLink/design-patterns/proxy"
	"github.com/MegalLink/design-patterns/singleton"
)

func main() {
	fLogger, err := logger.NewFastLogger()
	if err != nil {
		println("Error")
	}

	builderPatternTest(fLogger)
	factoryPatternTest(fLogger)
	prototypePatternTest(fLogger)
	singletonPatternTest(fLogger)
	adapterPatternTest()
	bridgePatternTest()
	compositePatternTest(fLogger)
	decoratorPatternTest(fLogger)
	facadePatternTest(fLogger)
	flyweightPatternTest(fLogger)
	proxyPatterTest()
	commandPatternTest(fLogger)
	iteratorPatterTest(fLogger)
	mediatorPatternTest()
}

func builderPatternTest(logger logger.IFastLogger) {
	// builder pattern with director
	houseBuilder1 := builder.NewHouseBuilder()
	director := builder.NewDirector(houseBuilder1)
	newHouse := director.BuildHouse()
	logger.Info("builderPatternTest", newHouse)
}

func factoryPatternTest(logger logger.IFastLogger) {
	//factory interface
	person := factory.NewDefaultPerson("Jeff", 2, logger)
	person.SayHello()

	//factory generator

	developerFactory := factory.NewEmployeeFactory("developer", 1000)
	managerFactory := factory.NewEmployeeFactory("manager", 1200)

	developer := developerFactory("Jeff")

	manager := managerFactory("Josue")
	logger.Info("factoryPatternTest | factory generator", developer)
	logger.Info("factoryPatternTest | factory generator", manager)

	bossFactory := factory.NewEmployeeFactory2("CEO", 2000)
	boss := bossFactory.Create("Majo")
	boss.AnualIncome = 1000000
	logger.Info("factoryPatternTest | factory generator", boss)

	//factory prototype
	prototypeManager := factory.NewEmployeePrototype(factory.Developer)
	prototypeManager.Name = "Auroplay"
	logger.Info("factoryPatternTest | factory prototype", prototypeManager)
}

func prototypePatternTest(logger logger.IFastLogger) {
	john := prototype.Person{
		Name: "Jhon",
		Address: &prototype.Address{
			StreetAddress: "123",
			City:          "London",
			Country:       "UK"},
		Friends: []string{"One", "Two", "Three"},
	}

	// deep coopying, including pointers and slices
	james := john.DeepCopy()
	james.Name = "James"
	james.Address.StreetAddress = "123456789"
	james.Friends = append(james.Friends, "Four")
	// deep copy with seralization example
	elena := john.DeepCopySerialization()
	elena.Name = "Elena"
	elena.Address.StreetAddress = "Some Stree adress"
	elena.Friends = append(elena.Friends, "Anastasia")

	logger.Info("prototypePatternTest ", john)
	logger.Info("prototypePatternTest ", james)
	logger.Info("prototypePatternTest ", elena)

	// prototype implementation
	// instead of making copies by hand, better build prototype function so its easier to use,
	// now we dont have to set name for example by name in the struct like in the examples above
	jeff := prototype.NewPrototypedPerson("Jeff")
	logger.Info("prototypePatternTest", jeff)
}

type dummyDatabase struct {
	dummyData map[string]int
}

func (d *dummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
		}
	}
	return d.dummyData[name]
}

func singletonPatternTest(logger logger.IFastLogger) {
	db := singleton.GetSingletonDatabase()
	population := db.GetPopulation("one")
	logger.Info("singletonPatternTest", population)
	//normal singleton we cant do unit tests
	totalPopulation := singleton.GetTotalPopulation([]string{"zero", "two"})
	logger.Info("singletonPatternTest | total population", totalPopulation == 2)
	//better singleton to do unit tests
	dummyDB := &dummyDatabase{}
	totalPopulationDummy := singleton.GetTotalPopulationEx(dummyDB, []string{"alpha", "beta"})
	logger.Info("singletonPatternTest | total population", totalPopulationDummy == 3)
}

func adapterPatternTest() {
	//api response this with this type but in draw point function i have other type so i need adapter
	vectorRectangle := adapter.NewRectangle(5, 4)
	// this has nothing special is just a cast from one type to other
	rasterRectangle := adapter.VectorToRaster(vectorRectangle)
	fmt.Print(adapter.DrawPoints(rasterRectangle))
}

func bridgePatternTest() {
	raster := bridge.RasterRenderer{}
	vector := bridge.VectorRenderer{}
	// the trick in this patter is implement bridge interface to send as a param
	circle := bridge.NewCircle(&raster, 5)
	circle.Draw()
	circle.Resize(2)
	circle.Draw()
	circle2 := bridge.NewCircle(&vector, 10)
	circle2.Draw()
	circle2.Resize(2)
	circle2.Draw()
}

func compositePatternTest(logger logger.IFastLogger) {

	// graphic example , a graphic object can contain childs of same type
	drawing := composite.GraphicObject{Name: "My Drawing", Color: ""}
	drawing.Children = append(drawing.Children, *composite.NewCircle("Red"))
	drawing.Children = append(drawing.Children, *composite.NewSquare("Blue"))

	group := composite.GraphicObject{Name: "Group 1", Color: ""}
	group.Children = append(group.Children, *composite.NewCircle("Yellow"))
	group.Children = append(group.Children, *composite.NewSquare("Yellow"))
	drawing.Children = append(drawing.Children, group)
	logger.Info("compositePatternTest", "")
	fmt.Println(drawing.String())

	// neuron example a Neruon layer can contain Neurons, and Neurons can connect to other neurons
	neuron1, neuron2 := &composite.Neuron{}, &composite.Neuron{}
	layer1, layer2 := composite.NewNeuronLayer(3), composite.NewNeuronLayer(4)

	// now we want to just connect layer an neurons with just one function so composite pattern makes the trick
	// note: this works because NeuronLayers have Neurons
	// objects can use other objects via composition, some composend and singular objects need similar/identical behaviors
	// this pattern let us threat both types of objects uniformly
	composite.Connect(neuron1, neuron2)
	composite.Connect(neuron1, layer1)
	composite.Connect(layer2, neuron1)
	composite.Connect(layer1, layer2)
}

func decoratorPatternTest(logger logger.IFastLogger) {
	// multiple aggregation example
	dragon := decorator.Dragon{}
	dragon.SetAge(10)
	dragon.Fly()
	dragon.Crawl()

	// decorator example
	// basic structure
	circle := decorator.Circle{Radius: 2}
	circle.Resize(10) // we can resize this circle
	logger.Info("decoratorPatternTest | circle", circle.Render())

	// decorate with color
	redCircle := decorator.ColoredShape{Shape: &circle, Color: "Red"}
	//redCircle.Resize(10) we have lost resize in decorator :(
	logger.Info("decoratorPatternTest | circle", redCircle.Render())
	// we can decorate the redColor with transparency with this pattern
	redCircleTransparent := decorator.TransparentShape{Shape: &redCircle, Transparency: 0.5}
	logger.Info("decoratorPatternTest | circle", redCircleTransparent.Render())
	// be carefull to decorate other structs with same field names like we saw in dragon example
}

func facadePatternTest(logger logger.IFastLogger) {
	// in facade we implement an abstraction from a complex system , and we give this abstraction to user so he can use this simple
	// give simple API to something complicated that contains lots of components
	// expose internal funcitons trhough facade
	console := facade.NewConsole()
	caracter := console.GetCharacterAt(1)
	logger.Info("facadePatternTest", caracter)
}

func flyweightPatternTest(logger logger.IFastLogger) {
	// text format example
	text := "This is a brave new world"
	result := flyweight.NewFormattedText(text)
	logger.Info("flyweightPatternTest", result)
	result.Capitalize(10, 15)
	logger.Info("flyweightPatternTest capitalize", result.String())

	// better formated text with flyweight patter to avoid using to much memory
	bft := flyweight.NewBetterFormattedText(text)
	bft.Range(16, 19).Capitalize = true
	bft.Range(10, 14).Mask = true

	logger.Info("flyweightPatternTest capitalize better formated text", bft.String())
}

func proxyPatterTest() {
	// protected Proxy
	driver := proxy.NewCarProxy(&proxy.Driver{Age: 12})
	driver.Drive()

	fmt.Println("PROXY VIRTUAL EXAMPLE PROBLEM")
	// we we call NewBitMap we call to load image even if we are not drawing the image
	bmp := proxy.NewBitMap("demo1.png")
	proxy.DrawImage(bmp)
	proxy.DrawImage(bmp)
	// virtual proxy loads image lazyly on call of Draw, so we dont need to do de process of load the image unless is necessary
	fmt.Println("PROXY VIRTUAL SOLUTION")
	bmpLazy := proxy.NewLazyBitmap("demo.jpg")
	proxy.DrawImage(bmpLazy)
	proxy.DrawImage(bmpLazy)
}

func commandPatternTest(logger logger.IFastLogger) {
	myAccount := command.BankAccount{}
	cmd := command.NewBankAccountCommand(&myAccount, command.Deposit, 100)
	cmd.Call()
	cmd2 := command.NewBankAccountCommand(&myAccount, command.Withdraw, 50)
	cmd2.Call()
	logger.Info("commandPatternTest bankAccount", myAccount.GetBalance())
	// transfer from one bank account to other

	from := command.NewBankAccount(200)
	to := command.NewBankAccount(100)
	logger.Info("commandPatternTest transfer from to", fmt.Sprintf("From actual:%d to actual: %d", from.GetBalance(), to.GetBalance()))
	validCommand := command.NewMoneyTransferCommand(from, to, 50)
	validCommand.Call()
	logger.Info("commandPatternTest valid transfer from to", fmt.Sprintf("From result:%d to result: %d", from.GetBalance(), to.GetBalance()))
	validCommand.Undo()
	logger.Info("commandPatternTest valid transfer undo from to", fmt.Sprintf("From result:%d to result: %d", from.GetBalance(), to.GetBalance()))
	invalidCommand := command.NewMoneyTransferCommand(from, to, 600)
	invalidCommand.Call()
	logger.Info("commandPatternTest invalid transfer from to", fmt.Sprintf("From result:%d to result: %d", from.GetBalance(), to.GetBalance()))
}

func iteratorPatterTest(logger logger.IFastLogger) {
	root := iterator.NewNode(
		1, iterator.NewTerminalNode(2),
		iterator.NewTerminalNode(3),
	)
	it := iterator.NewInOrderIterator(root)

	logger.Info("iteratorPatterTest", root)
	for it.MoveNext() {
		logger.Info("current value", it.Current.Value)
	}
	// Iterator is not idiomatic in go but we can build it

	logger.Info("iteratorPatterTest tree implementation", root)
	t := iterator.NewBinaryTree(root)
	for i := t.InOrder(); i.MoveNext(); {
		logger.Info("current value", i.Current.Value)
	}
}

func mediatorPatternTest() {
	chatRoom := mediator.ChatRoom{}
	john := mediator.NewPerson("John")
	jane := mediator.NewPerson("Jane")

	chatRoom.JoinRoom(john)
	chatRoom.JoinRoom(jane)

	john.SendChatRoomMessage("hi room")
	john.SendChatRoomMessage("oh, hey jhon")

	simon := mediator.NewPerson("Simon")
	chatRoom.JoinRoom(simon)
	simon.SendChatRoomMessage("hi everyone!")

	jane.PrivateMessage("Simon", "Glad you could join us")
}
