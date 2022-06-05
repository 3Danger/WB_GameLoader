package _interfaces

type IWallet interface {
	SendTo(float32, IWallet) error
	GetInfo() float32
	Receive(float32)
}
