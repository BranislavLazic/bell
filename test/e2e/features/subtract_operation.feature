Feature: Subtract arithmetic expressions
  Scenario: It should give negative number
    Given the program
    """
    (- 2)
    """
    Then the result is
    """
    -2
    """

  Scenario: It should subtract two numbers
    Given the program
    """
    (- 2 1)
    """
    Then the result is
    """
    1
    """

  Scenario: It should subtract two numbers and give negative value
    Given the program
    """
    (- 2 5)
    """
    Then the result is
    """
    -3
    """

  Scenario: It should subtract multiple numbers
    Given the program
    """
    (- 6 2 3)
    """
    Then the result is
    """
    1
    """

  Scenario: It should add two numbers and subtract them from one number
    Given the program
    """
    (- 15
       (+ 6 5))
    """
    Then the result is
    """
    4
    """

  Scenario: It should add two numbers and subtract them from another subtract expression
    Given the program
    """
    (- (+ 6 5)
       (- 7 3))
    """
    Then the result is
    """
    7
    """