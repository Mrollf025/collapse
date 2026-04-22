package skills

type AllSkills struct {
	WeaponSkills    SkillList
	CombatSkills    SkillList
	CraftingSkills  SkillList
	GatheringSkills SkillList
	SecondarySkills SkillList
}

type SkillList []Skill

type Skill struct {
	Name            string
	Value           int
	ProgressToLevel float64
	Description     string
	IsWeapon        bool
	BaseAttribute   string
}

var WeaponSkillStrings = []string{
	"Knife",
	"Sword",
	"Spear",
	"Ax",
	"Mace",
	"Bow",
	"Gun",
	"Polearm",
	"Staff",
	"Hand to Hand",
}

var CombatSkillStrings = []string{
	"Parry",
	"Dodge",
	"Black Magic",
	"White Magic",
	"Poision",
	"Tactics",
	"Light Armor",
	"Medium Armor",
	"Heavy Armor",
}

var CraftingSkillStrings = []string{
	"Cooking",
	"Scribe",
	"Alchemy",
	"Smithing",
	"Tailoring",
	"Tanning",
	"Fletching",
	"Carpentry",
}

var GatheringSkillStrings = []string{
	"Minning",
	"Skinning",
	"Foraging",
	"Lumberjack",
	"Husbandry",
	"Horticulture",
}

var SecondarySkillStrings = []string{
	"First Aid",
	"Tracking",
	"Survival",
	"Animal Taming",
	"Identify",
	"Pick Locks",
}

func createSkills(names []string, isWeapon bool) SkillList {
	list := make(SkillList, len(names))
	for i, name := range names {
		list[i] = Skill{
			Name:            name,
			Value:           1,
			ProgressToLevel: 0.0,
			IsWeapon:        isWeapon,
			BaseAttribute:   "", // Initialize as empty or set a default
		}
	}
	return list
}

func BuildSkillList() AllSkills {
	return AllSkills{
		WeaponSkills:    createSkills(WeaponSkillStrings, true),
		CombatSkills:    createSkills(CombatSkillStrings, false),
		CraftingSkills:  createSkills(CraftingSkillStrings, false),
		GatheringSkills: createSkills(GatheringSkillStrings, false),
		SecondarySkills: createSkills(SecondarySkillStrings, false),
	}
}
