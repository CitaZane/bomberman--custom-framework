/* @jsx jsx */
import jsx from "../../framework/vDom/jsx";
import { refs } from "../../framework/vDom/render";

function updateMovement(player, id) {
  if (!refs[`monster-${id}`]) return;
  refs[`monster-${id}`].style.setProperty("--y-movement", player.y);
  refs[`monster-${id}`].style.setProperty("--x-movement", player.x);

    updateAnimation(player,id)

}

function updateAnimation(player, id){
    refs[`monster-${id}`].classList.add(`monster-${player.movement}`)
}

export function MonsterSprite({ player, id }) {
  let ref = `monster-${id}`;
  //   updateMovement(player, id);

  return {
    template: <div ref={ref} class="monster" id={ref}></div>,
    onMounted: () => {
      updateMovement(player, id);
    },
  };
}
