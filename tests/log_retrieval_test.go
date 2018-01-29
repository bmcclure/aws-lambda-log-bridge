package tests

import "github.com/DATA-DOG/godog"

func thereIsALogFileContainingLines(arg1 int) error {
	return godog.ErrPending
}

func iRetrieveTheLogFromTheLocalFolder() error {
	return godog.ErrPending
}

func thereShouldBeLogEntriesRetrieved(arg1 int) error {
	return godog.ErrPending
}

func theLogIsAvailableOnAnSFTPServer() error {
	return godog.ErrPending
}

func iRetrieveTheLogFromTheSFTPServer() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there is a log file containing (\d+) lines$`, thereIsALogFileContainingLines)
	s.Step(`^I retrieve the log from the local folder$`, iRetrieveTheLogFromTheLocalFolder)
	s.Step(`^there should be (\d+) log entries retrieved$`, thereShouldBeLogEntriesRetrieved)
	s.Step(`^the log is available on an SFTP server$`, theLogIsAvailableOnAnSFTPServer)
	s.Step(`^I retrieve the log from the SFTP server$`, iRetrieveTheLogFromTheSFTPServer)
}
