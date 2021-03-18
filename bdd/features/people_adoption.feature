Feature: know people adoption
  In order to understand the level of adoption
  As a implementer of certain adoptable
  I need to be able to know how many people is using the adoptable

  @adoption
  Scenario: 50% of adoption
    Given there is a adoptable named Uzojoje
    And there is a person named Locakag
    And there is a person named Fujobme which have adopted adoptable Uzojoje
    When we ask for the level of adoption of the adoptable Uzojoje
    Then the adoption level of the adoptable Uzojoje should be 50 percent

  @adoption
  Scenario: 75% of adoption
    Given there is a adoptable named Uzojoje
    And there is a person named Locakag
    And there is a person named Fujobme which have adopted adoptable Uzojoje
    And there is a person named Dupilze which have adopted adoptable Uzojoje
    And there is a person named Jullakiko which have adopted adoptable Uzojoje
    When we ask for the level of adoption of the adoptable Uzojoje
    Then the adoption level of the adoptable Uzojoje should be 75 percent

  @adoption
  Scenario: Nobody have adopted
    Given there is a adoptable named Uzojoje
    When we ask for the level of adoption of the adoptable Uzojoje
    Then the adoption level of the adoptable Uzojoje should be 0 percent

  @adoption
  Scenario: Retrieve adopter list
    Given there is a adoptable named Uzojoje
    And there is a person named Fujobme which have adopted adoptable Uzojoje
    And there is a person named Dupilze which have adopted adoptable Uzojoje
    And there is a person named Jullakiko which have adopted adoptable Uzojoje
    When we ask for the level of adoption of the adoptable Uzojoje
    Then the list of adopters of the adoptable Uzojoje should contain to Fujobme
    And the list of adopters of the adoptable Uzojoje should contain to Dupilze
    And the list of adopters of the adoptable Uzojoje should contain to Jullakiko
    And the list of adopters of the adoptable Uzojoje should not contain to Kavdanah

  @adoption
  Scenario: Retrieve absentees list
    Given there is a adoptable named Uzojoje
    And there is a person named Fujobme
    And there is a person named Dupilze
    And there is a person named Jullakiko
    When we ask for the level of adoption of the adoptable Uzojoje
    Then the list of absentees of the adoptable Uzojoje should contain to Fujobme
    And the list of absentees of the adoptable Uzojoje should contain to Dupilze
    And the list of absentees of the adoptable Uzojoje should contain to Jullakiko
    And the list of absentees of the adoptable Uzojoje should not contain to Kavdanah
    