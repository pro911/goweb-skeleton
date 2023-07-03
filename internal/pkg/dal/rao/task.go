package rao

type Callback struct {
	TaskID         string `json:"task_id"`
	TaskName       string `json:"task_name"`
	VID            int64  `json:"vid"`
	VersionName    string `json:"version_name"`
	ExecAt         int64  `json:"exec_at"`
	ExecType       int64  `json:"exec_type"`
	ProjectID      int64  `json:"project_id"`
	LocalProjectID int64  `json:"local_project_id"`
	ProjectName    string `json:"project_name"`
	TeamID         int64  `json:"team_id"`
	LocalTeamID    string `json:"local_team_id"`
	TeamName       string `json:"team_name"`
}

type EnvInfo struct {
	EnvID     string        `json:"env_id"`
	ListSort  int64         `json:"list_sort"`
	Name      string        `json:"name"`
	PreURL    string        `json:"pre_url"`
	List      []interface{} `json:"list"`
	PreUrls   []interface{} `json:"pre_urls"`
	CreatedAt int64         `json:"created_at"`
}

type Option struct {
	TaskID         string        `json:"task_id"`
	TestData       []interface{} `json:"test_data"`
	EnableTestData int64         `json:"enable_test_data"`
	ExecuteCount   int64         `json:"execute_count"`
	IntervalTime   int64         `json:"interval_time"`
	EnvID          string        `json:"env_id"`
	EnvInfo        EnvInfo       `json:"envInfo"`
	ServersData    []interface{} `json:"serversData"`
	GlobalVars     GlobalVars    `json:"globalVars"`
	GlobalRequest  GlobalRequest `json:"globalRequest"`
	Collections    []Collections `json:"collections"`
}

type TaskSendResp struct {
	Callback   Callback     `json:"callback"`
	TestEvents []TestEvents `json:"test_events"`
	Option     Option       `json:"option"`
}

