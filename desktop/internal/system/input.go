package system

type Controller interface {
	Move(x, y int)
	Configure()
	CenterCursor()

}
