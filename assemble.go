package evmasm

import (
	"github.com/reiver/go-evmop"

	"bytes"
	"fmt"
	"io"
	"strings"
)

func Assemble(writer io.Writer, reader io.Reader) error {
	if nil == writer {
		return errNilWriter
	}
	if nil == reader {
		return errNilReader
	}


	var tokenizer tokenize

	if err := tokenizer.SetReader(reader); nil != err {
		return err
	}

	for {
		var token string
		{
			var buffer bytes.Buffer

			err := tokenizer.Token(&buffer)
			if io.EOF == err {
				if 0 < buffer.Len() {
					return errInternalError
				} else {
					break
				}
			}

			token = strings.ToUpper(buffer.String())
		}

		switch token {
		case "STOP":
			_, err := evmop.Stop{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "ADD":
			_, err := evmop.Add{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "MUL":
			_, err := evmop.Mul{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SUB":
			_, err := evmop.Sub{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DIV":
			_, err := evmop.Div{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SDIV":
			_, err := evmop.SDiv{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "MOD":
			_, err := evmop.Mod{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SMOD":
			_, err := evmop.SMod{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "ADDMOD":
			_, err := evmop.AddMod{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "MULMOD":
			_, err := evmop.MulMod{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "EXP":
			_, err := evmop.Exp{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SIGNEXTEND":
			_, err := evmop.SignExtend{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "LT":
			_, err := evmop.LT{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "GT":
			_, err := evmop.GT{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SLT":
			_, err := evmop.SLT{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SGT":
			_, err := evmop.SGT{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "EQ":
			_, err := evmop.Eq{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "ISZERO":
			_, err := evmop.IsZero{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "AND":
			_, err := evmop.And{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "OR":
			_, err := evmop.Or{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "XOR":
			_, err := evmop.Xor{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "NOT":
			_, err := evmop.Not{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "BYTE":
			_, err := evmop.Byte{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SHA3":
			_, err := evmop.SHA3{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "ADDRESS":
			_, err := evmop.Address{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "BALANCE":
			_, err := evmop.Balance{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "ORIGIN":
			_, err := evmop.Origin{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "CALLER":
			_, err := evmop.Caller{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "CALLVALUE":
			_, err := evmop.CallValue{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "CALLDATALOAD":
			_, err := evmop.CallDataLoad{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "CALLDATASIZE":
			_, err := evmop.CallDataSize{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "CALLDATACOPY":
			_, err := evmop.CallDataCopy{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "CODESIZE":
			_, err := evmop.CodeSize{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "CODECOPY":
			_, err := evmop.CodeCopy{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "GASPRICE":
			_, err := evmop.GasPrice{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "EXTCODESIZE":
			_, err := evmop.ExtCodeSize{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "EXTCODECOPY":
			_, err := evmop.ExtCodeCopy{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "BLOCKHASH":
			_, err := evmop.Blockhash{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "COINBASE":
			_, err := evmop.Coinbase{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "TIMESTAMP":
			_, err := evmop.Timestamp{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "NUMBER":
			_, err := evmop.Number{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DIFFICULTY":
			_, err := evmop.Difficulty{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "GASLIMIT":
			_, err := evmop.GasLimit{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "POP":
			_, err := evmop.Pop{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "MLOAD":
			_, err := evmop.MLoad{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "MSTORE":
			_, err := evmop.MStore{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "MSTORE8":
			_, err := evmop.MStore8{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SLOAD":
			_, err := evmop.SLoad{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SSTORE":
			_, err := evmop.SStore{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "JUMP":
			_, err := evmop.Jump{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "JUMPI":
			_, err := evmop.JumpI{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "PC":
			_, err := evmop.PC{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "MSIZE":
			_, err := evmop.MSize{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "GAS":
			_, err := evmop.Gas{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "JUMPDEST":
			_, err := evmop.JumpDest{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "PUSH1",
		     "PUSH2",
		     "PUSH3",
		     "PUSH4",
		     "PUSH5",
		     "PUSH6",
		     "PUSH7",
		     "PUSH8",
		     "PUSH9",
		     "PUSH10",
		     "PUSH11",
		     "PUSH12",
		     "PUSH13",
		     "PUSH14",
		     "PUSH15",
		     "PUSH16",
		     "PUSH17",
		     "PUSH18",
		     "PUSH19",
		     "PUSH20",
		     "PUSH21",
		     "PUSH22",
		     "PUSH23",
		     "PUSH24",
		     "PUSH25",
		     "PUSH26",
		     "PUSH27",
		     "PUSH28",
		     "PUSH29",
		     "PUSH30",
		     "PUSH31",
		     "PUSH32":

			var limit int

			switch token {
			case "PUSH1":
				limit = 1
			case "PUSH2":
				limit = 2
			case "PUSH3":
				limit = 3
			case "PUSH4":
				limit = 4
			case "PUSH5":
				limit = 5
			case "PUSH6":
				limit = 6
			case "PUSH7":
				limit = 7
			case "PUSH8":
				limit = 8
			case "PUSH9":
				limit = 9
			case "PUSH10":
				limit = 10
			case "PUSH11":
				limit = 11
			case "PUSH12":
				limit = 12
			case "PUSH13":
				limit = 13
			case "PUSH14":
				limit = 14
			case "PUSH15":
				limit = 15
			case "PUSH16":
				limit = 16
			case "PUSH17":
				limit = 17
			case "PUSH18":
				limit = 18
			case "PUSH19":
				limit = 19
			case "PUSH20":
				limit = 20
			case "PUSH21":
				limit = 21
			case "PUSH22":
				limit = 22
			case "PUSH23":
				limit = 23
			case "PUSH24":
				limit = 24
			case "PUSH25":
				limit = 25
			case "PUSH26":
				limit = 26
			case "PUSH27":
				limit = 27
			case "PUSH28":
				limit = 28
			case "PUSH29":
				limit = 29
			case "PUSH30":
				limit = 30
			case "PUSH31":
				limit = 31
			case "PUSH32":
				limit = 32
			default:
				return errInternalError
			}

			var b [32]byte

			for i:=0; i<limit; i++ {
				var buffer bytes.Buffer

				err := tokenizer.Token(&buffer)
				if io.EOF == err {
					if 0 < buffer.Len() {
						return errInternalError
					} else {
						return fmt.Errorf("Missing number #%d after %q", 1+i, token)
					}
				}

				n, err := numeral(buffer.Bytes())
				if nil != err {
					return err
				}

				b[i] = n
			}

			switch token {
			case "PUSH1":
				_, err := evmop.Push1{ b[0] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH2":
				_, err := evmop.Push2{ b[0], b[1] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH3":
				_, err := evmop.Push3{ b[0], b[1], b[2] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH4":
				_, err := evmop.Push4{ b[0], b[1], b[2], b[3] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH5":
				_, err := evmop.Push5{ b[0], b[1], b[2], b[3], b[4] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH6":
				_, err := evmop.Push6{ b[0], b[1], b[2], b[3], b[4], b[5] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH7":
				_, err := evmop.Push7{ b[0], b[1], b[2], b[3], b[4], b[5], b[6] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH8":
				_, err := evmop.Push8{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH9":
				_, err := evmop.Push9{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH10":
				_, err := evmop.Push10{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH11":
				_, err := evmop.Push11{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH12":
				_, err := evmop.Push12{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH13":
				_, err := evmop.Push13{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH14":
				_, err := evmop.Push14{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH15":
				_, err := evmop.Push15{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH16":
				_, err := evmop.Push16{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH17":
				_, err := evmop.Push17{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH18":
				_, err := evmop.Push18{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH19":
				_, err := evmop.Push19{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH20":
				_, err := evmop.Push20{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH21":
				_, err := evmop.Push21{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH22":
				_, err := evmop.Push22{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH23":
				_, err := evmop.Push23{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21], b[22] }.WriteTo(writer)
				if nil != err {
					return err
				}


			case "PUSH24":
				_, err := evmop.Push24{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH25":
				_, err := evmop.Push25{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23], b[24] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH26":
				_, err := evmop.Push26{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23], b[24], b[25] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH27":
				_, err := evmop.Push27{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23], b[24], b[25], b[26] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH28":
				_, err := evmop.Push28{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23], b[24], b[25], b[26], b[27] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH29":
				_, err := evmop.Push29{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23], b[24], b[25], b[26], b[27], b[28] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH30":
				_, err := evmop.Push30{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23], b[24], b[25], b[26], b[27], b[28], b[29] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH31":
				_, err := evmop.Push31{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23], b[24], b[25], b[26], b[27], b[28], b[29], b[30] }.WriteTo(writer)
				if nil != err {
					return err
				}

			case "PUSH32":
				_, err := evmop.Push32{ b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19], b[20], b[21], b[22], b[23], b[24], b[25], b[26], b[27], b[28], b[29], b[30], b[31] }.WriteTo(writer)
				if nil != err {
					return err
				}

			default:
				return errInternalError
			}

		case "DUP1":
			_, err := evmop.Dup1{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP2":
			_, err := evmop.Dup2{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP3":
			_, err := evmop.Dup3{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP4":
			_, err := evmop.Dup4{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP5":
			_, err := evmop.Dup5{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP6":
			_, err := evmop.Dup6{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP7":
			_, err := evmop.Dup7{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP8":
			_, err := evmop.Dup8{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP9":
			_, err := evmop.Dup9{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP10":
			_, err := evmop.Dup10{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP11":
			_, err := evmop.Dup11{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP12":
			_, err := evmop.Dup12{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP13":
			_, err := evmop.Dup13{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP14":
			_, err := evmop.Dup14{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP15":
			_, err := evmop.Dup15{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DUP16":
			_, err := evmop.Dup16{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP1":
			_, err := evmop.Swap1{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP2":
			_, err := evmop.Swap2{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP3":
			_, err := evmop.Swap3{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP4":
			_, err := evmop.Swap4{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP5":
			_, err := evmop.Swap5{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP6":
			_, err := evmop.Swap6{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP7":
			_, err := evmop.Swap7{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP8":
			_, err := evmop.Swap8{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP9":
			_, err := evmop.Swap9{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP10":
			_, err := evmop.Swap10{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP11":
			_, err := evmop.Swap11{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP12":
			_, err := evmop.Swap12{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP13":
			_, err := evmop.Swap13{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP14":
			_, err := evmop.Swap14{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP15":
			_, err := evmop.Swap15{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SWAP16":
			_, err := evmop.Swap16{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "LOG0":
			_, err := evmop.Log0{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "LOG1":
			_, err := evmop.Log1{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "LOG2":
			_, err := evmop.Log2{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "LOG3":
			_, err := evmop.Log3{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "LOG4":
			_, err := evmop.Log4{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "CREATE":
			_, err := evmop.Create{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "CALL":
			_, err := evmop.Call{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "CALLCODE":
			_, err := evmop.CallCode{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "RETURN":
			_, err := evmop.Return{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "DELEGATECALL":
			_, err := evmop.DelegateCall{}.WriteTo(writer)
			if nil != err {
				return err
			}

		case "SUICIDE":
			_, err := evmop.Suicide{}.WriteTo(writer)
			if nil != err {
				return err
			}
		}
	}

	return nil
}
