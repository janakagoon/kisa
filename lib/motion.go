package lib

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

type Switch struct {
	Board int64
	Relay int64
}

type HBridge struct {
	S1 Switch
	S2 Switch
}

type MaxMinFlex struct {
	Max AngleDegrees
	Min AngleDegrees
}

type AngleDegrees float64

type MotionController struct {
	MotorHBridge        map[Motor]*HBridge
	MotorParams         map[Motor]MotorParam
	MotorMovementSwitch map[Motor]map[Movement]int
	MotorSwitchMovement map[Motor]map[int]Movement
	Mock                bool
	MockRelayStatus     map[int64]map[int64]RelayStatus

	RelativeFlexPosition map[Motor]AngleDegrees
}

var RelayIds = [8]int64{1, 2, 3, 4, 5, 6, 7, 8}
var BoardIds = [5]int64{0, 1, 2, 3, 4}

type MotorParam struct {
	MinFlex                float64
	MaxFlex                float64
	FlexPerSecondEnergized float64
}

func NewMotionController(mock bool) (*MotionController, error) {
	result := MotionController{}

	result.MotorHBridge = make(map[Motor]*HBridge)

	result.MotorParams = make(map[Motor]MotorParam)

	result.MotorMovementSwitch = make(map[Motor]map[Movement]int)
	result.MotorSwitchMovement = make(map[Motor]map[int]Movement)

	result.RelativeFlexPosition = make(map[Motor]AngleDegrees)

	// ArmRotate
	// Not Impl
	result.MotorMovementSwitch[ArmRotate] = make(map[Movement]int)
	result.MotorMovementSwitch[ArmRotate][Clockwise] = 1
	result.MotorMovementSwitch[ArmRotate][CounterClockwise] = 2

	result.MotorSwitchMovement[ArmRotate] = make(map[int]Movement)
	result.MotorSwitchMovement[ArmRotate][1] = Clockwise
	result.MotorSwitchMovement[ArmRotate][2] = CounterClockwise

	// WristUpDown
	result.MotorMovementSwitch[WristUpDown] = make(map[Movement]int)
	result.MotorMovementSwitch[WristUpDown][Up] = 1
	result.MotorMovementSwitch[WristUpDown][Down] = 2

	result.MotorSwitchMovement[WristUpDown] = make(map[int]Movement)
	result.MotorSwitchMovement[WristUpDown][1] = Up
	result.MotorSwitchMovement[WristUpDown][2] = Down

	// WristLeftRight
	result.MotorMovementSwitch[WristLeftRight] = make(map[Movement]int)
	result.MotorMovementSwitch[WristLeftRight][Left] = 1
	result.MotorMovementSwitch[WristLeftRight][Right] = 2

	result.MotorSwitchMovement[WristLeftRight] = make(map[int]Movement)
	result.MotorSwitchMovement[WristLeftRight][1] = Left
	result.MotorSwitchMovement[WristLeftRight][2] = Right

	// IndexTip
	result.MotorMovementSwitch[IndexTip] = make(map[Movement]int)
	result.MotorMovementSwitch[IndexTip][Open] = 1
	result.MotorMovementSwitch[IndexTip][Close] = 2

	result.MotorSwitchMovement[IndexTip] = make(map[int]Movement)
	result.MotorSwitchMovement[IndexTip][1] = Open
	result.MotorSwitchMovement[IndexTip][2] = Close

	// IndexMiddle
	result.MotorMovementSwitch[IndexMiddle] = make(map[Movement]int)
	result.MotorMovementSwitch[IndexMiddle][Open] = 1
	result.MotorMovementSwitch[IndexMiddle][Close] = 2

	result.MotorSwitchMovement[IndexMiddle] = make(map[int]Movement)
	result.MotorSwitchMovement[IndexMiddle][1] = Open
	result.MotorSwitchMovement[IndexMiddle][2] = Close

	// IndexBase
	result.MotorMovementSwitch[IndexBase] = make(map[Movement]int)
	result.MotorMovementSwitch[IndexBase][Open] = 1
	result.MotorMovementSwitch[IndexBase][Close] = 2

	result.MotorSwitchMovement[IndexBase] = make(map[int]Movement)
	result.MotorSwitchMovement[IndexBase][1] = Open
	result.MotorSwitchMovement[IndexBase][2] = Close

	// IndexSpread
	result.MotorMovementSwitch[IndexSpread] = make(map[Movement]int)
	result.MotorMovementSwitch[IndexSpread][Open] = 1
	result.MotorMovementSwitch[IndexSpread][Close] = 2

	result.MotorSwitchMovement[IndexSpread] = make(map[int]Movement)
	result.MotorSwitchMovement[IndexSpread][1] = Open
	result.MotorSwitchMovement[IndexSpread][2] = Close

	// MiddleTip
	result.MotorMovementSwitch[MiddleTip] = make(map[Movement]int)
	result.MotorMovementSwitch[MiddleTip][Open] = 1
	result.MotorMovementSwitch[MiddleTip][Close] = 2

	result.MotorSwitchMovement[MiddleTip] = make(map[int]Movement)
	result.MotorSwitchMovement[MiddleTip][1] = Open
	result.MotorSwitchMovement[MiddleTip][2] = Close

	// MiddleMiddle
	result.MotorMovementSwitch[MiddleMiddle] = make(map[Movement]int)
	result.MotorMovementSwitch[MiddleMiddle][Open] = 1
	result.MotorMovementSwitch[MiddleMiddle][Close] = 2

	result.MotorSwitchMovement[MiddleMiddle] = make(map[int]Movement)
	result.MotorSwitchMovement[MiddleMiddle][1] = Open
	result.MotorSwitchMovement[MiddleMiddle][2] = Close

	// MiddleBase
	result.MotorMovementSwitch[MiddleBase] = make(map[Movement]int)
	result.MotorMovementSwitch[MiddleBase][Open] = 1
	result.MotorMovementSwitch[MiddleBase][Close] = 2

	result.MotorSwitchMovement[MiddleBase] = make(map[int]Movement)
	result.MotorSwitchMovement[MiddleBase][1] = Open
	result.MotorSwitchMovement[MiddleBase][2] = Close

	// RingTip
	result.MotorMovementSwitch[RingTip] = make(map[Movement]int)
	result.MotorMovementSwitch[RingTip][Open] = 1
	result.MotorMovementSwitch[RingTip][Close] = 2

	result.MotorSwitchMovement[RingTip] = make(map[int]Movement)
	result.MotorSwitchMovement[RingTip][1] = Open
	result.MotorSwitchMovement[RingTip][2] = Close

	// RingMiddle
	result.MotorMovementSwitch[RingMiddle] = make(map[Movement]int)
	result.MotorMovementSwitch[RingMiddle][Open] = 1
	result.MotorMovementSwitch[RingMiddle][Close] = 2

	result.MotorSwitchMovement[RingMiddle] = make(map[int]Movement)
	result.MotorSwitchMovement[RingMiddle][1] = Open
	result.MotorSwitchMovement[RingMiddle][2] = Close

	// RingBase
	result.MotorMovementSwitch[RingBase] = make(map[Movement]int)
	result.MotorMovementSwitch[RingBase][Open] = 1
	result.MotorMovementSwitch[RingBase][Close] = 2

	result.MotorSwitchMovement[RingBase] = make(map[int]Movement)
	result.MotorSwitchMovement[RingBase][1] = Open
	result.MotorSwitchMovement[RingBase][2] = Close

	// RingSpread
	result.MotorMovementSwitch[RingSpread] = make(map[Movement]int)
	result.MotorMovementSwitch[RingSpread][Open] = 1
	result.MotorMovementSwitch[RingSpread][Close] = 2

	result.MotorSwitchMovement[RingSpread] = make(map[int]Movement)
	result.MotorSwitchMovement[RingSpread][1] = Open
	result.MotorSwitchMovement[RingSpread][2] = Close

	// ThumbTip
	result.MotorMovementSwitch[ThumbTip] = make(map[Movement]int)
	result.MotorMovementSwitch[ThumbTip][Open] = 1
	result.MotorMovementSwitch[ThumbTip][Close] = 2

	result.MotorSwitchMovement[ThumbTip] = make(map[int]Movement)
	result.MotorSwitchMovement[ThumbTip][1] = Open
	result.MotorSwitchMovement[ThumbTip][2] = Close

	// ThumbMiddle
	result.MotorMovementSwitch[ThumbMiddle] = make(map[Movement]int)
	result.MotorMovementSwitch[ThumbMiddle][Open] = 1
	result.MotorMovementSwitch[ThumbMiddle][Close] = 2

	result.MotorSwitchMovement[ThumbMiddle] = make(map[int]Movement)
	result.MotorSwitchMovement[ThumbMiddle][1] = Open
	result.MotorSwitchMovement[ThumbMiddle][2] = Close

	// ThumbBase
	result.MotorMovementSwitch[ThumbBase] = make(map[Movement]int)
	result.MotorMovementSwitch[ThumbBase][Open] = 1
	result.MotorMovementSwitch[ThumbBase][Close] = 2

	result.MotorSwitchMovement[ThumbBase] = make(map[int]Movement)
	result.MotorSwitchMovement[ThumbBase][1] = Open
	result.MotorSwitchMovement[ThumbBase][2] = Close

	// ThumbSpread
	result.MotorMovementSwitch[ThumbSpread] = make(map[Movement]int)
	result.MotorMovementSwitch[ThumbSpread][Open] = 1
	result.MotorMovementSwitch[ThumbSpread][Close] = 2

	result.MotorSwitchMovement[ThumbSpread] = make(map[int]Movement)
	result.MotorSwitchMovement[ThumbSpread][1] = Open
	result.MotorSwitchMovement[ThumbSpread][2] = Close

	// When mocking out the physical relays.
	result.Mock = mock
	if mock {
		result.MockRelayStatus = make(map[int64]map[int64]RelayStatus)
		for _, board := range BoardIds {
			result.MockRelayStatus[board] = make(map[int64]RelayStatus)
			for _, relay := range RelayIds {
				result.MockRelayStatus[board][relay] = RelayOff
			}
		}
	}

	return &result, nil
}

