package user

const (
	StatusActive = "active"

	FieldID     = "id"
	FieldName   = "name"
	FieldStatus = "status"
)

type Model struct {
	dirtyValues map[string]any
}

func (m *Model) GetDirtyValues() map[string]any {
	return m.dirtyValues
}

func (m *Model) setDirtyValue(name string, value any) {
	if m.dirtyValues == nil {
		m.dirtyValues = make(map[string]any)
	}

	m.dirtyValues[name] = value
}

type User struct {
	Model
	ID     int    `db:"id"`
	Name   string `db:"name"`
	Status string `db:"status"`
}

func (u User) Table() string {
	return "user"
}

func (u User) SelectFields() []string {
	return []string{"*"}
}

func (u User) PK() string {
	return FieldID
}

func (u User) PKValue() any {
	return u.ID
}

func (u *User) SetName(name string) {
	u.Name = name
	u.setDirtyValue(FieldName, name)
}
