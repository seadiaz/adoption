Feature: manage people
  In order to incorporate new people to follow up them
  As a implementer of certain tool
  I need to be able to add/update/remove people

  Scenario: Idempotency
    Given there is a person named Dinpetor
    When we try to create a person named Dinpetor
    And we ask for the list of people
    Then the list of the people should have the length of 1