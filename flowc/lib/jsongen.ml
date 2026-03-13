open Ast

type step = {
  name: string;
  exec: string;
}

type parse_res = {
  name: string; 
  step: step list;
} 

let process_step steps = 
  List.filter_map (fun step -> (
    match step with 
    | EStep (_, name, exec, _) -> Some({name = name; exec = exec})
  )) steps

let step_to_json (steps: step list) =
  let res = List.map (fun (step: step) -> (
    `Assoc [
      ("name", `String step.name);
      ("exec", `String step.exec);
    ]
  )) steps in 
  res

let generate_json def = 
  let workflow = match def with 
  | DWorkFlow (_, name, step) -> (
    {name; step = process_step step}
  ) in
  let workflow_name = workflow.name in 
  let workflow_step = workflow.step in 
  let my_json : Yojson.Basic.t =
    `Assoc [
      ("workflow", `String workflow_name);
      ("steps", `List (step_to_json workflow_step));
    ] 
  in
  my_json
