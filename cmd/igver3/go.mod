module igver3

go 1.24.0

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/inspektor-gadget/inspektor-gadget v0.40.0
	k8s.io/client-go v0.33.1
)

replace github.com/inspektor-gadget/inspektor-gadget => github.com/alban/inspektor-gadget v0.40.1-0.20250519094622-ebaae6b65775
