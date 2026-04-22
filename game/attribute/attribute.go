package attribute

type AllAttribute []Attribute

type Attribute struct {
	Name            string
	Value           int
	ProgressToLevel float64
	Description     string
}

var attributeStrings = []string{
	"Strength",
	"Agility",
	"Vitality",
	"Wisdom",
	"Intelligence",
}

func InitAttributes() AllAttribute {
	allAttribute := make(AllAttribute, len(attributeStrings))
	for i, name := range attributeStrings {
		allAttribute[i] = Attribute{
			Name:            name,
			Value:           10,
			ProgressToLevel: 0.0,
		}
	}
	return allAttribute
}
