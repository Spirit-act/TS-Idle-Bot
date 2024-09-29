package teamspeak3

import "strconv"

const (
	CLIENT_USER  string = "0"
	CLIENT_QUERY string = "1"
)

type ClientList struct {
	Clients []Client `json:"body"`
	Status  Status   `json:"status"`
}

type Client struct {
	Id                                              string `json:"clid,omitempty"`
	Channel_id                                      string `json:"cid,omitempty"`
	Away                                            string `json:"client_away,omitempty"`
	Away_message                                    string `json:"client_away_message,omitempty"`
	Badges                                          string `json:"client_badges,omitempty"`
	Base64HashClientUID                             string `json:"client_base64HashClientUID,omitempty"`
	Channel_group_id                                string `json:"client_channel_group_id,omitempty"`
	Channel_group_inherited_channel_id              string `json:"client_channel_group_inherited_channel_id,omitempty"`
	Country                                         string `json:"client_country,omitempty"`
	Created                                         string `json:"client_created,omitempty"`
	Database_id                                     string `json:"client_database_id,omitempty"`
	Default_channel                                 string `json:"client_default_channel,omitempty"`
	Default_token                                   string `json:"client_default_token,omitempty"`
	Description                                     string `json:"client_description,omitempty"`
	Flag_avatar                                     string `json:"client_flag_avatar,omitempty"`
	Icon_id                                         string `json:"client_icon_id,omitempty"`
	Idle_time                                       string `json:"client_idle_time,omitempty"`
	Input_hardware                                  string `json:"client_input_hardware,omitempty"`
	Input_muted                                     string `json:"client_input_muted,omitempty"`
	Integrations                                    string `json:"client_integrations,omitempty"`
	Is_channel_commander                            string `json:"client_is_channel_commander,omitempty"`
	Is_priority_speaker                             string `json:"client_is_priority_speaker,omitempty"`
	Is_recording                                    string `json:"client_is_recording,omitempty"`
	Is_talker                                       string `json:"client_is_talker,omitempty"`
	Lastconnected                                   string `json:"client_lastconnected,omitempty"`
	Login_name                                      string `json:"client_login_name,omitempty"`
	Meta_data                                       string `json:"client_meta_data,omitempty"`
	Month_bytes_downloaded                          string `json:"client_month_bytes_downloaded,omitempty"`
	Month_bytes_uploaded                            string `json:"client_month_bytes_uploaded,omitempty"`
	Myteamspeak_avatar                              string `json:"client_myteamspeak_avatar,omitempty"`
	Myteamspeak_id                                  string `json:"client_myteamspeak_id,omitempty"`
	Needed_serverquery_view_power                   string `json:"client_needed_serverquery_view_power,omitempty"`
	Nickname                                        string `json:"client_nickname,omitempty"`
	Nickname_phonetic                               string `json:"client_nickname_phonetic,omitempty"`
	Output_hardware                                 string `json:"client_output_hardware,omitempty"`
	Output_muted                                    string `json:"client_output_muted,omitempty"`
	Outputonly_muted                                string `json:"client_outputonly_muted,omitempty"`
	Platform                                        string `json:"client_platform,omitempty"`
	Security_hash                                   string `json:"client_security_hash,omitempty"`
	Servergroups                                    string `json:"client_servergroups,omitempty"`
	Signed_badges                                   string `json:"client_signed_badges,omitempty"`
	Talk_power                                      string `json:"client_talk_power,omitempty"`
	Talk_request                                    string `json:"client_talk_request,omitempty"`
	Talk_request_msg                                string `json:"client_talk_request_msg,omitempty"`
	Total_bytes_downloaded                          string `json:"client_total_bytes_downloaded,omitempty"`
	Total_bytes_uploaded                            string `json:"client_total_bytes_uploaded,omitempty"`
	Totalconnections                                string `json:"client_totalconnections,omitempty"`
	Type                                            string `json:"client_type,omitempty"`
	Unique_identifier                               string `json:"client_unique_identifier,omitempty"`
	Version                                         string `json:"client_version,omitempty"`
	Version_sign                                    string `json:"client_version_sign,omitempty"`
	Connection_bandwidth_received_last_minute_total string `json:"connection_bandwidth_received_last_minute_total,omitempty"`
	Connection_bandwidth_received_last_second_total string `json:"connection_bandwidth_received_last_second_total,omitempty"`
	Connection_bandwidth_sent_last_minute_total     string `json:"connection_bandwidth_sent_last_minute_total,omitempty"`
	Connection_bandwidth_sent_last_second_total     string `json:"connection_bandwidth_sent_last_second_total,omitempty"`
	Connection_bytes_received_total                 string `json:"connection_bytes_received_total,omitempty"`
	Connection_bytes_sent_total                     string `json:"connection_bytes_sent_total,omitempty"`
	Connection_client_ip                            string `json:"connection_client_ip,omitempty"`
	Connection_connected_time                       string `json:"connection_connected_time,omitempty"`
	Connection_filetransfer_bandwidth_received      string `json:"connection_filetransfer_bandwidth_received,omitempty"`
	Connection_filetransfer_bandwidth_sent          string `json:"connection_filetransfer_bandwidth_sent,omitempty"`
	Connection_packets_received_total               string `json:"connection_packets_received_total,omitempty"`
	Connection_packets_sent_total                   string `json:"connection_packets_sent_total,omitempty"`
}

func (c *Client) ToIdle(threshold int) bool {
	idle_time, err := strconv.Atoi(c.Idle_time)
	if err != nil {
		return false
	}

	return idle_time > threshold
}

func (c *Client) RecentActive(threshold int) bool {
	idle_time, err := strconv.Atoi(c.Idle_time)
	if err != nil {
		return false
	}

	return idle_time <= threshold
}
