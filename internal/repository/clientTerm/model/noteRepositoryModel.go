package model

// Its also a model only for pero calls that can be changed without affecting
// 		the main service code

type Notepad struct {
	Username string
	Notes    map[string]string 
}
