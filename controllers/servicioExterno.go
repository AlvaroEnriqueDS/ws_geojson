package controllers

import (
        "encoding/json"
        "fmt"
        "github.com/alvaroenriqueds/ws_potencie/configuration"
        "github.com/alvaroenriqueds/ws_potencie/constans"
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
                fmt.Println(track)
                fmt.Println("Error al volcar la data entrante")
                log.Fatal(err)
                //recover()

                msg.Message = "No se pudo recibir el json"
                msg.ErrorCode = "Error code"

                return c.JSON(400, msg)
        }
        fmt.Println(track)

        //bd
        db:= configuration.GetConnectionPsql()
        defer db.Close()

        stmt, err := db.Prepare(constans.Insert_Tracking)
        if err != nil {
                fmt.Println("Error al preparar la querie")
                log.Fatal(err)

                msg.Message = "No se acceder a la querie con exito"
                msg.ErrorCode = "Error code"

                return c.JSON(500, msg)
        }
        stmt.QueryRow(track.Nickname, track.Latitude, track.Longitude, track.Acuraccy)
        //stmt.Exec(track.Nickname, track.Latitude, track.Longitude, track.Acuraccy)



        port := 5050
        //port2 := 9494
        origin := fmt.Sprintf(
                "http://localhost:%d/", port)
        url := fmt.Sprintf(
                "ws://localhost:%d/ws", port)
        ws, err := websocket.Dial(url, "", origin)
        if err != nil {
                fmt.Println("Error al crear el websocket")
                log.Fatal(err)

                msg.Message = "No se pudo crear la conexion WS"
                msg.ErrorCode = "Error code"

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
                fmt.Println("Error al convertir geojson a bytes")
                log.Fatal(err)

                msg.Message = "No se pudo convertir la respuesta a bytes"
                msg.ErrorCode = "Error code"

                return c.JSON(500, msg)
        }

        msg.Message = "Todo salio bien"
        msg.ErrorCode= "Error code"

        return c.JSON(200, geoj)
}
