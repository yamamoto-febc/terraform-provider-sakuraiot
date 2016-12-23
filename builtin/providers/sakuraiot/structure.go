package sakuraiot

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
)

// Takes the result of flatmap.Expand for an array of strings
// and returns a []string
func expandStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		vs = append(vs, string(v.(string)))
	}
	return vs
}

// Takes the result of schema.Set of strings and returns a []string
func expandStringSet(configured *schema.Set) []string {
	return expandStringList(configured.List())
}

// Takes list of pointers to strings. Expand to an array
// of raw strings and returns a []interface{}
// to keep compatibility w/ schema.NewSetschema.NewSet
func flattenStringList(list []string) []interface{} {
	vs := make([]interface{}, 0, len(list))
	for _, v := range list {
		vs = append(vs, v)
	}
	return vs
}

func forceString(target interface{}) string {
	if target == nil {
		return ""
	}

	return target.(string)
}

func toInt64ID(id string) int64 {
	v, _ := strconv.ParseInt(id, 10, 64)
	return v
}

func toStrID(id int64) string {
	return fmt.Sprintf("%d", id)
}
