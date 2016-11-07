package client

// Preset define the set of parameters of a given preset
type Preset struct {
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Container   string      `json:"container,omitempty"`
	RateControl string      `json:"rateControl,omitempty"`
	Video       VideoPreset `json:"video"`
	Audio       AudioPreset `json:"audio"`
}

// VideoPreset define the set of parameters for video on a given preset
type VideoPreset struct {
	Width         string `json:"width,omitempty"`
	Height        string `json:"height,omitempty"`
	Codec         string `json:"codec,omitempty"`
	Bitrate       string `json:"bitrate,omitempty"`
	GopSize       string `json:"gopSize,omitempty"`
	GopMode       string `json:"gopMode,omitempty"`
	Profile       string `json:"profile,omitempty"`
	ProfileLevel  string `json:"profileLevel,omitempty"`
	InterlaceMode string `json:"interlaceMode,omitempty"`
}

// AudioPreset define the set of parameters for audio on a given preset
type AudioPreset struct {
	Codec   string `json:"codec,omitempty"`
	Bitrate string `json:"bitrate,omitempty"`
}

// GetPresets returns a list of presets
func (c *Client) GetPresets() ([]Preset, error) {
	var result []Preset
	err := c.do("GET", "/presets", nil, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPreset return details of a given preset name
func (c *Client) GetPreset(presetName string) (*Preset, error) {
	var result *Preset
	err := c.do("GET", "/presets/"+presetName, nil, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreatePreset creates a new preset
func (c *Client) CreatePreset(preset *Preset) (*Preset, error) {
	var result *Preset
	err := c.do("POST", "/presets", preset, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeletePreset removes a preset based on its preset name
func (c *Client) DeletePreset(presetName string) error {
	return c.do("DELETE", "/presets/"+presetName, nil, nil)
}
