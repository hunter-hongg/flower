{
  open Parser
  exception SyntaxError of string
}

let digit = ['0'-'9']
let digits = digit+
let ident = ['a'-'z' 'A'-'Z' '_'] ['a'-'z' 'A'-'Z' '0'-'9' '_']*

rule token = parse
  | ","                 { COMMA }
  | '['                 { LBRACKET }
  | ']'                 { RBRACKET }
  | "workflow"          { WORKFLOW }
  | "step"              { STEP }
  | "doing"             { DOING }
  | "done"              { DONE }
  | "with"              { WITH }
  | "finish"            { FINISH }
  | "is"                { is_end "" lexbuf }
  | [' ' '\t' '\r']     { token lexbuf }
  | "#" [^'\n']* '\n'?  { Lexing.new_line lexbuf; token lexbuf }
  | '\n'                { Lexing.new_line lexbuf; token lexbuf }
  | ident as id         { IDENT id }
  | eof                 { EOF }
  | _ as c              { raise (SyntaxError ("Unexpected char: " ^ String.make 1 c)) }

and is_end acc = parse 
  | "end"               { IS (acc) }
  | _ as c              { is_end (acc ^ String.make 1 c) lexbuf }
