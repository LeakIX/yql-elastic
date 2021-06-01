package yql_elastic

import (
	"errors"
	"log"
)
// Default state
func lexText(lexer *Lexer) stateFn {
	for {
		// Skip whitespaces
		lexer.skipWhiteSpaces()
		// Check for conditions modifiers
		rune := lexer.next()
		if rune == eof {
			return nil
		}
		// ignore the next char if escaped
		if rune == escapeChar {
			lexer.next()
			continue
		}
		// Group
		if rune == leftParenthesis  {
			return lexQueryGroup(lexer)
		}
		// Must
		if rune == plusSign {
			lexer.nextCondition = queryMust
			lexer.Advance(itemMust)
			continue
		}
		// MustNot
		if rune == minusSign  {
			lexer.nextCondition = queryMustNot
			lexer.Advance(itemMustNot)
			continue
		}
		// Nothing, must be term
		lexer.backup()
		return lexTerm(lexer)
	}
}

func (lexer *Lexer) parseFieldType() {
	nextRune := lexer.peek()
	switch nextRune {
	case equalSign:
		lexer.matchType = keywordMatch
		lexer.next()
		lexer.Advance(itemKeyword)
	case tildeSign:
		lexer.matchType = regexMatch
		lexer.next()
		lexer.Advance(itemRegex)
	case greaterThan:
		lexer.matchType = upperMatch
		lexer.next()
		lexer.Advance(itemGreaterThan)
	case lowerThan:
		lexer.matchType = lowerMatch
		lexer.next()
		lexer.Advance(itemLowerThan)
	case leftParenthesis:
		lexer.matchType = simpleQueryMatch
		lexer.Advance(itemLeftGroupDelim)
		lexer.inQuote = rightParenthesis
	}
}

// a term can be a field+value a value or a phrase ( between " )
func lexTerm(lexer *Lexer) stateFn {
	// Quote state
	for {
		rune := lexer.next()
		// ignore the next char if escaped
		if rune == escapeChar {
			lexer.next()
			continue
		}
		// WARNING : we can look for rightParenthesis as closing because of parseFieldType() when in field
		if rune == doubleQuote ||  rune == singleQuote || rune == rightParenthesis  {
			if lexer.inQuote != noQuote && lexer.inQuote == rune{
				lexer.inQuote = noQuote
			} else if lexer.start == lexer.pos{
				lexer.inQuote = rune
			}
		}
		if rune == semiColon && lexer.pos != 0 && lexer.inQuote == noQuote {
			lexer.backup()
			lexer.field = lexer.value()
			lexer.next()
			lexer.Advance(itemField)
			lexer.parseFieldType()
		}
		if rune == eof || (rune == whiteSpace && lexer.inQuote == noQuote ) {
			lexer.commitQuery()
			return lexText(lexer)
		}
	}
}

//Create new lexer with inside () contents to merge to this lexer
func lexQueryGroup(lexer *Lexer) stateFn {
	lexer.Advance(itemLeftGroupDelim)
	openDepth := 1
	for {
		rune := lexer.next()
		log.Printf("at %s", string(rune))
		if rune == eof {
			lexer.lastError = errors.New("unclosed " + string(leftParenthesis))
			return nil
		}
		// ignore the next char if escaped
		if rune == escapeChar {
			lexer.next()
			continue
		}
		if rune == doubleQuote ||  rune == singleQuote {
			if lexer.inQuote != noQuote && lexer.inQuote == rune{
				lexer.inQuote = noQuote
			} else {
				lexer.inQuote = rune
			}
		}
		if rune == leftParenthesis && lexer.inQuote == noQuote {
			log.Printf("opendepth++ %d", openDepth)
			openDepth++
			continue
		}
		if rune == rightParenthesis && lexer.inQuote == noQuote {
			log.Printf("opendepth-- %d", openDepth)
			openDepth--
			if openDepth > 0 {
				continue
			}
			lexer.backup()
			groupQuery := lexer.input[lexer.start:lexer.pos]
			subQuery, err := Parse(groupQuery, lexer.options...)
			if err != nil {
				lexer.lastError = err
				return nil
			}
			lexer.addQuery(subQuery)
			lexer.next()
			lexer.Advance(itemRightGroupDelim)
			return lexText(lexer)
		}
	}
	return nil
}