func (mc *MotionController) SetToNormal() {
	for m := range mc.RelativeFlexPosition {
		mc.RelativeFlexPosition[m] = AngleDegrees(0)
	}
}

func (mc *MotionController) setRelay(board, relay int64, status RelayStatus) error {
	if mc.Mock {
		mc.MockRelayStatus[board][relay] = status
		return nil
	} else {
		return UnsafeSetRelay(board, relay, status)
	}
}

func (mc *MotionController) getRelay(board, relay int64) (RelayStatus, error) {
	if mc.Mock {
		return mc.MockRelayStatus[board][relay], nil
	} else {
		return GetRelay(board, relay)
	}
}

func (mc *MotionController) MovementSwitch(motor Motor, movement Movement, bridge *HBridge) (on *Switch, off *Switch, err error) {
	if movementSwitch, ok1 := mc.MotorMovementSwitch[motor]; ok1 {
		if s, ok2 := movementSwitch[movement]; ok2 {
			if s == 1 {
				on = &bridge.S1
				off = &bridge.S2
			} else if s == 2 {
				on = &bridge.S2
				off = &bridge.S1
			}
		} else {
			err = errors.New("invalid movement")
		}
	} else {
		err = errors.New("invalid motor")
	}

	return
}

func (mc *MotionController) MakeHBridge(board int64, R1 int64, R2 int64) *HBridge {
	return &HBridge{
		S1: Switch{
			Board: board,
			Relay: R1,
		},
		S2: Switch{
			Board: board,
			Relay: R2,
		},
	}
}

