import jsx from "../../framework/vDom/jsx"
import { refs } from "../../framework/vDom/render";
import {store} from "../app"

function updateFire(fire, ref, explosionId){
    if (!refs[ref]) return;
    // set position on screen
    refs[ref].style.setProperty("--y", fire.y);
    refs[ref].style.setProperty("--x", fire.x);

    // configure class list
    refs[ref].classList.remove(...refs[ref].classList);
    var fireClass = `fire-${fire.type}`
    refs[ref].classList.add(fireClass)

    // set animation start time 
    // if time already saved, then use that
    let startTime = store.state.explosionTime[explosionId]
    if (!startTime){
        let time = refs[ref].getAnimations()[0].timeline.currentTime
        store.dispatch('addStartTime', {time, explosionId})
        setTimeout(()=>{store.dispatch('removeStartTime', explosionId)}, 900)
    }else{
        refs[ref].getAnimations()[0].startTime = startTime
    }

}

export function FireSprite({fire, id, parent}) {
    let ref = `${parent}-${id}`;
    return {
        template: (
            <div ref={ref}></div>
        ),
        onMounted: () => {
        updateFire(fire, ref, parent);
        },
    }
}