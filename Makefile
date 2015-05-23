all:
	cd ${GOPATH} && go get -v ...mirrorbraind && go build ...mirrorbraind

curl:
	curl -v http://localhost:8080/congress/2012/mp4-h264-LQ-iProd/29c3-5044-en-time_is_not_on_your_side_h264-iprod.mp4

