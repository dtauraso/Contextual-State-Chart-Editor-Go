package StarbucksTree

// import (
// 	csc "github.com/dtauraso/Contextual-State-Chart-Editor-Go/ContextualStateChart"
// 	u "github.com/dtauraso/Contextual-State-Chart-Editor-Go/Utility"
// )

// type IStateNamePartTree struct {
// 	NPT       map[string]IStateNamePartTree
// 	Atom     IState
// 	DataTable Idatabase
// }

// type IState struct {
// 	FunctionCode        func(csc.Graph) bool
// 	EdgeKinds           map[string]IEdges
// 	Parents             []string
// 	Children            map[string]IStateNamePartTree
// 	HaveStartChildren   bool
// 	Variable            any
// 	Variables           map[string]any
// 	LockedByStates      map[string]bool
// 	LockedByStatesCount int
// 	Database            map[string]IStateNamePartTree
// }

// type Idatabase struct {
// 	Array []any
// 	Map   map[string]any
// }
// type IEdges struct {
// 	Edges       [][]string
// 	AreParallel bool
// }

// var Customer = IStateNamePartTree{
// 	NPT: map[string]IStateNamePartTree{
// 		"Cashier": {
// 			Atom: IState{
// 				FunctionCode: u.ReturnTrue,
// 				EdgeKinds: map[string]IEdges{
// 					"StartChildren": {
// 						Edges:       [][]string{{"Place order"}},
// 						AreParallel: false,
// 					},
// 				},
// 				HaveStartChildren: true,
// 				Children: map[string]IStateNamePartTree{
// 					"Place order": {
// 						Atom: IState{
// 							FunctionCode: u.ReturnTrue,
// 							EdgeKinds: map[string]IEdges{
// 								"Next": {
// 									Edges: [][]string{
// 										{"Dig up money"},
// 										{"Sip coffee"},
// 									},
// 									AreParallel: true,
// 								},
// 							},
// 							HaveStartChildren: false,
// 						},
// 					},
// 					"Dig up money": {
// 						Atom: IState{
// 							FunctionCode: u.ReturnTrue,
// 							EdgeKinds: map[string]IEdges{
// 								"Next": {
// 									Edges: [][]string{
// 										{"Put away change"},
// 									},
// 									AreParallel: true,
// 								},
// 							},
// 							HaveStartChildren:   false,
// 							LockedByStates:      map[string]bool{"Compute Price": true},
// 							LockedByStatesCount: 1,
// 						},
// 					},
// 					"Put away change": {
// 						Atom: IState{
// 							FunctionCode:        u.ReturnTrue,
// 							LockedByStates:      map[string]bool{"Compute change": true},
// 							LockedByStatesCount: 1,
// 						},
// 					},
// 					"Sip coffee": {
// 						Atom: IState{
// 							FunctionCode:        u.ReturnTrue,
// 							LockedByStates:      map[string]bool{"Output buffer": true},
// 							LockedByStatesCount: 1,
// 						},
// 					},
// 				},
// 				Variables: map[string]any{"drink": "frap choco"},
// 			}},
// 		"Barista": {
// 			Atom: IState{
// 				FunctionCode: u.ReturnTrue,
// 			},
// 		},
// 	},
// }