func (mc *MotionController) ConnectMotor(motor Motor, hBridge *HBridge) {
	mc.MotorHBridge[motor] = hBridge
}

func (mc *MotionController) ConfigureMotor(motor Motor, minFlex, maxFlex, flexPerSecondEnergized float64) {
	mc.MotorParams[motor] = MotorParam{
		MinFlex:                minFlex,
		MaxFlex:                maxFlex,
		FlexPerSecondEnergized: flexPerSecondEnergized,
	}
}

func (mc *MotionController) ReversePoles(motor Motor) {
	mc.MotorHBridge[motor] = &HBridge{
		S1: mc.MotorHBridge[motor].S2,
		S2: mc.MotorHBridge[motor].S1,
	}
}

func (mc *MotionController) Flex(motor Motor, movement Movement, duration time.Duration) AngleDegrees {
	polarity := 0

	/* The rest are unimplemented
	Nil              Movement = "Nil"
	Open             Movement = "Open"
	Close            Movement = "Close"
	Left             Movement = "Left"
	Right            Movement = "Right"
	Up               Movement = "Up"
	Down             Movement = "Down"
	Clockwise        Movement = "Clockwise"
	CounterClockwise Movement = "CounterClockwise"
	*/

	if movement == Close {
		polarity = 1
	}

	if movement == Open {
		polarity = -1
	}

	return AngleDegrees(mc.MotorParams[motor].FlexPerSecondEnergized * duration.Seconds() * float64(polarity))
}

