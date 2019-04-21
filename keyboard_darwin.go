/*
   Copyright 2017, Yoshiki Shibukawa

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package gotomation

import (
	"time"
	"unicode/utf16"
	"unsafe"
)

/*
#cgo LDFLAGS: -framework Carbon -framework IOKit
#include <ApplicationServices/ApplicationServices.h>
#import <IOKit/hidsystem/IOHIDLib.h>
#import <IOKIt/hidsystem/IOLLEvent.h>
#include <Carbon/Carbon.h>
#include <mach/mach_traps.h>

static io_connect_t _getAuxiliaryKeyDriver(void)
{
	static mach_port_t sEventDrvrRef = 0;
	mach_port_t masterPort, service, iter;
	kern_return_t kr;

	if (!sEventDrvrRef) {
		kr = IOMasterPort( bootstrap_port, &masterPort );
		assert(KERN_SUCCESS == kr);
		kr = IOServiceGetMatchingServices(masterPort, IOServiceMatching( kIOHIDSystemClass), &iter );
		assert(KERN_SUCCESS == kr);
		service = IOIteratorNext( iter );
		assert(service);
		kr = IOServiceOpen(service, mach_task_self(), kIOHIDParamConnectType, &sEventDrvrRef );
		assert(KERN_SUCCESS == kr);
		IOObjectRelease(service);
		IOObjectRelease(iter);
	}
	return sEventDrvrRef;
}
*/
import "C"

