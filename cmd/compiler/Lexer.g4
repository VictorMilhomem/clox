lexer grammar Lexer;

// Tokens
TOKEN_LEFT_PAREN : '(' ;
TOKEN_RIGHT_PAREN : ')' ;
TOKEN_LEFT_BRACE : '{' ;
TOKEN_RIGHT_BRACE : '}' ;
TOKEN_COMMA : ',' ;
TOKEN_DOT : '.' ;
TOKEN_MINUS : '-' ;
TOKEN_PLUS : '+' ;
TOKEN_SEMICOLON : ';' ;
TOKEN_SLASH : '/' ;
TOKEN_STAR : '*' ;
TOKEN_BANG : '!' ;
TOKEN_BANG_EQUAL : '!=' ;
TOKEN_EQUAL : '=' ;
TOKEN_EQUAL_EQUAL : '==' ;
TOKEN_GREATER : '>' ;
TOKEN_GREATER_EQUAL : '>=' ;
TOKEN_LESS : '<' ;
TOKEN_LESS_EQUAL : '<=' ;
TOKEN_IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]* ;
TOKEN_STRING : '"' (~["])* '"' ;
TOKEN_NUMBER : DIGIT+ ('.' DIGIT+)? ;
TOKEN_AND : 'and' ;
TOKEN_CLASS : 'class' ;
TOKEN_ELSE : 'else' ;
TOKEN_FALSE : 'false' ;
TOKEN_FOR : 'for' ;
TOKEN_FUN : 'fun' ;
TOKEN_IF : 'if' ;
TOKEN_NIL : 'nil' ;
TOKEN_OR : 'or' ;
TOKEN_PRINT : 'print' ;
TOKEN_RETURN : 'return' ;
TOKEN_SUPER : 'super' ;
TOKEN_THIS : 'this' ;
TOKEN_TRUE : 'true' ;
TOKEN_VAR : 'var' ;
TOKEN_WHILE : 'while' ;
TOKEN_ERROR : .; // Ignorar caracteres desconhecidos
TOKEN_EOF : EOF ;

// Defina fragmentos
fragment DIGIT : [0-9] ;

