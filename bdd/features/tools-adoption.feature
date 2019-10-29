Feature: know people adoption
  In order to understand the level of adoption
  As a implementer of certain tool
  I need to be able to know how many people is using the tool

  Scenario: 50% of adoption
    Given there is a tool named Uzojoje
    And a person named Locakag
    And a person named Fujobme which have adopted tool Uzojoje
    When we ask for the level of adoption of the tool Uzojoje
    Then the adoption level of tool Uzojoje should be 50 percent

  Scenario: Nobody have adopted
    Given there is a tool named Uzojoje
    When we ask for the level of adoption of the tool Uzojoje
    Then the adoption level of tool Uzojoje should be 0 percent
    