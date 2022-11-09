/* @jsx jsx */
import jsx from "../../framework/vDom/jsx";
import { MonsterSprite } from "../components/MonsterSprite";
import { BombSprite } from "../components/BombSprite";
import { GameMap } from "../components/GameMap";
import { store } from "../app";

export function GameView() {
  let players = store.state.players;

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
      </div>
    ),
  };
}
