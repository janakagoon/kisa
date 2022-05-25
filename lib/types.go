package lib

type Motor string

const (
	ArmRotate      Motor = "ArmRotate"
	WristUpDown    Motor = "WristUpDown"
	WristLeftRight Motor = "WristLeftRight"
	IndexTip       Motor = "IndexTip"
	IndexMiddle    Motor = "IndexMiddle"
	IndexBase      Motor = "IndexBase"
	IndexSpread    Motor = "IndexSpread"
	MiddleTip      Motor = "MiddleTip"
	MiddleMiddle   Motor = "MiddleMiddle"
	MiddleBase     Motor = "MiddleBase"
	RingTip        Motor = "RingTip"
	RingMiddle     Motor = "RingMiddle"
	RingBase       Motor = "RingBase"
	RingSpread     Motor = "RingSpread"
	ThumbTip       Motor = "ThumbTip"
	ThumbMiddle    Motor = "ThumbMiddle"
	ThumbBase      Motor = "ThumbBase"
	ThumbSpread    Motor = "CounterClockwise"
)

type Movement string

const (
	Nil              Movement = "Nil"
	Open             Movement = "Open"
	Close            Movement = "Close"
	Left             Movement = "Left"
	Right            Movement = "Right"
	Up               Movement = "Up"
	Down             Movement = "Down"
	Clockwise        Movement = "Clockwise"
	CounterClockwise Movement = "CounterClockwise"
)
