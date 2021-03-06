package gobot

type testDriver struct {
	Driver
}

func (me *testDriver) Start() bool { return true }

type testAdaptor struct {
	Adaptor
}

func (me *testAdaptor) Finalize() bool   { return true }
func (me *testAdaptor) Connect() bool    { return true }
func (me *testAdaptor) Disconnect() bool { return true }
func (me *testAdaptor) Reconnect() bool  { return true }

func newTestDriver(name string) *testDriver {
	d := new(testDriver)
	d.Name = name
	d.Commands = []string{
		"DriverCommand1",
		"DriverCommand2",
		"DriverCommand3",
	}
	return d
}
func newTestAdaptor(name string) *testAdaptor {
	a := new(testAdaptor)
	a.Name = name
	return a
}

func newTestRobot(name string) Robot {
	return Robot{
		Name:        name,
		Connections: []Connection{newTestAdaptor("Connection 1"), newTestAdaptor("Connection 2"), newTestAdaptor("Connection 3")},
		Devices:     []Device{newTestDriver("Device 1"), newTestDriver("Device 2"), newTestDriver("Device 3")},
		Commands: map[string]interface{}{
			"Command1": func() {},
			"Command2": func() {},
		},
	}
}
