import jsx from "../../framework/vDom/jsx";
import { FireSprite } from "./FireSprite";

export function ExplosionSprite({ explosion, id }) {
  console.log("Explosion:", explosion, id);
  let ref = `explosion-${id}`;
  return {
    template: (
      <div ref={ref} class="explosion">
        {explosion.map((fire, id) => (
          <FireSprite fire={fire} id={id} parent={ref} />
        ))}
      </div>
    ),
  };
}
