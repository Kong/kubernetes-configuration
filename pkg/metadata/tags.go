package metadata

import (
	"strings"
)

type ObjectWithAnnotations interface {
	GetAnnotations() map[string]string
}

// ExtractTags extracts a set of tags from a comma-separated string.
// Copy pasted from: https://github.com/Kong/kubernetes-ingress-controller/blob/eb80ec2c58f4d53f8c6d7c997bcfb1f334b801e1/internal/annotations/annotations.go#L407-L416
func ExtractTags(obj ObjectWithAnnotations) []string {
	anns := obj.GetAnnotations()
	val := anns[annotationPrefix+tagsKey]
	// If the annotation is not present, the map provides an empty value,
	// and splitting that will create a slice containing a single empty string tag.
	// These aren't valid, hence this special case.
	if len(val) == 0 {
		return nil
	}
	return strings.Split(val, ",")
}
