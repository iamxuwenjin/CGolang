package structure

import (
	"fmt"
	"testing"
)

type Actor struct {
}

func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func TestCallEvent(t *testing.T) {
	actor := new(Actor)
	RegisterEvent("OnSkill", actor.OnEvent)
	RegisterEvent("OnSkill", GlobalEvent)
	CallEvent("OnSkill", 100)
}
