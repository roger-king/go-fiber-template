setup:
	go get -u github.com/valyala/quicktemplate
	go get -u github.com/valyala/quicktemplate/qtc
	
start:
	go run $(GO_FILES) & echo $$! > $(PID_FILE)
  
# .PHONY is used for reserving tasks words
.PHONY: start before stop restart serve