/* @jsx jsx */
import jsx from "../../framework/vDom/jsx";
import { refs } from "../../framework/vDom/render";

const MONSTER_SIZE = 64;

function updateMovement(player, id) {
  if (!refs[`monster-${id}`]) return;
  refs[`monster-${id}`].style.setProperty("--y-movement", player.y);
  refs[`monster-${id}`].style.setProperty("--x-movement", player.x);

  refs[`monster-${id}`].style.setProperty("--y", player.state * MONSTER_SIZE);
  refs[`monster-${id}`].style.setProperty("--x", player.frame * MONSTER_SIZE);
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
