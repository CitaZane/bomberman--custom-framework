import jsx from "../../framework/vDom/jsx"
import { refs } from "../../framework/vDom/render";
import {store} from "../app"

function updateFire(fire, ref){
    if (!refs[ref]) return;

    refs[ref].style.setProperty("--y", fire.y);
    refs[ref].style.setProperty("--x", fire.x);
}

export function FireSprite({fire, id, parent}) {
    let ref = `${parent}-${id}`;
    var fireClass = `fire-${fire.type}`
    return {
        template: (
            <div ref={ref} class={fireClass}></div>
        ),
        onMounted: () => {
        updateFire(fire, ref);
        },
    }
}