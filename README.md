# release notes URL annotation kustomize plugin

This repository contains a kustomize plugin to add an release notes URL annotation to release CRs.


## How to use

Create new file `releaseNotesTransformer.yaml` (File name does not matter)

```yaml
apiVersion: giantswarm.io/v1
kind: releaseNotesURLAnnotationTransformer
metadata:
  name: myReleaesNotesURLAnnotationTransformer
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

Make the plugin binary available to kustomize:

```bash
go get github.com/giantswarm/kustomize-plugin-releasenotesurlannotationtransformer
mkdir -p "${XDG_CONFIG_HOME:-$HOME/.cache}/kustomize/plugin/giantswarm.io/v1/releasenotesurlannotationtransformer"
cp "$(go env GOPATH)/bin/kustomize-plugin-releasenotesurlannotationtransformer" "${XDG_CONFIG_HOME:-$HOME/.cache}/kustomize/plugin/giantswarm.io/v1/releasenotesurlannotationtransformer/releaseNotesURLAnnotationTransformer"
```

When executing `kustomize build` add the `--enable_alpha_plugins` flag.
