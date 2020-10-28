Feature: Let expression
  Scenario: It should evaluate and assign a result
    Given the program
      """
      (let x 3)
      """
    Then the result is
      """
      3
      """

  Scenario: It should evaluate and assign a result in combination with arithmetic expression
    Given the program
      """
      (let x 3)
      (+ x 5)
      """
    Then the result is
      """
      8
      """

  Scenario: It should evaluate and assign a result of an arithmetic expression
    Given the program
      """
      (let x (+ 3 5 6))
      """
    Then the result is
      """
      14
      """

  Scenario: It should evaluate and assign a result of an logical expression
    Given the program
      """
      (let x (or true false (not false)))
      """
    Then the result is
      """
      true
      """


  Scenario: It should evaluate to a function
    Given the program
      """
      (let sq [x] (* x x))
      """
    Then the result is
      """
      (sq x)
      """

  Scenario: It should evaluate to a function and call the function
    Given the program
      """
      (let sq [x] (* x x))
      (sq 3)
      """
    Then the result is
      """
      9
      """

  Scenario: It should evaluate to a function with multiple args and call the function
    Given the program
      """
      (let max [a b]
        (if (> a b) a b))
      (max 3 4)
      """
    Then the result is
      """
      4
      """

  Scenario: It should evaluate to a function and call another function as an arg
    Given the program
      """
      (let sq [x] (* x x))
      (let add [a b] (+ a b))
      (add (sq 3) 5)
      """
    Then the result is
      """
      14
      """

  Scenario: It should evaluate to a function and call the function with complex args
    Given the program
      """
      (let add [a b] (+ a b))
      (add (+ 3 6) (* 5 7))
      """
    Then the result is
      """
      44
      """

  Scenario: It should evaluate to a function and call the function with an argument from the outer scope
    Given the program
      """
      (let x 7)
      (let multiply [a] (* a x))
      (multiply 5)
      """
    Then the result is
      """
      35
      """

  Scenario: It should evaluate and support higher order functions
    Given the program
      """
      (let map [arg func] (func arg))
      (map 10 (let a [a] (* a 2)))
      """
    Then the result is
      """
      20
      """

  Scenario: It should evaluate and use a variable from the inner scope
    Given the program
      """
      (let x 3)
      (let map [a] (let x 5) (+ a x))
      (map 2)
      """
    Then the result is
      """
      7
      """

  Scenario: It should evaluate and use a variable from the outer scope
    Given the program
      """
      (let x 3)
      (let map [a] (+ a x))
      (map 2)
      """
    Then the result is
      """
      5
      """