// var Cashier = IStateNamePartTree{
// 	Atom: IState{
// 		FunctionCode: u.ReturnTrue,
// 		EdgeKinds: map[string]IEdges{
// 			"StartChildren": {
// 				Edges: [][]string{
// 					{"Take order", "from customer"},
// 				},
// 				AreParallel: true,
// 			},
// 		},
// 		HaveStartChildren: true,
// 		Children: map[string]IStateNamePartTree{
// 			"Take order": {
// 				NPT: map[string]IStateNamePartTree{
// 					"from customer": {
// 						Atom: IState{
// 							FunctionCode: u.ReturnTrue,
// 							EdgeKinds: map[string]IEdges{
// 								"Next": {
// 									Edges: [][]string{
// 										{"Compute Price"},
// 									},
// 									AreParallel: false,
// 								},
// 							},
// 							HaveStartChildren:   false,
// 							LockedByStates:      map[string]bool{"Place order": true},
// 							LockedByStatesCount: 1,
// 						},
// 					}},
// 			},
// 			"Compute Price": {
// 				Atom: IState{
// 					FunctionCode: u.ReturnTrue,
// 					EdgeKinds: map[string]IEdges{
// 						"Next": {
// 							Edges: [][]string{
// 								{"Compute change"},
// 							},
// 							AreParallel: true,
// 						},
// 					},
// 					HaveStartChildren: false,
// 				},
// 			},
// 			"Compute change": {
// 				Atom: IState{
// 					FunctionCode: u.ReturnTrue,
// 					EdgeKinds: map[string]IEdges{
// 						"Next": {
// 							Edges: [][]string{
// 								{"No change"},
// 							},
// 							AreParallel: false,
// 						},
// 					},
// 					HaveStartChildren:   false,
// 					LockedByStates:      map[string]bool{"Dig up money": true},
// 					LockedByStatesCount: 1,
// 				},
// 			},
// 			"No change": {
// 				Atom: IState{
// 					FunctionCode: u.ReturnTrue,
// 				},
// 			},
// 		},
// 		Variables: map[string]any{
// 			"currentOrder": 23456,
// 			"price":        0,
// 		},
// 	},
// }

