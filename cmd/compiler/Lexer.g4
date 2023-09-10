lexer grammar Lexer;

// Ignorar espaços em branco (whitespace)
WS : [ \t\r\n]+ -> skip ;

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

// Tokens para strings e identificadores
TOKEN_STRING : '"' (~["])* '"' ;
TOKEN_NUMBER : DIGIT+ ('.' DIGIT+)? ;
TOKEN_IDENTIFIER : [a-zA-Z_][a-zA-Z_0-9]* ;

// Token EOF
TOKEN_EOF : EOF ;

// Token para caracteres desconhecidos (ERROR)
TOKEN_ERROR : . -> skip ;

// Defina fragmentos
fragment DIGIT : [0-9] ;
