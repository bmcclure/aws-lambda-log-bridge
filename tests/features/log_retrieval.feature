Feature: retrieve logs
  In order to push logs to a logging service
  As a user
  I need to be able to retrieve logs from multiple sources.

  Scenario: Retrieve a local log file
    Given there is a log file containing 10 lines
    When I retrieve the log from the local folder
    Then there should be 10 log entries retrieved

  Scenario: Retrieve a log file from SFTP
    Given there is a log file containing 10 lines
    And the log is available on an SFTP server
    When I retrieve the log from the SFTP server
    Then there should be 10 log entries retrieved
