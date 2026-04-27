package player

import (
	"github.com/Mrollf025/collapse/game/attribute"
	"github.com/Mrollf025/collapse/game/skills"
)

type Player struct {
	ID         int
	Name       string
	Attributes attribute.AllAttribute
	Skills     skills.AllSkills
	CurZone    int
	CurX       int
	CurY       int
}
