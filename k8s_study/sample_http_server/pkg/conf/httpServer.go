package conf

type ServiceSettingS struct {
	BindIP  string `yaml:"bindIP"`
	Port    int    `yaml:"port"`
	LogPath string `yaml:"logPath"`
}

type HttpServerS struct {
	ServiceSetting ServiceSettingS
}

func (s *Setting) ReadHttpServer(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
