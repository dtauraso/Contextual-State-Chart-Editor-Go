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
	// "golang.org/x/exp/slices"
	"reflect"
	"sort"
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
	editActive  []bool
}

func (sc *StateComponent) Render() app.UI {
	return sc.StateComponent()
}
func (sc *StateComponent) StateComponent2() app.UI {

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
func (sc *StateComponent) StateComponent() app.UI {
	editFlagIds := []int{
		0, 1,
	}
	return app.Ul().Body(
		&AppComponent2{},
		app.Range(editFlagIds).Slice(func(i int) app.UI {
			if len(ss.Names) == 0 {
				return app.Div()
			}
			if sc.editActive[i] {
				return app.Div().Body(
					app.Ul().
						Style("padding-left", "1rem").
						Body(),
					app.Input().
						Type("text").
						Value(ss.Names[i]).
						Placeholder("enter Atom name").
						AutoFocus(true).
						OnChange(func(ctx app.Context, e app.Event) {
							ss.Names[i] = ctx.JSSrc().Get("value").String()
							sc.editActive[i] = false
							sc.saveData2()
						}).
						Style("width", fmt.Sprintf("%dpx", len(ss.Name)*6)))
			}
			return app.Div().Body(
				app.Li().
					Style("list-style", "none").
					Style("width", fmt.Sprintf("%dpx", len(ss.Name)*6)).
					Text(ss.Names[i]).
					OnClick(func(ctx app.Context, e app.Event) { sc.UpdateEditFlag(i) }))
		}))

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
	ss.Names = make([]string, 2)
	ss.Names[0] = ss.Name
	ss.Names[1] = string(ss.SavedStates2[3].StringValue)
	sc.editActive = make([]bool, 2)
	sc.editActive[0] = false
	sc.editActive[1] = false
	sc.editActive1 = false
	sc.editActive2 = false
	sc.editActive3 = false

}
func (sc *StateComponent) UpdateEditFlag(flagId int) {

	sc.editActive3 = true
	if flagId >= 0 && flagId <= 1 {

		sc.editActive[flagId] = true
	}
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

func (sc *StateComponent) saveData2() app.UI {

	if reflect.DeepEqual(ss.SavedStates2, reflect.ValueOf(ss.SavedStates2).IsZero()) {
		return nil
	}
	t.SaveString(ss.SavedStates2, 2, ss.Names[0])
	binaryOutput, err := json.Marshal(ss.SavedStates2[2])
	if err != nil {
		panic(err)
	}

	output := url.Values{"Atom": {string(binaryOutput)}, "fileId": {"2"}}
	http.PostForm("/save", output)

	t.SaveString(ss.SavedStates2, 3, ss.Names[1])
	binaryOutput2, err2 := json.Marshal(ss.SavedStates2[3])
	if err2 != nil {
		panic(err2)
	}

	output2 := url.Values{"Atom": {string(binaryOutput2)}, "fileId": {"3"}}
	http.PostForm("/save", output2)

	return nil
}

/*
name           delete

	add

add

name

	nested name input map
	nested name input string
	nested name input int
	nested name input bool

next name input map
next name input string
next name input int
next name input int


Parents
	-1

Name

	this is a test child state name 4

	StarbucksMachine

FunctionName
	ReturnTrue

StartChildren

	Edges


			state1 name1

			state1 name2


			state2 name1

			state2 name2


	AreParallel
		true

Values

	drinkOrder
		[]
	orderQueue
		[]

*/

type AtomForm struct {
	app.Compo
	Graph              t.Graph
	Key                string
	IsKeyMapKey        bool
	AtomId             int
	IsEditActive       bool
	isAddChildActive   bool
	newChildMap        string
	newChildString     string
	newChildInt        string
	newChildBool       string
	isAddSiblingActive bool
	newSiblingMap      string
	newSiblingString   string
	newSiblingInt      string
	newSiblingBool     string
	siblingAtomType    string
	isDeleteActive     bool
	ParentAtom         int
	childAtoms         []int
}

func (a *AtomForm) Render() app.UI {

	atomId, _, returnType := a.Graph.GetAtom(
		a.AtomId,
		[]string{"AtomForm", "IsEditActive"})
	isEditActive := false
	if returnType == t.FOUND {
		isEditActive = a.Graph.Atoms[atomId].BoolValue
	}

	if isEditActive {
		return app.Li().
			Body(
				app.Span().Text(a.DisplayText()),
				app.Span().Text("delete"),
			).
			Style("display", "flex").
			Style("flex-direction", "row").
			OnClick(func(ctx app.Context, e app.Event) { a.UpdateEditFlag() })

	}
	return app.
		Li().
		Text(a.DisplayText()).
		OnClick(func(ctx app.Context, e app.Event) { a.UpdateEditFlag() })

}

func (a *AtomForm) UpdateEditFlag() {

	atomId, currentPath, returnType := a.Graph.GetAtom(
		a.AtomId,
		[]string{"AtomForm", "IsEditActive"})
	if returnType == t.FOUND {
		entry := a.Graph.Atoms[atomId]
		entry.BoolValue = !entry.BoolValue
		a.Graph.Atoms[atomId] = entry

	} else if len(currentPath) == 0 {
		atomFormId := len(a.Graph.Atoms)
		isEditActiveId := atomFormId + 1
		if entry, ok := a.Graph.Atoms[a.AtomId]; ok {
			if len(a.Graph.Atoms[a.AtomId].MapValues) == 0 {
				entry.MapValues = make(map[string]int)
				a.Graph.Atoms[a.AtomId] = entry
			}
		}
		a.Graph.Atoms[a.AtomId].MapValues["AtomForm"] = atomFormId
		a.Graph.Atoms[atomFormId] = t.Atom{
			Id:           atomFormId,
			MapValues:    map[string]int{"IsEditActive": isEditActiveId},
			TypeValueSet: "MapValues"}
		a.Graph.Atoms[isEditActiveId] = t.Atom{
			Id:           isEditActiveId,
			BoolValue:    true,
			TypeValueSet: "BoolValue"}

	}
	a.Update() // Manual updated trigger

}

func (a *AtomForm) DisplayText() any {
	if a.IsKeyMapKey {
		return a.Key
	}
	atom := a.Graph.Atoms[a.AtomId]
	return atom.Value()

}

type AtomUI struct {
	app.Compo
	AtomForms map[int]AtomForm
	Graph     t.Graph
}

func (a *AtomUI) OnMount(ctx app.Context) {

	res, err := http.Get("/loadAllStates")
	if err != nil {

		panic(err)
	}
	defer res.Body.Close()
	body, err2 := io.ReadAll(res.Body)
	if err2 != nil {

		panic(err2)
	}
	err3 := json.Unmarshal(body, &a.Graph.Atoms)
	if err3 != nil {
		panic(err3)
	}

}
func makeTreeHelper(atomId int, graph t.Graph) app.UI {

	atom := graph.Atoms[atomId]
	if atom.TypeValueSet == "MapValues" ||
		atom.TypeValueSet == "ArrayValues" {
		keys := []string{}
		for key := range atom.MapValues {
			if key != "AtomForm" {
				keys = append(keys, key)
			}
		}
		if len(keys) == 0 {
			return app.Ul().
				Body(app.Li().Text("[]").
					Style("list-style-type", "none").
					Style("margin-bottom", "20px"))
		}
		sort.Strings(keys)

		return app.Ul().
			Body(
				app.Range(keys).Slice(func(i int) app.UI {
					key := keys[i]
					// fmt.Println("range here")
					return app.Div().Body(
						app.If(atom.TypeValueSet == "MapValues",
							&AtomForm{
								Key:         key,
								IsKeyMapKey: true,
								AtomId:      atomId,
								Graph:       graph}),
						makeTreeHelper(atom.MapValues[key], graph),
					)
				}),
				app.Li().
					Text(" ").
					Style("list-style-type", "none").
					Style("margin-bottom", "20px")).
			Style("list-style-type", "none").
			Style("padding-left", "10px").
			Style("margin-bottom", "40px")

	}
	return app.Ul().
		Body(
			&AtomForm{
				IsKeyMapKey: false,
				AtomId:      atomId,
				Graph:       graph},
		).
		Style("list-style-type", "none").
		Style("margin-bottom", "10px")

}
func makeTree(a *AtomUI) app.UI {
	if len(a.Graph.Atoms) == 0 {
		return app.Div().Body(
			app.P().Text("no data"),
			&AtomForm{ParentAtom: 1})
	}

	return makeTreeHelper(0, a.Graph)
}
func (a *AtomUI) Render() app.UI {
	return makeTree(a)
}

type AppComponent2 struct {
	app.Compo
}

func MakeStyles(part app.HTMLDiv, styles [][]string, i int) app.HTMLDiv {
	if i >= len(styles) {
		return part
	}
	return MakeStyles(part.Style(styles[i][0], styles[i][1]), styles, i+1)
}

func (c *AppComponent2) Render() app.UI {
	return MakeStyles(
		app.Div().
			Body(app.P().Text("test")),
		[][]string{{"position", "absolute"},
			{"top", "1000px"}}, 0)

}
