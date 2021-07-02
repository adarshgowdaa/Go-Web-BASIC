package main

import "github.com/sirupsen/logrus"

func main() {

	config := NewAppConfig()
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	s := NewServer(config,logger)

	s.Initilize()
	s.Listen()

}
