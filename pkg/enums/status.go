package enums

const (
	Normal = iota
	Frozen
	Deleted
)

var statusFlags = map[int]string{
	Normal:  "Normal",
	Frozen:  "Frozen",
	Deleted: "Deleted",
}

func GetStatusDesc(value int) string {
	desc, ok := statusFlags[value]
	if ok {
		return desc
	}

	return statusFlags[Fail]
}
