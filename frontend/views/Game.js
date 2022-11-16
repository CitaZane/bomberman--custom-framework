/* @jsx jsx */
import jsx from "../../framework/vDom/jsx";
import { MonsterSprite } from "../components/MonsterSprite";
import { ExplosionSprite } from "../components/ExplosionSprite";
import { BombSprite } from "../components/BombSprite";
import {GameMap} from "../components/GameMap";
import { store } from "../app";

export function GameView() {
  let players = store.state.players;
  // let allExplosions = store.state.explosions;
  const allExplosions = players.reduce((prev, current) => {
    return prev.concat(current?.explosions);
  }, []);
  const allBombs = players.reduce((prev, current) => {
    return prev.concat(current?.bombs);
  }, []);

  return {
    template: (
      <div id="home">
        <GameMap />
        
        <div id="pu-1"></div>
        <div id="pu-2"></div>
        <div id="pu-3"></div>
        {players.map((player, i) => (
          <MonsterSprite player={player} id={i} />
        ))}

        {allExplosions.map((explosion, i) => (
          <ExplosionSprite explosion={explosion} />))}
        
        {allBombs.map((bomb, i) => (
          <BombSprite bomb={bomb} id={i} />
        ))}

         

      </div>
    ),
  };
}
