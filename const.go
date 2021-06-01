package yql_elastic

//go:generate stringer -type=itemType
type itemType int

const (
	itemTerm itemType = iota
	itemLeftGroupDelim
	itemRightGroupDelim
	itemField
	itemFieldValue
	itemMust
	itemMustNot
	itemSkipWhitespace
	itemLowerThan
	itemGreaterThan
	itemKeyword
	itemRegex
	itemOpenQuote
	itemCloseQuote
)

//go:generate stringer -type=conditionType
type conditionType int

const (
	queryShould conditionType = iota
	queryMust
	queryMustNot
)

//go:generate stringer -type=matchType
type matchType int

const (
	autoMatch matchType = iota
	phraseMatch
	keywordMatch
	regexMatch
	lowerMatch
	upperMatch
	simpleQueryMatch
)

const (
	noQuote          rune = 0
	singleQuote           = '\''
	doubleQuote           = '"'
	plusSign              = '+'
	minusSign             = '-'
	whiteSpace            = ' '
	leftParenthesis       = '('
	rightParenthesis      = ')'
	escapeChar            = '\\'
	semiColon             = ':'
	lowerThan             = '<'
	greaterThan           = '>'
	equalSign             = '='
	tildeSign             = '~'
)

const eof = -1
