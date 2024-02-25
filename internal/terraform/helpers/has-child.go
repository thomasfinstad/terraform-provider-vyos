package helpers

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// HasChild takes a resource model and checks if it has children configured
// This is useful to halt deleting of resources
func HasChild(ctx context.Context, value VyosResourceDataModel) bool {
	valueReflection := reflect.ValueOf(value).Elem()
	typeReflection := valueReflection.Type()

	for i := 0; i < valueReflection.NumField(); i++ {
		fName := typeReflection.Field(i).Name
		fType := typeReflection.Field(i).Type
		fValue := valueReflection.Field(i)
		fTags := strings.Split(typeReflection.Field(i).Tag.Get("vyos"), ",")
		tflog.Debug(ctx, "processing field", map[string]interface{}{"Field": fName, "Type": fType, "Tags": fTags})
		fmt.Printf("Field: %s\tType: %s\tTags: %v\n", fName, fType, fTags)

		// Set flags based on tags, first tag must be the vyos field name, the rest are bools with default of false
		// TODO create struct of valid options
		flags := map[string]any{
			"name":      fTags[0],
			"self-id":   false,
			"parent-id": false,
			"omitempty": false,
			"child":     false,
		}
		for _, tag := range fTags[1:] {
			flags[tag] = true
		}
		if !flags["child"].(bool) {
			fmt.Printf("\tNot a child, skipping field: %s\n", fName)
			continue
		}

		if fValue.Equal(reflect.ValueOf(true)) {
			return true
		}
	}

	// No child was configured
	return false
}