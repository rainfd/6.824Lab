package mapreduce

import (
	"fmt"
	"strconv"
)

// Debugging enabled?
const debugEnabled = true

// debug() will only print if debugEnabled is true
func debug(format string, a ...interface{}) (n int, err error) {
	if debugEnabled {
		n, err = fmt.Printf(format, a...)
	}
	return
}

// jobPhase indicates whether a task is scheduled as a map or reduce task.
type jobPhase string

const (
	mapPhase    jobPhase = "mapPhase"
	reducePhase          = "reducePhase"
)

// KeyValue is a type used to hold the key/value pairs passed to the map and
// reduce functions.
type KeyValue struct {
	Key   string
	Value string
}

// reduceName constructs the name of the intermediate file which map task
// <mapTask> produces for reduce task <reduceTask>.
func reduceName(jobName string, mapTask int, reduceTask int) string {
	return "mrtmp." + jobName + "-" + strconv.Itoa(mapTask) + "-" + strconv.Itoa(reduceTask)
}

// mergeName constructs the name of the output file of reduce task <reduceTask>
func mergeName(jobName string, reduceTask int) string {
	return "mrtmp." + jobName + "-res-" + strconv.Itoa(reduceTask)
}

// KeyValues is a type used to sort KeyValue
type KeyValues []*KeyValue

func (kv KeyValues) Len() int { return len(kv) }
func (kv KeyValues) Less(i, j int) bool {
	in, _ := strconv.Atoi(kv[i].Key)
	jn, _ := strconv.Atoi(kv[j].Key)
	return in < jn
}
func (kv KeyValues) Swap(i, j int) { kv[i], kv[j] = kv[j], kv[i] }
