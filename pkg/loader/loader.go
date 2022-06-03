package loader

import (
	"GameLoaders/pkg/wallet"
	"errors"
	"math/rand"
)

type Loader struct {
	wallet.IWallet
	maxWeightTrans float32 //5-30kg
	salary         int     //ЗП
	fatigue        float32 //усталость
	drunk          bool
}

func NewLoader() *Loader {
	drunk := rand.Int()&1 == 0
	return &Loader{
		IWallet:        wallet.NewWallet(0),
		maxWeightTrans: rand.Float32()*25 + 5,
		salary:         rand.Intn(20) + 10,
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

type IUnload interface {
	Unload(unload float32)
}

func (l *Loader) Unload(task IUnload) error {
	//- формула рассчета переносимого веса
	// (вес*(100 - усталость/100)*(пьянство/100))
	canMoveWeight := l.CanMoveWeight()
	if canMoveWeight <= 0. {
		return errors.New("the loader is very tired")
	}
	task.Unload(canMoveWeight)
	l.fatigue += 0.2
	return nil
}
