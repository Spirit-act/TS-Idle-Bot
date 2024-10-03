package teamspeak3

type ChannelList struct {
	Channels []Channel `json:"body"`
	Status   Status    `json:"status"`
}

type Channel struct {
	Id                              string `json:"cid,omitempty"`
	Banner_gfx_url                  string `json:"channel_banner_gfx_url,omitempty"`
	Banner_mode                     string `json:"channel_banner_mode,omitempty"`
	Codec                           string `json:"channel_codec,omitempty"`
	Codec_is_unencrypted            string `json:"channel_codec_is_unencrypted,omitempty"`
	Codec_latency_factor            string `json:"channel_codec_latency_factor,omitempty"`
	Codec_quality                   string `json:"channel_codec_quality,omitempty"`
	Delete_delay                    string `json:"channel_delete_delay,omitempty"`
	Description                     string `json:"channel_description,omitempty"`
	Filepath                        string `json:"channel_filepath,omitempty"`
	Flag_default                    string `json:"channel_flag_default,omitempty"`
	Flag_maxclients_unlimited       string `json:"channel_flag_maxclients_unlimited,omitempty"`
	Flag_maxfamilyclients_inherited string `json:"channel_flag_maxfamilyclients_inherited,omitempty"`
	Flag_maxfamilyclients_unlimited string `json:"channel_flag_maxfamilyclients_unlimited,omitempty"`
	Flag_password                   string `json:"channel_flag_password,omitempty"`
	Flag_permanent                  string `json:"channel_flag_permanent,omitempty"`
	Flag_semi_permanent             string `json:"channel_flag_semi_permanent,omitempty"`
	Forced_silence                  string `json:"channel_forced_silence,omitempty"`
	Icon_id                         string `json:"channel_icon_id,omitempty"`
	Maxclients                      string `json:"channel_maxclients,omitempty"`
	Maxfamilyclients                string `json:"channel_maxfamilyclients,omitempty"`
	Name                            string `json:"channel_name,omitempty"`
	Name_phonetic                   string `json:"channel_name_phonetic,omitempty"`
	Needed_talk_power               string `json:"channel_needed_talk_power,omitempty"`
	Order                           string `json:"channel_order,omitempty"`
	Password                        string `json:"channel_password,omitempty"`
	Security_salt                   string `json:"channel_security_salt,omitempty"`
	Topic                           string `json:"channel_topic,omitempty"`
	Unique_identifier               string `json:"channel_unique_identifier,omitempty"`
	Pid                             string `json:"pid,omitempty"`
	Seconds_empty                   string `json:"seconds_empty,omitempty"`
}
