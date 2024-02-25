package helpers

import (
	"context"
	"fmt"
	"slices"
	"strconv"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GenerateVyosOps takes a resource data model and converts it to vyos api operation compatible list of lists of strings
func GenerateVyosOps(ctx context.Context, vyosPath []string, vyosData map[string]interface{}) [][]string {
	tflog.Trace(ctx, "GenerateVyosOps Input", map[string]interface{}{"vyosPath": vyosPath, "vyosData": vyosData})
	vyosOps := iron(ctx, vyosPath, vyosData)
	tflog.Trace(ctx, "GenerateVyosOps return", map[string]interface{}{"vyosOps": vyosOps})
	return vyosOps
}

// Do not understand why I must clone here, but it seems the slice is some times a pointer, so it duplicates elements of the last one in the slice instead of adding all the uniqeue values otherwise
// Is this because it comes from the marshalling function?
// TODO deep investigation into origins of the bug requiring the cloning

func iron(ctx context.Context, vyosPath []string, values map[string]interface{}) [][]string {

	ret := [][]string{}

	for key, value := range values {
		fmt.Printf("Path: %#v, Key: %#v, Value: %#v\n", vyosPath, key, value)

		cVyosPath := append(vyosPath, key)
		switch value := value.(type) {

		// LeafNodes
		case string:
			fmt.Printf("type: string\n")
			tflog.Trace(ctx, "ironing string value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			val := append(cVyosPath, value)
			val = slices.Clone(val)
			tflog.Trace(ctx, "appending to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
			fmt.Printf("1 ret: %#v, val: %#v\n", ret, val)
			ret = append(ret, val)
			fmt.Printf("2 ret: %#v, val: %#v\n", ret, val)
		// LeafNodes
		case bool:
			fmt.Printf("type: bool\n")
			tflog.Trace(ctx, "ironing bool value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			if value {
				val := append(cVyosPath, "{}")
				tflog.Trace(ctx, "appending to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
				ret = append(ret, val)
			} else {
				tflog.Trace(ctx, "NOT appending to ret because bool is false", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret)})
			}
		// LeafNodes
		case float64:
			fmt.Printf("type: float64\n")
			tflog.Trace(ctx, "ironing float value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			val := append(cVyosPath, strconv.FormatFloat(value, 'f', -1, 64))
			val = slices.Clone(val)
			tflog.Trace(ctx, "appending to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
			ret = append(ret, val)
		// LeafNodes multi value
		case []string:
			fmt.Printf("type: []string]\n")
			tflog.Trace(ctx, "ironing slice of strings value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			for _, element := range value {
				val := append(cVyosPath, element)
				val = slices.Clone(val)
				tflog.Trace(ctx, "appending to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
				ret = append(ret, val)
				tflog.Trace(ctx, "appended to ret", map[string]interface{}{"ret": fmt.Sprintf("%#v", ret), "val": fmt.Sprintf("%#v", val)})
			}

		// TagNodes and Nodes
		case map[string]interface{}:
			fmt.Printf("type: map[string]interface{}\n")
			tflog.Trace(ctx, "ironing nested map value", map[string]interface{}{"current-vyos-path": cVyosPath, "type": fmt.Sprintf("%T", value), "value": fmt.Sprintf("%#v", value)})
			tflog.Trace(ctx, "recursing for ret", map[string]interface{}{"cVyosPath": fmt.Sprintf("%#v", cVyosPath)})
			val := iron(ctx, cVyosPath, value)
			val = slices.Clone(val)
			fmt.Printf("3 ret: %#v, val: %#v\n", ret, val)
			ret = append(
				ret,
				val...,
			)
			fmt.Printf("4 ret: %#v, val: %#v\n", ret, val)

		// ERROR
		default:
			tflog.Error(ctx, "No handling of value type", map[string]interface{}{"type": fmt.Sprintf("%T", value), "key": key, "cVyosPath": cVyosPath})
			panic("unhandled type see last log entry")
		}

		fmt.Printf("CurrentReturnValue: %#v\n", ret)
	}

	tflog.Trace(ctx, "ironing result", map[string]interface{}{"result": fmt.Sprintf("%#v", ret)})
	fmt.Printf("returning: %#v\n", ret)
	return ret
}