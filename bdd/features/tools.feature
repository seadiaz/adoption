Feature: manage tools
  In order to incorporate new tools to follow up them
  As a implementer of certain tool
  I need to be able to add/update/remove tools

  @tools
  Scenario: idempotency
    Given there is a tool named Catelnuw
    When we try to create a tool named Catelnuw
    And we ask for the list of managed tools
    Then the list of the tool should have the length of 1

  @tools
  Scenario: filter by label
    Given there is a tool named Oboguame
    And there is a tool named Dabefimo
    And the tool Oboguame is marked with team label as Ruabov
    When we ask for the list of managed tools filter by label team equals Ruabov
    Then the list of the tool should have the length of 1
    And the list of the tool should contains to Oboguame