func (mc *MotionController) IsSafe(motor Motor, movement Movement, duration time.Duration) (isSafe bool, delta AngleDegrees, newPosition AngleDegrees) {

	log.Printf("IsSafe %s %s %+v\r\n", motor, movement, duration)

	delta = mc.Flex(motor, movement, duration)

	newPosition = mc.RelativeFlexPosition[motor] + delta

	log.Printf("IsSafe transtion %f -> %f range= [%f, %f]\r\n", mc.RelativeFlexPosition[motor], newPosition, mc.MotorParams[motor].MaxFlex, mc.MotorParams[motor].MinFlex)

	if float64(newPosition) < mc.MotorParams[motor].MaxFlex && float64(newPosition) > mc.MotorParams[motor].MinFlex {
		isSafe = true
	} else {
		isSafe = false
	}

	return
}

func (mc *MotionController) IsWithinBounds(motor Motor, newPosition AngleDegrees) bool {
	if float64(newPosition) < mc.MotorParams[motor].MaxFlex && float64(newPosition) > mc.MotorParams[motor].MinFlex {
		return true
	}

	return false
}

func (mc *MotionController) Commit(updated map[Motor]AngleDegrees) {
	for motor, flex := range updated {
		mc.RelativeFlexPosition[motor] = flex
	}
}

