Feature: Shopping Cart
    As a customer
    I want to add and remove items from my shopping cart
    So that I can manage my intended purchases

    Scenario: Add item to empty cart
        Given the cart is empty
        When I add a "book" with price 29.99
        Then the cart should contain 1 item
        And the total should be 29.99

    Scenario: Remove item from cart
        Given I have a "book" in my cart with price 29.99
        When I remove the "book" from the cart
        Then the cart should be empty
        And the total should be 0.00
