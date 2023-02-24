package resources

type ListResponse struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	TglParsed    string `json:"tgl_parsed"`
	PriceUSD     string `json:"price_usd"`
}

type RegisterResponse struct {
	ID       int    `json:"id"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
