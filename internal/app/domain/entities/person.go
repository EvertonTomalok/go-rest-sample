package entities

type Person struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
