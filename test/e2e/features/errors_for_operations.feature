Feature: Errors for operations
  Scenario: It should return "Unexpected EOF" when closing parentheses is missing
    Given the program
      """
      (+ 6 3
      """
    Then the error is
      """
      Unexpected EOF at index 6.
      """

  Scenario: It should return "No expression given" when opening parentheses is missing
    Given the program
      """
      + 6 3
      """
    Then the error is
      """
      No expression given.
      """

  Scenario: It should return "Illegal use of operator" when operator is used as an expression
    Given the program
      """
      (+ 6 3 + 5)
      """
    Then the error is
      """
      Illegal use of operator '+' at index 7.
      """


  Scenario: It should return "Illegal character found" when non-existing operator is used as an expression
    Given the program
      """
      (& 6 3)
      """
    Then the error is
      """
      Illegal character '&' found at index 1.
      """


  Scenario: It should return an error when more than one expression is found for logical "not" expression
    Given the program
      """
      (not true false)
      """
    Then the error is
      """
      'not' operation contains more than one expression or lacks a closing parentheses.
      """


  Scenario: It should return an error when boolean is found in an arithmetic expression
    Given the program
      """
      (+ 2 false)
      """
    Then the result is
      """
      Operation (+ 2 false) cannot be performed for types: INTEGER and BOOLEAN
      """

  Scenario: It should return an error when number is found in a logical expression
    Given the program
      """
      (and true (+ 2 3))
      """
    Then the result is
      """
      Operation (and true (+ 2 3)) cannot be performed for types: BOOLEAN and INTEGER
      """

  Scenario: It should return an error when bool is found in a relational expression
    Given the program
      """
      (> (+ 2 3) true)
      """
    Then the result is
      """
      Operation (> (+ 2 3) true) cannot be performed for types: INTEGER and BOOLEAN
      """
