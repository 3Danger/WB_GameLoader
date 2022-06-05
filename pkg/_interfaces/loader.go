package _interfaces

type ILoader interface {
	IWallet
	Salary() float32
	Unload(task Unloadable) error
}
