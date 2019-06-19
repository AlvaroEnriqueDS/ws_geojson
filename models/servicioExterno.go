package models


type ServicioExterno struct {
        Nickname  string  `json:"nickname"`
        Latitude  float64 `json:"latitude"`
        Longitude float64 `json:"longitude"`
        Acuraccy  float64 `json:"acuraccy"`
        Data      Data    `json:"data"`
}

type Data struct {
        Name string `json:"name"`
}
