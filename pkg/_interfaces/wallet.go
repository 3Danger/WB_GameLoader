package _interfaces

type IWallet interface {
	SendTo(float32, IWallet) error
	Get() float32
	Receive(float32)
}
