// Code generated by ros-gen-go.
// source: PointHeadActionFeedback.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/ppg/rosgo/msgs/actionlib_msgs"
	"github.com/ppg/rosgo/msgs/std_msgs"
	"github.com/ppg/rosgo/ros"
)

type _MsgPointHeadActionFeedback struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPointHeadActionFeedback) Text() string {
	return t.text
}

func (t *_MsgPointHeadActionFeedback) Name() string {
	return t.name
}

func (t *_MsgPointHeadActionFeedback) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPointHeadActionFeedback) NewMessage() ros.Message {
	m := new(PointHeadActionFeedback)

	return m
}

var (
	MsgPointHeadActionFeedback = &_MsgPointHeadActionFeedback{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

Header header
actionlib_msgs/GoalStatus status
PointHeadFeedback feedback
`,
		"control_msgs/PointHeadActionFeedback",
		"9130b1e1f7b30a434f83565d195ebfb5",
	}
)

type PointHeadActionFeedback struct {
	Header   std_msgs.Header
	Status   actionlib_msgs.GoalStatus
	Feedback PointHeadFeedback
}

func (m *PointHeadActionFeedback) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "actionlib_msgs/GoalStatus", &m.Status); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "PointHeadFeedback", &m.Feedback); err != nil {
		return err
	}

	return
}

func (m *PointHeadActionFeedback) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Status
	if err = ros.DeserializeMessageField(r, "actionlib_msgs/GoalStatus", &m.Status); err != nil {
		return err
	}

	// Feedback
	if err = ros.DeserializeMessageField(r, "PointHeadFeedback", &m.Feedback); err != nil {
		return err
	}

	return
}