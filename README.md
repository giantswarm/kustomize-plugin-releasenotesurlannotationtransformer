# release notes URL annotation kustomize plugin

This repository contains a kustomize plugin to add an release notes URL annotation to release CRs.


## How to use

Create new file `releaseNotesTransformer.yaml` (File name does not matter)

```yaml
apiVersion: giantswarm.io/v1
kind: releaseNotesURLAnnotationTransformer
metadata:
  name: awsReleaesNotesURLAnnotationTransformer
# `argsOneLiner` contains the name of the provider.
# The value is being used for the generation of the release
# notes URL
argsOneLiner: aws
```

In your `kustomization.yaml` add:

```yaml
transformers:
  - releaseNotesTransformer.yaml
```

