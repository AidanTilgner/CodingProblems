open Ocaml_problems

let () =
  print_endline "\n";

  print_endline "Last of [ 'a'; 'b'; 'c'; 'd' ]";
  Beginner.last [ "a"; "b"; "c"; "d" ] |> Utils.print_none_safe;
  print_endline "\n";

  print_endline "Last two of [ 'a'; 'b'; 'c'; 'd'; 'e' ]";
  Beginner.last_two [ "a"; "b"; "c"; "d"; "e" ] |> Utils.print_none_safe_tuple;
  print_endline "\n";

  print_endline "4rd of [ 'a'; 'b'; 'c'; 'd'; 'e' ]";
  Beginner.at 3 [ "a"; "b"; "c"; "d"; "e" ] |> Utils.print_none_safe;
  print_endline "\n";

  print_endline "Length of [ 'a'; 'b'; 'c'; 'd'; 'e' ]";
  Beginner.length [ "a"; "b"; "c"; "d"; "e" ] |> string_of_int |> print_endline;
  print_endline "\n";

  print_endline "Reverse of [ 'a'; 'b'; 'c'; ]";
  Beginner.rev [ "a"; "b"; "c" ] |> Utils.print_list;
  print_endline "\n";

  print_endline "Is palindrome: [ 'a'; 'b'; 'a'; ]";
  Beginner.is_palindrome [ "a"; "b"; "a" ] |> string_of_bool |> print_endline;
  print_endline "\n"
