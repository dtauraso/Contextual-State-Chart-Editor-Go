package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	f "github.com/sa-/slicefunk"
	x "github.com/dtauraso/Contextual-State-Chart-Editor-Go/Starbucks"
	csc "github.com/dtauraso/Contextual-State-Chart-Editor-Go/ContextualStateChart"
)

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo
}

type person struct {
	name string
	age int
}

func TestState(state x.IState) {
	fmt.Println(state)
}

func ReturnTrue(test csc.Graph) bool {return true}
// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	return app.H1().Text("Hello World!!")
}

func GetFunctionName(temp interface{}) string {
    strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
    return strs[len(strs)-1]
}
func GetType(myvar interface{}) string {
    if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
        return "*" + t.Elem().Name()
    } else {
        return t.Name()
    }
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
	fmt.Println(GetFunctionName(ReturnTrue), GetFunctionName(ReturnTrue) == "ReturnTrue")
	fmt.Println(GetType(x.StateTree), GetType(x.StateTree) == "IStateNamePart")
	fmt.Println(x.StateTree.State.Variables["test"])
	// fmt.Println(x.StateTree.State)
	// fmt.Println(x.StateTree.FunctionCode("I pass"))
	// fmt.Println(x.StateTree.State.FunctionCode("I pass again"))
	test := x.IState{FunctionCode: ReturnTrue, Variables: map[string]any{"test":x.IState{FunctionCode: ReturnTrue, EdgeKinds:nil}}}	// fmt.Println(test.INamePart)
	// fmt.Println(test.FunctionCode("I pass again 2"))
	fmt.Println(test)
	fmt.Println(test.Variables["test"])

	// fmt.Println(test.State.FunctionCode("I pass again 3"))
	// TestState(test)
		// TestState(*test.IState)

	original := []int{1, 2, 3, 4, 5}
    newArray := f.Map(original, func(item int) int { return item + 1 })
    newArray = f.Map(newArray, func(item int) int { return item * 3 })
    newArray = f.Filter(newArray, func(item int) bool { return item%2 == 0 })
    fmt.Println(newArray)

	people := []person{
		person{"test", 5},
		person{"test again", 20},
	}
	newPeople := f.Map(people, func(aPerson person) person {return person{aPerson.name, aPerson.age + 2}})
	fmt.Println(people)
	fmt.Println(newPeople)

	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	app.Route("/", &hello{})

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
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
