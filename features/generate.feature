Feature: generate
  In order to be happy
  As a hungry gopher
  I need to be able to generate

  Scenario: Start with 10 and number is 3
    Given the start number 10
    When follow a number as 3
    Then the result should be 30
  