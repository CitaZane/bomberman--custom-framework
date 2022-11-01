import jsx from "../../framework/vDom/jsx"
import {refs} from '../../framework/vDom/render'

const FIRE_SIZE = 320;
let gameFrame = 0;

let fire ={
    frame:0,
    frameCount: 3, //total frames per animation
    staggerFrames:10, //slow animation down
    blow:1, //if blow = 1 then blow active, if 0, fire goes out
}

function animate(){
    if (refs.fire){
        // change fire animation
        refs.fire.style.setProperty("--x", fire.frame*FIRE_SIZE)
    }
    if(gameFrame%fire.staggerFrames == 0){
        if(fire.blow == 1){
            fire.frame +=1;
        }else{
            fire.frame -=1;
        }
        // reverse blow
        if (fire.frame == fire.frameCount){
            fire.blow = 0;
        }
        if (fire.frame == 0){
            fire.blow = 1;
        }
    }
    gameFrame++;

    requestAnimationFrame(animate);
}

export function FireSprite() {
    animate()
    return {
        template: (
            <div ref="fire" class="fire" ></div>
        )
    }
}