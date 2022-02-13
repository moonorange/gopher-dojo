package dict

func Keys(m map[string] bool) []string {
	ks := []string{}
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
