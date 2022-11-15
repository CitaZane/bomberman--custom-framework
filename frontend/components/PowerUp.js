import jsx from "../../framework/vDom/jsx";

export function PowerUp({ powerUp, id }) {
  let ref = `powerup-${powerUp.type}-${id}`;
  // let className = `powerup ${powerUp.type}`;
  let className = `powerup increase-bombs`;

  return {
    template: <div class={className} ref={ref}></div>,

    onMounted(refs) {
      refs[ref].style.setProperty("--x-pos", `${powerUp.x}px`);
      refs[ref].style.setProperty("--y-pos", `${powerUp.y}px`);
    },
  };
}
