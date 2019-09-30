module github.com/Limard/msgpack

require (
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/vmihailenco/tagparser v0.1.0
	golang.org/x/net v0.0.0-20190926025831-c00fd9afed17 // indirect
	google.golang.org/appengine v1.6.1
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
)

replace golang.org/x/net v0.0.0-20190926025831-c00fd9afed17 => github.com/golang/net v0.0.0-20190926025831-c00fd9afed17

go 1.13
