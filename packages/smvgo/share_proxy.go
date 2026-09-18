package smvgo

type ShareProxy interface {
	GetValue(key string, defaultValue any) any
	AddListener(key string, callback func(event any)) error
	DeleteListener(key string) error
	NotifyListener(key string, value any) error
	NotifyListeners(key string, value any) error
	DeleteListeners(key string)
	Update(key string, value any)
	Delete(key string)
	final()
}

type shareProxy struct {
	listenerId   string
	shareService ShareService
}

func (sp *shareProxy) GetValue(key string, defaultValue any) any {
	return sp.shareService.GetValue(key, defaultValue)
}

func (sp *shareProxy) AddListener(key string, callback func(event any)) error {
	return sp.shareService.AddListener(key, sp.listenerId, callback)
}

func (sp *shareProxy) DeleteListener(key string) error {
	return sp.shareService.DeleteListener(key, sp.listenerId)
}

func (sp *shareProxy) NotifyListener(key string, value any) error {
	return sp.shareService.NotifyListener(key, sp.listenerId, value)
}

func (sp *shareProxy) NotifyListeners(key string, value any) error {
	return sp.shareService.NotifyListeners(key, value)
}

func (sp *shareProxy) DeleteListeners(key string) {
	sp.shareService.DeleteListeners(key)
}

func (sp *shareProxy) Update(key string, value any) {
	sp.shareService.Update(key, value)
}

func (sp *shareProxy) Delete(key string) {
	sp.shareService.Delete(key)
}

func (sp *shareProxy) final() {
}

func NewShareProxy() ShareProxy {
	return &shareProxy{
		listenerId:   InstanceIterationService().Next(),
		shareService: InstanceShareService(),
	}
}
