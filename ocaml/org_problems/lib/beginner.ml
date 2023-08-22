(** [last l] returns the last element of [l] if [l] is not empty. *)
let rec last = function [] -> None | [ x ] -> Some x | _ :: t -> last t

(** [last_two l] returns the last two elements of [l] if [l] is not empty and contains more than one element. *)
let rec last_two = function
  | [] -> None
  | [ _ ] -> None
  | [ x; y ] -> Some (x, y)
  | _ :: t -> last_two t

(** [at l] returns the element in [l] at index [l]. *)
let rec at k = function
  | [] -> None
  | h :: t -> if k = 0 then Some h else at (k - 1) t

(** [length l] returns the length of the provided list *)
let length l =
  let rec aux n = function [] -> n | _ :: t -> aux (n + 1) t in
  aux 0 l

(** [rev l] returns a reversed version of the provided list *)
let rev l =
  let rec aux l2 = function [] -> l2 | h :: t -> aux (h :: l2) t in
  aux [] l

(** [is_palindrome l] returns a boolean representing whether or not the list is a palindrome *)
let is_palindrome l =
  let list = rev l in
  list = l

(** [run_length_encode l] returns a list of tuples which represent a run length encoding (https://en.wikipedia.org/wiki/Run-length_encoding) *)
let run_length_encode list =
  let rec aux count acc = function
    | [] -> []
    | [ x ] -> (count + 1, x) :: acc
    | a :: (b :: _ as t) ->
        if a = b then aux (count + 1) acc t else aux 0 ((count + 1, a) :: acc) t
  in

  aux 0 [] list
