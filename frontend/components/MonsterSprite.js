/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import {refs} from '../../framework/vDom/render'
import {store} from "../app"

const MONSTER_SIZE = 64
let gameFrame = 0;

let monster ={
    type:0, 
    frame:0,
    frameCount: 5, //total frames per animation
    staggerFrames:5, //slow animation down
    x:0,
    y:0,
    speed:1
}

const monsterStates = {ArrowLeft:0, ArrowDown:1,ArrowRight:2,ArrowUp:3}


function animate(){
    if (refs.monster){
        // change monster animation

        let inputs = store.state.inputs
        if (inputs['ArrowLeft']){
            monster.type = monsterStates['ArrowLeft']
        }else if(inputs['ArrowDown']){
             monster.type = monsterStates['ArrowDown']
        } else if(inputs['ArrowUp']){
            monster.type = monsterStates['ArrowUp']
        }  else if(inputs['ArrowRight']){
            monster.type = monsterStates['ArrowRight']
        }
        // x and y moves through sprite sheet
        refs.monster.style.setProperty("--y", monster.type*MONSTER_SIZE)
        refs.monster.style.setProperty("--x", monster.frame*MONSTER_SIZE)

        // move around screen
        switch (monster.type) {
            case 0:
                monster.x = monster.x  - monster.speed
                break;
            case 1:
                monster.y = monster.y  + monster.speed
                break;
            case 2:
                monster.x = monster.x +monster.speed
                break;
            case 3:
                monster.y = monster.y  - monster.speed
                
                break;
            default:
                break;
        }
        refs.monster.style.setProperty("--y-movement", monster.y)
        refs.monster.style.setProperty("--x-movement", monster.x)

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
            <div ref="monster" class="monster"></div>
        )
    }
}