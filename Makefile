GO:=go
topic1:; $(GO) run src/topic_1/hello_world.go
topic2:; $(GO) run src/topic_2/slices.go
topic3:; $(GO) run src/topic_3/generics.go
topic4:; $(GO) run src/topic_4/concurrency.go
topic5:; $(GO) run src/topic_5/webserver.go
list:
	@echo Available: topic1 topic2 topic3 topic4 topic5
