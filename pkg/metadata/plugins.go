package metadata

import (
	"strings"
)

// ExtractPluginsWithNamespaces extracts plugin namespaced names from the given object's
// konghq.com/plugins annotation.
// This function trims the whitespace from the plugin names.
//
// For example, for KongConsumer in namespace default, having the "konghq.com/plugins"
// annotation set to "p1,p2" this will return []string{"default/p1", "default/p2"}
func ExtractPluginsWithNamespaces(obj ObjectWithAnnotationsAndNamespace) []string {
	ann, ok := obj.GetAnnotations()[AnnotationKeyPlugins]
	if !ok || len(ann) == 0 {
		return nil
	}

	namespace := obj.GetNamespace()
	split := strings.Split(ann, ",")
	ret := make([]string, 0, len(split))
	for _, p := range split {
		trimmed := strings.TrimSpace(p)
		if trimmed == "" {
			continue
		}
		ret = append(ret, namespace+"/"+trimmed)
	}
	return ret
}

// ExtractPlugins extracts plugin names from the given object's
// konghq.com/plugins annotation.
// This function trims the whitespace from the plugin names.
//
// For example, for KongConsumer in namespace default, having the "konghq.com/plugins"
// annotation set to "p1,p2" this will return []string{"p1", "p2"}
func ExtractPlugins(obj ObjectWithAnnotations) []string {
	ann, ok := obj.GetAnnotations()[AnnotationKeyPlugins]
	if !ok || len(ann) == 0 {
		return nil
	}

	split := strings.Split(ann, ",")
	ret := make([]string, 0, len(split))
	for _, p := range split {
		trimmed := strings.TrimSpace(p)
		if trimmed == "" {
			continue
		}
		ret = append(ret, trimmed)
	}
	return ret
}
