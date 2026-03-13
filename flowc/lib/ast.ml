(* 位置信息类型 *)
type loc = {
  line: int;  (* 行号，从1开始 *)
  col: int;   (* 列号，从1开始 *)
}

type typ =
  | TInt
  | TBool
  | TFloat
  | TString

type expr =
  | EStep of loc * string * string * with_expr option

and binop =
  | Add | Sub | Mul | Eq | Neq

and def =
  | DWorkFlow of loc * string * expr list

and simple_expr = 
  | EArray of loc * string list

and with_atomic = 
  | EWithAtom of loc * string * simple_expr

and with_expr = 
  | EWith of loc * with_atomic list
