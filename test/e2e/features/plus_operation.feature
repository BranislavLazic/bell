Feature: Plus arithmetic expressions
  Scenario: It should add two numbers
    Given the program
    """
    (+ 6 3)
    """
    Then the result is
    """
    9
    """

  Scenario: It should add two numbers in two expressions but print only the result of the last expression
    Given the program
    """
    (+ 6 3)
    (+ 5 3)
    """
    Then the result is
    """
    8
    """

  Scenario: It should add multiple numbers
    Given the program
    """
    (+ 2 3 4 5)
    """
    Then the result is
    """
    14
    """

  Scenario: It should add multiple numbers and expressions
    Given the program
    """
    (+ 2 (+ 3 4) 5)
    """
    Then the result is
    """
    14
    """

  Scenario: It should evaluate a single integer
    Given the program
    """
    (+ 3)
    """
    Then the result is
    """
    3
    """

  Scenario: It should two plus expressions
    Given the program
    """
    (+ (+ 1 2)
       (+ 3 4))
    """
    Then the result is
    """
    10
    """

  Scenario: It should add one plus expressions and a literal
    Given the program
    """
    (+ (+ 1 2)
       5)
    """
    Then the result is
    """
    8
    """

  Scenario: It should add a literal and one plus expression
    Given the program
    """
    (+ 5
       (+ 1 2))
    """
    Then the result is
    """
    8
    """