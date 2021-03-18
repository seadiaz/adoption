Feature: know team adoption
  In order to understand the level of adoption
  As a implementer of certain adoptable
  I need to be able to know how many people is using the adoptable

  @adoption
  Scenario: Retrieve team adopter list
    Given there is a adoptable named Uzojoje
    And there is a person named Fujobme which have adopted adoptable Uzojoje
    And there is a person named Dupilze which have adopted adoptable Uzojoje
    And there is a team named Mozpakkek
    And the person Fujobme is member of the team Mozpakkek
    And there is a team named Hunolbu
    And the person Dupilze is member of the team Hunolbu
    When we ask for the level of adoption of the adoptable Uzojoje
    Then the list of team adopters of the adoptable Uzojoje should contain to Mozpakkek
    And the list of team adopters of the adoptable Uzojoje should contain to Hunolbu

  @adoption
  Scenario: Retrieve team absentees list
    Given there is a adoptable named Uzojoje
    And there is a person named Fujobme
    And there is a person named Dupilze
    And there is a team named Mozpakkek
    And the person Fujobme is member of the team Mozpakkek
    And there is a team named Hunolbu
    And the person Dupilze is member of the team Hunolbu
    When we ask for the level of adoption of the adoptable Uzojoje
    Then the list of team absentees of the adoptable Uzojoje should contain to Mozpakkek
    And the list of team absentees of the adoptable Uzojoje should contain to Hunolbu

  @adoption
  Scenario: Retrieve team adoption
    Given there is a adoptable named Uzojoje
    And there is a person named Fujobme which have adopted adoptable Uzojoje
    And there is a person named Etuwime
    And there is a team named Mozpakkek
    And the person Fujobme is member of the team Mozpakkek
    And the person Etuwime is member of the team Mozpakkek
    And there is a person named Dupilze
    And there is a team named Hunolbu
    And the person Dupilze is member of the team Hunolbu
    When we ask for the level of adoption of the adoptable Uzojoje
    Then the team adoption level of the adoptable Uzojoje should be 50 percent

  @adoption
  Scenario: Retrieve adoption per team
    Given there is a adoptable named Uzojoje
    And there is a person named Fujobme which have adopted adoptable Uzojoje
    And there is a person named Etuwime
    And there is a team named Mozpakkek
    And the person Fujobme is member of the team Mozpakkek
    And the person Etuwime is member of the team Mozpakkek
    And there is a person named Dupilze
    And there is a team named Hunolbu
    And the person Dupilze is member of the team Hunolbu
    When we ask for the level of adoption of the adoptable Uzojoje
    Then the team adoption level for the team Mozpakkek of the adoptable Uzojoje should be 50 percent
    And the team adoption level for the team Hunolbu of the adoptable Uzojoje should be 0 percent
