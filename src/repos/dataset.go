package repos

type Dataset struct {
	values map[string]interface{}
}

func NewDataset() *Dataset {
	dataset := new(Dataset)
	dataset.values = make(map[string]interface{})
	return dataset
}

func (d *Dataset) Set(key string, value interface{}) {
	d.values[key] = value
}

func (d *Dataset) Get(key string) interface{} {
	return d.values[key]
}

func (d *Dataset) GetInt(key string) int {
	return d.values[key].(int)
}

func (d *Dataset) GetString(key string) string {
	return d.values[key].(string)
}
