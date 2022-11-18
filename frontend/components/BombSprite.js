/* @jsx jsx */
import jsx from "../framework/vDom/jsx";

export function BombSprite({ bomb, id }) {
  return {
    template: <div ref={`bomb-${id}`} class="bomb"></div>,
    onMounted(refs) {
      refs[`bomb-${id}`].style.setProperty("--x-pos", bomb.x);
      refs[`bomb-${id}`].style.setProperty("--y-pos", bomb.y);
    },
  };
}
