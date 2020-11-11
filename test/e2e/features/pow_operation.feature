Feature: Pow arithmetic expressions
  Scenario: It should raise a number to the power of another number
    Given the program
      """
      (^ 2 3)
      """
    Then the result is
      """
      8
      """

  Scenario: It should raise a number to the power of a result of the complex expression
    Given the program
      """
      (^ 2 (+ 3 4))
      """
    Then the result is
      """
      128
      """

  Scenario: It should raise a number to the power of another number and again to the power of third number
    Given the program
      """
      (^ 2 3 4)
      """
    Then the result is
      """
      4096
      """