func (mc *MotionController) CheckUpdate(motor Motor, movement Movement, duration time.Duration) (isWithinBounds bool, updated map[Motor]AngleDegrees) {
	delta := mc.Flex(motor, movement, duration)

	log.Printf("CheckUpdate delta= %+v\r\n", delta)

	updated = make(map[Motor]AngleDegrees)

	// Update this motor
	updated[motor] = mc.RelativeFlexPosition[motor] + delta

	log.Printf("CheckUpdate delta= %+v, update= %+v \r\n", delta, updated[motor])

	isWithinBounds = mc.IsWithinBounds(motor, updated[motor])
	if !isWithinBounds {
		return
	}

	// Update side effects

	// Thumb
	if motor == ThumbMiddle {
		ToTipC := -1.0
		updated[ThumbTip] = mc.RelativeFlexPosition[ThumbTip] + delta*AngleDegrees(ToTipC)

		isWithinBounds = mc.IsWithinBounds(ThumbTip, updated[ThumbTip])
		if !isWithinBounds {
			return
		}
	}

	if motor == ThumbBase {
		ToMiddleC := -1.0
		updated[ThumbMiddle] = mc.RelativeFlexPosition[ThumbMiddle] + delta*AngleDegrees(ToMiddleC)

		isWithinBounds = mc.IsWithinBounds(ThumbMiddle, updated[ThumbMiddle])
		if !isWithinBounds {
			return
		}
	}

	if motor == ThumbSpread {
		ToTipC := -1.0
		updated[ThumbTip] = mc.RelativeFlexPosition[ThumbTip] + delta*AngleDegrees(ToTipC)

		isWithinBounds = mc.IsWithinBounds(ThumbTip, updated[ThumbTip])
		if !isWithinBounds {
			return
		}

		ToMiddleC := -1.0
		updated[ThumbMiddle] = mc.RelativeFlexPosition[ThumbMiddle] + delta*AngleDegrees(ToMiddleC)

		isWithinBounds = mc.IsWithinBounds(ThumbMiddle, updated[ThumbMiddle])
		if !isWithinBounds {
			return
		}

		ToBaseC := -1.0
		updated[ThumbBase] = mc.RelativeFlexPosition[ThumbBase] + delta*AngleDegrees(ToBaseC)

		isWithinBounds = mc.IsWithinBounds(ThumbBase, updated[ThumbBase])
		if !isWithinBounds {
			return
		}
	}

	// Ring
	if motor == RingMiddle {
		ToTipC := -1.0
		updated[RingTip] = mc.RelativeFlexPosition[RingTip] + delta*AngleDegrees(ToTipC)

		isWithinBounds = mc.IsWithinBounds(RingTip, updated[RingTip])
		if !isWithinBounds {
			return
		}
	}

	if motor == RingBase {
		ToMiddleC := -1.0
		updated[RingMiddle] = mc.RelativeFlexPosition[RingMiddle] + delta*AngleDegrees(ToMiddleC)

		isWithinBounds = mc.IsWithinBounds(RingMiddle, updated[RingMiddle])
		if !isWithinBounds {
			return
		}
	}

	if motor == RingSpread {
		ToTipC := -1.0
		updated[RingTip] = mc.RelativeFlexPosition[RingTip] + delta*AngleDegrees(ToTipC)

		isWithinBounds = mc.IsWithinBounds(RingTip, updated[RingTip])
		if !isWithinBounds {
			return
		}

		ToMiddleC := -1.0
		updated[RingMiddle] = mc.RelativeFlexPosition[RingMiddle] + delta*AngleDegrees(ToMiddleC)

		isWithinBounds = mc.IsWithinBounds(RingMiddle, updated[RingMiddle])
		if !isWithinBounds {
			return
		}

		ToBaseC := -1.0
		updated[RingBase] = mc.RelativeFlexPosition[RingBase] + delta*AngleDegrees(ToBaseC)

		isWithinBounds = mc.IsWithinBounds(RingBase, updated[RingBase])
		if !isWithinBounds {
			return
		}
	}

	//Middle
	if motor == MiddleMiddle {
		ToTipC := -1.0
		updated[MiddleTip] = mc.RelativeFlexPosition[MiddleTip] + delta*AngleDegrees(ToTipC)

		isWithinBounds = mc.IsWithinBounds(MiddleTip, updated[MiddleTip])
		if !isWithinBounds {
			return
		}
	}

	if motor == MiddleBase {
		ToMiddleC := -1.0
		updated[MiddleMiddle] = mc.RelativeFlexPosition[MiddleMiddle] + delta*AngleDegrees(ToMiddleC)

		isWithinBounds = mc.IsWithinBounds(MiddleMiddle, updated[MiddleMiddle])
		if !isWithinBounds {
			return
		}
	}

	//Index
	if motor == IndexMiddle {
		ToTipC := -1.0
		updated[IndexTip] = mc.RelativeFlexPosition[IndexTip] + delta*AngleDegrees(ToTipC)

		isWithinBounds = mc.IsWithinBounds(IndexTip, updated[IndexTip])
		if !isWithinBounds {
			return
		}
	}

	if motor == IndexBase {
		ToMiddleC := -1.0
		updated[IndexMiddle] = mc.RelativeFlexPosition[IndexMiddle] + delta*AngleDegrees(ToMiddleC)

		isWithinBounds = mc.IsWithinBounds(IndexMiddle, updated[IndexMiddle])
		if !isWithinBounds {
			return
		}
	}

	if motor == IndexSpread {
		ToTipC := -1.0
		updated[IndexTip] = mc.RelativeFlexPosition[IndexTip] + delta*AngleDegrees(ToTipC)

		isWithinBounds = mc.IsWithinBounds(IndexTip, updated[IndexTip])
		if !isWithinBounds {
			return
		}

		ToMiddleC := -1.0
		updated[IndexMiddle] = mc.RelativeFlexPosition[IndexMiddle] + delta*AngleDegrees(ToMiddleC)

		isWithinBounds = mc.IsWithinBounds(IndexMiddle, updated[IndexMiddle])
		if !isWithinBounds {
			return
		}

		ToBaseC := -1.0
		updated[IndexBase] = mc.RelativeFlexPosition[IndexBase] + delta*AngleDegrees(ToBaseC)

		isWithinBounds = mc.IsWithinBounds(IndexBase, updated[IndexBase])
		if !isWithinBounds {
			return
		}
	}

	return
}

