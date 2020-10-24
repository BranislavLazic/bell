@wip
Feature: If expression
  Scenario: It should evaluate
    Given the program
      """
      (if (> 4 3) true false)
      """
    Then the result is
      """
      true
      """