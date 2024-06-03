package game

import input "github.com/quasilyte/ebitengine-input"

const (
	ActionMoveLeft input.Action = iota
	ActionMoveRight
	ActionMoveUp
	ActionMoveDown
	ActionConfirm
	ActionPrintState
)

var keymap = input.Keymap{
	ActionMoveLeft:   {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
	ActionMoveRight:  {input.KeyGamepadRight, input.KeyRight, input.KeyD},
	ActionMoveUp:     {input.KeyGamepadUp, input.KeyUp, input.KeyW},
	ActionMoveDown:   {input.KeyGamepadDown, input.KeyDown, input.KeyS},
	ActionConfirm:    {input.KeyGamepadA, input.KeyGamepadB, input.KeyEnter},
	ActionPrintState: {input.KeyPause, input.KeyGamepadB, input.KeyP},
}
