let a = [3; 4; 2; 1; 3; 3]
let b = [4; 3; 5; 3; 9 ;3]

let diffsum x y : int = 
  let x_sorted = List.sort compare x in
  let y_sorted = List.sort compare y in
  
  let xy = List.combine x_sorted y_sorted in
  
  let diff d = List.map (fun (k, l) -> abs (k - l)) d in
  let s = List.fold_left (fun acc o -> acc + o) 0 in
  
  let dd = diff xy in
  let ss = s dd in
  ss

let () = Printf.printf "%d\n" (diffsum a b)