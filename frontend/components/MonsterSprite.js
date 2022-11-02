/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import { refs } from '../../framework/vDom/render'



export function MonsterSprite({ player }) {
    let ref = `monster-${player.name}`

    return {
        template: (
            <div ref={ref} class="monster"></div>
        ),
        onMounted: ()=>{
            refs[`monster-${player.name}`].style.setProperty("--y-movement", player.y)
        }
    }
}


