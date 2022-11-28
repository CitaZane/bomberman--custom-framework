/* @jsx jsx */
import jsx from "../framework/vDom/jsx";
import { MonsterSprite } from "../components/MonsterSprite";
import { ExplosionSprite } from "../components/ExplosionSprite";
import { BombSprite } from "../components/BombSprite";
import { GameMap } from "../components/GameMap";
import { store } from "../app";
import { PowerUp } from "../components/PowerUp";
import { ChatRoom } from "../components/ChatRoom";
import { GamePlayers } from "../components/gamePlayers/GamePlayers";
import { ws } from "../websocket";

function leaveGame() {
  window.location.href = window.location.origin + "/";
}

export function GameView() {
  if (!ws) {
    window.location.href = window.location.origin + "/";
    return;
  }
  let players = store.state.players;
  let powerUps = store.state.powerUps;

  const allExplosions = players.reduce((prev, current) => {
    return prev.concat(current?.explosions);
  }, []);
  let allBombs = players.reduce((prev, current) => {
    return prev.concat(current?.bombs);
  }, []);

  return {
    template: (
      <div id="game-layout">
        <div id="game-wrapper">
          <div id="game">
            <GameMap />

            {players.map((player, i) => (
              <MonsterSprite player={player} id={i} />
            ))}

            {allExplosions.map((explosion, i) => (
              <ExplosionSprite explosion={explosion} />
            ))}

            {allBombs.map((bomb, i) => (
              <BombSprite bomb={bomb} id={i} />
            ))}

            {powerUps.map((powerUp, i) => (
              <PowerUp powerUp={powerUp} id={i} />
            ))}

            {store.state.winner && <h2> WINNER IS : {store.state.winner}</h2>}
          </div>

          {/* Position element relative to the game-wrapper */}
          <div class="left-sidebar">
            <GamePlayers />
            <button id="quit" class="btn" onClick={leaveGame}>
              Quit
            </button>
          </div>

          {/* Position element relative to the game-wrapper */}
          <ChatRoom />
        </div>
      </div>
    ),
  };
}
