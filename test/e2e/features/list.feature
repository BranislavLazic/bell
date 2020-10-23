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