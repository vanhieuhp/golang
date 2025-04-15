package simple

type simplePlugin struct {
	name string
}

func NewSimplePlugin(name string) *simplePlugin {
	return &simplePlugin{
		name: name,
	}
}

func (s *simplePlugin) GetPrefix() string {
	//TODO implement me
	panic("implement me")
}

func (s *simplePlugin) Get() interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *simplePlugin) Name() string {
	//TODO implement me
	panic("implement me")
}

func (s *simplePlugin) InitFlags() {
	//TODO implement me
	panic("implement me")
}

func (s *simplePlugin) Configure() error {
	//TODO implement me
	panic("implement me")
}

func (s *simplePlugin) Run() error {
	//TODO implement me
	panic("implement me")
}

func (s *simplePlugin) Stop() <-chan bool {
	//TODO implement me
	panic("implement me")
}
