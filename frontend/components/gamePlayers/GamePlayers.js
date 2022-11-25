/* @jsx jsx */
import jsx from "../../framework/vDom/jsx";
import { store } from "../../app";
import { PlayerCard } from "./PlayerCard";

export function GamePlayers() {
  const players = store.state.players;
  const placeholderPlayers = [];
  if (players.length < 4) {
    let playersToAdd = 4 - players.length;
    for (let i = 0; i < playersToAdd; i++) {
      placeholderPlayers.push(<PlayerCard />);
    }
  }

  return {
    template: (
      <div class="players">
        {players.map((player, i) => (
          <PlayerCard player={player} id={i} />
        ))}
        {placeholderPlayers}
      </div>
    ),
  };
}
