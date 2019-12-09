Feature: manage adoptables
  In order to incorporate new adoptables to follow up them
  As a implementer of certain adoptable
  I need to be able to add/update/remove adoptables

  @adoptables
  Scenario: idempotency
    Given there is a adoptable named Catelnuw
    When we try to create a adoptable named Catelnuw
    And we ask for the list of managed adoptables
    Then the list of the adoptable should have the length of 1

  @adoptables
  Scenario: filter by label
    Given there is a adoptable named Oboguame
    And there is a adoptable named Dabefimo
    And the adoptable Oboguame is marked with team label as Ruabov
    When we ask for the list of managed adoptables filter by label team equals Ruabov
    Then the list of the adoptable should have the length of 1
    And the list of the adoptable should contains to Oboguame