const (
	VK_NOT_A_KEY KeyCode = 9999

	VK_A            = C.kVK_ANSI_A
	VK_S            = C.kVK_ANSI_S
	VK_D            = C.kVK_ANSI_D
	VK_F            = C.kVK_ANSI_F
	VK_H            = C.kVK_ANSI_H
	VK_G            = C.kVK_ANSI_G
	VK_Z            = C.kVK_ANSI_Z
	VK_X            = C.kVK_ANSI_X
	VK_C            = C.kVK_ANSI_C
	VK_V            = C.kVK_ANSI_V
	VK_B            = C.kVK_ANSI_B
	VK_Q            = C.kVK_ANSI_Q
	VK_W            = C.kVK_ANSI_W
	VK_E            = C.kVK_ANSI_E
	VK_R            = C.kVK_ANSI_R
	VK_Y            = C.kVK_ANSI_Y
	VK_T            = C.kVK_ANSI_T
	VK_1            = C.kVK_ANSI_1
	VK_2            = C.kVK_ANSI_2
	VK_3            = C.kVK_ANSI_3
	VK_4            = C.kVK_ANSI_4
	VK_6            = C.kVK_ANSI_6
	VK_5            = C.kVK_ANSI_5
	VK_EQUAL        = C.kVK_ANSI_Equal
	VK_9            = C.kVK_ANSI_9
	VK_7            = C.kVK_ANSI_7
	VK_MINUS        = C.kVK_ANSI_Minus
	VK_8            = C.kVK_ANSI_8
	VK_0            = C.kVK_ANSI_0
	VK_RIGHTBRACKET = C.kVK_ANSI_RightBracket
	VK_O            = C.kVK_ANSI_O
	VK_U            = C.kVK_ANSI_U
	VK_LEFTBRACKET  = C.kVK_ANSI_LeftBracket
	VK_I            = C.kVK_ANSI_I
	VK_P            = C.kVK_ANSI_P
	VK_L            = C.kVK_ANSI_L
	VK_J            = C.kVK_ANSI_J
	VK_QUOTE        = C.kVK_ANSI_Quote
	VK_K            = C.kVK_ANSI_K
	VK_SEMICOLON    = C.kVK_ANSI_Semicolon
	VK_BACKSLASH    = C.kVK_ANSI_Backslash
	VK_COMMA        = C.kVK_ANSI_Comma
	VK_SLASH        = C.kVK_ANSI_Slash
	VK_N            = C.kVK_ANSI_N
	VK_M            = C.kVK_ANSI_M
	VK_PERIOD       = C.kVK_ANSI_Period
	VK_GRAVE        = C.kVK_ANSI_Grave

	VK_BACKSPACE = C.kVK_Delete
	VK_DELETE    = C.kVK_ForwardDelete
	VK_RETURN    = C.kVK_Return
	VK_TAB       = C.kVK_Tab
	VK_ESCAPE    = C.kVK_Escape
	VK_UP        = C.kVK_UpArrow
	VK_DOWN      = C.kVK_DownArrow
	VK_RIGHT     = C.kVK_RightArrow
	VK_LEFT      = C.kVK_LeftArrow
	VK_HOME      = C.kVK_Home
	VK_END       = C.kVK_End
	VK_PAGEUP    = C.kVK_PageUp
	VK_PAGEDOWN  = C.kVK_PageDown
	VK_F1        = C.kVK_F1
	VK_F2        = C.kVK_F2
	VK_F3        = C.kVK_F3
	VK_F4        = C.kVK_F4
	VK_F5        = C.kVK_F5
	VK_F6        = C.kVK_F6
	VK_F7        = C.kVK_F7
	VK_F8        = C.kVK_F8
	VK_F9        = C.kVK_F9
	VK_F10       = C.kVK_F10
	VK_F11       = C.kVK_F11
	VK_F12       = C.kVK_F12
	VK_F13       = C.kVK_F13
	VK_F14       = C.kVK_F14
	VK_F15       = C.kVK_F15
	VK_F16       = C.kVK_F16
	VK_F17       = C.kVK_F17
	VK_F18       = C.kVK_F18
	VK_F19       = C.kVK_F19
	VK_F20       = C.kVK_F20
	VK_META      = C.kVK_Command
	VK_LMETA     = C.kVK_Command
	VK_RMETA     = C.kVK_Command
	VK_LWIN      = C.kVK_Command
	VK_RWIN      = C.kVK_Command
	VK_LCOMMAND  = C.kVK_Command
	VK_RCOMMAND  = C.kVK_Command
	VK_ALT       = C.kVK_Option
	VK_CONTROL   = C.kVK_Control
	VK_LCONTROL  = C.kVK_Control
	VK_RCONTROL  = C.kVK_Control
	VK_SHIFT     = C.kVK_Shift
	VK_LSHIFT    = C.kVK_Shift
	VK_RSHIFT    = C.kVK_RightShift
	VK_CAPSLOCK  = C.kVK_CapsLock
	VK_SPACE     = C.kVK_Space
	VK_INSERT    = VK_NOT_A_KEY
	VK_SNAPSHOT  = VK_NOT_A_KEY

	VK_NUMPAD_0       = C.kVK_ANSI_Keypad0
	VK_NUMPAD_1       = C.kVK_ANSI_Keypad1
	VK_NUMPAD_2       = C.kVK_ANSI_Keypad2
	VK_NUMPAD_3       = C.kVK_ANSI_Keypad3
	VK_NUMPAD_4       = C.kVK_ANSI_Keypad4
	VK_NUMPAD_5       = C.kVK_ANSI_Keypad5
	VK_NUMPAD_6       = C.kVK_ANSI_Keypad6
	VK_NUMPAD_7       = C.kVK_ANSI_Keypad7
	VK_NUMPAD_8       = C.kVK_ANSI_Keypad8
	VK_NUMPAD_9       = C.kVK_ANSI_Keypad9
	VK_NUMPAD_DECIMAL = C.kVK_ANSI_KeypadDecimal
	VK_NUMPAD_PLUS    = C.kVK_ANSI_KeypadPlus
	VK_NUMPAD_MINUS   = C.kVK_ANSI_KeypadMinus
	VK_NUMPAD_MUL     = C.kVK_ANSI_KeypadMultiply
	VK_NUMPAD_DIV     = C.kVK_ANSI_KeypadDivide
	VK_NUMPAD_CLEAR   = C.kVK_ANSI_KeypadClear
	VK_NUMPAD_ENTER   = C.kVK_ANSI_KeypadEnter
	VK_NUMPAD_EQUAL   = C.kVK_ANSI_KeypadEquals

	VK_AUDIO_VOLUME_MUTE = 1007
	VK_AUDIO_VOLUME_DOWN = 1001
	VK_AUDIO_VOLUME_UP   = 1000
	VK_MEDIA_PLAY_PAUSE  = 1016
	VK_AUDIO_STOP        = VK_NOT_A_KEY
	VK_AUDIO_PREV        = 1018
	VK_AUDIO_NEXT        = 1017

	VK_LIGHTS_MON_UP     = 1002
	VK_LIGHTS_MON_DOWN   = 1003
	VK_LIGHTS_KBD_TOGGLE = 1023
	VK_LIGHTS_KBD_UP     = 1021
	VK_LIGHTS_KBD_DOWN   = 1022

	VK_YEN          = C.kVK_JIS_Yen
	VK_UNDERSCORE   = C.kVK_JIS_Underscore
	VK_KEYPAD_COMMA = C.kVK_JIS_KeypadComma
	VK_EISU         = C.kVK_JIS_Eisu
	VK_KANA         = C.kVK_JIS_Kana
)

