import jsx from "../../framework/vDom/jsx";
import { store } from "../app";

export function GameMap(){
    let map_data = store.state.map
    return {
        
        template: (
            <div id="gamemap">              
              {map_data.map((tile) => {
                switch (tile){
                  case 2 :
                    return <div class="wall_block"></div>;
                  case 1:
                    return <div class="breakable_block"></div>;
                  default:
                    return <div class="floor_block"></div>;
                }          
            })}
            </div>
          ),
    }
}