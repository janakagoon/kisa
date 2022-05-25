package main

import (
	"encoding/json"
	"errors"
	"github.com/pkg/term"
	"log"
	"motorcontroller/lib"
	"os"
	"time"
)

func ReadChar(t *term.Term) (c string, err error) {
	bytes := make([]byte, 3)

	count, err := t.Read(bytes)

	if count == 1 {
		// Hande special case for control c.
		if bytes[0] == 3 {
			c = "Control-C"
		} else if bytes[0] == 19 {
			c = "Control-S"
		} else if bytes[0] == 1 {
			c = "Control-A"
		} else if bytes[0] == 17 {
			c = "Control-Q"
		} else if bytes[0] == 23 {
			c = "Control-W"
		} else if bytes[0] == 5 {
			c = "Control-E"
		} else if bytes[0] == 26 {
			c = "Control-Z"
		} else if bytes[0] == 24 {
			c = "Control-X"
		} else if bytes[0] == 16 {
			c = "Control-P"
		} else {
			c = string(bytes[0:1])
		}

	} else {
		return "", errors.New("un-parsable character")
	}

	return
}

type Mode string

const (
	Interactive Mode = "interactive"
	Recording        = "recording"
)

var AllOff = make(map[lib.Motor]lib.Movement)

const stepDuration = 100 * time.Millisecond

func KeyToMotorMovement(c string) (lib.Motor, lib.Movement, error) {
	switch c {
	case "7":
		return lib.RingTip, lib.Open, nil
	case "&":
		return lib.RingTip, lib.Close, nil
	case "y":
		return lib.RingMiddle, lib.Open, nil
	case "Y":
		return lib.RingMiddle, lib.Close, nil
	case "g":
		return lib.RingBase, lib.Open, nil
	case "G":
		return lib.RingBase, lib.Close, nil
	case "v":
		return lib.RingSpread, lib.Open, nil
	case "V":
		return lib.RingSpread, lib.Close, nil

	case "8":
		return lib.MiddleTip, lib.Open, nil
	case "*":
		return lib.MiddleTip, lib.Close, nil
	case "u":
		return lib.MiddleMiddle, lib.Open, nil
	case "U":
		return lib.MiddleMiddle, lib.Close, nil
	case "h":
		return lib.MiddleBase, lib.Open, nil
	case "H":
		return lib.MiddleBase, lib.Close, nil

	case "9":
		return lib.IndexTip, lib.Open, nil
	case "(":
		return lib.IndexTip, lib.Close, nil
	case "i":
		return lib.IndexMiddle, lib.Open, nil
	case "I":
		return lib.IndexMiddle, lib.Close, nil
	case "j":
		return lib.IndexBase, lib.Open, nil
	case "J":
		return lib.IndexBase, lib.Close, nil
	case "n":
		return lib.IndexSpread, lib.Open, nil
	case "N":
		return lib.IndexSpread, lib.Close, nil

	case "0":
		return lib.ThumbTip, lib.Open, nil
	case ")":
		return lib.ThumbTip, lib.Close, nil
	case "o":
		return lib.ThumbMiddle, lib.Open, nil
	case "O":
		return lib.ThumbMiddle, lib.Close, nil
	case "k":
		return lib.ThumbBase, lib.Open, nil
	case "K":
		return lib.ThumbBase, lib.Close, nil
	case "m":
		return lib.ThumbSpread, lib.Open, nil
	case "M":
		return lib.ThumbSpread, lib.Close, nil

	case "left":
		return lib.WristLeftRight, lib.Left, nil
	case "right":
		return lib.WristLeftRight, lib.Right, nil
	case "up":
		return lib.WristUpDown, lib.Up, nil
	case "down":
		return lib.WristUpDown, lib.Down, nil
	}

	return lib.ArmRotate, lib.Clockwise, errors.New("unknown input")
}

func Opposite(m lib.Movement) lib.Movement {
	switch m {
	case lib.Nil:
		return lib.Nil
	case lib.Open:
		return lib.Close
	case lib.Close:
		return lib.Open
	case lib.Left:
		return lib.Right
	case lib.Right:
		return lib.Left
	case lib.Up:
		return lib.Down
	case lib.Down:
		return lib.Up
	case lib.Clockwise:
		return lib.CounterClockwise
	case lib.CounterClockwise:
		return lib.Clockwise
	default:
		return lib.Nil
	}
}

func D6(m lib.Movement) string {
	switch m {
	case lib.Nil:
		return "     "
	case lib.Open:
		return "OPEN "
	case lib.Close:
		return "CLOSE"
	default:
		return "ERROR"
	}
}

