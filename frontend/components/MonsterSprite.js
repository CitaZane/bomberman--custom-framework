/* @jsx jsx */
import jsx from "../../framework/vDom/jsx"
import { refs } from '../../framework/vDom/render'



export function MonsterSprite({ y }) {
    // refs.monster.style.setProperty("--y-movement", y)

    return {
        template: (
            <div ref="monster" class="monster"></div>
        )
    }
}


