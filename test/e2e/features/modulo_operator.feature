Feature: Modulo expressions
  Scenario: It should give a remainder for division of two numbers
    Given the program
      """
      (% 4 3)
      """
    Then the result is
      """
      1
      """

  Scenario: It should give a remainder for division of multiple numbers
    Given the program
      """
      (% 10 3 2)
      """
    Then the result is
      """
      1
      """