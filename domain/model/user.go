package model

type (
	User interface {
		ID() string
		Name() string

		SetID(id string)
		SetName(name string)
	}

	user struct {
		id   string
		name string
	}
)

func NewUser() User {
	return newUser()
}

func newUser() *user {
	return &user{}
}

func (m *user) ID() string {
	return m.id
}

func (m *user) Name() string {
	return m.name
}

func (m *user) SetID(id string) {
	m.id = id
}

func (m *user) SetName(name string) {
	m.name = name
}
