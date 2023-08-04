Feature: Adding items to a cart

  Background:
    Given an empty cart

    And the following items are available in the pool:
      | Name       | Price | Quantity |
      | Product 1  | 10    | 5        |
      | Product 2  | 20    | 3        |
      | Product 3  | 15    | 4        |

  Scenario: Adding items to an empty cart
    When I add the item "Product 1" with quantity 2
    And I add the item "Product 2" with quantity 1
    Then the cart should contain 2 items
    And the total price of the cart should be 40

  Scenario: Adding items to a non-empty cart
    Given a cart with the item "Product 1" with quantity 2
    When I add the item "Product 3" with quantity 3
    Then the cart should contain 2 items
    And the total price of the cart should be 55