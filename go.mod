module github.com/akishitara/cronjob-operator

go 1.12

require (
	github.com/ghodss/yaml v1.0.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-logr/logr v0.1.0
	github.com/onsi/ginkgo v1.8.0
	github.com/onsi/gomega v1.5.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.1
	golang.org/x/net v0.0.0-20200625001655-4c5254603344
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	sigs.k8s.io/controller-runtime v0.2.2
)
