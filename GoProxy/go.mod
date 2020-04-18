module github.com/el10savio/Kube-Chain-Replication/GoProxy

go 1.14

// Refer https://github.com/kubernetes/client-go/blob/master/INSTALL.md#go-modules

require (
	github.com/gorilla/mux v1.7.4
	github.com/sirupsen/logrus v1.5.0
	k8s.io/apimachinery v0.0.0-20191004115801-a2eda9f80ab8
	k8s.io/client-go v0.0.0-20191016111102-bec269661e48
)
