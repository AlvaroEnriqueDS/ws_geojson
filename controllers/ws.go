package controllers

import (
        "github.com/labstack/echo"
        "github.com/olahol/melody"
)


var mel *melody.Melody

func init()  {
        mel = melody.New()
}

func WebSocket(c echo.Context) error {
        mel.HandleRequest(c.Response().Writer, c.Request())
        //mel.HandleConnect(hConnect)
        //mel.HandleDisconnect(hDisconnect)
        mel.HandleMessage(hMessage)
        return nil
}

func hMessage(s *melody.Session, msg []byte) {
        //mel.Broadcast(msg)
        mel.BroadcastFilter(msg, func(q *melody.Session) bool {
                return q.Request.URL.Path == s.Request.URL.Path
        })
}
