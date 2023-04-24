package entity

type Model struct {
	Name      string
	Maxtokens int
}

func NewModel(name string, maxtokens int) *Model {
	return &Model{
		Name:      name,
		Maxtokens: maxtokens,
	}
}

func (m *Model) GetMaxTokens() int {
	return m.Maxtokens
}

func (m *Model) GetModelName() string {
	return m.Name
}
