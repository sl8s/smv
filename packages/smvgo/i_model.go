package smvgo

type IModel interface {
	Id() string
	Clone() IModel
	ToMap() map[string]any
}
