package mao

type TaskOption struct {
	TaskID         string        `bson:"task_id" json:"task_id"`
	TestData       []interface{} `bson:"test_data" json:"test_data"`
	EnableTestData int64         `bson:"enable_test_data" json:"enable_test_data"`
	ExecuteCount   int64         `bson:"execute_count" json:"execute_count"`
	IntervalTime   int64         `bson:"interval_time" json:"interval_time"`
	EnvID          string        `bson:"env_id" json:"env_id"`
	EnvInfo        EnvInfo       `bson:"envInfo" json:"envInfo"`
	ServersData    interface{}   `bson:"serversData" json:"serversData"`
	GlobalVars     interface{}   `bson:"globalVars" json:"globalVars"`
	GlobalRequest  GlobalRequest `bson:"globalRequest" json:"globalRequest"`
}

type EnvInfo struct {
	EnvID     string            `bson:"env_id" json:"env_id"`
	ListSort  int64             `bson:"list_sort" json:"list_sort"`
	Name      string            `bson:"name" json:"name"`
	PreURL    string            `bson:"pre_url" json:"pre_url"`
	List      interface{}       `bson:"list" json:"list,omitempty"`
	PreUrls   map[string]string `bson:"pre_urls" json:"pre_urls"`
	CreatedAt int64             `bson:"created_at" json:"created_at"`
}

type EnvInfoList struct {
	Value        string `bson:"value" json:"value"`
	CurrentValue string `bson:"current_value" json:"current_value"`
	Description  string `bson:"description" json:"description"`
}

type Header struct {
	IsChecked   string `bson:"is_checked" json:"is_checked"`
	Type        string `bson:"type" json:"type"`
	Key         string `bson:"key" json:"key"`
	Value       string `bson:"value" json:"value"`
	Description string `bson:"description" json:"description"`
}
type Kv struct {
	Key   string `bson:"key" json:"key"`
	Value string `bson:"value" json:"value"`
}
type Bearer struct {
	Key string `bson:"key" json:"key"`
}
type Basic struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}
type Auth struct {
	Type   string `bson:"type" json:"type"`
	Kv     Kv     `bson:"kv" json:"kv"`
	Bearer Bearer `bson:"bearer" json:"bearer"`
	Basic  Basic  `bson:"basic" json:"basic"`
}
type GlobalRequest struct {
	Header []Header      `bson:"header" json:"header"`
	Query  []interface{} `bson:"query" json:"query"`
	Body   []interface{} `bson:"body" json:"body"`
	Auth   Auth          `bson:"auth" json:"auth"`
}
