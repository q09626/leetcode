package main

type State int
type CharType int

const (
	StateInitial         State = iota // 起始空格
	StateIntSign                      // 符号位
	StateInteger                      // 整数部分
	StatePoint                        // 左侧有整数的小数点
	StatePointWithoutInt              // 左侧无整数的小数点
	StateFraction                     // 小数部分
	StateExp                          // 字符e
	StateExpSign                      // 指数部分的符号位
	StateExpNumber                    // 指数部分的整合部分
	StateEnd                          // 末尾的空格
)

const (
	CharNumber CharType = iota
	CharExp
	CharPoint
	CharSign
	CharSpace
	CharIllegal
)

func toCharType(c byte) CharType {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharNumber
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case '+', '-':
		return CharSign
	case ' ':
		return CharSpace
	default:
		return CharIllegal
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		StateInitial: map[CharType]State{
			CharSpace:  StateInitial,
			CharSign:   StateIntSign,
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
		},
		StateIntSign: map[CharType]State{
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
		},
		StateInteger: map[CharType]State{
			CharNumber: StateInteger,
			CharExp:    StateExp,
			CharPoint:  StatePoint,
			CharSpace:  StateEnd,
		},
		StatePoint: map[CharType]State{
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StatePointWithoutInt: map[CharType]State{
			CharNumber: StateFraction,
		},
		StateFraction: map[CharType]State{
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StateExp: map[CharType]State{
			CharNumber: StateExpNumber,
			CharSign:   StateExpSign,
		},
		StateExpSign: map[CharType]State{
			CharNumber: StateExpNumber,
		},
		StateExpNumber: map[CharType]State{
			CharNumber: StateExpNumber,
			CharSpace:  StateEnd,
		},
		StateEnd: map[CharType]State{
			CharSpace: StateEnd,
		},
	}
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		}
		state = transfer[state][typ]
	}
	return state == StateInteger || state == StatePoint || state == StateFraction || state == StateExpNumber || state == StateEnd
}
