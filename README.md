# `pnp-go`

This work is heavily inspired by [yarnpkg/pnp-rs](https://github.com/yarnpkg/pnp-rs).  
Thanks to the pnp-rs maintainers for the original implementation.

This crate implements the Yarn Plug'n'Play [resolution algorithms](https://yarnpkg.com/advanced/pnp-spec) for Go so that it can be easily reused within Go-based tools. It also includes utilities allowing to transparently read files from within zip archives.

## Resolution

```go
func Example() {
	manifest, err := loadPNPManifest(".pnp.cjs")
	if err != nil {
		return
	}

	manifestCopy := manifest
	host := ResolutionHost{
		FindPnPManifest: func(_ string) (*Manifest, error) {
			m := manifestCopy
			return &m, nil
		},
	}

	config := ResolutionConfig{
		Host: host,
	}

	parent := filepath.FromSlash("/path/to/index.js")
	resolution, err := resolveToUnqualified("lodash/cloneDeep", parent, config)

	if err != nil {
		return
	}

	switch resolution.Kind {
	case ResolutionResolved:
		// path  = "/path/to/lodash.zip"
		// subpath = "cloneDeep"
		_ = resolution.Path
		_ = resolution.Subpath
	case ResolutionSkipped:
		// This is returned when the PnP resolver decides that it shouldn't
		// handle the resolution for this particular specifier. In that case,
		// the specifier should be forwarded to the default resolver.
	}
}
```