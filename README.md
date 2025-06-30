## Liveness Observer for Services

Lightweight Go library that implements a simple Liveness Probe HTTP endpoint for Kubernetes. Easily plug it into your Go applications to expose health status checks for Google Cloud and any K8s-based deployments. Helps keep your pods healthy and auto-recovering with minimal setup. Ideal for microservices and highload apps that need reliable self-healing in production.

### Usage:

    // probe will be failed after 15 seconds
    probe := liveness.NewProbe(time.Second*15)

    service1 := usefulService1.New()
    service2 := usefulService2.New()

    probe.Watch(service1, service2)

    if !probe.IsAlive() {
        panic("observable services are down")
    }


    
