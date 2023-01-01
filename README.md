# libp2p-relay-daemon

[![Build](https://github.com/masih/libp2p-relay-daemon/actions/workflows/build.yml/badge.svg)](https://github.com/masih/libp2p-relay-daemon/actions/workflows/build.yml)

This repository provides a Cloud Native ready to deploy libp2p relay daemon, including:

* Containerised [`go-libp2p-relay-daemon`](https://github.com/libp2p/go-libp2p-relay-daemon)
* Kubernetes manifests to run the daemon as a deployment object

Additionally, it provides:

* Example deployment manifest to run the daemon with stable libp2p identity
* Development utilities to run the daemon on a local Kubernetes cluster

See:

* [`examples/stable-identity`](examples/stable-identity) - example manifests to run relay in a self
  contained Kubernetes namespace, with stable libp2p identity and default JSON configuration file.

## Running relay locally

Install the following prerequisites:

* **Docker Desktop** -- to build and run containers
* **[`kind`](https://kind.sigs.k8s.io)** - to run a local Kubernetes cluster
    * Alternatively, you may use the local
      cluster [provided by Docker Desktop](https://docs.docker.com/desktop/kubernetes/). The
      instructions here use `kind`.
* **[`skaffold`](https://skaffold.dev)** - to aid build and deployment orchestration.
* **[`kustomize`](https://kustomize.io)** - to render Kubernetes manifests, used internally
  by `skaffold`.
* **[`kubectl`](https://kustomize.io)** - to interact with the local Kubernetes cluster.

Run:

1. `kind create cluster`
2. `skaffold dev --profile local`

The commands above will:

* creates a new local Kubernetes cluster using `kind`.
* configures the current `kubectl` context to the local `kind` cluster.
* builds the [`Dockerfile`](Dockerfile)
* generates a random libp2p identity
* renders the example manifests at [`examples/stable-identity`](examples/stable-identity)
* deploys the manifests onto the local Kubernetes cluster
* tails the output from the running container onto the terminal session.
* port-forwards the running instance onto `localhost`
  * Two ports are forwarded: 
    1. `4001` - the libp2p host running the relay, and
    2. `6060` - the debug pprof HTTP server reachable at `http://localhost:6060/debug/pprof/`.

To stop the running instance and delete the deployment, interrupt the session by pressing `Ctrl-C`.

To destroy the local cluster along with the deployed instance, run:

* `kind delete cluster`

