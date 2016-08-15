// Code generated by protoc-gen-go.
// source: executor.proto
// DO NOT EDIT!

/*
Package executorproto is a generated protocol buffer package.

It is generated from these files:
	executor.proto

It has these top-level messages:
	Event
	Call
*/
package executorproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import mesos_v1 "github.com/jimenez/mesoscon-demo/lib/mesosproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible event types, followed by message definitions if
// applicable.
type Event_Type int32

const (
	// This must be the first enum value in this list, to
	// ensure that if 'type' is not set, the default value
	// is UNKNOWN. This enables enum values to be added
	// in a backwards-compatible way. See: MESOS-4997.
	Event_UNKNOWN      Event_Type = 0
	Event_SUBSCRIBED   Event_Type = 1
	Event_LAUNCH       Event_Type = 2
	Event_KILL         Event_Type = 3
	Event_ACKNOWLEDGED Event_Type = 4
	Event_MESSAGE      Event_Type = 5
	Event_ERROR        Event_Type = 6
	// Received when the agent asks the executor to shutdown/kill itself.
	// The executor is then required to kill all its active tasks, send
	// `TASK_KILLED` status updates and gracefully exit. The executor
	// should terminate within a `MESOS_EXECUTOR_SHUTDOWN_GRACE_PERIOD`
	// (an environment variable set by the agent upon executor startup);
	// it can be configured via `ExecutorInfo.shutdown_grace_period`. If
	// the executor fails to do so, the agent will forcefully destroy the
	// container where the executor is running. The agent would then send
	// `TASK_LOST` updates for any remaining active tasks of this executor.
	//
	// NOTE: The executor must not assume that it will always be allotted
	// the full grace period, as the agent may decide to allot a shorter
	// period and failures / forcible terminations may occur.
	//
	// TODO(alexr): Consider adding a duration field into the `Shutdown`
	// message so that the agent can communicate when a shorter period
	// has been allotted.
	Event_SHUTDOWN Event_Type = 7
)

var Event_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUBSCRIBED",
	2: "LAUNCH",
	3: "KILL",
	4: "ACKNOWLEDGED",
	5: "MESSAGE",
	6: "ERROR",
	7: "SHUTDOWN",
}
var Event_Type_value = map[string]int32{
	"UNKNOWN":      0,
	"SUBSCRIBED":   1,
	"LAUNCH":       2,
	"KILL":         3,
	"ACKNOWLEDGED": 4,
	"MESSAGE":      5,
	"ERROR":        6,
	"SHUTDOWN":     7,
}

func (x Event_Type) Enum() *Event_Type {
	p := new(Event_Type)
	*p = x
	return p
}
func (x Event_Type) String() string {
	return proto.EnumName(Event_Type_name, int32(x))
}
func (x *Event_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Event_Type_value, data, "Event_Type")
	if err != nil {
		return err
	}
	*x = Event_Type(value)
	return nil
}
func (Event_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Possible call types, followed by message definitions if
// applicable.
type Call_Type int32

const (
	// See comments above on `Event::Type` for more details on this enum value.
	Call_UNKNOWN   Call_Type = 0
	Call_SUBSCRIBE Call_Type = 1
	Call_UPDATE    Call_Type = 2
	Call_MESSAGE   Call_Type = 3
)

var Call_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUBSCRIBE",
	2: "UPDATE",
	3: "MESSAGE",
}
var Call_Type_value = map[string]int32{
	"UNKNOWN":   0,
	"SUBSCRIBE": 1,
	"UPDATE":    2,
	"MESSAGE":   3,
}

