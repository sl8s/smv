package smvgo

type IModelRepository[T IModel, Y IModels[T]] interface {
	IDispose
	FromMap(m map[string]any) T
	FromMaps(ms []map[string]any) Y
}

func GetSafeValue[T any](m map[string]any, key string, defaultValue T) T {
	value, existsValue := m[key]
	if !existsValue {
		return defaultValue
	}
	typedValue, existsTypedValue := value.(T)
	if !existsTypedValue {
		return defaultValue
	}
	return typedValue
}
