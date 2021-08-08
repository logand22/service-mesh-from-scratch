# service mesh from scratch

Practice programming by building a service mesh from scratch

## Controller 

The controller is in charge of watching the kubernetes API for changes and updating the data plane proxies. 

## The injector 

The sidecar injector watches the kubernetes API and injects envoy sidecars into containers when it sees the 
correct labels added. 

## Testing

Currently creating local test instances using [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)