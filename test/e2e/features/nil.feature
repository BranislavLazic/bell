Feature: Nil
  Scenario: It should check equality with nil
    Given the program
      """
      (= nil nil)
      """
    Then the result is
      """
      true
      """

  Scenario: It should give false when compared with integer
    Given the program
      """
      (= nil 1)
      """
    Then the result is
      """
      false
      """
    Given the program
      """
      (= 1 nil)
      """
    Then the result is
      """
      false
      """

  Scenario: It should give false when compared with boolean
    Given the program
      """
      (= nil false)
      """
    Then the result is
      """
      false
      """
    Given the program
      """
      (= true nil)
      """
    Then the result is
      """
      false
      """

  Scenario: It should result in an error when arithmetic expression is applied
    Given the program
      """
      (+ nil 1)
      """
    Then the result is
      """
      Operation (+ nil 1) cannot be performed for types: NIL and INTEGER
      """

  Scenario: It should result in an error when logical expression is applied
    Given the program
      """
      (> nil true)
      """
    Then the result is
      """
      Operation (> nil true) cannot be performed for types: NIL and BOOLEAN
      """