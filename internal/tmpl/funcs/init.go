package funcs

type Maps map[string]any

func (m Maps) Add(key string, value any) map[string]any {
	m[key] = value
	return m
}

func (m Maps) AddMaps(map2 Maps) Maps {
	for key, value := range map2 {
		m[key] = value
	}
	return m
}
