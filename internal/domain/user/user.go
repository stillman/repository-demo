package user

const (
	StatusActive   = "active"
	StatusDisabled = "disabled"

	FieldID     = "id"
	FieldName   = "name"
	FieldStatus = "status"
)

// Model is a base model structure. it is used for embedding into other models
type Model struct {
	dirtyValues map[string]any
}

func (m *Model) DirtyValues() map[string]any {
	return m.dirtyValues
}

func (m *Model) ClearDirtyValues() {
	m.dirtyValues = make(map[string]any)
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

func (u *User) SetID(id int) {
	u.ID = id
	// do not update dirty value here because it's a PK
}

func (u *User) SetName(name string) {
	u.Name = name
	u.setDirtyValue(FieldName, name)
}

func (u *User) SetStatus(status string) {
	u.Status = status
	u.setDirtyValue(FieldStatus, status)
}
