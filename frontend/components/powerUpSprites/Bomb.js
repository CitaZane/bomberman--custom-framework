import jsx from "../../../framework/vDom/jsx";

export function PowerUpBombSprite() {
  return {
    template: <div class="powerup increase-bombs"></div>,
  };
}

// export function PowerUpBombSprite({ powerup, id }) {
//   let ref = `explosion-${id}`;
//   return {
//     template: (
//       <div ref={ref} class="explosion">
//         {explosion.map((fire, id) => (
//           <FireSprite fire={fire} id={id} parent={ref} />
//         ))}
//       </div>
//     ),
//   };
// }
