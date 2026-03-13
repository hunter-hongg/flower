comp: 
  cd flowc && dune clean && dune build 
cli: 
  cd flow && go run main.go 
exec: 
  cd flowe && cargo run
front: 
  cd flowd-front && pnpm dev
back: 
  cd flowd-back && dotnet run 
pyexec: 
  cd flowd && python3 main.py
