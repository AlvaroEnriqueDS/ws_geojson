package models

type GeoJson struct {
        Type       string     `json:"type"`
        Properties properties `json:"properties"`
        Geometry   Geometry   `json:"geometry"`
}

type properties struct {
}

type Geometry struct {
        Type        string    `json:"type"`
        Coordinates []float64 `json:"coordinates"`
}
