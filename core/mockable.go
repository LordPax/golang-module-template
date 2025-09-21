package core

type MockMethod struct {
	Method string
	Params []any
}

type IMockable[T any] interface {
	SetItems(items []T)
	SetItem(item T)
	GetItems() []T
	MockMethod(method string)
	MethodCalled(method string, params ...any)
	IsMethodCalled(method string) bool
	GetMethodParams(method string) []any
}

type Mockable[T any] struct {
	stubedItems  []T
	mockedMethod map[string]*MockMethod
}

func NewMockable[T any]() *Mockable[T] {
	return &Mockable[T]{
		mockedMethod: make(map[string]*MockMethod),
	}
}

func (m *Mockable[T]) SetItems(items []T) {
	m.stubedItems = items
}

func (m *Mockable[T]) AddItem(items T) {
	m.stubedItems = append(m.stubedItems, items)
}

func (m *Mockable[T]) GetItems() []T {
	return m.stubedItems
}

func (m *Mockable[T]) MethodCalled(method string, params ...any) {
	m.mockedMethod[method] = &MockMethod{method, params}
}

func (m *Mockable[T]) IsMethodCalled(method string) bool {
	_, ok := m.mockedMethod[method]
	return ok
}

func (m *Mockable[T]) GetMethodParams(method string) []any {
	called, ok := m.mockedMethod[method]
	if !ok {
		return nil
	}
	return called.Params
}

func (m *Mockable[T]) IsParamsEqual(method string, params ...any) bool {
	called, ok := m.mockedMethod[method]
	if !ok {
		return false
	}
	if len(called.Params) != len(params) {
		return false
	}
	for i, param := range params {
		if called.Params[i] != param {
			return false
		}
	}
	return true
}
