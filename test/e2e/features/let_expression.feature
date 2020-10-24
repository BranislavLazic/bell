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