import createRouter from "../framework/router";
import createStore from "../framework/store";
import routes from "./router";
import storeObj from "./store/index";
import { refs } from '../framework/vDom/render'

const store = createStore(storeObj);
const router = createRouter(routes);


document.addEventListener("keyup", (e) => {
    store.dispatch('registerKeyUp', e.code)
})
document.addEventListener("keydown", (e) => {
    store.dispatch('registerKeyDown', e.code)
})



const MONSTER_SIZE = 64
let gameFrame = 0;

let monster = {
    type: 0,
    frame: 0,
    frameCount: 5, //total frames per animation
    staggerFrames: 5, //slow animation down
    x: 0,
    y: 0,
    speed: 1
}
const BOMB_SIZE = 64

let bomb = {
    frame: 0,
    frameCount: 2,
    staggerFrames: 5,
}

const monsterStates = { ArrowLeft: 0, ArrowDown: 1, ArrowRight: 2, ArrowUp: 3 }


function animate() {
    let inputs = store.state.inputs
    if (refs.monster) {
        // change monster animation

        if (inputs['ArrowLeft']) {
            monster.type = monsterStates['ArrowLeft']
        } else if (inputs['ArrowDown']) {
            monster.type = monsterStates['ArrowDown']
        } else if (inputs['ArrowUp']) {
            monster.type = monsterStates['ArrowUp']
        } else if (inputs['ArrowRight']) {
            monster.type = monsterStates['ArrowRight']
        }
        // x and y moves through sprite sheet
        refs.monster.style.setProperty("--y", monster.type * MONSTER_SIZE)
        refs.monster.style.setProperty("--x", monster.frame * MONSTER_SIZE)

        // move around screen
        let x = store.state.monster.x
        let y = store.state.monster.y
        switch (monster.type) {
            case 0:
                monster.x = x - monster.speed
                break;
            case 1:
                monster.y = y + monster.speed
                break;
            case 2:
                monster.x = x + monster.speed
                break;
            case 3:
                monster.y = y - monster.speed
                break;
            default:
                break;
        }
        refs.monster.style.setProperty("--y-movement", monster.y)
        refs.monster.style.setProperty("--x-movement", monster.x)
        store.dispatch('updateMonsterX', monster.x)
        store.dispatch('updateMonsterY', monster.y)

    }
    if (gameFrame % monster.staggerFrames == 0) {
        // loop for aniation based on frames
        monster.frame = monster.frame >= monster.frameCount ? 0 : monster.frame += 1;
    }
    // update bomb
    if (inputs['Space']) {
        let bombActive = store.state.bomb.drop;
        if (!bombActive) {
            store.dispatch('updateBombDrop', true)
            setTimeout(() => {
                store.dispatch('updateBombDrop', false)

            }, 2000)
            // // change bomb animation 
            refs.bomb.style.setProperty("--y-pos", monster.y)
            refs.bomb.style.setProperty("--x-pos", monster.x)
        }
    }
    if (refs.bomb) {
        // change bomb animation 
        refs.bomb.style.setProperty("--x", bomb.frame * BOMB_SIZE)

        if (gameFrame % bomb.staggerFrames == 0) {
            // loop for aniation based on frames
            bomb.frame = bomb.frame >= bomb.frameCount ? 0 : bomb.frame += 1;
        }

    }
    gameFrame++;

    requestAnimationFrame(animate);
}
animate()

export { store, router };
