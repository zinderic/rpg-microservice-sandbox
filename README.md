# RPG Microservice Sandbox

This sandbox explores microservice architecture, deployment patterns and concepts by building a simple RPG game.

# Configure

```
export KUBECONFIG=path/to/kubeconfig
```

# Build

```
skaffold build
```

# Run

Locally:

```
skaffold run
```

Remotely:

```
KUBECONFIG=path/to/kubeconfig skaffold run
```

# Debug

Locally:

```
skaffold debug
```

Remotely:

```
KUBECONFIG=path/to/kubeconfig skaffold debug
```
