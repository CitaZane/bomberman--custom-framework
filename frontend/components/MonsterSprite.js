/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import { refs } from '../../framework/vDom/render'

const MONSTER_SIZE = 64

function updateMovement(player){
    if(!refs[`monster-${player.id}`])return
    refs[`monster-${player.id}`].style.setProperty("--y-movement", player.y)
    refs[`monster-${player.id}`].style.setProperty("--x-movement", player.x)

    refs[`monster-${player.id}`].style.setProperty("--y", player.state*MONSTER_SIZE)
    refs[`monster-${player.id}`].style.setProperty("--x", player.frame*MONSTER_SIZE)
}

export function MonsterSprite({ player }) {
    let ref = `monster-${player.id}`
     updateMovement(player)

    return {
        template: (
            <div ref={ref} class="monster" id={ref}></div>
        ),
        onMounted: ()=>{
            updateMovement(player)
        }
    }
}


