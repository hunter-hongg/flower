%{
  open Ast
%}

%token WORKFLOW
%token STEP
%token WITH FINISH
%token DOING DONE
%token COMMA LBRACKET RBRACKET
%token <string> IDENT
%token <string> IS
%token EOF

%start program
%type <Ast.def> program

%%

program:
  | defs = def EOF { defs }

def:
  | WORKFLOW name = IDENT DOING steps = list(step) DONE { 
    DWorkFlow (
      { 
        line = $startpos.Lexing.pos_lnum; 
        col = $startpos.Lexing.pos_cnum - $startpos.Lexing.pos_bol + 1;
      },
      name, steps
    )
  }

step: 
  | STEP name = IDENT steps = IS with_expr = with_opt {
    EStep (
      { 
        line = $startpos.Lexing.pos_lnum; 
        col = $startpos.Lexing.pos_cnum - $startpos.Lexing.pos_bol + 1;
      },
      name, steps, with_expr
    )
  }

with_opt: 
  /* None */ { None }
  | WITH list(with_expr) FINISH { 
    Some (
      EWith (
        {
          line = $startpos.Lexing.pos_lnum; 
          col = $startpos.Lexing.pos_cnum - $startpos.Lexing.pos_bol + 1;
        },
      $2
      )
    )
  }

with_expr: 
  | IDENT simple_expr COMMA { 
    EWithAtom(
      {
        line = $startpos.Lexing.pos_lnum; 
        col = $startpos.Lexing.pos_cnum - $startpos.Lexing.pos_bol + 1;
      },
      $1, $2
    )
  }

simple_expr: 
  | LBRACKET separated_list(COMMA, IDENT) RBRACKET { EArray(
      { 
        line = $startpos.Lexing.pos_lnum; 
        col = $startpos.Lexing.pos_cnum - $startpos.Lexing.pos_bol + 1;
      },
      $2
  ) }
