Feature: String expressions
  Scenario: It should parse a string literal
    Given the program
      """
      ("brown fox")
      """
    Then the result is
      """
      brown fox
      """

  Scenario: It should concatenate three strings
    Given the program
      """
      (+ "brown" " " "fox")
      """
    Then the result is
      """
      brown fox
      """

  Scenario: It should concatenate a string and an integer
    Given the program
      """
      (+ 1 " brown")
      """
    Then the result is
      """
      1 brown
      """

  Scenario: It should concatenate a string and a boolean
    Given the program
      """
      (+ true " brown")
      """
    Then the result is
      """
      true brown
      """

  Scenario: It should concatenate a string and a nil
    Given the program
      """
      (+ nil " brown")
      """
    Then the result is
      """
      nil brown
      """

  Scenario: It should evaluate a length of a string
    Given the program
      """
      (size "brown")
      """
    Then the result is
      """
      5
      """

  Scenario: It should evaluate a length of a string concatenation
    Given the program
      """
      (size (+ "brown " "fox"))
      """
    Then the result is
      """
      9
      """

  Scenario: It should return a first character by using 'head' function
    Given the program
      """
      (head (+ "brown " "fox"))
      """
    Then the result is
      """
      b
      """

  Scenario: It should return a "tail" of a string by using 'tail' function
    Given the program
      """
      (tail (+ "brown " "fox"))
      """
    Then the result is
      """
      rown fox
      """