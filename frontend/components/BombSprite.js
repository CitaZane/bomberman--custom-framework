/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import {refs} from '../../framework/vDom/render'

const BOMB_SIZE = 64
let gameFrame = 0;

let bomb = {
    frame:0,
    frameCount:2,
    staggerFrames:10,
}

function animate(){
    if (refs.bomb){
        // change monster animation 
        refs.bomb.style.setProperty("--x", bomb.frame*BOMB_SIZE)
    }
    if(gameFrame%bomb.staggerFrames == 0){
        // loop for aniation based on frames
        bomb.frame = bomb.frame >= bomb.frameCount? 0:bomb.frame +=1;
    }
    gameFrame++;

    requestAnimationFrame(animate);
}

export function BombSprite() {
    animate()
    return {
        template: (
            <div ref="bomb" class="bomb" ></div>
        )
    }
}