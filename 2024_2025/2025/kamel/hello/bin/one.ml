let _diffsum x y : int =
  let x_sorted = List.sort compare x in
  let y_sorted = List.sort compare y in

  let xy = List.combine x_sorted y_sorted in

  let diff d = List.map (fun (k, l) -> abs (k - l)) d in
  let s = List.fold_left (fun acc o -> acc + o) 0 in

  let dd = diff xy in
  let ss = s dd in
  ss

let _parse_input input_str =
  let lines = String.split_on_char '\n' input_str in
  let pairs = List.map(fun line ->
    let parts = String.split_on_char ' ' line
      |> List.filter (fun s -> String.length s > 0)
    in
    match parts with
    | [left; right] -> (int_of_string left, int_of_string right)
    | _ -> failwith "Invalid format"
  ) lines in

  let left_list = List.map fst pairs in
  let right_list = List.map snd pairs in
  (left_list, right_list)

let rec count_occurences lst =
  match lst with
    | [] -> [] (* base case, return empty list *)
    | x :: xs -> (* if we have at least one element, recurse on it *)
      let rest = count_occurences xs in (* recurse first so we get to the bottom *)
      (* here x is now the last element in the list
        and rest is a list of tuples with counts
      *)
      match List.assoc_opt x rest with (* matches x to the key of tuples in rest, with value n *)
       (* filter out existing tuple, and replace with value incremented *) 
        | Some n -> (x, n + 1) :: List.filter (fun (k, _) -> x <> k) rest
        | None -> (x, 1) :: rest (* return *)

let _print_tuple_list lst =
  List.iter (fun (num, count) ->
    Printf.printf "%d appears %d times\n" num count
  ) lst

let _part2 leftl rightl =
  (* count the number of times each number appears in the right list *)
  let occs = count_occurences rightl in
  
  let rec sum_similarity acc = function
      | [] -> acc
      | x :: xs ->
        let count = List.assoc_opt x occs in
        let product = match count with
          | Some n -> x * n
          | None -> 0
        in
        sum_similarity (acc + product) xs
  in
  
  sum_similarity 0 leftl

let _rotate_dial1 v char numi =
  (* v is here e.g. 50, num e.g. 45 or -46 *)
  let new_val = match char with
    | 'L' -> v - numi
    | 'R' -> v + numi
    | _ -> v
  in
  (new_val mod 100 + 100) mod 100

let rotate_dial v char numi =
  let direction = match char with
    | 'L' -> -numi
    | 'R' -> numi
    | _ -> 0
  in

  let new_val = v + direction in
  let wrapped = (new_val mod 100 + 100) mod 100 in

  (* Count how many individual clicks land on 0 *)
  let zeros_hit =
    if direction > 0 then
      let count = ref 0 in
      for i = 1 to direction do
        if (v + i) mod 100 = 0 then incr count
      done;
      !count
    else if direction < 0 then
      let distance = abs direction in
      let count = ref 0 in
      for i = 1 to distance do
        if (v - i) mod 100 = 0 then incr count
      done;
      !count
    else
      0
  in
  (wrapped, zeros_hit)
  
(* 
let rec process_directions acc = function
  | [] -> acc
  | direction_str :: rest ->
    let char = String.get direction_str 0 in
    let num = String.sub direction_str 1 (String.length direction_str - 1) in
    let numi = int_of_string num in
    
    let current = match List.rev acc with
      | x :: _ -> x
      | [] -> 50
    in
    
    let next_val = rotate_dial current char numi in
    
    process_directions (acc @ [next_val]) rest *)

let rec process_directions (current_pos, zero_count) = function
  | [] -> (current_pos, zero_count)
  | direction_str :: rest ->
    let char = String.get direction_str 0 in
    let num = String.sub direction_str 1 (String.length direction_str - 1) in
    let numi = int_of_string num in

    let (next_pos, zeros_hit) = rotate_dial current_pos char numi in

    process_directions (next_pos, zero_count + zeros_hit) rest

let () =
  let input = Hello.Utils.load_input "2025" "1" in
  let lines = String.split_on_char '\n' input
    |> List.filter (fun s -> String.length s > 0)
  in
  let (_final_pos, zero_count) = process_directions (50, 0) lines in
  Printf.printf "Number of zeros: %d\n" zero_count