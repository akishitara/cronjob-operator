module github.com/akishitara/cronjob-operator

go 1.12

require (
	github.com/emicklei/go-restful v2.11.0+incompatible // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-logr/logr v0.1.0
	github.com/kubernetes/dashboard v1.10.1
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	golang.org/x/net v0.0.0-20190812203447-cdfb69ac37fc
	gopkg.in/square/go-jose.v2 v2.4.0 // indirect
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/controller-runtime v0.2.2
)
