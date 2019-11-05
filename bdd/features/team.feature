Feature: group people in teams
  In order to visualize the adoption in a grouped way
  As a implementer of certain tool
  I need to be able to add/remove people to/from teams

  @wip
  Scenario: Creating team
    Given there is a team named Dinpetor
    And there is a person named Uppukmid
    And there is a person named Puwtuvwo
    And there is a person named Donagi
    And the person Uppukmid is member of the team Dinpetor
    And the person Puwtuvwo is member of the team Dinpetor
    And the person Donagi is member of the team Dinpetor
    When we ask for the members of team Dinpetor
    Then the list of the members should have the length of 3
    And the list of the members should contains to Uppukmid
    And the list of the members should contains to Puwtuvwo
    And the list of the members should contains to Donagi