func PrintHand(mc *lib.MotionController) {
	state := mc.GetInternalFlex()
	pos, err := mc.GetAllMotorStatus()

	if err != nil {
		log.Printf("PrintHand error mc.GetAllMotorStatus(). error= %+v\r\n", err)
		return
	}

	log.Printf("\r\n")
	log.Printf("         RING    MID    INDEX   THUMB\r\n")

	log.Printf("       [%s] [%s] [%s] [%s]\r\n", D6(pos[lib.RingTip]), D6(pos[lib.MiddleTip]), D6(pos[lib.IndexTip]), D6(pos[lib.ThumbTip]))
	log.Printf("   TIP [%05.2f] [%05.2f] [%05.2f] [%05.2f]\r\n", state[lib.RingTip], state[lib.MiddleTip], state[lib.IndexTip], state[lib.ThumbTip])

	log.Printf("       [%s] [%s] [%s] [%s]\r\n", D6(pos[lib.RingMiddle]), D6(pos[lib.MiddleMiddle]), D6(pos[lib.IndexMiddle]), D6(pos[lib.ThumbMiddle]))
	log.Printf("MIDDLE [%05.2f] [%05.2f] [%05.2f] [%05.2f]\r\n", state[lib.RingMiddle], state[lib.MiddleMiddle], state[lib.IndexMiddle], state[lib.ThumbMiddle])

	log.Printf("       [%s] [%s] [%s] [%s]\r\n", D6(pos[lib.RingBase]), D6(pos[lib.MiddleBase]), D6(pos[lib.IndexBase]), D6(pos[lib.ThumbBase]))
	log.Printf("  BASE [%05.2f] [%05.2f] [%05.2f] [%05.2f]\r\n", state[lib.RingBase], state[lib.MiddleBase], state[lib.IndexBase], state[lib.ThumbBase])

	log.Printf("       [%s] [     ] [%s] [%s]\r\n", D6(pos[lib.RingSpread]), D6(pos[lib.IndexSpread]), D6(pos[lib.ThumbSpread]))
	log.Printf("   ROT [%05.2f] [     ] [%05.2f] [%05.2f]\r\n", state[lib.RingSpread], state[lib.IndexSpread], state[lib.ThumbSpread])
	log.Printf("\r\n")
}

