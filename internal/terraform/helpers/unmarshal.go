package helpers

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UnmarshalVyos takes unmarshalled json data and a Terraform resource model pointer and populates it with the data
func UnmarshalVyos(ctx context.Context, data map[string]any, value VyosResourceDataModel) error {
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
		if flags["self-id"].(bool) || flags["parent-id"].(bool) {
			fmt.Printf("\tNot configuring field: %s\n", fName)
			tflog.Debug(ctx, "Not configuring field", map[string]interface{}{"field-name": fName})
			continue
		}

		if flags["child"].(bool) {
			if fieldName, ok := flags["name"].(string); ok {
				if data, ok := data[fieldName]; ok {
					if data != nil {
						tfValue := basetypes.NewBoolValue(true)
						tfValueRefection := reflect.ValueOf(tfValue)
						fValue.Set(tfValueRefection)
					}
				}
			}

			continue
		}

		switch fValue.Interface().(type) {
		case basetypes.StringValue:
			var tfval basetypes.StringValuable

			if !KeyInMap(flags["name"].(string), data) {
				fmt.Printf("\tNo data for field: %s, setting to empty string\n", fName)
				tflog.Debug(ctx, "No data for field, setting to empty string", map[string]interface{}{"field-name": fName})

				tfval = basetypes.NewStringNull()
			} else {
				dataValue := data[flags["name"].(string)]

				fmt.Printf("\tUnmarshalling String Field: %s\t%s=%s\n", fName, flags["name"].(string), dataValue)
				tflog.Debug(ctx, "Unmarshalling String Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

				tfval = basetypes.NewStringValue(dataValue.(string))
			}
			tfValueRefection := reflect.ValueOf(tfval)
			fValue.Set(tfValueRefection)

		case basetypes.BoolValue:
			fmt.Printf("\tUnmarshalling Bool Field: %s\n", fName)
			tflog.Debug(ctx, "Unmarshalling Bool Field", map[string]interface{}{"field-name": fName})
			tfValue := basetypes.NewBoolValue(KeyInMap(flags["name"].(string), data))
			tfValueRefection := reflect.ValueOf(tfValue)
			fValue.Set(tfValueRefection)
		case basetypes.NumberValue:
			var tfval basetypes.NumberValuable

			if !KeyInMap(flags["name"].(string), data) {
				fmt.Printf("\tNo data for field: %s, skipping\n", fName)
				tflog.Debug(ctx, "No data for field, skipping", map[string]interface{}{"field-name": fName})

				tfval = basetypes.NewNumberNull()
			} else {
				dataValue := data[flags["name"].(string)]

				fmt.Printf("\tUnmarshalling Number Field: %s\t%s=%s\n", fName, flags["name"].(string), dataValue)
				tflog.Debug(ctx, "Unmarshalling Number Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

				var dataValueF float64
				switch dataValue := dataValue.(type) {
				case string:
					var err error
					dataValueF, err = strconv.ParseFloat(dataValue, 64)
					if err != nil {
						tflog.Error(ctx, "cannot convert string to float64", map[string]interface{}{"value": dataValue})
						return fmt.Errorf("cannot convert string '%#v' to float64", dataValue)
					}
				case int:
					dataValueF = float64(dataValue)
				case int8:
					dataValueF = float64(dataValue)
				case int16:
					dataValueF = float64(dataValue)
				case int32:
					dataValueF = float64(dataValue)
				case int64:
					dataValueF = float64(dataValue)
				case uint:
					dataValueF = float64(dataValue)
				case uint8:
					dataValueF = float64(dataValue)
				case uint16:
					dataValueF = float64(dataValue)
				case uint32:
					dataValueF = float64(dataValue)
				case uint64:
					dataValueF = float64(dataValue)
				case float32:
					dataValueF = float64(dataValue)
				case float64:
					dataValueF = float64(dataValue)
				default:
					tflog.Error(ctx, "cannot convert UNHANDLED to float64", map[string]interface{}{"value": dataValue})
					return fmt.Errorf("cannot convert UNHANDLED '%#v' to float64", dataValue)
				}

				bf := big.NewFloat(dataValueF)
				tfval = basetypes.NewNumberValue(bf)
			}
			tfValueRefection := reflect.ValueOf(tfval)
			fValue.Set(tfValueRefection)
			tflog.Trace(ctx, "Setting Value", map[string]interface{}{"tfval": tfval, "tfValueRefection": tfValueRefection, "fValue": fValue})

		case basetypes.ListValue:
			var tfval basetypes.ListValuable
			tfSchemaName := strings.Split(typeReflection.Field(i).Tag.Get("tfsdk"), ",")[0]
			schemaAttr := value.ResourceSchemaAttributes()[tfSchemaName]
			listAttr := schemaAttr.(schema.ListAttribute)
			elemType := listAttr.ElementType

			if !KeyInMap(flags["name"].(string), data) {
				fmt.Printf("\tNo data for field: %s, skipping\n", fName)
				tflog.Debug(ctx, "No data for field, skipping", map[string]interface{}{"field-name": fName})

				tfval = basetypes.NewListNull(elemType)
			} else {
				dataValue := data[flags["name"].(string)]

				fmt.Printf("\tUnmarshalling List Field: %s\t%s=%s\n", fName, flags["name"].(string), dataValue)
				tflog.Debug(ctx, "Unmarshalling List Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

				if reflect.TypeOf(dataValue).Kind() != reflect.Slice {
					fmt.Printf("\tdataValue is not a slice, wrapping the value in list: %#v\n", dataValue)
					tflog.Debug(ctx, "dataValue is not a slice, wrapping the value in list", map[string]interface{}{flags["name"].(string): dataValue})
					dataValue = []interface{}{dataValue}
				}

				var diags diag.Diagnostics
				tfval, diags = basetypes.NewListValueFrom(ctx, elemType, dataValue)
				if diags != nil {
					return fmt.Errorf("ERROR: %v", diags)
				}
			}
			tflog.Trace(ctx, "setting list value", map[string]interface{}{"tfval": tfval})
			tfValueRefection := reflect.ValueOf(tfval)
			fValue.Set(tfValueRefection)

		default:
			//var tfval VyosResourceDataModel

			if !KeyInMap(flags["name"].(string), data) {
				fmt.Printf("\tNo data for field: %s, skipping\n", fName)
				tflog.Debug(ctx, "No data for field, skipping", map[string]interface{}{"field-name": fName})

				// TODO set nil / null for terraform state on sub-objects
				//tfval = basetypes.NewObjectNull()
				continue
			}
			dataValue := data[flags["name"].(string)]

			fmt.Printf("\tUnmarshalling Struct Field: %s\t%s=%s\n", fName, flags["name"].(string), dataValue)
			tflog.Debug(ctx, "Unmarshalling Struct Field", map[string]interface{}{"field-name": fName, flags["name"].(string): dataValue})

			subSctructReflection := reflect.New(fType.Elem())
			subStructIf := subSctructReflection.Interface()
			subStruct := subStructIf.(VyosResourceDataModel)
			err := UnmarshalVyos(ctx, dataValue.(map[string]any), subStruct)
			if err != nil {
				return err
			}

			fValue.Set(subSctructReflection)
		}
	}

	// Return no error
	return nil
}
