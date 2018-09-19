module Main.Msg exposing (Msg(..))

import Common.Json exposing (EmptyResponse)
import Homepage.Homepage
import Http
import Main.User exposing (User)
import Material
import Navigation exposing (Location)
import Portfolio.Msg
import Tokens.Msg
import Window


type Msg
    = OnRouteChange Location
    | ToggleMobileNav
    | OnCheckSessionResponse (Result Http.Error User)
    | Mdl (Material.Msg Msg)
    | OnWindowResize Window.Size
    | Homepage Homepage.Homepage.Msg
    | Tokens Tokens.Msg.Msg
    | Portfolio Portfolio.Msg.Msg