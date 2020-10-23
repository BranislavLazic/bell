Feature: Multiply arithmetic expressions
  Scenario: It should multiply two numbers
    Given the program
      """
      (* 6 3)
      """
    Then the result is
      """
      18
      """

  Scenario: It should multiply multiple numbers
    Given the program
      """
      (* 2 3 4)
      """
    Then the result is
      """
      24
      """

  Scenario: It should multiply multiple numbers and expressions
    Given the program
      """
      (* 2 (* 1 3) 4)
      """
    Then the result is
      """
      24
      """

  Scenario: It should evaluate a single integer
    Given the program
      """
      (* 3)
      """
    Then the result is
      """
      3
      """

  Scenario: It should multiply two expressions
    Given the program
      """
      (* (* 2 3)
      (* 4 5))
      """
    Then the result is
      """
      120
      """

  Scenario: It should multiply one expressions and a literal
    Given the program
      """
      (* (* 1 2)
      5)
      """
    Then the result is
      """
      10
      """

  Scenario: It should multiply a literal and one expression
    Given the program
      """
      (* 5
      (* 1 2))
      """
    Then the result is
      """
      10
      """