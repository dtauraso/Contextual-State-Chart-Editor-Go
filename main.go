package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	x "github.com/dtauraso/Contextual-State-Chart-Editor-Go/Starbucks"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

type person struct {
	name string
	age  int
}

func TestState(state x.IState) {
	fmt.Println(state)
}

// IEEE_Software_Design_2PC.pdf
// IEEE Software Blog_ Your Local Coffee Shop Performs Resource Scaling.pdf
// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
// func (h *hello) Render() app.UI {
// 	return app.Div().Body(
// 		app.H1().Text("Hello World!!"),
// 		app.H1().Text("Hello World2!!"))

// }
type myCompo struct {
	app.Compo

	Number int
	Data   []string
}

var myTest = "test pass 3"

// func (c *myCompo) Render() app.UI {
// 	return app.Div().Text(c.Number)
// }

func (c *myCompo) customTrigger(ctx app.Context, e app.Event) {
	c.Number = rand.Intn(42)
	if c.Number < 30 {
		myTest = "test 2"
	} else {
		c.Data = []string{
			"test x 1",
			"test x 2",
		}
	}
	// fo, err := os.Create("output.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer func() {
	// 	if err := fo.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// err1 := os.WriteFile("output.txt", []byte("test"), 0644)
	// if err1 != nil {
	// 	panic(err1)
	// }
	c.Update() // Manual updated trigger
}

func (c *myCompo) x() app.UI {
	return app.Ul().Body(
		app.Range(c.Data).Slice(func(i int) app.UI {
			return app.Li().Text(c.Data[i])
		}),
	)
}
func (c *myCompo) Render() app.UI {
	return app.Div().Body(
		app.P().Text(c.Number),
		c.x(),
	).OnClick(c.customTrigger)
}

func (c *myCompo) onClick(ctx app.Context, e app.Event) {
	fmt.Println("onClick is called")
}

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {
	// fmt.Println(x.SayHello())
	// fmt.Println(x.X)
	// fmt.Println(x.StateTree)
	// fmt.Println(x.StateTree.NamePart)
	// fmt.Println(x.StateTree.State)
	// fmt.Println(u.GetFunctionName(u.ReturnTrue), u.GetFunctionName(u.ReturnTrue) == "ReturnTrue")
	// fmt.Println(reflect.TypeOf(x.StateTree), reflect.TypeOf(x.StateTree).String() == "map[string]StarbucksTree.IStateNamePartTree")
	// fmt.Println(x.StateTree.State.Variables["test"])
	// fmt.Println(x.StateTree.State)
	// fmt.Println(x.StateTree.FunctionCode("I pass"))
	// fmt.Println(x.StateTree.State.FunctionCode("I pass again"))
	// test := x.IState{FunctionCode: u.ReturnTrue, Variables: map[string]any{"test": x.IState{FunctionCode: u.ReturnTrue, EdgeKinds: nil}}} // fmt.Println(test.INamePart)
	// fmt.Println(test.FunctionCode("I pass again 2"))
	// fmt.Println(test)
	// fmt.Println(test.Variables["test"])

	// namesTrie := tt.TrieTree{}
	// namesTrie = tt.TrieTree.
	// 	Insert(namesTrie, tt.InsertParameters{Name: []string{"test"}, StateID: 0}).
	// 	Insert(tt.InsertParameters{Name: []string{"test", "test2"}, StateID: 1}).
	// 	Insert(tt.InsertParameters{Name: []string{"test", "test2", "test3"}, StateID: 2}).
	// 	Insert(tt.InsertParameters{Name: []string{"testx", "test2", "test3"}, StateID: 3})
	// fmt.Println(namesTrie)

	// TrieTree.search([]string{"test"}) = 0
	// TrieTree.search([]string{"test", "test2"}) = 1

	// expected output from trie tree from test inputs (lines: [58, 64])
	// expected output of trie tree from the starbucks tree
	// expected output of array of states from starbucks tree after init conversion
	// expected output of states run (async and non async) and variable changes

	// fmt.Println(test.State.FunctionCode("I pass again 3"))
	// TestState(test)
	// TestState(*test.IState)

	// original := []int{1, 2, 3, 4, 5}
	// newArray := f.Map(original, func(item int) int { return item + 1 })
	// newArray = f.Map(newArray, func(item int) int { return item * 3 })
	// newArray = f.Filter(newArray, func(item int) bool { return item%2 == 0 })
	// fmt.Println(newArray)

	// people := []person{
	// 	{"test", 5},
	// 	{"test again", 20},
	// }
	// newPeople := f.Map(people, func(aPerson person) person { return person{aPerson.name, aPerson.age + 2} })
	// fmt.Println(people)
	// fmt.Println(newPeople)

	// fmt.Print("\n")
	// fmt.Print(x.StateTree)

	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &myCompo{})

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example222",
	})
	http.HandleFunc("/test", func(http.ResponseWriter, *http.Request) {
		err1 := os.WriteFile("ContextualStateChart/TrieTree/output.txt", []byte(myTest), 0644)
		if err1 != nil {
			panic(err1)
		}

	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
