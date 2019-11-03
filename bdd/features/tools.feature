Feature: manage tools
  In order to incorporate new tools to follow up them
  As a implementer of certain tool
  I need to be able to add/update/remove tools

  Scenario: Idempotency
    Given there is a tool named Catelnuw
    And there is a tool named Catelnuw
    When we ask for the list of managed tools
    Then the list of the tool should have the length of 1