package StarbucksTree

// import (
// 	csc "github.com/dtauraso/Contextual-State-Chart-Editor-Go/ContextualStateChart"
// 	u "github.com/dtauraso/Contextual-State-Chart-Editor-Go/Utility"
// )

// type IStateNamePartTree struct {
// 	NPT       map[string]IStateNamePartTree
// 	State     IState
// 	DataTable IDatabase
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

// type IDatabase struct {
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
// 			State: IState{
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
// 						State: IState{
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
// 						State: IState{
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
// 						State: IState{
// 							FunctionCode:        u.ReturnTrue,
// 							LockedByStates:      map[string]bool{"Compute change": true},
// 							LockedByStatesCount: 1,
// 						},
// 					},
// 					"Sip coffee": {
// 						State: IState{
// 							FunctionCode:        u.ReturnTrue,
// 							LockedByStates:      map[string]bool{"Output buffer": true},
// 							LockedByStatesCount: 1,
// 						},
// 					},
// 				},
// 				Variables: map[string]any{"drink": "frap choco"},
// 			}},
// 		"Barista": {
// 			State: IState{
// 				FunctionCode: u.ReturnTrue,
// 			},
// 		},
// 	},
// }

// var Cashier = IStateNamePartTree{
// 	State: IState{
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
// 						State: IState{
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
// 				State: IState{
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
// 				State: IState{
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
// 				State: IState{
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
// 	State: IState{
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
// 				State: IState{
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
// 				State: IState{
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
// 				State: IState{
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
// 							State: IState{
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
// 									State: IState{
// 										EdgeKinds: map[string]IEdges{
// 											"Next": {
// 												Edges: [][]string{
// 													{"Pistachio ID"},
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
// 													State: IState{
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
// 											State: IState{
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
// 													State: IState{
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
// 							DataTable: IDatabase{
// 								Array: []any{
// 									"size",
// 									"flavors",
// 									"toppings",
// 								},
// 							},
// 						},
// 						"sizes": {
// 							DataTable: IDatabase{
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
// 									DataTable: IDatabase{
// 										Array: []any{
// 											"Dark Caramel Sauce",
// 											"Mocha Sauce",
// 										},
// 									},
// 								},
// 								"syrups": {
// 									DataTable: IDatabase{
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
// 									DataTable: IDatabase{
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
// 							State: IState{
// 								EdgeKinds: map[string]IEdges{
// 									"Next": {
// 										Edges: [][]string{
// 											{"Pistachio ID"},
// 										},
// 									},
// 								},
// 							},
// 						},
// 						"Pistachio ID": {
// 							State: IState{
// 								Database: map[string]IStateNamePartTree{
// 									"name": {
// 										State: IState{
// 											Variable: "Pistachio",
// 										},
// 									},
// 									"sizes": {
// 										State: IState{
// 											EdgeKinds: map[string]IEdges{
// 												"Next": {
// 													Edges: [][]string{
// 														{"id of sizes state"},
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
