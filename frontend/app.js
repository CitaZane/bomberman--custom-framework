import createRouter from "../framework/router";
import createStore from "../framework/store";
import routes from "./router";
import storeObj from "./store/index";
import { refs } from '../framework/vDom/render'

const store = createStore(storeObj);
const router = createRouter(routes);

export class Player {
    constructor(name) {
        this.name = name
        this.type = 0
        this.frame = 0
        this.frameCount = 5 //total frames per animation
        this.staggerFrames = 5 //slow animation down
        this.x = 0
        this.y = 0
        this.speed = 1
    }
}
export class Game {
    constructor() {
        this.players = []
    }

    addPlayer(player) {
        this.players.push(player)
    }
}


export function setupGame() {
    document.addEventListener("keyup", (e) => {
        store.dispatch('registerKeyUp', e.code)
    })
    document.addEventListener("keydown", (e) => {
        if (!gameStarted) {
            gameStarted = true
            // animate()
        }
        store.dispatch('registerKeyDown', e.code)
    })
}


export { store, router };
