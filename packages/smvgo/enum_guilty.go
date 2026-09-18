package smvgo

type EnumGuilty interface {
	String() string
	final()
}

type enumGuilty struct {
	value string
}

func (eg enumGuilty) String() string {
	return eg.value
}

func (eg enumGuilty) final() {
}

func DeveloperByEnumGuilty() EnumGuilty {
	return enumGuilty{value: "Developer"}
}

func DeviceByEnumGuilty() EnumGuilty {
	return enumGuilty{value: "Device"}
}

func UserByEnumGuilty() EnumGuilty {
	return enumGuilty{value: "User"}
}
