package day05

type Almanac struct {
	Seeds      []int        `parser:"'seeds' ':' ( @Int+)*"`
	Converters []*Converter `parser:"@@*"`
}

type Converter struct {
	From  string  `parser:"@Ident '-'"`
	To    string  `parser:"'to' '-' @Ident 'map' ':'"`
	Rules []*Rule `parser:"@@*"`
}

type Rule struct {
	ToStart   int `parser:"@Int"`
	FromStart int `parser:"@Int"`
	Length    int `parser:"@Int"`
}

func (rule *Rule) IsApplicable(value int) bool {
	return value >= rule.FromStart && value < rule.FromStart+rule.Length
}

func (rule *Rule) Apply(value int) int {
	diff := value - rule.FromStart
	return rule.ToStart + diff
}

func (conv *Converter) Convert(value int) int {
	for _, rule := range conv.Rules {
		if rule.IsApplicable(value) {
			return rule.Apply(value)
		}
	}
	return value
}
