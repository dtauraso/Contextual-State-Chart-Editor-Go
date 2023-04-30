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

	// "fmt"
	"reflect"
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

	res, err := http.Get("/load?id=0")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {

		panic(err2)
	}
	err3 := json.Unmarshal(body, &ss.SavedState)
	if err3 != nil {
		panic(err3)
	}

	ss.SavedStates = append(ss.SavedStates, ss.SavedState)
	res2, err4 := http.Get("/load?id=1")
	if err4 != nil {
		panic(err4)
	}
	defer res2.Body.Close()
	body2, err5 := io.ReadAll(res2.Body)
	if err5 != nil {
		panic(err5)
	}
	err6 := json.Unmarshal(body2, &ss.SavedState2)
	if err6 != nil {
		panic(err6)
	}
	ss.SavedStates = append(ss.SavedStates, ss.SavedState2)

	res3, err7 := http.Get("/load?id=2")
	if err7 != nil {
		panic(err7)
	}
	defer res3.Body.Close()
	body3, err8 := io.ReadAll(res3.Body)
	if err8 != nil {
		panic(err8)
	}
	err9 := json.Unmarshal(body3, &ss.SavedState3)
	if err9 != nil {
		panic(err9)
	}
	ss.SavedStates = append(ss.SavedStates, ss.SavedState3)

	ss.Name = string(ss.SavedStates[2].StringValue)

	c.Update() // Manual updated trigger

}

func save() app.UI {

	if reflect.ValueOf(ss.SavedState).IsZero() {
		return nil
	}

	ss.SavedStates[2].StringValue = ss.Name
	fmt.Println(ss.SavedStates[2])
	binaryOutput, err := json.Marshal(ss.SavedStates[2])
	if err != nil {
		panic(err)
	}
	output := url.Values{"test": {string(binaryOutput)}, "fileID": {"2"}}
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
