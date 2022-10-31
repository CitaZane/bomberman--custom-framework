/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import {refs} from '../../framework/vDom/render'
import {store} from "../app"

const MONSTER_SIZE = 266
let gameFrame = 0;

let monster ={
    type:0, 
    frame:0,
    frameCount: 5, //total frames per animation
    staggerFrames:5, //slow animation down
}

function animate(){
    if (refs.monster){
        // change monster animation
        monster.type = store.state.monster.type;     
        refs.monster.style.setProperty("--y", monster.type*MONSTER_SIZE)
        refs.monster.style.setProperty("--x", monster.frame*MONSTER_SIZE)
    }
    if(gameFrame%monster.staggerFrames == 0){
        // loop for aniation based on frames
        monster.frame = monster.frame >= monster.frameCount? 0:monster.frame +=1;
    }
    gameFrame++;

    requestAnimationFrame(animate);
}

export function MonsterSprite() {
    animate()
    return {
        template: (
            <div ref="monster" class="monster" ></div>
        )
    }
}