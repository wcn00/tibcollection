package tibcollection

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/util"
)

const (
	GET    = "get"
	APPEND = "append"
	DELETE = "delete"

	//model constants
	OPERATION  = "operation"
	KEY        = "key"
	OBJ        = "object"
	COLLECTION = "collection"
	SIZE       = "size"
)

type Collection struct {
	colmap map[string][]interface{}
}

var col *Collection

func init() {
	col = new(Collection)
	col.colmap = make(map[string][]interface{})
}

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata  *activity.Metadata
	generator *util.Generator
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (myActivity *MyActivity) Metadata() *activity.Metadata {
	return myActivity.metadata
}

func IsZero(v interface{}) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.String:
		return len(v.([]interface{})) == 0
	case reflect.Bool:
		return v.(bool)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.(int) == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.(uint) == 0
	case reflect.Float32, reflect.Float64:
		return v.(float32) == 0
	case reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return false //v.IsNil()
	}
	return false
}

var collectionCacheMutex sync.Mutex

// Eval implements activity.Activity.Eval
func (myActivity *MyActivity) Eval(context activity.Context) (done bool, err error) {
	collectionCacheMutex.Lock()
	defer collectionCacheMutex.Unlock()

	op, ok := context.GetSetting(OPERATION)
	if !ok {
		return false, fmt.Errorf("Operation not specified")
	}
	key := context.GetInput(KEY)
	obj := context.GetInput(OBJ)

	switch op {
	case APPEND:
		if key == nil || len(key.(string)) == 0 {
			key, err = myActivity.newKey()
			if err != nil {
				return false, fmt.Errorf("Append with no key failed to create dynamic key for reason [%s]", err)
			}
		}
		if obj != nil && !reflect.ValueOf(obj).IsNil() { // len(obj.(map[string]interface{})) > 0 {
			col.colmap[key.(string)] = append(col.colmap[key.(string)], obj)
		}

		context.SetOutput(SIZE, len(col.colmap[key.(string)]))
		context.SetOutput(KEY, key)
		return true, nil

	case GET:
		if key == nil {
			return false, fmt.Errorf("Get operation called with no key")
		}
		array, ok := col.colmap[key.(string)]
		if !ok {
			return false, fmt.Errorf("Get operation called for invalid key: %s", key.(string))
		}
		context.SetOutput(SIZE, len(col.colmap[key.(string)]))
		context.SetOutput(KEY, key)
		context.SetOutput(COLLECTION, array)
		return true, nil

	case "delete":
		if key == nil {
			return false, fmt.Errorf("Get operation called with no key")
		}
		delete(col.colmap, key.(string))
		context.SetOutput(SIZE, -1)
		return true, nil
	default:
		return false, fmt.Errorf("Get operation called with invalid operation [%s]", op)
	}
}

func (myActivity *MyActivity) newKey() (res string, err error) {
	if myActivity.generator == nil {
		myActivity.generator, err = util.NewGenerator()
		if err != nil {
			return "", fmt.Errorf("Failed to generate a dynamic key for collection for reason [%s]", err)
		}
	}
	return myActivity.generator.NextAsString(), nil
}
