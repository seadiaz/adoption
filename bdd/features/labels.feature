Feature: labels support
  In order to filter only the tools on which I'm interested in
  As a visualizer of the levels of adoption
  I need to be able to filter tools by different criteria

  @wip
  Scenario: Assign labels
    Given there is a tool named Jenkata
    And the tool Jenkata is marked with label Ruabov
    And the tool Jenkata is marked with label Afeubaji
    When we ask for the tool Jenkata
    Then the list of the labels should have the length of 2
    And the list of the labels should contains to Ruabov
    And the list of the labels should contains to Afeubaji
