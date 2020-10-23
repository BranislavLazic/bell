Feature: Errors for operations
  Scenario: It should return "Unexpected EOF" when closing parentheses is missing
    Given the program
    """
    (+ 6 3
    """
    Then the error is
    """
    Unexpected EOF at index 6.
    """

  Scenario: It should return "No expression given" when opening parentheses is missing
    Given the program
    """
    + 6 3
    """
    Then the error is
    """
    No expression given.
    """

  Scenario: It should return "Illegal use of operator" when operator is used as an expression
    Given the program
    """
    (+ 6 3 + 5)
    """
    Then the error is
    """
    Illegal use of operator '+' at index 7.
    """


  Scenario: It should return "Illegal character found" when non-existing operator is used as an expression
    Given the program
    """
    (& 6 3)
    """
    Then the error is
    """
    Illegal character '&' found at index 1.
    """