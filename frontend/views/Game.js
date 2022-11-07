/* @jsx jsx */
import jsx from "../../framework/vDom/jsx";
import { MonsterSprite } from "../components/MonsterSprite";
import { GameMap } from "../components/GameMap";
import { store } from "../app";
// import { players } from "../websocket";
export function GameView() {
  let players = store.state.players;
  return {
    template: (
      <div id="home">
        {/* <h1>Game View</h1> */}
        <GameMap />
        {players.map((player, i) => (
          <MonsterSprite player={player} id={i} />
        ))}

        {bombs.map(<BombSprite />)}
      </div>
    ),
  };
}