// var Barista = IStateNamePartTree{
// 	Atom: IState{
// 		FunctionCode: u.ReturnTrue,
// 		EdgeKinds: map[string]IEdges{
// 			"StartChildren": {
// 				Edges: [][]string{
// 					{"Make drink"},
// 				},
// 				AreParallel: true,
// 			},
// 		},
// 		HaveStartChildren: true,
// 		Children: map[string]IStateNamePartTree{
// 			"Make drink": {
// 				Atom: IState{
// 					FunctionCode: u.ReturnTrue,
// 					EdgeKinds: map[string]IEdges{
// 						"Next": {
// 							Edges: [][]string{
// 								{"Output buffer"},
// 							},
// 							AreParallel: false,
// 						},
// 					},
// 					HaveStartChildren: false,
// 				},
// 			},
// 			"Output buffer": {
// 				Atom: IState{
// 					FunctionCode: u.ReturnTrue,
// 					EdgeKinds: map[string]IEdges{
// 						"Next": {
// 							Edges: [][]string{},
// 						},
// 					},
// 					HaveStartChildren: false,
// 				},
// 			},
// 		},
// 	},
// }
// var StateTree = map[string]IStateNamePartTree{
// 	"machine": {
// 		NPT: map[string]IStateNamePartTree{
// 			"StarbucksMachine": {
// 				Atom: IState{
// 					FunctionCode: u.ReturnTrue,
// 					EdgeKinds: map[string]IEdges{
// 						"StartChildren": {
// 							Edges: [][]string{
// 								{"Register"},
// 								{"Barista"},
// 							},
// 							AreParallel: true,
// 						},
// 					},
// 					HaveStartChildren: true,
// 					Children: map[string]IStateNamePartTree{
// 						"Register": {
// 							Atom: IState{
// 								FunctionCode: u.ReturnTrue,
// 								EdgeKinds: map[string]IEdges{
// 									"StartChildren": {
// 										Edges: [][]string{
// 											{"Customer", "Cashier"},
// 											{"Cashier"},
// 										},
// 										AreParallel: true,
// 									},
// 									"Next": {
// 										Edges: [][]string{
// 											{"Customer", "Barista"},
// 										},
// 										AreParallel: false,
// 									},
// 								},
// 								HaveStartChildren: true,
// 								Children: map[string]IStateNamePartTree{
// 									"Customer": Customer,
// 									"Cashier":  Cashier,
// 								},
// 								Variables: map[string]any{
// 									"drinkPrice": 0,
// 									"change":     0,
// 								},
// 							},
// 						},
// 						"Barista": Barista,
// 					},
// 					Variables: map[string]any{
// 						"orderQueue":   []string{},
// 						"drinkOrder":   []string{},
// 						"outputBuffer": []string{},
// 					},
// 					Database: map[string]IStateNamePartTree{
// 						"names": {
// 							NPT: map[string]IStateNamePartTree{
// 								"Pistachio": {
// 									Atom: IState{
// 										EdgeKinds: map[string]IEdges{
// 											"Next": {
// 												Edges: [][]string{
// 													{"Pistachio Id"},
// 												},
// 											},
// 										},
// 									},
// 								},
// 								"Dark Caramel Sauce": {
// 									NPT: map[string]IStateNamePartTree{
// 										"flavor": {
// 											NPT: map[string]IStateNamePartTree{
// 												"Sauces": {
// 													Atom: IState{
// 														EdgeKinds: map[string]IEdges{
// 															"Next": {
// 																Edges: [][]string{
// 																	{"Dark Caramel Sauce", "flavor", "Sauces"},
// 																},
// 															},
// 														},
// 													},
// 												},
// 											},
// 										},
// 									},
// 								},
// 								"size": {
// 									NPT: map[string]IStateNamePartTree{
// 										"options": {
// 											Atom: IState{
// 												EdgeKinds: map[string]IEdges{
// 													"Next": {
// 														Edges: [][]string{
// 															{"size", "options"},
// 														},
// 													},
// 												},
// 											},
// 										},
// 									},
// 								},
// 								"Chocolate Cream Cold Foam": {
// 									NPT: map[string]IStateNamePartTree{
// 										"toppings": {
// 											NPT: map[string]IStateNamePartTree{
// 												"cold foam": {
// 													Atom: IState{
// 														EdgeKinds: map[string]IEdges{
// 															"Next": {
// 																Edges: [][]string{
// 																	{"Chocolate Cream Cold Foam", "toppings", "cold foam"},
// 																},
// 															},
// 														},
// 													},
// 												},
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 						"prices": {},
// 						"options": {
// 							DataTable: Idatabase{
// 								Array: []any{
// 									"size",
// 									"flavors",
// 									"toppings",
// 								},
// 							},
// 						},
// 						"sizes": {
// 							DataTable: Idatabase{
// 								Map: map[string]any{
// 									"large":  3,
// 									"grande": 2,
// 									"vente":  0,
// 								},
// 							},
// 						},
// 						"flavors": {
// 							NPT: map[string]IStateNamePartTree{
// 								"Sauces": {
// 									DataTable: Idatabase{
// 										Array: []any{
// 											"Dark Caramel Sauce",
// 											"Mocha Sauce",
// 										},
// 									},
// 								},
// 								"syrups": {
// 									DataTable: Idatabase{
// 										Array: []any{
// 											"Brown Sugar Syrup",
// 											"Caramel Syrup",
// 										},
// 									},
// 								},
// 							},
// 						},
// 						"toppings": {
// 							NPT: map[string]IStateNamePartTree{
// 								"cold foam": {
// 									DataTable: Idatabase{
// 										Map: map[string]any{
// 											"value":    "Chocolate Cream Cold Foam",
// 											"servings": 5,
// 											"price":    1,
// 										},
// 									},
// 								},
// 							},
// 						},
// 						"drinks": {
// 							Atom: IState{
// 								EdgeKinds: map[string]IEdges{
// 									"Next": {
// 										Edges: [][]string{
// 											{"Pistachio Id"},
// 										},
// 									},
// 								},
// 							},
// 						},
// 						"Pistachio Id": {
// 							Atom: IState{
// 								Database: map[string]IStateNamePartTree{
// 									"name": {
// 										Atom: IState{
// 											Variable: "Pistachio",
// 										},
// 									},
// 									"sizes": {
// 										Atom: IState{
// 											EdgeKinds: map[string]IEdges{
// 												"Next": {
// 													Edges: [][]string{
// 														{"id of sizes Atom"},
// 													},
// 												},
// 											},
// 										},
// 									},
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	},
// }

// func Test(input string) string {
// 	return input
// }
// func SayHello() string {
// 	return "Hi from package dir1"
// }
