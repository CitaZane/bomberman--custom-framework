/* @jsx jsx */
import jsx from "../../framework/vDom/jsx";
import { MonsterSprite } from "../components/MonsterSprite";
import { ExplosionSprite } from "../components/ExplosionSprite";
import { BombSprite } from "../components/BombSprite";
import {GameMap} from "../components/GameMap";
import { store } from "../app";

export function GameView() {
  let players = store.state.players;
  let explosions = store.state.explosions;
  const allBombs = players.reduce((prev, current) => {
    return prev.concat(current?.bombs);
  }, []);

  // console.log("AllBombs", allBombs)

  return {
    template: (
      <div id="home">
        <GameMap />
        {players.map((player, i) => (
          <MonsterSprite player={player} id={i} />
        ))}
        {allBombs.map((bomb, i) => (
          <BombSprite bomb={bomb} id={i} />
        ))}

         {explosions.map((explosion, i) => (
          <ExplosionSprite explosion={explosion} id={i} />))}
      </div>
    ),
  };
}
