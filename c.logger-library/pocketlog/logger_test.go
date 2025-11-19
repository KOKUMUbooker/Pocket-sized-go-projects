package pocketlog_test

// Go normally doesn't allow multiple packages defined in the same directory, but
// defining a testing package inside an existing package directory is allowed
// This is done to keep the testing file close to whatever file its testing

import (
	"logger/pocketlog"
	"testing"
)

// testWriter is a struct that implements io.Writer.
// We use it to validate that we can write to a specific output.
type testWriter struct {
	contents string
}
// Write implements the io.Writer interface.
func (tw *testWriter) Write(p []byte) (n int, err error) { 
	tw.contents = tw.contents + string(p) 
	return len(p), nil
}

const ( 
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage = "And keep invention in a noted weed,"
	errorMessage = "That every word doth almost tell my name,"
)

// func ExampleLogger_Debugf(){
// 	logger := pocketlog.New(pocketlog.LevelDebug,pocketlog.WithOutput(os.Stdout));
// 	logger.Debugf("Hello world");
// 	// Output: Hello world
// }

func  TestLogger_DebugfInfofErrorf(t *testing.T) {
	type testCase struct {
		level pocketlog.Level
		expected string
	}
	concat := pocketlog.GetFormatConcat;

	tt := map[string]testCase{
		"debug": {
			level: pocketlog.LevelDebug,
			expected: concat(pocketlog.LevelDebug,debugMessage) +
					  concat(pocketlog.LevelDebug,infoMessage) +
					  concat(pocketlog.LevelDebug,errorMessage),
		},
		"info": {
			level: pocketlog.LevelInfo,
			expected: concat(pocketlog.LevelInfo,infoMessage) +
					  concat(pocketlog.LevelInfo,errorMessage),
		}, 
		"error": {
			level: pocketlog.LevelError,
			expected: concat(pocketlog.LevelError,errorMessage),
		},
	}
	
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw));

			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)
			if tw.contents != tc.expected {
				t.Errorf("invalid contents, expected %q, got %q", tc.expected,tw.contents)
			}
		});
	}
	}