func (mc *MotionController) mockSetOnContinue(motor Motor, movement Movement, duration time.Duration) error {
	isWithinBounds, update := mc.CheckUpdate(motor, movement, duration)

	if isWithinBounds {
		if _, ok := mc.MotorHBridge[motor]; ok {
			mc.Commit(update)
			return nil
		} else {
			return errors.New("MockSet: motor not connected")
		}
	} else {
		_ = mc.Off(motor)
		return errors.New(fmt.Sprintf("MockSet: unsafe motor movement. motor= %s movement= %s duration= %s. update= %+v", motor, movement, duration, update))
	}
}

func (mc *MotionController) Set(motor Motor, movement Movement, duration time.Duration) error {
	isWithinBounds, update := mc.CheckUpdate(motor, movement, duration)

	log.Printf("Set isWithinBounds= %+v, update=%+v\r\n", isWithinBounds, update)

	if isWithinBounds {
		if hb, ok := mc.MotorHBridge[motor]; ok {
			on, off, err := mc.MovementSwitch(motor, movement, hb)

			if err != nil {
				return err
			}

			err = mc.setRelay(on.Board, on.Relay, RelayOn)
			if err != nil {
				return err
			}

			err = mc.setRelay(off.Board, off.Relay, RelayOff)
			if err != nil {
				return err
			}

			mc.Commit(update)

			return nil
		} else {
			return errors.New("motor not connected")
		}
	} else {
		return errors.New(fmt.Sprintf("unsafe motor movement. motor= %s movement= %s duration= %s. update= %+v", motor, movement, duration, update))
	}
}

func (mc *MotionController) Get(motor Motor) (Movement, error) {
	if hb, ok := mc.MotorHBridge[motor]; ok {
		s1Status, err := mc.getRelay(hb.S1.Board, hb.S1.Relay)

		if err != nil {
			return Nil, err
		}

		s2Status, err := mc.getRelay(hb.S2.Board, hb.S2.Relay)

		if err != nil {
			return Nil, err
		}

		if s1Status == RelayOn && s2Status == RelayOff {
			return mc.MotorSwitchMovement[motor][1], nil
		} else if s2Status == RelayOn && s1Status == RelayOff {
			return mc.MotorSwitchMovement[motor][2], nil
		} else {
			return Nil, nil
		}
	}

	return Nil, errors.New("motor not connected")
}

