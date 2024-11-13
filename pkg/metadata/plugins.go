package metadata

import (
	"strings"

	"k8s.io/apimachinery/pkg/types"
)

// ExtractPluginsWithNamespaces extracts plugin namespaced names from the given object's
// konghq.com/plugins annotation.
// This function trims the whitespace from the plugin names.
//
// For example, for KongConsumer in namespace default, having the "konghq.com/plugins"
// annotation set to "p1,p2" this will return []string{"default/p1", "default/p2"}
func ExtractPluginsWithNamespaces(obj ObjectWithAnnotationsAndNamespace) []string {
	return extractPlugins(obj, nsOptWithNamespace)
}

// ExtractPlugins extracts plugin names from the given object's
// konghq.com/plugins annotation.
// This function trims the whitespace from the plugin names.
//
// For example, for KongConsumer in namespace default, having the "konghq.com/plugins"
// annotation set to "p1,p2" this will return []string{"p1", "p2"}
func ExtractPlugins(obj ObjectWithAnnotationsAndNamespace) []string {
	return extractPlugins(obj, nsOptWithoutNamespace)
}

// ExtractPluginsNamespacedNames extracts plugin namespaced names from the given object's
// konghq.com/plugins annotation. Plugins can optionally specify the namespace using the
// "<namespace>:<plugin-name>" format.
// This function trims the whitespace from the plugin names.
//
// For example, for an object having the "konghq.com/plugins" annotation set to "default:p1,p2"
// this will return:
//
//	 []types.NamespacedName{
//			types.NamespacedName{Namespace: "default", Name: "p1"},
//			types.NamespacedName{Namespace: "", Name: "p2"},
//		}
func ExtractPluginsNamespacedNames(obj ObjectWithAnnotationsAndNamespace) []types.NamespacedName {
	ann, ok := obj.GetAnnotations()[AnnotationKeyPlugins]
	if !ok || len(ann) == 0 {
		return nil
	}

	ann = strings.Trim(ann, ",")
	commas := strings.Count(ann, ",")
	plugins := make([]types.NamespacedName, 0, commas+1)

	for i, idx := 0, 0; i < commas+1; i++ {
		idx = strings.IndexByte(ann, ',')
		if idx == -1 {
			idx = len(ann)
		}

		s := strings.TrimSpace(ann[:idx])
		if s == "" {
			if idx == len(ann) {
				break
			}
			ann = ann[idx+1:]
			continue
		}

		idxColon := strings.Index(s, ":")
		if idxColon == len(s)-1 || idxColon == 0 {
			// invalid plugin name or namespace
			if idx >= len(ann) {
				break
			}
			ann = ann[idx+1:]
			continue
		}

		plugin := types.NamespacedName{}
		if idxColon != -1 {
			plugin.Namespace = strings.TrimSpace(s[0:idxColon])
			plugin.Name = strings.TrimSpace(s[idxColon+1:])
		} else {
			plugin.Name = strings.TrimSpace(s)
		}
		plugins = append(plugins, plugin)
		if idx >= len(ann) {
			break
		}
		ann = ann[idx+1:]
	}
	return plugins
}

type extractPluginsNamespaceOpt byte

const (
	nsOptWithNamespace extractPluginsNamespaceOpt = iota
	nsOptWithoutNamespace
)

func extractPlugins(obj ObjectWithAnnotationsAndNamespace, nsOpt extractPluginsNamespaceOpt) []string {
	if obj == nil {
		return nil
	}

	ann, ok := obj.GetAnnotations()[AnnotationKeyPlugins]
	if !ok || len(ann) == 0 {
		return nil
	}

	namespace := obj.GetNamespace()
	ann = strings.Trim(ann, ",")
	commas := strings.Count(ann, ",")
	plugins := make([]string, 0, commas+1)

	for i, idx := 0, 0; i < commas+1 && idx <= len(ann); i++ {
		idx = strings.IndexByte(ann, ',')
		if idx == -1 {
			idx = len(ann)
		}

		s := strings.TrimSpace(ann[:idx])
		if s == "" {
			if idx == len(ann) {
				break
			}
			ann = ann[idx+1:]
			continue
		}

		v := s
		if nsOpt == nsOptWithNamespace {
			v = namespace + "/" + s
		}
		plugins = append(plugins, v)
		if idx >= len(ann) {
			break
		}
		ann = ann[idx+1:]
	}

	return plugins
}
