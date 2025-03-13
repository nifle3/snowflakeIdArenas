package snowflakeid

type Model struct {
	Base int64

	Text   string
	Binary int64
	Base64 string
}

func (m *Model) ToText() {
}

func (m *Model) ToBinary() {

}

func (m *Model) ToBase64() {

}