func main() {
	file := "/tmp/saved.txt"

	mc, err := lib.NewMotionController(true)

	if err != nil {
		log.Fatalf("error creating motor controller. error = %+v\n", err)
	}

	// WARNING: If physically connected, then these must be set.
	// There is a bug in the EC2 bridging software, where sometimes relays will
	// spontaneously come on. This software will detect such spurious ONs and
	// turn off the relay, but only if the motor configuration is available via the
	// configuration below. Else the hand can get damaged.

	mc.ConnectMotor(lib.RingTip, mc.MakeHBridge(4, 1, 2))
	mc.ConnectMotor(lib.RingMiddle, mc.MakeHBridge(3, 7, 8))
	mc.ConnectMotor(lib.RingBase, mc.MakeHBridge(1, 3, 4))
	mc.ReversePoles(lib.RingBase)
	mc.ConnectMotor(lib.RingSpread, mc.MakeHBridge(0, 3, 4))
	mc.ReversePoles(lib.RingSpread)

	mc.ConnectMotor(lib.MiddleTip, mc.MakeHBridge(3, 3, 4))
	mc.ReversePoles(lib.MiddleTip)
	mc.ConnectMotor(lib.MiddleMiddle, mc.MakeHBridge(1, 1, 2))
	mc.ConnectMotor(lib.MiddleBase, mc.MakeHBridge(2, 1, 2))

	mc.ConnectMotor(lib.IndexTip, mc.MakeHBridge(3, 5, 6))
	mc.ReversePoles(lib.IndexTip)
	mc.ConnectMotor(lib.IndexMiddle, mc.MakeHBridge(0, 5, 6))
	mc.ConnectMotor(lib.IndexBase, mc.MakeHBridge(3, 1, 2))
	mc.ConnectMotor(lib.IndexSpread, mc.MakeHBridge(2, 3, 4))

	mc.ConnectMotor(lib.ThumbTip, mc.MakeHBridge(1, 7, 8))
	mc.ConnectMotor(lib.ThumbMiddle, mc.MakeHBridge(0, 7, 8))
	mc.ConnectMotor(lib.ThumbBase, mc.MakeHBridge(1, 5, 6))
	mc.ConnectMotor(lib.ThumbSpread, mc.MakeHBridge(2, 5, 6))

	/* MOTOR CONFIGURATION */

	smallCoupler := 20.0
	mediumCoupler := 30.0
	largeCoupler := 35.0

	spreaderCouplerMultiplier := 4.8

	mc.ConfigureMotor(lib.RingTip, -90, 90, smallCoupler)
	mc.ConfigureMotor(lib.RingMiddle, -90, 90, smallCoupler)
	mc.ConfigureMotor(lib.RingBase, -45, 90, mediumCoupler)
	mc.ConfigureMotor(lib.RingSpread, -20, 5, smallCoupler/spreaderCouplerMultiplier)

	mc.ConfigureMotor(lib.MiddleTip, -90, 90, smallCoupler)
	mc.ConfigureMotor(lib.MiddleMiddle, -90, 90, mediumCoupler)
	mc.ConfigureMotor(lib.MiddleBase, -90, 90, largeCoupler)

	mc.ConfigureMotor(lib.IndexTip, -90, 90, smallCoupler)
	mc.ConfigureMotor(lib.IndexMiddle, -90, 90, smallCoupler)
	mc.ConfigureMotor(lib.IndexBase, -45, 90, mediumCoupler)
	mc.ConfigureMotor(lib.IndexSpread, -20, 5, smallCoupler/spreaderCouplerMultiplier)

	mc.ConfigureMotor(lib.ThumbTip, -90, 90, mediumCoupler)
	mc.ConfigureMotor(lib.ThumbMiddle, -90, 90, largeCoupler)
	mc.ConfigureMotor(lib.ThumbBase, -45, 90, smallCoupler)
	mc.ConfigureMotor(lib.ThumbSpread, -20, 20, largeCoupler/spreaderCouplerMultiplier)

	t, err := term.Open("/dev/tty")

	if err != nil {
		log.Fatalf("error creating term. error = %+v\n", err)
	}

	err = term.RawMode(t)

	if err != nil {
		log.Fatalf("error setting term to raw mode. error = %+v\n", err)
	}

	defer func() {
		_ = t.Restore()
		_ = t.Close()
		_ = mc.StopAllMotors()
	}()

	mode := Interactive

	currentStep := make(map[lib.Motor]lib.Movement)
	accumulator := make([]map[lib.Motor]lib.Movement, 0)

	log.Printf("controller started. mode= %s\n\r", mode)

	for {
		c, err := ReadChar(t)

		if err != nil {
			log.Printf("error reading input. err= %s\n\r", err.Error())
			continue
		}

		if c == "?" {
			log.Printf("Commands:\r\n")
			log.Printf("Control-C: quit\r\n")
			log.Printf("<space>: print the current hand configuration\r\n")
			log.Printf("Control-S: enter recording mode OR save previous step and start new recording\r\n")
			log.Printf("Control-A: exit recording mode, saving step to accumulator.\r\n")
			log.Printf("Control-Q: playback accumulator.\r\n")
			log.Printf("Control-W: rewind accumulator.\r\n")
			log.Printf("Control-E: clear accumulator.\r\n")
			log.Printf("Control-Z: load accumulator.\r\n")
			log.Printf("Control-X: save accumulator.\r\n")
			log.Printf("Control-P: print accumulator.\r\n")
			log.Printf("\";\": set hand to normal position (WARNING, incorrect use will dmanage hand).\r\n")
			log.Printf("\":\": insert an empty step.\r\n")
			continue
		}

		if c == "Control-P" {
			accumulatorJson, err := json.MarshalIndent(accumulator, "\r", " ")

			if err != nil {
				log.Printf("error json.Unmarshal(accumulatorJson, &accumulator). error= %+v\r\n", err)
				break
			}

			log.Printf("accumulator= %s\r\n", accumulatorJson)
			continue
		}

		if c == "Control-C" {
			log.Printf("received termination key. bye!\r\n")
			break
		}

		if c == "Control-S" {
			if mode == Interactive {

				mode = Recording
				currentStep = make(map[lib.Motor]lib.Movement)

				log.Printf("entered recording mode.\r\n")
			} else if mode == Recording {
				if len(currentStep) > 0 {
					// record the current step
					accumulator = append(accumulator, currentStep)

					// clear the current step
					currentStep = make(map[lib.Motor]lib.Movement)

					log.Printf("step saved. new recording started.\r\n")
				} else {
					log.Printf("nothing in current step. noop.\r\n")
				}
			} else {
				log.Printf("no-op. already recording mode. press Control-A to exit recoridng mode.\r\n")
			}

			continue
		}

		if c == "Control-A" {
			if mode == Recording {

				// record the current step
				accumulator = append(accumulator, currentStep)

				// clear the current step
				currentStep = make(map[lib.Motor]lib.Movement)

				log.Printf("step saved. entered interactive mode.\r\n")
				mode = Interactive

			} else {
				log.Printf("no-op. not in recording mode. press Control-S to enter recording mode.\r\n")
			}

			continue
		}

		if c == "Control-Q" {
			// Playback accumulator
			for i, s := range accumulator {
				log.Printf("playing back step= %d.\r\n", i)
				err = mc.MultiStep(s, stepDuration)

				if err != nil {
					log.Printf("error during playback of accumulator. err= %+v\r\n", err)
					break
				}

				// Clear screen, in order to animate
				log.Printf("\033c\r\n")
				PrintHand(mc)
			}

			err = mc.StopAllMotors()

			continue
		}

		if c == "Control-W" {
			// Playback accumulator in reverse
			steps := len(accumulator)

			for i := steps - 1; i >= 0; i-- {
				step := accumulator[i]
				reversedStep := make(map[lib.Motor]lib.Movement)

				for m, s := range step {
					reversedStep[m] = Opposite(s)
				}

				log.Printf("playing back inverse of step= %d.\r\n", i)
				err = mc.MultiStep(reversedStep, stepDuration)

				// Clear screen, in order to animate
				log.Printf("\033c\r\n")
				PrintHand(mc)
			}

			err = mc.StopAllMotors()

			continue
		}

		if c == "Control-E" {
			currentStep = make(map[lib.Motor]lib.Movement)
			accumulator = make([]map[lib.Motor]lib.Movement, 0)

			mode = Interactive
			log.Printf("accumulator is cleared. mode is interactive.\r\n")

			continue
		}

		if c == "Control-Z" {
			log.Printf("loading accumulator.\r\n")

			accumulatorJson, err := os.ReadFile(file)

			if err != nil {
				log.Printf("error os.ReadFile. error= %+v\r\n", err)
				break
			}

			err = json.Unmarshal(accumulatorJson, &accumulator)
			if err != nil {
				log.Printf("error json.Unmarshal(accumulatorJson, &accumulator). error= %+v\r\n", err)
				break
			}

			log.Printf("loaded accumulator.\r\n")
			continue
		}

		if c == "Control-X" {
			log.Printf("saving accumulator.\r\n")

			accumulatorJson, err := json.Marshal(accumulator)

			if err != nil {
				log.Printf("error json.Marshal. error= %+v\r\n", err)
				break
			}

			err = os.WriteFile(file, accumulatorJson, 0644)

			if err != nil {
				log.Printf("error os.WriteFile(file, accumulatorJson, 0). error= %+v\r\n", err)
				break
			}

			log.Printf("saved accumulator.\r\n")

			continue
		}

		if c == ";" {
			log.Printf("adding empty step.\r\n")
			// ; key means pause all motors in the next step.
			if mode == Recording {
				accumulator = append(accumulator, currentStep)
			}

			accumulator = append(accumulator, AllOff)

			currentStep = make(map[lib.Motor]lib.Movement)
			continue
		}

		if c == ":" {
			log.Printf("WARNING! WARNING! WARNING! incorrect use can damange hand. setting internal state to normal.\r\n")

			currentStep = make(map[lib.Motor]lib.Movement)
			accumulator = make([]map[lib.Motor]lib.Movement, 0)

			log.Printf("accumulator is cleared.\r\n")

			mc.SetToNormal()
			continue
		}

		if c == " " {
			PrintHand(mc)
			continue
		}

		motor, movement, err := KeyToMotorMovement(c)

		if err != nil {
			log.Printf("KeyToMotorMovement error= %+v\r\n", err)
			continue
		}

		if mode == Recording {
			if _, ok := currentStep[motor]; ok {
				log.Printf("already motion is set for this motor. assuming exit to interactive mode.")
				// record the current step
				accumulator = append(accumulator, currentStep)

				log.Printf("step saved. entered interactive mode.\r\n")
				// clear the current step
				currentStep = make(map[lib.Motor]lib.Movement)
				currentStep[motor] = movement
				mode = Interactive
			} else {
				currentStep[motor] = movement
			}
		}

		log.Printf("Calling mc.Step(motor=%s, movement=%s)\r\n", motor, movement)
		err = mc.Step(motor, movement, stepDuration)

		if err != nil {
			log.Printf("error stepping %v, %v. error= %+v\r\n", motor, movement, err)
			continue
		}

		if mode == Interactive {
			currentStep[motor] = movement

			accumulator = append(accumulator, currentStep)
			currentStep = make(map[lib.Motor]lib.Movement)
		}

		PrintHand(mc)

	}
}
