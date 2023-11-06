package main

type TokenType int

const (
	EOF TokenType = iota

	// Literals
	Identifer
	// Integer64
	// Integer32
	// Integer16
	// Integer8
	Integer

	// UnsignedInteger64
	// UnsignedInteger32
	// UnsignedInteger16
	// UnsignedInteger8

	// Float64
	// Float32
	Float

	String
	Boolean
	Null
	Void

	// Comparison and Equality
	Equals
	NotEquals
	GreaterThan
	GreaterThanEqualTo
	LessThan
	LessThanEqualTo

	// Binary Operators
	Plus
	Minus
	Modulo
	Divide
	Multiply

	Exponent
	Square_Root

	And
	Or

	Logic_And
	Logic_Or

	// Assignment
	Basic_Assign
	Plus_Assign
	Minus_Assign
	Multiply_Assign
	Divide_Assign
	Modulo_Assign

	// Prefix/Postfix
	Decrement
	Increment
	Logic_Not

	Xor
	Not

	// Symbols
	Comma
	Dot
	QuesionMark
	Colon
	SemiColon
	Arrow
	Spread

	OpenParen
	CloseParen

	OpenBrace
	CloseBrace

	OpenBracket
	CloseBracket

	// Keywords
	Function
	Import
	Package
	Variable
	Constant
	If
	Else
	Return
)

type Token struct {
	Line     int
	Position int
	Literal  string
	Type     TokenType
}

func GetKeyword(value string) TokenType {
	m := map[string]TokenType{
		"string": String,
		"bool":   Boolean,

		"int": Integer,
		// "int64": Integer64,
		// "int32": Integer32,
		// "int16": Integer16,
		// "int8":  Integer8,

		// "uint":   UnsignedInteger32,
		// "uint64": UnsignedInteger64,
		// "uint32": UnsignedInteger32,
		// "uint16": UnsignedInteger16,
		// "uint8":  UnsignedInteger8,

		// "byte": Integer8,
		// "char": UnsignedInteger8,

		"float": Float,
		"void":  Void,
		// "float64": Float64,
		// "float32": Float32,

		"fn":      Function,
		"using":   Import,
		"package": Package,
		"const":   Constant,
		"var":     Variable,
		"true":    Boolean,
		"false":   Boolean,
		"null":    Null,
		"if":      If,
		"else":    Else,
		"return":  Return,
	}
	if _, ok := m[value]; ok {
		return m[value]
	}
	return Identifer
}