(** [none_safe default] returns [default] if the argument is [None], and the value inside the [Some] otherwise. *)
let none_safe default = function None -> default | Some x -> x

(** [print_none_safe] prints the value inside the [Some] if the argument is [Some], and "None" otherwise. *)
let print_none_safe = function
  | None -> print_endline "None"
  | Some x -> print_endline x

(** [print_none_safe_tuple] prints out a tuple only if it's not None *)
let print_none_safe_tuple = function
  | None -> print_endline "None"
  | Some (x, y) -> print_endline @@ "(" ^ x ^ "," ^ y ^ ")"

(** [print_list] prints every element in a list *)
let rec print_list = function
  | [] -> print_endline ""
  | [ x ] -> print_endline x
  | h :: t ->
      print_endline h;
      print_list t
