package yql_elastic

type ParserOption func(lexer *Lexer) error

type FieldCallBack func(text string) (string, error)

type FieldCallBackError struct {
	error
	FieldName  string
	FieldValue string
}

func WithFieldCallBack(fieldName string, back FieldCallBack) ParserOption {
	return func(lexer *Lexer) (err error) {
		if lexer.fieldCallbacks == nil {
			lexer.fieldCallbacks = make(map[string]FieldCallBack)
		}
		lexer.fieldCallbacks[fieldName] = back
		return nil
	}
}

func WithDefaultFields(fields []string) ParserOption {
	return func(lexer *Lexer) (err error) {
		lexer.defaultFields = fields
		return nil
	}
}

func WithNestedPaths(nestedPaths []string) ParserOption {
	return func(lexer *Lexer) (err error) {
		lexer.nestedPaths = nestedPaths
		return nil
	}
}

func WithFieldMapping(mappings map[string]string) ParserOption {
	return func(lexer *Lexer) error {
		lexer.mappings = mappings
		return nil
	}
}
