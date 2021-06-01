package yql_elastic

import (
	"github.com/olivere/elastic/v7"
	"strings"
	"unicode/utf8"
)

type Lexer struct {
	input string
	// Starting position of the currently parsed operation
	start int
	// Position of the currently parsed operation
	pos   int
	// Length of the next char in bytes ( UFT-8 )
	width int
	// Accumulated elasticsearch Query storing conditions
	query *elastic.BoolQuery
	// next condition
	nextCondition conditionType
	// field we're in if any
	field string
	// field match type
	matchType     matchType
	// Default fields for no-field terms
	defaultFields []string
	// Last error
	lastError     error
	// Lexer/Parser options
	options       []ParserOption
	// Nested path
	nestedPaths   []string
	// Static alias
	mappings      map[string]string
	// Is set to the next closing delimiter : noQuote, doubleQuote, singleQuote or rightParenthesis
	inQuote       rune
}

// State function type for moving between states
type stateFn func(lexer *Lexer) stateFn

// Parse
// Main glue to parse a query
func Parse(input string, opts ...ParserOption) (elastic.Query, error) {
	l := &Lexer{
		input:         input,
		query:         elastic.NewBoolQuery(),
		nextCondition: queryShould,
		inQuote: noQuote,
		options:       opts,
		matchType:     autoMatch,
		field:         "",
	}
	return l.Run()
}

// resets conditions for the next term, likely after adding a new query
func (lexer *Lexer) resetConditions() {
	lexer.nextCondition = queryShould
	lexer.field = ""
	lexer.matchType = autoMatch
}

// use accumulated information to create an elasticsearch query and adds it to the parser
func (lexer *Lexer) commitQuery() *elastic.BoolQuery {
	var query elastic.Query
	// Single term
	if len(lexer.field) < 1 {
		query = elastic.NewSimpleQueryStringQuery(lexer.value())
		for _, field := range lexer.defaultFields {
			query.(*elastic.SimpleQueryStringQuery).Field(field)
		}
		lexer.Advance(itemTerm)
		return lexer.addQuery(query)
	}
	// Field logic
	// Remap
	for sourceName, targetName := range lexer.mappings {
		if lexer.field == sourceName {
			lexer.field = targetName
		}
		if strings.HasPrefix(lexer.field, sourceName+".") {
			lexer.field = strings.Replace(lexer.field, sourceName, targetName, 1)
		}
	}
	switch lexer.matchType {
	case upperMatch:
		query = elastic.NewRangeQuery(lexer.field).Gt(lexer.value())
	case lowerMatch:
		query = elastic.NewRangeQuery(lexer.field).Lt(lexer.value())
	case keywordMatch:
		query = elastic.NewTermQuery(lexer.field+".keyword", lexer.value())
	case regexMatch:
		query = elastic.NewRegexpQuery(lexer.field, lexer.value())
	case simpleQueryMatch:
		query = elastic.NewSimpleQueryStringQuery(lexer.value()).Field(lexer.field)
	case autoMatch:
		query = elastic.NewMatchQuery(lexer.field, lexer.value())
	default:
		query = elastic.NewMatchQuery(lexer.field, lexer.value())
	}
	for _, nestedPath := range lexer.nestedPaths {
		if strings.HasPrefix(lexer.field, nestedPath+".") {
			query = elastic.NewNestedQuery(nestedPath, query)
		}
	}
	lexer.Advance(itemFieldValue)
	return lexer.addQuery(query)
}

// addQuery adds the query to the correct lexer.query set ( Must, MustNot, Should ) and resets conditions
func (lexer *Lexer) addQuery(query elastic.Query) *elastic.BoolQuery {
	defer lexer.resetConditions()
	switch lexer.nextCondition {
	case queryMustNot:
		return lexer.query.MustNot(query)
	case queryMust:
		return lexer.query.Must(query)
	}
	return lexer.query.Should(query)
}

// Run configures our options and run through our states functions
func (lexer *Lexer) Run() (elastic.Query, error) {
	for _, opt := range lexer.options {
		err := opt(lexer)
		if err != nil {
			return nil, err
		}
	}
	for state := lexText; state != nil; {
		state = state(lexer)
	}
	if lexer.lastError != nil {
		return lexer.query, lexer.lastError
	}
	return lexer.query, nil
}

// Advance commits the parsing we just did to move to the next step.
func (lexer *Lexer) Advance(emitedType itemType) {
	//log.Printf("Done with %s : %s", emitedType, lexer.value())
	lexer.start = lexer.pos
}

// Skips ... whitespaces
func (lexer *Lexer) skipWhiteSpaces() {
	for {
		rune := lexer.peek()
		if rune == ' ' {
			lexer.next()
			lexer.Advance(itemSkipWhitespace)
		} else {
			break
		}
	}
}

// next returns the next rune in the input.
func (lexer *Lexer) next() rune {
	if lexer.pos >= len(lexer.input) {
		lexer.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(lexer.input[lexer.pos:])
	lexer.width = w
	lexer.pos += lexer.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (lexer *Lexer) peek() rune {
	r := lexer.next()
	lexer.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (lexer *Lexer) backup() {
	lexer.pos -= lexer.width
}

// value gets the value for our current position since the last Advance
func (lexer *Lexer) value() string {
	return strings.TrimSpace(lexer.input[lexer.start:lexer.pos])
}
