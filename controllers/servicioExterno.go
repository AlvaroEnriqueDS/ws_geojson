package controllers

import (
        "encoding/json"
        "fmt"
        "github.com/alvaroenriqueds/ws_potencie/models"
        "github.com/labstack/echo"
        "github.com/labstack/gommon/log"
        "golang.org/x/net/websocket"
)

func Tracking(c echo.Context) error  {
        track := models.ServicioExterno{}
        msg := models.Error{}


        err := c.Bind(&track)
        if err != nil {
                fmt.Println("Error al aceptar la data entrante :/track")
                log.Error(err)

                msg.Message = "Error al aceptar la data entrante"
                msg.ErrorCode = "Error code"
                msg.Error = err.Error()

                return c.JSON(400, msg)
        }

        if track.Nickname == "" {
                msg.Message = "El campo nickname no puede estar vacio"
                msg.ErrorCode = "Error code"
                msg.Error = ""

                return c.JSON(400, msg)
        }

        //control := control(track)

        port := 5050
        //port2 := 9494
        origin := fmt.Sprintf(
                "http://localhost:%d/", port)
        url := fmt.Sprintf(
                "ws://localhost:%d/ws", port)
        ws, err := websocket.Dial(url, "", origin)
        if err != nil {
                fmt.Println("Error al crear el websocket :/track ")
                log.Error(err)

                msg.Message = "No se pudo crear la conexion WS"
                msg.ErrorCode = "Error code"
                msg.Error = err.Error()

                return c.JSON(500, msg)
        }

        geoj := models.GeoJson{
              Type: "Feature",
              Geometry: models.Geometry{
                      Type: "Point",
                      Coordinates: []float64{track.Longitude, track.Latitude}},
        }


        j, err := json.Marshal(&geoj)
        if _, err := ws.Write(j); err != nil {
                fmt.Println("Error al convertir geojson a bytes :/track")
                log.Error(err)

                msg.Message = "No se pudo convertir la data a GeoJson"
                msg.ErrorCode = "Error code"
                msg.Error = err.Error()

                return c.JSON(500, msg)
        }

        return c.JSON(200, geoj)
}

func control(track models.ServicioExterno)  (regreso []models.Error){
        var msg models.Error
        if track.Nickname ==  "" {
                msg.Message = "El campo nickname no puede entrar vacio"
                msg.ErrorCode = "Error code"

                regreso = append(regreso, msg)

        }
        /*
        if track.Latitude == 0 {
                msg.Message = ""
                msg.ErrorCode = "Error code"
                msg.Error = ""

                regreso = append(regreso, msg)
        }
        if track.Longitude == 0 {
                msg.Message = "El campo nickname no puede entrar vacio"
                msg.ErrorCode = "Error code"
                msg.Error = ""

                regreso = append(regreso, msg)
        }
        */

        return regreso
}
