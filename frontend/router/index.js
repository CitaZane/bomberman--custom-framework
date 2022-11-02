import { QueueView } from "../views/Queue"
import { HomeView } from "../views/Home"
import { GameView } from "../views/Game"

export default [{
    path: "/",
    component: HomeView
},
{
    path: "/queue",
    component: QueueView
},

{
    path: "/game",
    component: GameView
}]