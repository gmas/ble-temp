package dal

import (
	"./"
	"bytes"
	"encoding/json"
	_ "fmt"
	"testing"
)

func TestNewReadoutFromJson(t *testing.T) {
	testReadout := &dal.Readout{"aaaaa-bbbbb", "2015-04-05T23:14:36.865Z", 0.2, 0.0}
	t.Logf("%s\n", testReadout)
	//t.Errorf("testReadout: %v\n", testReadout)

	var testJsonBuffer bytes.Buffer
	jsonEnc := json.NewEncoder(&testJsonBuffer)
	jsonEnc.Encode(testReadout)
	testJson := testJsonBuffer.String()
	t.Logf("%v\n", testJson)

	readout, err := dal.NewReadoutFromJson([]byte(testJson))
	t.Logf("got readout: %v", readout)

	if err != nil {
		t.Errorf("Failed to read from JSON")
	}

	if *testReadout != *readout {
		t.Logf("Expected %v == %v", testReadout, readout)
		t.Error("Expected did not match value")
	}

}
