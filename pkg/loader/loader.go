package loader

import (
	. "GameLoaders/pkg/_interfaces"
	. "GameLoaders/pkg/wallet"
	"errors"
	"math/rand"
)

type Loader struct {
	IWallet
	name           string
	maxWeightTrans float32 //5-30kg
	salary         float32 //ЗП
	fatigue        float32 //усталость
	drunk          bool
}

func (l Loader) Salary() float32 { return l.salary }

func NewLoader(name string) *Loader {
	drunk := rand.Int()&1 == 0
	return &Loader{
		name:           name,
		IWallet:        NewWallet(0),
		maxWeightTrans: rand.Float32()*25 + 5,
		salary:         rand.Float32()*20 + 10,
		fatigue:        0.0,
		drunk:          drunk,
	}
}

func (l Loader) CanMoveWeight() float32 {
	fatigue := l.fatigue
	if l.drunk {
		fatigue += 0.5
	}
	if fatigue > 1.0 {
		fatigue = 1.0
	}
	return l.maxWeightTrans * (1.0 - fatigue)
}

func (l *Loader) Unload(task Unloadable) error {
	//- формула рассчета переносимого веса
	// (вес*(100 - усталость/100)*(пьянство/100))
	canMoveWeight := l.CanMoveWeight()
	if canMoveWeight <= 0. {
		return errors.New(l.name + " is very tired")
	}
	task.Unload(canMoveWeight)
	l.fatigue += 0.2
	return nil
}
