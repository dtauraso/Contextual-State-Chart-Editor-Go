package app

import (
	// "encoding/json"
	// "fmt"
	// "io"
	// "log"
	// "math/rand"
	// "net/http"
	// "net/url"
	// "os"
	"encoding/json"
	"fmt"
	"reflect"

	// "time"
	// "fmt"
	// "fmt"

	// "fmt"
	// "reflect"
	// "strconv"

	// "fmt"
	"io"
	"net/http"
	"net/url"

	// "net/url"

	ss "Contextual-State-Chart-Editor-Go/SavedStates"

	t "Contextual-State-Chart-Editor-Go/ContextualStateChart"
	// x "github.com/dtauraso/Contextual-State-Chart-Editor-Go/Starbucks"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type AppComponent struct {
	app.Compo
}

func (c *AppComponent) Render() app.UI {
	return app.Div().Body(
		app.P().Text("test"),
	)
}

type Hello struct {
	app.Compo

	name string
}

func (c *Hello) customTrigger(ctx app.Context, e app.Event) {

	res, err := http.Get("/loadAllStates")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {

		panic(err2)
	}
	err3 := json.Unmarshal(body, &ss.SavedStates)
	if err3 != nil {
		panic(err3)
	}

	ss.Name = string(ss.SavedStates[2].StringValue)

	c.Update() // Manual updated trigger

}

func save() app.UI {

	if len(ss.SavedStates) == 0 {
		return nil
	}

	ss.SavedStates[2].StringValue = ss.Name
	binaryOutput, err := json.Marshal(ss.SavedStates[2])
	if err != nil {
		panic(err)
	}
	output := url.Values{"Atom": {string(binaryOutput)}, "fileId": {"2"}}
	http.PostForm("/save", output)
	return nil
}

func (h *Hello) Render() app.UI {

	return app.Div().Body(
		app.H1().Body(
			app.Text("Hello, "),
			app.If(ss.Name != "",
				app.Text(ss.Name),
				save(),
			).Else(
				app.Text("World!"),
			),
		),
		app.P().Body(
			app.Input().
				Type("text").
				Value(ss.Name).
				Placeholder("enter Atom name").
				AutoFocus(true).
				OnChange(h.ValueTo(&ss.Name)),
		).OnClick(h.customTrigger),
	)
}

// basic structure for displaying 1 Atom
type StateComponent struct {
	app.Compo
	StateId     int
	editActive1 bool
	editActive2 bool
	editActive3 bool
}

func (sc *StateComponent) Render() app.UI {
	return sc.StateComponent()
}
func (sc *StateComponent) OnMount(ctx app.Context) {
	res, err := http.Get("/loadAllStates")
	if err != nil {

		panic(err)
	}
	defer res.Body.Close()
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {

		panic(err2)
	}
	err3 := json.Unmarshal(body, &ss.SavedStates2)
	if err3 != nil {
		panic(err3)
	}

	ss.Name = string(ss.SavedStates2[2].StringValue)
	sc.editActive1 = false
	sc.editActive2 = false
	sc.editActive3 = false
}
func (sc *StateComponent) UpdateEditFlag(flagId int) {

	sc.editActive3 = true
	sc.Update() // Manual updated trigger

}

func (sc *StateComponent) saveData() app.UI {

	if reflect.DeepEqual(ss.SavedStates2, reflect.ValueOf(ss.SavedStates2).IsZero()) {
		return nil
	}
	t.SaveString(ss.SavedStates2, 2, ss.Name)
	binaryOutput, err := json.Marshal(ss.SavedStates2[2])
	if err != nil {
		panic(err)
	}

	output := url.Values{"Atom": {string(binaryOutput)}, "fileId": {"2"}}
	http.PostForm("/save", output)
	return nil
}

/*
name

	nested name input map
	nested name input string
	nested name input int
	nested name input bool

next name input map
next name input string
next name input int
next name input int
*/
func (sc *StateComponent) StateComponent() app.UI {

	return app.Div().
		Body(
			app.Ul().
				Style("position", "absolute").
				Style("top", "1000px").
				Style("left", "100px").
				Style("width", "200px").
				Style("height", "400px").
				Style("background-color", "#f39c12").
				Body(
					app.Li().
						Style("list-style", "none").
						Text("parent state name"),
					app.Ul().
						Style("padding-left", "1rem").
						Body(app.Li().
							Style("list-style", "none").
							Text("child state name 1"),
							app.If(sc.editActive3,
								app.Ul().
									Style("padding-left", "1rem").
									Body(),
								app.Input().
									Type("text").
									Value(ss.Name).
									Placeholder("enter Atom name").
									AutoFocus(true).
									OnChange(func(ctx app.Context, e app.Event) {
										ss.Name = ctx.JSSrc().Get("value").String()
										sc.editActive3 = false
										sc.saveData()
									}).
									Style("width", fmt.Sprintf("%dpx", len(ss.Name)*6)),
							).Else(
								app.Li().
									Style("list-style", "none").
									Style("width", fmt.Sprintf("%dpx", len(ss.Name)*6)).
									Text(ss.Name).
									OnClick(func(ctx app.Context, e app.Event) { sc.UpdateEditFlag(3) }),
							)),
				),
		)
}
