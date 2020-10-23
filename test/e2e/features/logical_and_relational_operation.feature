Feature: Evaluate logical expressions
  Scenario: It should give a result for a logical "and" expression
    Given the program
      """
      (and true false)
      """
    Then the result is
      """
      false
      """

  Scenario: It should give a result for a logical "or" expression
    Given the program
      """
      (or true false)
      """
    Then the result is
      """
      true
      """

  Scenario: It should give a result for a logical "not" expression
    Given the program
      """
      (not true)
      """
    Then the result is
      """
      false
      """

  Scenario: It should give a result for a complex logical expression
    Given the program
      """
      (not (or true false (and true false)))
      """
    Then the result is
      """
      false
      """

  Scenario: It should give a result for greater than expression
    Given the program
      """
      (> 3 4)
      """
    Then the result is
      """
      false
      """

  Scenario: It should give a result for greater than expression
    Given the program
      """
      (< 3 4)
      """
    Then the result is
      """
      true
      """

  Scenario: It should give a result for greater than or equal expression
    Given the program
      """
      (>= 3 3)
      """
    Then the result is
      """
      true
      """

  Scenario: It should give a result for less than or equal expression
    Given the program
      """
      (<= 4 3)
      """
    Then the result is
      """
      false
      """

  Scenario: It should give a result for equal expression
    Given the program
      """
      (= 4 3)
      """
    Then the result is
      """
      false
      """

  Scenario: It should give a result for not equal expression
    Given the program
      """
      (not= 4 3)
      """
    Then the result is
      """
      true
      """

  Scenario: It should give a result for a logical "not" expression followed by comparative expression
    Given the program
      """
      (not (> 4 5))
      """
    Then the result is
      """
      true
      """

  Scenario: It should give a result for  alogical "and" expression followed by comparative expressions
    Given the program
      """
      (and (> 4 5) (< 3 4))
      """
    Then the result is
      """
      false
      """

  Scenario: It should give a result for a logical "or" expression followed by comparative expressions
    Given the program
      """
      (or (> 4 5) (< 3 4))
      """
    Then the result is
      """
      true
      """