func (x Call_Type) Enum() *Call_Type {
	p := new(Call_Type)
	*p = x
	return p
}
func (x Call_Type) String() string {
	return proto.EnumName(Call_Type_name, int32(x))
}
func (x *Call_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Call_Type_value, data, "Call_Type")
	if err != nil {
		return err
	}
	*x = Call_Type(value)
	return nil
}
func (Call_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// *
// Executor event API.
//
// An event is described using the standard protocol buffer "union"
// trick, see https://developers.google.com/protocol-buffers/docs/techniques#union.
type Event struct {
	// Type of the event, indicates which optional field below should be
	// present if that type has a nested message definition.
	// Enum fields should be optional, see: MESOS-4997.
	Type             *Event_Type         `protobuf:"varint,1,opt,name=type,enum=mesos.v1.executor.Event_Type" json:"type,omitempty"`
	Subscribed       *Event_Subscribed   `protobuf:"bytes,2,opt,name=subscribed" json:"subscribed,omitempty"`
	Acknowledged     *Event_Acknowledged `protobuf:"bytes,3,opt,name=acknowledged" json:"acknowledged,omitempty"`
	Launch           *Event_Launch       `protobuf:"bytes,4,opt,name=launch" json:"launch,omitempty"`
	Kill             *Event_Kill         `protobuf:"bytes,5,opt,name=kill" json:"kill,omitempty"`
	Message          *Event_Message      `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	Error            *Event_Error        `protobuf:"bytes,7,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Event) GetType() Event_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Event_UNKNOWN
}

func (m *Event) GetSubscribed() *Event_Subscribed {
	if m != nil {
		return m.Subscribed
	}
	return nil
}

func (m *Event) GetAcknowledged() *Event_Acknowledged {
	if m != nil {
		return m.Acknowledged
	}
	return nil
}

func (m *Event) GetLaunch() *Event_Launch {
	if m != nil {
		return m.Launch
	}
	return nil
}

func (m *Event) GetKill() *Event_Kill {
	if m != nil {
		return m.Kill
	}
	return nil
}

func (m *Event) GetMessage() *Event_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Event) GetError() *Event_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// First event received when the executor subscribes.
// The 'id' field in the 'framework_info' will be set.
type Event_Subscribed struct {
	ExecutorInfo     *mesos_v1.ExecutorInfo  `protobuf:"bytes,1,req,name=executor_info" json:"executor_info,omitempty"`
	FrameworkInfo    *mesos_v1.FrameworkInfo `protobuf:"bytes,2,req,name=framework_info" json:"framework_info,omitempty"`
	AgentInfo        *mesos_v1.AgentInfo     `protobuf:"bytes,3,req,name=agent_info" json:"agent_info,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *Event_Subscribed) Reset()                    { *m = Event_Subscribed{} }
func (m *Event_Subscribed) String() string            { return proto.CompactTextString(m) }
func (*Event_Subscribed) ProtoMessage()               {}
func (*Event_Subscribed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Event_Subscribed) GetExecutorInfo() *mesos_v1.ExecutorInfo {
	if m != nil {
		return m.ExecutorInfo
	}
	return nil
}

func (m *Event_Subscribed) GetFrameworkInfo() *mesos_v1.FrameworkInfo {
	if m != nil {
		return m.FrameworkInfo
	}
	return nil
}

func (m *Event_Subscribed) GetAgentInfo() *mesos_v1.AgentInfo {
	if m != nil {
		return m.AgentInfo
	}
	return nil
}

// Received when the framework attempts to launch a task. Once
// the task is successfully launched, the executor must respond with
// a TASK_RUNNING update (See TaskState in v1/mesos.proto).
type Event_Launch struct {
	Task             *mesos_v1.TaskInfo `protobuf:"bytes,1,req,name=task" json:"task,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Event_Launch) Reset()                    { *m = Event_Launch{} }
func (m *Event_Launch) String() string            { return proto.CompactTextString(m) }
func (*Event_Launch) ProtoMessage()               {}
func (*Event_Launch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *Event_Launch) GetTask() *mesos_v1.TaskInfo {
	if m != nil {
		return m.Task
	}
	return nil
}

// Received when the scheduler wants to kill a specific task. Once
// the task is terminated, the executor should send a TASK_KILLED
// (or TASK_FAILED) update. The terminal update is necessary so
// Mesos can release the resources associated with the task.
type Event_Kill struct {
	TaskId *mesos_v1.TaskID `protobuf:"bytes,1,req,name=task_id" json:"task_id,omitempty"`
	// If set, overrides any previously specified kill policy for this task.
	// This includes 'TaskInfo.kill_policy' and 'Executor.kill.kill_policy'.
	// Can be used to forcefully kill a task which is already being killed.
	KillPolicy       *mesos_v1.KillPolicy `protobuf:"bytes,2,opt,name=kill_policy" json:"kill_policy,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *Event_Kill) Reset()                    { *m = Event_Kill{} }
func (m *Event_Kill) String() string            { return proto.CompactTextString(m) }
func (*Event_Kill) ProtoMessage()               {}
func (*Event_Kill) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *Event_Kill) GetTaskId() *mesos_v1.TaskID {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *Event_Kill) GetKillPolicy() *mesos_v1.KillPolicy {
	if m != nil {
		return m.KillPolicy
	}
	return nil
}

// Received when the agent acknowledges the receipt of status
// update. Schedulers are responsible for explicitly acknowledging
// the receipt of status updates that have 'update.status().uuid()'
// field set. Unacknowledged updates can be retried by the executor.
// They should also be sent by the executor whenever it
// re-subscribes.
type Event_Acknowledged struct {
	TaskId           *mesos_v1.TaskID `protobuf:"bytes,1,req,name=task_id" json:"task_id,omitempty"`
	Uuid             []byte           `protobuf:"bytes,2,req,name=uuid" json:"uuid,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Event_Acknowledged) Reset()                    { *m = Event_Acknowledged{} }
func (m *Event_Acknowledged) String() string            { return proto.CompactTextString(m) }
func (*Event_Acknowledged) ProtoMessage()               {}
func (*Event_Acknowledged) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

func (m *Event_Acknowledged) GetTaskId() *mesos_v1.TaskID {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *Event_Acknowledged) GetUuid() []byte {
	if m != nil {
		return m.Uuid
	}
	return nil
}

// Received when a custom message generated by the scheduler is
// forwarded by the agent. Note that this message is not
// interpreted by Mesos and is only forwarded (without reliability
// guarantees) to the executor. It is up to the scheduler to retry
// if the message is dropped for any reason.
type Event_Message struct {
	Data             []byte `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Event_Message) Reset()                    { *m = Event_Message{} }
func (m *Event_Message) String() string            { return proto.CompactTextString(m) }
func (*Event_Message) ProtoMessage()               {}
func (*Event_Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 4} }

func (m *Event_Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Received in case the executor sends invalid calls (e.g.,
// required values not set).
// TODO(arojas): Remove this once the old executor driver is no
// longer supported. With HTTP API all errors will be signaled via
// HTTP response codes.
type Event_Error struct {
	Message          *string `protobuf:"bytes,1,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Event_Error) Reset()                    { *m = Event_Error{} }
func (m *Event_Error) String() string            { return proto.CompactTextString(m) }
func (*Event_Error) ProtoMessage()               {}
func (*Event_Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 5} }

func (m *Event_Error) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

// *
// Executor call API.
//
// Like Event, a Call is described using the standard protocol buffer
// "union" trick (see above).
type Call struct {
	// Identifies the executor which generated this call.
	ExecutorId  *mesos_v1.ExecutorID  `protobuf:"bytes,1,req,name=executor_id" json:"executor_id,omitempty"`
	FrameworkId *mesos_v1.FrameworkID `protobuf:"bytes,2,req,name=framework_id" json:"framework_id,omitempty"`
	// Type of the call, indicates which optional field below should be
	// present if that type has a nested message definition.
	// In case type is SUBSCRIBED, no message needs to be set.
	// See comments on `Event::Type` above on the reasoning behind this field being optional.
	Type             *Call_Type      `protobuf:"varint,3,opt,name=type,enum=mesos.v1.executor.Call_Type" json:"type,omitempty"`
	Subscribe        *Call_Subscribe `protobuf:"bytes,4,opt,name=subscribe" json:"subscribe,omitempty"`
	Update           *Call_Update    `protobuf:"bytes,5,opt,name=update" json:"update,omitempty"`
	Message          *Call_Message   `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Call) Reset()                    { *m = Call{} }
func (m *Call) String() string            { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()               {}
func (*Call) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Call) GetExecutorId() *mesos_v1.ExecutorID {
	if m != nil {
		return m.ExecutorId
	}
	return nil
}

func (m *Call) GetFrameworkId() *mesos_v1.FrameworkID {
	if m != nil {
		return m.FrameworkId
	}
	return nil
}

func (m *Call) GetType() Call_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Call_UNKNOWN
}

func (m *Call) GetSubscribe() *Call_Subscribe {
	if m != nil {
		return m.Subscribe
	}
	return nil
}

func (m *Call) GetUpdate() *Call_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *Call) GetMessage() *Call_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

// Request to subscribe with the agent. If subscribing after a disconnection,
// it must include a list of all the tasks and updates which haven't been
// acknowledged by the scheduler.
type Call_Subscribe struct {
	UnacknowledgedTasks   []*mesos_v1.TaskInfo `protobuf:"bytes,1,rep,name=unacknowledged_tasks" json:"unacknowledged_tasks,omitempty"`
	UnacknowledgedUpdates []*Call_Update       `protobuf:"bytes,2,rep,name=unacknowledged_updates" json:"unacknowledged_updates,omitempty"`
	XXX_unrecognized      []byte               `json:"-"`
}

func (m *Call_Subscribe) Reset()                    { *m = Call_Subscribe{} }
func (m *Call_Subscribe) String() string            { return proto.CompactTextString(m) }
func (*Call_Subscribe) ProtoMessage()               {}
func (*Call_Subscribe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Call_Subscribe) GetUnacknowledgedTasks() []*mesos_v1.TaskInfo {
	if m != nil {
		return m.UnacknowledgedTasks
	}
	return nil
}

func (m *Call_Subscribe) GetUnacknowledgedUpdates() []*Call_Update {
	if m != nil {
		return m.UnacknowledgedUpdates
	}
	return nil
}

// Notifies the scheduler that a task has transitioned from one
// state to another. Status updates should be used by executors
// to reliably communicate the status of the tasks that they
// manage. It is crucial that a terminal update (see TaskState
// in v1/mesos.proto) is sent to the scheduler as soon as the task
// terminates, in order for Mesos to release the resources allocated
// to the task. It is the responsibility of the scheduler to
// explicitly acknowledge the receipt of a status update. See
// 'Acknowledged' in the 'Events' section above for the semantics.
type Call_Update struct {
	Status           *mesos_v1.TaskStatus `protobuf:"bytes,1,req,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *Call_Update) Reset()                    { *m = Call_Update{} }
func (m *Call_Update) String() string            { return proto.CompactTextString(m) }
func (*Call_Update) ProtoMessage()               {}
func (*Call_Update) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

func (m *Call_Update) GetStatus() *mesos_v1.TaskStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

// Sends arbitrary binary data to the scheduler. Note that Mesos
// neither interprets this data nor makes any guarantees about the
// delivery of this message to the scheduler.
// See 'Message' in the 'Events' section.
type Call_Message struct {
	Data             []byte `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Call_Message) Reset()                    { *m = Call_Message{} }
func (m *Call_Message) String() string            { return proto.CompactTextString(m) }
func (*Call_Message) ProtoMessage()               {}
func (*Call_Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 2} }

func (m *Call_Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "mesos.v1.executor.Event")
	proto.RegisterType((*Event_Subscribed)(nil), "mesos.v1.executor.Event.Subscribed")
	proto.RegisterType((*Event_Launch)(nil), "mesos.v1.executor.Event.Launch")
	proto.RegisterType((*Event_Kill)(nil), "mesos.v1.executor.Event.Kill")
	proto.RegisterType((*Event_Acknowledged)(nil), "mesos.v1.executor.Event.Acknowledged")
	proto.RegisterType((*Event_Message)(nil), "mesos.v1.executor.Event.Message")
	proto.RegisterType((*Event_Error)(nil), "mesos.v1.executor.Event.Error")
	proto.RegisterType((*Call)(nil), "mesos.v1.executor.Call")
	proto.RegisterType((*Call_Subscribe)(nil), "mesos.v1.executor.Call.Subscribe")
	proto.RegisterType((*Call_Update)(nil), "mesos.v1.executor.Call.Update")
	proto.RegisterType((*Call_Message)(nil), "mesos.v1.executor.Call.Message")
	proto.RegisterEnum("mesos.v1.executor.Event_Type", Event_Type_name, Event_Type_value)
	proto.RegisterEnum("mesos.v1.executor.Call_Type", Call_Type_name, Call_Type_value)
}

func init() { proto.RegisterFile("executor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x94, 0xdb, 0x4e, 0xdb, 0x4e,
	0x10, 0xc6, 0xff, 0x89, 0x1d, 0x87, 0x4c, 0x42, 0xfe, 0xee, 0x96, 0x82, 0x65, 0x51, 0x0e, 0xe9,
	0x99, 0x0a, 0x17, 0x50, 0xa5, 0x5e, 0x20, 0xb5, 0x4a, 0x62, 0x17, 0x10, 0x21, 0xa0, 0x1c, 0xd4,
	0xcb, 0xc8, 0x24, 0x0b, 0x58, 0x09, 0x76, 0xea, 0x03, 0x94, 0x8b, 0x3e, 0x46, 0x9f, 0xa1, 0x0f,
	0xd7, 0x97, 0xe8, 0xec, 0xda, 0xc6, 0x29, 0xb1, 0xd5, 0xde, 0x7a, 0x7e, 0xdf, 0xee, 0xec, 0x37,
	0xdf, 0x18, 0xaa, 0xf4, 0x1b, 0x1d, 0x06, 0xbe, 0xe3, 0x6a, 0x53, 0xd7, 0xf1, 0x1d, 0xf2, 0xe8,
	0x9a, 0x7a, 0x8e, 0xa7, 0xdd, 0xec, 0x6a, 0x71, 0x41, 0x2d, 0x87, 0x9f, 0x78, 0xbd, 0xf6, 0x4b,
	0x82, 0x82, 0x71, 0x43, 0x6d, 0x9f, 0xbc, 0x05, 0xd1, 0xbf, 0x9b, 0x52, 0x25, 0xb7, 0x91, 0x7b,
	0x5d, 0xdd, 0x7b, 0xaa, 0xcd, 0x09, 0x35, 0xce, 0x69, 0x3d, 0x84, 0xc8, 0x07, 0x00, 0x2f, 0x38,
	0xf7, 0x86, 0xae, 0x75, 0x4e, 0x47, 0x4a, 0x1e, 0x25, 0xe5, 0xbd, 0x67, 0x99, 0x92, 0xee, 0x3d,
	0x4a, 0xf6, 0xa1, 0x62, 0x0e, 0xc7, 0xb6, 0x73, 0x3b, 0xa1, 0xa3, 0x4b, 0x94, 0x0a, 0x5c, 0xfa,
	0x22, 0x53, 0x5a, 0x9f, 0x81, 0xc9, 0x3b, 0x90, 0x26, 0x66, 0x60, 0x0f, 0xaf, 0x14, 0x91, 0xcb,
	0xd6, 0x33, 0x65, 0x2d, 0x8e, 0xb1, 0x37, 0x8d, 0xad, 0xc9, 0x44, 0x29, 0x70, 0x3c, 0xfb, 0x4d,
	0xc7, 0x08, 0x91, 0x5d, 0x28, 0x62, 0xdd, 0x33, 0x2f, 0xa9, 0x22, 0x71, 0x7e, 0x23, 0x93, 0x3f,
	0x09, 0x39, 0xb2, 0x0d, 0x05, 0xea, 0xba, 0x8e, 0xab, 0x14, 0xb9, 0x60, 0x2d, 0x53, 0x60, 0x30,
	0x4a, 0xfd, 0x91, 0x03, 0x98, 0xf1, 0x62, 0x1b, 0x16, 0x63, 0x6c, 0x60, 0xd9, 0x17, 0x0e, 0x5a,
	0x9f, 0xc7, 0x53, 0x96, 0x93, 0x53, 0x8c, 0xa8, 0x7c, 0x84, 0x55, 0x7c, 0x7d, 0xf5, 0xc2, 0x35,
	0xaf, 0xe9, 0xad, 0xe3, 0x8e, 0x43, 0x3e, 0xcf, 0xf9, 0x95, 0x84, 0xff, 0x1c, 0xd7, 0xb9, 0xe0,
	0x15, 0x00, 0x36, 0x69, 0xfb, 0x21, 0x2c, 0x70, 0xf8, 0x71, 0x02, 0xd7, 0x59, 0x8d, 0x81, 0xea,
	0x16, 0x48, 0x91, 0x61, 0x1b, 0x18, 0x02, 0xd3, 0x1b, 0x47, 0x9d, 0x90, 0x04, 0xee, 0xe1, 0x57,
	0xce, 0xf6, 0x40, 0xe4, 0x6e, 0x6d, 0x42, 0x91, 0x91, 0x03, 0x6b, 0x14, 0xc1, 0xf2, 0x03, 0x58,
	0x27, 0x6f, 0xa0, 0xcc, 0xdc, 0x1f, 0x4c, 0x9d, 0x89, 0x35, 0xbc, 0x8b, 0x52, 0xb2, 0x94, 0x60,
	0xec, 0x9c, 0x33, 0x5e, 0x53, 0x3f, 0x41, 0xe5, 0x8f, 0x49, 0xff, 0xc3, 0xe9, 0x15, 0x10, 0x83,
	0xc0, 0x1a, 0x71, 0x13, 0x2a, 0xea, 0x0a, 0x14, 0xe3, 0xa1, 0x60, 0x61, 0x64, 0xfa, 0x26, 0x17,
	0x56, 0x54, 0x05, 0xf3, 0xcd, 0xcc, 0x27, 0xff, 0x27, 0xe3, 0x65, 0x95, 0x52, 0xed, 0x2b, 0x88,
	0x3c, 0xcb, 0x65, 0x28, 0xf6, 0xdb, 0xc7, 0xed, 0xd3, 0x2f, 0x6d, 0xf9, 0x3f, 0x52, 0xc5, 0x09,
	0xf5, 0x1b, 0xdd, 0x66, 0xe7, 0xa8, 0x61, 0xe8, 0x72, 0x8e, 0x00, 0x5a, 0x53, 0xef, 0xb7, 0x9b,
	0x87, 0x72, 0x9e, 0x2c, 0xe0, 0xd3, 0x8f, 0x5a, 0x2d, 0x59, 0x20, 0x32, 0xb6, 0xdb, 0x64, 0x92,
	0x96, 0xa1, 0x1f, 0x20, 0x27, 0xb2, 0x43, 0x4e, 0x8c, 0x6e, 0xb7, 0x7e, 0x60, 0xc8, 0x05, 0x52,
	0xc2, 0x3b, 0x3b, 0x9d, 0xd3, 0x8e, 0x2c, 0x61, 0x33, 0x0b, 0xdd, 0xc3, 0x7e, 0x4f, 0x67, 0xa7,
	0x17, 0x6b, 0x3f, 0x45, 0x10, 0x9b, 0x26, 0xba, 0x87, 0xd6, 0x24, 0xa3, 0x8f, 0xdf, 0xb8, 0x94,
	0x32, 0x78, 0x1d, 0x33, 0x5c, 0x99, 0x19, 0xfb, 0x28, 0x1a, 0xfa, 0x93, 0xb4, 0xa1, 0xeb, 0x64,
	0x2b, 0x5a, 0x62, 0x81, 0x2f, 0xf1, 0x6a, 0x4a, 0x1e, 0xd9, 0xf5, 0xe1, 0x0e, 0xbf, 0x87, 0xd2,
	0xfd, 0x0e, 0x47, 0x0b, 0xb5, 0x99, 0x25, 0xb8, 0x4f, 0x2d, 0xd1, 0x40, 0x0a, 0xa6, 0xe8, 0x2f,
	0x8d, 0x96, 0x6a, 0x2d, 0x4b, 0xd2, 0xe7, 0x14, 0xd9, 0x79, 0xb8, 0x55, 0xeb, 0x59, 0x82, 0x68,
	0x7e, 0xea, 0x77, 0x28, 0x25, 0xd7, 0xed, 0xc0, 0x52, 0x60, 0xcf, 0xfe, 0x31, 0x06, 0x2c, 0x17,
	0x1e, 0x3a, 0x26, 0xa4, 0x07, 0x94, 0x7c, 0x84, 0xe5, 0x07, 0x8a, 0xb0, 0x5f, 0x0f, 0x9d, 0x13,
	0xfe, 0xde, 0xb0, 0x8a, 0x0f, 0x8c, 0x5a, 0x7f, 0x0e, 0x92, 0xe7, 0x9b, 0x7e, 0xe0, 0xcd, 0xcf,
	0x87, 0xdd, 0xd6, 0xe5, 0xb5, 0xb4, 0xe4, 0xf1, 0x48, 0xd6, 0xf6, 0xd3, 0xf2, 0xb5, 0x88, 0x8f,
	0x8b, 0xf3, 0x15, 0xc6, 0xab, 0x7f, 0xa6, 0xd7, 0x7b, 0x06, 0xc6, 0x6b, 0x26, 0x42, 0x42, 0xe3,
	0x25, 0xac, 0x3a, 0xee, 0xa5, 0x66, 0x4e, 0xcd, 0xe1, 0x15, 0x9d, 0xef, 0xb8, 0x21, 0x9d, 0xb1,
	0xdf, 0xb7, 0xf7, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xde, 0x81, 0x0f, 0x72, 0xf0, 0x05, 0x00, 0x00,
}
