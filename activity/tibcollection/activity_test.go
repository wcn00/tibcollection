package tibcollection

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getObj() interface{} {
	objectJSON := []byte(`{
		"obj":{
		"name":"walter",
		"age" : 45 }
		}`)
	obj := make(map[string]interface{})
	json.Unmarshal(objectJSON, &obj)
	return obj

}
func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetSetting(OPERATION, APPEND)
	tc.SetInput(KEY, "blarg")
	act.Eval(tc)
	key := tc.GetOutput(KEY).(string)
	size := tc.GetOutput(SIZE).(int)
	if !assert.Equal(t, 0, size) {
		t.Errorf("Activity should have returned size 0")
		t.Fail()
	}
	if !assert.NotNil(t, key) {
		t.Errorf("Activity should have returned a key")
		t.Fail()
	}

	//check result attr
}

func TestEvalEndToEnd(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetSetting(OPERATION, APPEND)
	tc.SetInput(KEY, "blarg")

	ok, err := act.Eval(tc)
	assert.Nil(t, err)
	if err != nil {
		t.Errorf("Could not execute activty:  %s", err)
		t.Fail()
	}
	if !ok {
		t.Errorf("Activity returned false")
	}
	key := tc.GetOutput(KEY).(string)
	size := tc.GetOutput(SIZE).(int)
	if !assert.Equal(t, 0, size) {
		t.Errorf("Activity should have returned size 0")
		t.Fail()
	}
	if !assert.NotNil(t, key) {
		t.Errorf("Activity should have returned a key")
		t.Fail()
	}

	//Append an obj
	tc.SetInput(OBJ, getObj)
	tc.SetInput(KEY, key)
	ok, err = act.Eval(tc)
	assert.Nil(t, err)
	if err != nil {
		t.Errorf("Could not execute activty:  %s", err)
		t.Fail()
	}
	if !ok {
		t.Errorf("Activity returned false")
	}
	size = tc.GetOutput("size").(int)
	if !assert.Equal(t, 1, size) {
		t.Errorf("Activity should have returned size 1")
		t.Fail()
	}
	if !assert.Equal(t, key, tc.GetInput("key").(string)) {
		t.Errorf("Activity should have returned a key")
		t.Fail()
	}

	//Append second obj
	tc.SetInput(OBJ, getObj)
	tc.SetInput(KEY, key)
	ok, err = act.Eval(tc)
	assert.Nil(t, err)
	if err != nil {
		t.Errorf("Could not execute activty:  %s", err)
		t.Fail()
	}
	if !ok {
		t.Errorf("Activity returned false")
	}
	size = tc.GetOutput("size").(int)
	if !assert.Equal(t, 2, size) {
		t.Errorf("Activity should have returned size 2")
		t.Fail()
	}
	if !assert.Equal(t, key, tc.GetInput("key").(string)) {
		t.Errorf("Activity should have returned a key")
		t.Fail()
	}

	// GET
	tc.SetSetting(OPERATION, GET)
	tc.SetInput(KEY, key)
	ok, err = act.Eval(tc)
	assert.Nil(t, err)
	if err != nil {
		t.Errorf("Could not execute activty:  %s", err)
		t.Fail()
	}
	if !ok {
		t.Errorf("Activity returned false")
	}
	collection := tc.GetOutput("collection")
	if !assert.Equal(t, 2, len(collection.([]interface{}))) {
		t.Errorf("Returned collection length should be 2 not %d", len(collection.([]interface{})))
	}
	size = tc.GetOutput("size").(int)
	if !assert.Equal(t, 2, size) {
		t.Errorf("Activity should have returned size 2")
		t.Fail()
	}

	//Delete the collection
	tc.SetSetting(OPERATION, DELETE)
	tc.SetInput(KEY, key)
	ok, err = act.Eval(tc)
	assert.Nil(t, err)
	if err != nil {
		t.Errorf("Could not execute activty:  %s", err)
		t.Fail()
	}
	if !ok {
		t.Errorf("Activity returned false")
	}
	size = tc.GetOutput("size").(int)
	if !assert.Equal(t, -1, size) {
		t.Errorf("Activity should have returned size -1")
		t.Fail()
	}
}
