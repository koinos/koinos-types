grammar SchemaLang;

ModuleName : [a-z]([a-z.][a-z])* ;
VariableName : [a-z]([a-z_][a-z])* ;
IntLiteral : [0-9]+ | '0x' [a-f0-9]+ ;
Whitespace : [ \t]+ ;
OptionalNewline : [ \t\n]+ ;
Module : 'module' ;
Struct : 'struct' ;
Enum : 'enum' ;
Variant : 'variant' ;
Opaque : 'opaque' ;
Less : '<' ;
Greater : '>' ;
LeftBracket : '{' ;
RightBracket : '}' ;
LeftSquareBracket : '[' ;
RightSquareBracket : ']' ;
Equals : '=' ;
IsA : ':' ;
Comma : ',' ;
Dot: '.' ;

schmea_decl : module_decl type_list* ;

module_decl : Module ModuleName ;

type_list : type_declaration type_list* ;
type_declaration : VariableName Equals type_body ;
type_body : struct_decl | variant_decl | enum_decl | VariableName ;

struct_decl : Struct LeftBracket struct_body RightBracket ;
field_decl : VariableName IsA field_type ;
field_type : (opaque_decl | type_name) array_decl? ;
type_name : (ModuleName Dot)? VariableName (Less IntLiteral Greater)? ;
opaque_decl : Opaque Less type_name Greater ;
array_decl : LeftSquareBracket RightSquareBracket ;
struct_body : field_decl (Comma struct_body)? ;

variant_decl : Variant LeftBracket variant_body RightBracket ;
variant_body : VariableName (Comma variant_body)? ;

enum_decl : Enum LeftBracket enum_body RightBracket ;
enum_body : enum_body_values | enum_body_no_values ;
enum_value_decl : VariableName Equals IntLiteral ;
enum_body_values : enum_value_decl (Comma enum_body_values)? ;
enum_no_value_decl : VariableName Comma ;
enum_body_no_values : enum_no_value_decl (Comma enum_body_no_values)? ;
