module github.com/sssvip/goutil

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/gin-contrib/sse v0.0.0-20170109093832-22d885f9ecc7 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-sql-driver/mysql v1.3.0
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/golang/sys v0.0.0-20180830151530-49385e6e1522 // indirect
	github.com/google/uuid v0.0.0-20161128191214-064e2069ce9c
	github.com/json-iterator/go v0.0.0-20180701071628-ab8a2e0c74be
	github.com/mattn/go-isatty v0.0.4 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/olebedev/config v0.0.0-20180625110059-ed90d2035b81
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/sssvip/qr v0.2.1 // indirect
	github.com/sssvip/qrterminal v1.0.1
	github.com/stretchr/testify v1.2.3-0.20181224173747-660f15d67dbb
	github.com/ugorji/go v1.1.1 // indirect
	github.com/xcltapestry/xclpkg v0.0.0-20150203092146-4bb35f81c878
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/yaml.v2 v2.2.1 // indirect
)

replace (
	github.com/golang/sys => github.com/gofile/sys v0.0.1
	github.com/mattn/go-isatty => github.com/sssvip/go-isatty v0.0.4
	golang.org/x/sys => github.com/gofile/sys v0.0.1
)

go 1.13
