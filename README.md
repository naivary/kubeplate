# kubeplate

kubeplate is a template system for your kubernetes manifests. It's much simpler
to use compared to `kustomize` and doesn't rely on merging multiple config
files. It's built to be extensible and easily useable by anyone who has
knowledge about the go templating.

## Planned features

- [] Variable injection using YAML or JSON files or by ENV variables
- [] Extensible Design to add new functions if needed
- [] Extensible Design to Input Format of the Variables
- [] Extensible Design of the Output Destination (s3, local etc.)
- [] Automatic generation of needed manifests e.g. namespaces, services for
  deployment etc.
- [] Possibility to generate multiple manifest from one array e.g. each element
  = one manifest