func (mc *MotionController) Off(motor Motor) error {
	log.Printf("mc.Off(motor= %s)\r\n", motor)
	hb := mc.MotorHBridge[motor]

	if hb != nil {
		s1Status, err := mc.getRelay(hb.S1.Board, hb.S1.Relay)

		if err != nil {
			return err
		}

		s2Status, err := mc.getRelay(hb.S2.Board, hb.S2.Relay)

		if err != nil {
			return err
		}

		if s1Status == RelayOff && s2Status == RelayOff {
			return nil
		}

		// Although you can just toggle on the switches in a hbridge to turn it off,
		// the Sequent 8 relay board, relays draw less power from the raspberry pi, in NC position. So we will
		// reset the on relays to off.

		if s1Status == RelayOn {
			err = mc.setRelay(hb.S1.Board, hb.S1.Relay, RelayOff)

			if err != nil {
				return err
			}
		}

		if s2Status == RelayOn {
			err = mc.setRelay(hb.S2.Board, hb.S2.Relay, RelayOff)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (mc *MotionController) Step(motor Motor, movement Movement, duration time.Duration) error {
	defer func() {
		_ = mc.Off(motor)
	}()

	g := errgroup.Group{}

	g.Go(func() error {
		return mc.Set(motor, movement, duration)
	})

	time.Sleep(duration)

	return g.Wait()
}

func (mc *MotionController) GetAllMotorStatus() (map[Motor]Movement, error) {
	result := make(map[Motor]Movement)

	// Go over all the connected bridges, and get their statues
	for motor := range mc.MotorHBridge {
		m, err := mc.Get(motor)

		if err != nil {
			return nil, err
		}

		result[motor] = m
	}
	return result, nil
}

func (mc *MotionController) GetInternalFlex() map[Motor]AngleDegrees {
	return mc.RelativeFlexPosition
}

var multistepOverhead = 70 * time.Millisecond

func (mc *MotionController) MultiStep(next map[Motor]Movement, duration time.Duration) error {
	if duration < multistepOverhead {
		return errors.New("duration is too small. must allow more time that the overhead")
	}

	start := time.Now()

	c, err := mc.GetAllMotorStatus()

	if err != nil {
		return err
	}

	// Turn off all motors not specified in the next step
	for motor, movement := range c {
		if movement != Nil {
			if nextMovement, ok := next[motor]; ok && nextMovement != Nil && nextMovement == movement {
				continue
			} else {
				log.Printf("MuliStep mc.Off(motor= %s)\r\n", motor)
				err = mc.Off(motor)

				if err != nil {
					break
				}
			}
		}
	}

	if err != nil {
		return err
	}

	adjust := time.Since(start)

	if adjust > multistepOverhead {
		return errors.New(fmt.Sprintf("overhead buffer exceeded. excess= %+v", adjust-multistepOverhead))
	}

	time.Sleep(multistepOverhead - adjust)

	for motor, nextMovement := range next {
		// If the current motor statues is the same as the next step, then do nothing
		if currentMovement, ok := c[motor]; ok && currentMovement != Nil && nextMovement == currentMovement {
			log.Printf("MuliStep motor= %s state unchanged. no update to %s\r\n", motor, nextMovement)
			err = mc.mockSetOnContinue(motor, nextMovement, duration)

			if err != nil {
				break
			}
		} else {
			log.Printf("MuliStep mc.Off(motor= %s)\r\n", motor)
			err = mc.Set(motor, nextMovement, duration)

			if err != nil {
				break
			}

		}
	}

	if err != nil {
		_ = mc.StopAllMotors()
		return err
	}

	manualTune := 30 * time.Millisecond

	time.Sleep(duration - multistepOverhead - manualTune)

	return nil
}

func (mc *MotionController) StopAllMotors() error {
	log.Printf("mc.StopAllMotors()\r\n")
	for motor := range mc.MotorHBridge {
		_ = mc.Off(motor)
	}

	log.Printf("StopAllMotors \r\n")
	return nil
}
