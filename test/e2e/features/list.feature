Feature: List evaluation
  Scenario: It should evaluate to list of numbers and booleans
    Given the program
      """
      (list 2 (+ 2 2) (not false) 3 false)
      """
    Then the result is
      """
      2 4 true 3 false
      """

  Scenario: It should evaluate a builtin function 'head' applied to a list
    Given the program
      """
      (head (list (+ 2 3)))
      """
    Then the result is
      """
      5
      """

  Scenario: It should evaluate a builtin function 'head' for different types in a list
    Given the program
      """
      (head (list true 1 nil))
      """
    Then the result is
      """
      true
      """

  Scenario: It should not evaluate a builtin function 'head' applied to an integer type
    Given the program
      """
      (head 2)
      """
    Then the result is
      """
      Function is not applicable for INTEGER type.
      """

  Scenario: It should not evaluate a builtin function 'head' if more than one argument is passed
    Given the program
      """
      (head (list 1 2 3) 2)
      """
    Then the result is
      """
      Function is not applicable for INTEGER type.
      """

  Scenario: It should evaluate a builtin function 'tail'
    Given the program
      """
      (tail (list 1 2 3))
      """
    Then the result is
      """
      2 3
      """

  Scenario: It should evaluate a builtin function 'tail' for a list with a single element
    Given the program
      """
      (tail (list 1))
      """
    Then the result is
      """
      nil
      """

  Scenario: It should evaluate a builtin function 'tail' for a list with complex expressions
    Given the program
      """
      (tail (list (+ 1 2) (* 5 6)))
      """
    Then the result is
      """
      30
      """