package logging

type LogRetriever interface {
	Retrieve() []LogFile
}
