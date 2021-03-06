module Actions.Update exposing (update)

import Actions.Command exposing (loadActionsCmd)
import Actions.Model exposing (Model)
import Actions.Msg exposing (Msg(..))
import Debug
import Main.Context exposing (Context)
import Material
import Model.Actions exposing (actionsDecoder)


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        OnLoadActionsResponse resp ->
            case resp of
                Ok actions ->
                    { model | actions = Just actions, error = Nothing } ! []

                Err error ->
                    Debug.log (toString error)
                        { model
                            | actions = Nothing
                            , error = Just "Error loading your actions, please try again!"
                        }
                        ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
