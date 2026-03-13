module Lexer = Flowc_lib.Lexer
module Parser = Flowc_lib.Parser
module Jsongen = Flowc_lib.Jsongen
module Ast = Flowc_lib.Ast
module Yojson = Yojson.Basic

let read_file filename =
  let channel = open_in filename in
  let content = really_input_string channel (in_channel_length channel) in
  close_in channel;
  content

let write_file filename content =
  let channel = open_out filename in  (* 打开文件，如果存在则覆盖 *)
  output_string channel content;
  close_out channel 

let () = 
  let file = Sys.argv.(1) in
  let content = read_file file in 
  let lexbuf = Lexing.from_string content in
  let ast = Parser.program Lexer.token lexbuf in
  let temp_path = Filename.temp_file "FlowCompilerPlan" ".json" in
  let json_res = Jsongen.generate_json ast in
  write_file temp_path (Yojson.pretty_to_string json_res);
  Printf.printf "%s\n" (temp_path);
  ()