type Data struct {
	ParentID string `json:"parent_id"`
	TargetID string `json:"target_id"`
	Name     string `json:"name"`
	Method   string `json:"method"`
	URL      string `json:"url"`
}
type Kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type Bearer struct {
	Key string `json:"key"`
}
type Basic struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Digest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Realm     string `json:"realm"`
	Nonce     string `json:"nonce"`
	Algorithm string `json:"algorithm"`
	Qop       string `json:"qop"`
	Nc        string `json:"nc"`
	Cnonce    string `json:"cnonce"`
	Opaque    string `json:"opaque"`
}
type Hawk struct {
	AuthID             string `json:"authId"`
	AuthKey            string `json:"authKey"`
	Algorithm          string `json:"algorithm"`
	User               string `json:"user"`
	Nonce              string `json:"nonce"`
	ExtraData          string `json:"extraData"`
	App                string `json:"app"`
	Delegation         string `json:"delegation"`
	Timestamp          string `json:"timestamp"`
	IncludePayloadHash int64  `json:"includePayloadHash"`
}
type Awsv4 struct {
	AccessKey          string `json:"accessKey"`
	SecretKey          string `json:"secretKey"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	SessionToken       string `json:"sessionToken"`
	AddAuthDataToQuery int64  `json:"addAuthDataToQuery"`
}
type Ntlm struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Domain              string `json:"domain"`
	Workstation         string `json:"workstation"`
	DisableRetryRequest int64  `json:"disableRetryRequest"`
}
type Edgegrid struct {
	AccessToken   string `json:"accessToken"`
	ClientToken   string `json:"clientToken"`
	ClientSecret  string `json:"clientSecret"`
	Nonce         string `json:"nonce"`
	Timestamp     string `json:"timestamp"`
	BaseURi       string `json:"baseURi"`
	HeadersToSign string `json:"headersToSign"`
}
type Oauth1 struct {
	ConsumerKey          string `json:"consumerKey"`
	ConsumerSecret       string `json:"consumerSecret"`
	SignatureMethod      string `json:"signatureMethod"`
	AddEmptyParamsToSign int64  `json:"addEmptyParamsToSign"`
	IncludeBodyHash      int64  `json:"includeBodyHash"`
	AddParamsToHeader    int64  `json:"addParamsToHeader"`
	Realm                string `json:"realm"`
	Version              string `json:"version"`
	Nonce                string `json:"nonce"`
	Timestamp            string `json:"timestamp"`
	Verifier             string `json:"verifier"`
	Callback             string `json:"callback"`
	TokenSecret          string `json:"tokenSecret"`
	Token                string `json:"token"`
}
type Auth struct {
	Type     string   `json:"type"`
	Kv       Kv       `json:"kv"`
	Bearer   Bearer   `json:"bearer"`
	Basic    Basic    `json:"basic"`
	Digest   Digest   `json:"digest"`
	Hawk     Hawk     `json:"hawk"`
	Awsv4    Awsv4    `json:"awsv4"`
	Ntlm     Ntlm     `json:"ntlm"`
	Edgegrid Edgegrid `json:"edgegrid"`
	Oauth1   Oauth1   `json:"oauth1"`
}
type RawSchema struct {
	Type string `json:"type"`
}
type Body struct {
	Mode      string        `json:"mode"`
	Parameter []interface{} `json:"parameter"`
	Raw       string        `json:"raw"`
	RawPara   []interface{} `json:"raw_para"`
	RawSchema RawSchema     `json:"raw_schema"`
}
type Event struct {
	PreScript string `json:"pre_script"`
	Test      string `json:"test"`
}
type Header struct {
	Parameter []interface{} `json:"parameter"`
}
type Parameter struct {
	Description string `json:"description"`
	IsChecked   int64  `json:"is_checked"`
	Key         string `json:"key"`
	Type        string `json:"type"`
	NotNull     int64  `json:"not_null"`
	FieldType   string `json:"field_type"`
	Value       string `json:"value"`
}
type Query struct {
	Parameter []Parameter `json:"parameter"`
}
type Cookie struct {
	Parameter []interface{} `json:"parameter"`
}
type Resful struct {
	Parameter []interface{} `json:"parameter"`
}
type Request struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	Auth        Auth   `json:"auth"`
	Body        Body   `json:"body"`
	Event       Event  `json:"event"`
	Header      Header `json:"header"`
	Query       Query  `json:"query"`
	Cookie      Cookie `json:"cookie"`
	Resful      Resful `json:"resful"`
}
type AiExpect struct {
	List             []interface{} `json:"list"`
	NoneMathExpectID string        `json:"none_math_expect_id"`
}
type APIData struct {
	TargetID         string        `json:"target_id"`
	ParentID         string        `json:"parent_id"`
	ProjectID        string        `json:"project_id"`
	AutoImportID     string        `json:"auto_import_id"`
	Mark             string        `json:"mark"`
	TargetType       string        `json:"target_type"`
	ExampleType      string        `json:"example_type"`
	Name             string        `json:"name"`
	Method           string        `json:"method"`
	Sort             int64         `json:"sort"`
	TypeSort         int64         `json:"type_sort"`
	UpdateDay        int64         `json:"update_day"`
	UpdateDtime      int64         `json:"update_dtime"`
	Status           int64         `json:"status"`
	BakID            int64         `json:"bak_id"`
	Version          int64         `json:"version"`
	IsPublish        int64         `json:"is_publish"`
	Publisher        int64         `json:"publisher"`
	PublishDtime     int64         `json:"publish_dtime"`
	Hash             interface{}   `json:"hash"`
	IsChanged        int64         `json:"is_changed"`
	CreateDtime      int64         `json:"create_dtime"`
	IsDoc            int64         `json:"is_doc"`
	Vid              int64         `json:"vid"`
	CommitID         int64         `json:"commit_id"`
	IsJoin           int64         `json:"is_join"`
	ModifierID       string        `json:"modifier_id"`
	IsExample        int64         `json:"is_example"`
	Tags             []interface{} `json:"tags"`
	Request          Request       `json:"request"`
	AiExpect         AiExpect      `json:"ai_expect"`
	EnableAiExpect   int64         `json:"enable_ai_expect"`
	MockURL          string        `json:"mock_url"`
	IsFirstMockPath  int64         `json:"is_first_mock_path"`
	EnableServerMock int64         `json:"enable_server_mock"`
	IsLocked         int64         `json:"is_locked"`
	AttributeInfo    []interface{} `json:"attribute_info"`
	ServerID         string        `json:"server_id"`
}
type TestEvents struct {
	EventID       string  `json:"event_id"`
	ParentEventID string  `json:"parent_event_id"`
	ProjectID     string  `json:"project_id"`
	Type          string  `json:"type"`
	Enabled       int64   `json:"enabled"`
	Sort          int64   `json:"sort"`
	Data          Data    `json:"data"`
	APIData       APIData `json:"apiData"`
	TaskID        string  `json:"task_id"`
}
type Tttt struct {
	Value        string `json:"value"`
	CurrentValue string `json:"current_value"`
	Description  string `json:"description"`
}

type GlobalVars struct {
	Key string `json:"key"`
}
type GlobalHeader struct {
	IsChecked   string `json:"is_checked"`
	Type        string `json:"type"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type GlobalRequest struct {
	Header []GlobalHeader `json:"header"`
	Query  []interface{}  `json:"query"`
	Body   []interface{}  `json:"body"`
	Auth   Auth           `json:"auth"`
}
type Collections struct {
	TargetID         string        `json:"target_id"`
	ParentID         string        `json:"parent_id"`
	ProjectID        string        `json:"project_id,omitempty"`
	AutoImportID     string        `json:"auto_import_id,omitempty"`
	Mark             string        `json:"mark,omitempty"`
	TargetType       string        `json:"target_type"`
	ExampleType      string        `json:"example_type,omitempty"`
	Name             string        `json:"name"`
	Method           string        `json:"method"`
	Sort             int64         `json:"sort"`
	TypeSort         int64         `json:"type_sort,omitempty"`
	UpdateDay        int64         `json:"update_day,omitempty"`
	UpdateDtime      int64         `json:"update_dtime,omitempty"`
	Status           int64         `json:"status,omitempty"`
	BakID            int64         `json:"bak_id,omitempty"`
	Version          int64         `json:"version,omitempty"`
	IsPublish        int64         `json:"is_publish,omitempty"`
	Publisher        int64         `json:"publisher,omitempty"`
	PublishDtime     int64         `json:"publish_dtime,omitempty"`
	Hash             interface{}   `json:"hash,omitempty"`
	IsChanged        int64         `json:"is_changed,omitempty"`
	CreateDtime      int64         `json:"create_dtime"`
	IsDoc            int64         `json:"is_doc,omitempty"`
	Vid              int64         `json:"vid,omitempty"`
	CommitID         int64         `json:"commit_id,omitempty"`
	IsJoin           int64         `json:"is_join,omitempty"`
	ModifierID       string        `json:"modifier_id,omitempty"`
	IsExample        int64         `json:"is_example,omitempty"`
	Tags             []interface{} `json:"tags,omitempty"`
	Request          Request       `json:"request,omitempty"`
	AiExpect         AiExpect      `json:"ai_expect,omitempty"`
	EnableAiExpect   int64         `json:"enable_ai_expect,omitempty"`
	MockURL          string        `json:"mock_url,omitempty"`
	IsFirstMockPath  int64         `json:"is_first_mock_path,omitempty"`
	EnableServerMock int64         `json:"enable_server_mock,omitempty"`
	IsLocked         int64         `json:"is_locked,omitempty"`
	AttributeInfo    []interface{} `json:"attribute_info,omitempty"`
	ServerID         string        `json:"server_id"`
	TaskID           string        `json:"task_id"`
}
