Feature: If expression
  Scenario: It should evaluate to true
    Given the program
      """
      (if (> 4 3) true false)
      """
    Then the result is
      """
      true
      """

  Scenario: It should evaluate to false
    Given the program
      """
      (if (<= 4 3) true false)
      """
    Then the result is
      """
      false
      """

  Scenario: It should evaluate to an error when condition is an arithmetic expression
    Given the program
      """
      (if (+ 4 3) true false)
      """
    Then the result is
      """
      Condition for if expression should evaluate to BOOLEAN type. Found INTEGER type.
      """

  Scenario: It should evaluate and give INTEGER type
    Given the program
      """
      (if (not= 7 6) 3 4)
      """
    Then the result is
      """
      3
      """

  Scenario: It should evaluate and give INTEGER type for arithmetic expressions
    Given the program
      """
      (if (not true) (+ 6 7) (+ 9 12))
      """
    Then the result is
      """
      21
      """

  Scenario: It should evaluate in combination with LET expression
    Given the program
      """
      (let x 3)
      (if (not true) (+ 6 7) x))
      """
    Then the result is
      """
      3
      """

  Scenario: It should evaluate in combination with LET expression and identifier as a condition
    Given the program
      """
      (let x true)
      (if x 4 3)
      """
    Then the result is
      """
      4
      """

  Scenario: It should not evaluate if condition is missing
    Given the program
      """
      (if)
      """
    Then the error is
      """
      If expression is missing condition.
      """

  Scenario: It should not evaluate if "then" expression is missing
    Given the program
      """
      (if true)
      """
    Then the error is
      """
      If expression is missing then expression.
      """

  Scenario: It should not evaluate if "else" expression is missing
    Given the program
      """
      (if true 4)
      """
    Then the error is
      """
      If expression is missing else expression.
      """
