Feature: Divide arithmetic expressions
  Scenario: It should divide two numbers
    Given the program
      """
      (/ 6 3)
      """
    Then the result is
      """
      2
      """

  Scenario: It should divide multiple numbers
    Given the program
      """
      (/ 12 2 3)
      """
    Then the result is
      """
      2
      """

  Scenario: It should divide multiple numbers and expressions
    Given the program
      """
      (/ 10 (/ 2 1) 5)
      """
    Then the result is
      """
      1
      """

  Scenario: It should evaluate a single integer
    Given the program
      """
      (/ 3)
      """
    Then the result is
      """
      3
      """

  Scenario: It should divide two expressions
    Given the program
      """
      (/ (+ 10 4)
      (- 4 2))
      """
    Then the result is
      """
      7
      """

  Scenario: It should divide one expressions and a literal
    Given the program
      """
      (/ (/ 10 2)
      1)
      """
    Then the result is
      """
      5
      """

  Scenario: It should divide a literal and one expression
    Given the program
      """
      (/ 16
      (+ 4 2))
      """
    Then the result is
      """
      2
      """