type Keyboard struct {
	waitBetweenChars time.Duration // delay
}

func (k *Keyboard) Type(str string) error {
	for _, char := range str {
		err := k.tap(char)
		if err != nil {
			return err
		}
		time.Sleep(k.waitBetweenChars)
	}
	return nil
}

func (k *Keyboard) TypeQuickly(str string) error {
	for _, char := range str {
		err := k.tap(char)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k *Keyboard) tap(char rune) error {
	err := k.toggleKeyByRune(char, true)
	if err != nil {
		return err
	}
	return k.toggleKeyByRune(char, false)
}

func (k *Keyboard) toggleKeyByRune(char rune, down bool) error {
	if char >= 0x10000 {
		r1, r2 := utf16.EncodeRune(char)
		chars := [2]C.UniChar{C.UniChar(r1), C.UniChar(r2)}

		event := C.CGEventCreateKeyboardEvent(nil, 0, C._Bool(down))
		defer C.CFRelease(C.CFTypeRef(event))
		C.CGEventKeyboardSetUnicodeString(event, 2, (*C.UniChar)(unsafe.Pointer(&chars[0])))
		C.CGEventPost(C.kCGSessionEventTap, event)
	} else {
		uniChar := C.UniChar(char)
		event := C.CGEventCreateKeyboardEvent(nil, 0, C._Bool(down))
		defer C.CFRelease(C.CFTypeRef(event))
		C.CGEventKeyboardSetUnicodeString(event, 1, (*C.UniChar)(unsafe.Pointer(&uniChar)))
		C.CGEventPost(C.kCGSessionEventTap, event)
	}
	return nil
}

type nxEventData struct { /* For window-changed, sys-defined, and app-defined events */
	reserved int16
	subType  int16      /* event subtype for compound events */
	L        [11]uint32 /* for use in compound events */
}

func (k *Keyboard) toggleKeyByCode(code KeyCode, down bool, modifiers []KeyModifier) error {
	if code > 1000 {
		code = code - 1000 /* Get the real keycode. */
		var loc C.IOGPoint
		evtInfo := uint32(code) << 16
		if down {
			evtInfo |= C.NX_KEYDOWN << 8
		} else {
			evtInfo |= C.NX_KEYUP << 8
		}
		event := nxEventData{
			subType: C.NX_SUBTYPE_AUX_CONTROL_BUTTONS,
		}
		event.L[0] = evtInfo
		kr := C.IOHIDPostEvent(C._getAuxiliaryKeyDriver(), C.NX_SYSDEFINED, loc, (*C.NXEventData)(unsafe.Pointer(&event)), C.kNXEventDataVersion, 0, 0)
		if C.KERN_SUCCESS != kr {
			panic("Carbon API error")
		}
	} else {
		event := C.CGEventCreateKeyboardEvent(nil, C.CGKeyCode(code), C._Bool(down))
		defer C.CFRelease(C.CFTypeRef(event))
		if down {
			C.CGEventSetType(event, C.kCGEventKeyDown)
		} else {
			C.CGEventSetType(event, C.kCGEventKeyUp)
		}
		C.CGEventPost(C.kCGSessionEventTap, event)
	}
	return nil
}
