// Code generated by ros-gen-go.
// source: SmachContainerStatus.msg
// DO NOT EDIT!
package smach_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ppg/rosgo/msgs/std_msgs"
	"github.com/ppg/rosgo/ros"
)

type _MsgSmachContainerStatus struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgSmachContainerStatus) Text() string {
	return t.text
}

func (t *_MsgSmachContainerStatus) Name() string {
	return t.name
}

func (t *_MsgSmachContainerStatus) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgSmachContainerStatus) NewMessage() ros.Message {
	m := new(SmachContainerStatus)

	return m
}

var (
	MsgSmachContainerStatus = &_MsgSmachContainerStatus{
		`Header header

# The path to this node in the server
string path

# The initial state description
# Effects an arc from the top state to each one
string[] initial_states

# The current state description
string[] active_states

# A pickled user data structure
# i.e. the UserData's internal dictionary
string local_data

# Debugging info string
string info
`,
		"smach_msgs/SmachContainerStatus",
		"2a07c0f13a55f9c0e61b861363c73741",
	}
)

type SmachContainerStatus struct {
	Header        std_msgs.Header
	Path          string
	InitialStates []string
	ActiveStates  []string
	LocalData     string
	Info          string
}

func (m *SmachContainerStatus) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Path); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.InitialStates)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.InitialStates {
		if err = ros.SerializeMessageField(w, "string", &elem); err != nil {
			return err
		}
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.ActiveStates)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.ActiveStates {
		if err = ros.SerializeMessageField(w, "string", &elem); err != nil {
			return err
		}
	}

	if err = ros.SerializeMessageField(w, "string", &m.LocalData); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Info); err != nil {
		return err
	}

	return
}

func (m *SmachContainerStatus) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Path
	if err = ros.DeserializeMessageField(r, "string", &m.Path); err != nil {
		return err
	}

	// InitialStates
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for InitialStates: %s", err)
		}
		m.InitialStates = make([]string, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "string", &m.InitialStates[i]); err != nil {
				return err
			}
		}
	}

	// ActiveStates
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for ActiveStates: %s", err)
		}
		m.ActiveStates = make([]string, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "string", &m.ActiveStates[i]); err != nil {
				return err
			}
		}
	}

	// LocalData
	if err = ros.DeserializeMessageField(r, "string", &m.LocalData); err != nil {
		return err
	}

	// Info
	if err = ros.DeserializeMessageField(r, "string", &m.Info); err != nil {
		return err
	}

	return
}