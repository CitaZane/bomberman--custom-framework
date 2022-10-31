import { GameView } from "../views/Game"
import { HomeView } from "../views/Home"

export default [{
    path: "/",
    component: HomeView
},

{
    path: "/game",
    component: GameView
}]