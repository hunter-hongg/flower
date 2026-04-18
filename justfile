comp: 
  cd flowc && dune clean && dune build 
cli: 
  cd flow && go build -o ../target/flow
exec: 
  cd flowe && cargo build && cp target/debug/flowe ../target/
front: 
  cd flowd-front && pnpm dev
back: 
  cd flowd-back && go build -o ../target/flowd-back
  target/flowd-back
pyexec: 
  cd flowd && python3 main.py
