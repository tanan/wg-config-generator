package handler

type MockCommand struct {
	CreatePrivateKeyFunc   func() (string, error)
	CreatePreSharedKeyFunc func() (string, error)
	CreatePublicKeyFunc    func(privateKey string) (string, error)
}

func (m *MockCommand) CreatePrivateKey() (string, error) {
	return m.CreatePrivateKeyFunc()
}

func (m *MockCommand) CreatePreSharedKey() (string, error) {
	return m.CreatePreSharedKeyFunc()
}

func (m *MockCommand) CreatePublicKey(privateKey string) (string, error) {
	return m.CreatePublicKeyFunc(privateKey)
}
