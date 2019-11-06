Feature: know people adoption
  In order to understand the level of adoption
  As a implementer of certain tool
  I need to be able to know how many people is using the tool

  Scenario: 50% of adoption
    Given there is a tool named Uzojoje
    And there is a person named Locakag
    And there is a person named Fujobme which have adopted tool Uzojoje
    When we ask for the level of adoption of the tool Uzojoje
    Then the adoption level of the tool Uzojoje should be 50 percent

  Scenario: 75% of adoption
    Given there is a tool named Uzojoje
    And there is a person named Locakag
    And there is a person named Fujobme which have adopted tool Uzojoje
    And there is a person named Dupilze which have adopted tool Uzojoje
    And there is a person named Jullakiko which have adopted tool Uzojoje
    When we ask for the level of adoption of the tool Uzojoje
    Then the adoption level of the tool Uzojoje should be 75 percent

  Scenario: Nobody have adopted
    Given there is a tool named Uzojoje
    When we ask for the level of adoption of the tool Uzojoje
    Then the adoption level of the tool Uzojoje should be 0 percent

  Scenario: Retrieve adopter list
    Given there is a tool named Uzojoje
    And there is a person named Fujobme which have adopted tool Uzojoje
    And there is a person named Dupilze which have adopted tool Uzojoje
    And there is a person named Jullakiko which have adopted tool Uzojoje
    When we ask for the level of adoption of the tool Uzojoje
    Then the list of adopters of the tool Uzojoje should contain to Fujobme
    And the list of adopters of the tool Uzojoje should contain to Dupilze
    And the list of adopters of the tool Uzojoje should contain to Jullakiko
    And the list of adopters of the tool Uzojoje should not contain to Kavdanah

  Scenario: Retrieve absentees list
    Given there is a tool named Uzojoje
    And there is a person named Fujobme
    And there is a person named Dupilze
    And there is a person named Jullakiko
    When we ask for the level of adoption of the tool Uzojoje
    Then the list of absentees of the tool Uzojoje should contain to Fujobme
    And the list of absentees of the tool Uzojoje should contain to Dupilze
    And the list of absentees of the tool Uzojoje should contain to Jullakiko
    And the list of absentees of the tool Uzojoje should not contain to Kavdanah

  Scenario: Retrieve team adopter list
    Given there is a tool named Uzojoje
    And there is a person named Fujobme which have adopted tool Uzojoje
    And there is a person named Dupilze which have adopted tool Uzojoje
    And there is a team named Mozpakkek
    And the person Fujobme is member of the team Mozpakkek
    And there is a team named Hunolbu
    And the person Dupilze is member of the team Hunolbu
    When we ask for the level of adoption of the tool Uzojoje
    Then the list of team adopters of the tool Uzojoje should contain to Mozpakkek
    And the list of team adopters of the tool Uzojoje should contain to Hunolbu

  @wip
  Scenario: Retrieve team adopter list
    Given there is a tool named Uzojoje
    And there is a person named Fujobme
    And there is a person named Dupilze
    And there is a team named Mozpakkek
    And the person Fujobme is member of the team Mozpakkek
    And there is a team named Hunolbu
    And the person Dupilze is member of the team Hunolbu
    When we ask for the level of adoption of the tool Uzojoje
    Then the list of team absentees of the tool Uzojoje should contain to Mozpakkek
    And the list of team absentees of the tool Uzojoje should contain to Hunolbu