package lexer

// func (l *Lexer) execTokenize() *token.Token {
//
// 	if l.ch == ")" {
// 		if l.currentDepth == 0 {
// 			l.currentState = lexerStateNormal
// 			return l.token(token.EEND, ")")
// 		} else {
// 			l.currentDepth--
// 			return l.token(token.EMIDDLE, ")")
// 		}
// 	}
//
// 	if l.ch == "(" {
// 		l.currentDepth++
// 		return l.token(token.EMIDDLE, "(")
// 	}
//
//     if l.ch == '{'
// }
