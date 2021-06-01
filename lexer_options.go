package yql_elastic

type ParserOption func(lexer *Lexer) error

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
