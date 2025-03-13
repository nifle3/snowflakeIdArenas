package snowflakeid

type TypeToGenerate int32

const (
	Text TypeToGenerate = iota
	Base64
	Binary
)

type MachineId interface {
	Get() int16
}

type Generator struct {
	machineId MachineId
}

func NewGenerator(machine MachineId) *Generator {
	return &Generator{
		machineId: machine,
	}
}

func (g Generator) Generate(typeToGenerate TypeToGenerate) {
	g.generate()
	g.convert(Model{})
}

func (Generator) generate() {

}

func (Generator) convert(m Model) {
}
