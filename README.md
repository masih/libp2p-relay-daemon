# libp2p-relay-daemon

This repository provides a Cloud Native ready to deploy libp2p relay daemon, including:

* Containerised [`go-libp2p-relay-daemon`](https://github.com/libp2p/go-libp2p-relay-daemon)
* Kubernetes manifests to run the daemon as a deployment object

Additionally, it provides:

* Example deployment manifest to run the daemon with stable libp2p identity
* Development utilities to run the daemon on a local Kubernetes cluster
