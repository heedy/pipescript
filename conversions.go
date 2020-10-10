package pipescript

import (
	"encoding/json"
	"strconv"
)

// Int takes an interface that was unmarshalled with the json package,
// and returns it as an int, or returns false
func Int(v interface{}) (int64, bool) {
	switch n := v.(type) {
	case float64:
		res := int64(n)
		return res, float64(res) == n
	case int64:
		return n, true
	case int:
		return int64(n), true
	case bool:
		if n {
			return 1, true
		}
		return 0, true
	default:
		return 0, false
	}
}

// Float takes an interface that was marshalled with the json package
// and returns it as a float, or gives an error
func Float(v interface{}) (float64, bool) {
	switch n := v.(type) {
	case float64:
		return n, true
	case int64:
		return float64(n), true
	case int:
		return float64(n), true
	case bool:
		if n {
			return 1.0, true
		}
		return 0, true
	default:
		return 0, false
	}
}

// String extracts the data as a string, or returns an error
func String(v interface{}) (string, bool) {
	switch n := v.(type) {
	case string:
		return n, true
	default:
		return "", false
	}
}

func ToString(v interface{}) string {
	if v == nil {
		return "null"
	}
	switch n := v.(type) {
	case string:
		return n
	case int:
		return strconv.Itoa(n)
	case int64:
		return strconv.FormatInt(n, 10)
	case bool:
		if n {
			return "true"
		}
		return "false"
	case float64:
		return strconv.FormatFloat(n, 'g', -1, 64)
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

func Bool(v interface{}) (bool, bool) {
	switch n := v.(type) {
	case bool:
		return n, true
	case float64:
		if n == 1.0 {
			return true, true
		} else if n == 0 {
			return false, true
		}
		return false, false
	case int64:
		if n == 1 {
			return true, true
		} else if n == 0 {
			return false, true
		}
		return false, false
	case int:
		if n == 1 {
			return true, true
		} else if n == 0 {
			return false, true
		}
		return false, false
	default:
		return false, false
	}

}

func FloatNoBool(v interface{}) (float64, bool) {
	switch n := v.(type) {
	case float64:
		return n, true
	case float32:
		return float64(n), true
	case int64:
		return float64(n), true
	case int:
		return float64(n), true
	default:
		return 0, false
	}
}

func IntNoBool(v interface{}) (int64, bool) {
	switch n := v.(type) {
	case float64:
		res := int64(n)
		return res, float64(res) == n
	case float32:
		res := int64(n)
		return res, float32(res) == n
	case int64:
		return n, true
	case int:
		return int64(n), true
	default:
		return 0, false
	}
}

func Equal(o1, o2 interface{}) bool {
	if o1 == nil || o2 == nil {
		return o1 == o2
	}
	switch v1 := o1.(type) {
	case string:
		v2, ok := o2.(string)
		return ok && v1 == v2
	case float64:
		v2, ok := FloatNoBool(o2)
		return ok && v1 == v2
	case float32:
		v2, ok := FloatNoBool(o2)
		return ok && float64(v1) == v2
	case int:
		v2, ok := IntNoBool(o2)
		return ok && int64(v1) == v2
	case int64:
		v2, ok := IntNoBool(o2)
		return ok && v1 == v2
	case bool:
		v2, ok := o2.(bool)
		return ok && v1 == v2
	case map[string]interface{}:
		v2, ok := o2.(map[string]interface{})
		if !ok || len(v1) != len(v2) {
			return false
		}
		for k, vv1 := range v1 {
			vv2, ok := v2[k]
			if !ok || !Equal(vv1, vv2) {
				return false
			}
		}
		return true
	case []interface{}:
		v2, ok := o2.([]interface{})
		if !ok || len(v1) != len(v2) {
			return false
		}
		for i := range v1 {
			if !Equal(v1[i], v2[i]) {
				return false
			}
		}
		return true
	default:
		panic("Unknown interface type for Equal")
	}
}

/*
//Equal attempts to check equality between two interfaces. If the values
//are not directly comparable thru DeepEqual, tries to do a "duck" comparison.
//	true true -> true
//	"true" true -> true
//	1 true -> true
//	1.0 1 -> true
//	1.345 "1.345" -> true
//	50.0 true -> true
//	0.0 false -> true
func Equal(arg1 interface{}, arg2 interface{}) bool {
	if arg1 == arg2 {
		return true
	}
	// Comparing things to nil should work
	if arg1 == nil || arg2 == nil {
		// The first if statement should have been equal if they are both not nil
		return false
	}

	s1, ok1 := arg1.(string)
	s2, ok2 := arg2.(string)
	if ok1 && ok2 {
		return s1 == s2
	}

	// Neither is a string. Let's see if one of them is a number.
	// That way we know we can compare as numbers
	f1, ok := Float(arg1)
	if ok {
		f2, ok := Float(arg2)
		if ok {
			if math.IsNaN(f1) && math.IsNaN(f2) {
				return true
			}
			return f1 == f2
		}

	}

	// using this slows doewn the entire function. It is the same issue
	// that required forking parseFloat to be faster.
	return reflect.DeepEqual(arg1, arg2)
}
*/
