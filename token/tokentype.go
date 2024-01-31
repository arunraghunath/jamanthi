package token

type tokenType int

const (
	//Errors
	INVALID tokenType = iota
	EOF

	//Primitive types
	CHAR
	STRING
	INT
	BOOL

	//Blocks
	LPAREN // (
	RPAREN // )
	LBRACE // {
	RBRACE // }

	//Operators
	ASSIGN        // =
	PLUS          //+
	MINUS         //-
	ASTERIK       //* Multiply
	SLASH         // / Divide
	MODULO        // % Remainder
	INCREMENT     // ++
	DECREMENT     // --
	PLUSASSIGN    // +=
	MINUSASSIGN   // -=
	ASTERIKASSIGN //*=
	SLASHASSIGN   // /=

	//Comparison Operators
	GT        // > greater than
	LT        // < less than
	GTET      // greater than equal to
	LTET      // less than equal to
	EQUALS    // ==
	NOTEQUALS // !=

	//Logical Operators
	AND // &&
	OR  // ||
	NOT // !

	//Keywords
	VAR
	FUNCTION // func
	TYPE     // type
	IF       //
	ELSE     //
	FOR      //
	TRUE     // Boolean true
	FALSE    // Boolean false
)
