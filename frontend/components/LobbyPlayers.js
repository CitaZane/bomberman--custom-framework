/* @jsx jsx */
import { store } from "../app";
import jsx from "../framework/vDom/jsx";

function leaveLobby() {
  window.location.href = window.location.origin + "/";
}

export function LobbyPlayers() {
  const lobbyPlayersNames = store.state.lobbyPlayersNames;
  return {
    template: (
      <div id="lobby-players">
        <h3>Waiting players to join...</h3>
        <ul id="lobby-players_list">
          {lobbyPlayersNames.map((playerName, i) => {
            console.log("playerName", playerName);
            return (
              <li>
                <div class="player-monster" id={`monster-${i}`}></div>
                <p class={`player-name monster-${i}__color`}>{playerName}</p>
              </li>
            );
          })}
        </ul>

        <button class="btn" onClick={leaveLobby}>
          Leave lobby
        </button>
      </div>
    ),
  };
}
