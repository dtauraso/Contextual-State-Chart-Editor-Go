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

	ss "github.com/dtauraso/Contextual-State-Chart-Editor-Go/SavedStates"

	// t "github.com/dtauraso/Contextual-State-Chart-Editor-Go/ContextualStateChart"
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
	output := url.Values{"state": {string(binaryOutput)}, "fileID": {"2"}}
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
				Placeholder("enter state name").
				AutoFocus(true).
				OnChange(h.ValueTo(&ss.Name)),
		).OnClick(h.customTrigger),
	)
}

// basic structure for displaying 1 state
type StateComponent struct {
	app.Compo
	StateID int
}

func (sc *StateComponent) Render() app.UI {
	return sc.StateComponent()
}
func (sc *StateComponent) StateComponent() app.UI {
	return app.Div().Body(
		app.Ul().Body(
			app.Li().
				Style("list-style", "none").
				Text("list element"),
			app.Ul().Body(
				app.Li().Style("list-style", "none").Text("test"),
				app.Ul().Body(
					app.Li().Style("list-style", "none").Text("test 2"))),
		),
	)
}
