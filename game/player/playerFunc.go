package player

import (
	"github.com/Mrollf025/collapse/game/attribute"
	"github.com/Mrollf025/collapse/game/skills"
)

func newPlayer(name string) Player {
	return Player{
		ID:         1,
		Name:       name,
		Attributes: attribute.InitAttributes(),
		Skills:     skills.BuildSkillList(),
		CurZone:    0,
		CurX:       0,
		CurY:       0,
	}
}
